package userservice
import (
	"code.isstream.com/stream/db"
	log "github.com/Sirupsen/logrus"
	"code.isstream.com/stream/lib"
	"code.isstream.com/stream/s-user/usermodel"
	"code.isstream.com/stream/idg"
)

const (
	QUERY_USER_BY_ID = `select username, nickname, email, mcc, mobile, avatar_id from app_user where status = 1 and id= $1`
	QUERY_OAUTH_BY_USER_ID = `select openid, access_token, expire_at, source, status from user_oauth_bind where status = 1 and user_id = $1`
	QUERY_USER_BY_USERNAME_OR_EMAIL = `select * from app_user where username = $1 or email = $1`
	QUERY_USER_BY_MOBILE = `select * from app_user where mobile = $1`
)


func CreateUser(user *usermodel.User) error {
	var err error
	user.Id, err = idg.Id(); if err != nil {
		return err
	}

	_, err = db.Engine.Omit("status").Insert(user)

	return err
}

func GetUser(id int64) (user *usermodel.User, err error) {
	user = &usermodel.User{}

	isExist, err := db.Engine.Id(id).Omit("password").Get(user); if err != nil {
		return nil, err
	}

	log.Debug("get user by id ", id, user)

	if !isExist {
		return nil, nil
	}

	return user, nil
}

func GetUserByMobile(mobile int64) (has bool, user *usermodel.User, err error) {
	user = &usermodel.User{}

	has, err = db.Engine.Table(user.TableName()).Where("mobile = ?", mobile).Get(user); if err != nil {
		return false, nil, err
	}

	log.Debug("get user by mobile ", mobile, user)

	return has, user, nil
}

func GetUserByEmail(email string) (has bool, user *usermodel.User, err error) {
	user = &usermodel.User{}

	has, err = db.Engine.Table(user.TableName()).Where("email = ?", email).Get(user); if err != nil {
		return false, nil, err
	}

	log.Debug("get user by email ", email, user)

	return has, user, nil
}

/**
    1.说明：根据用户名、手机号、邮箱获取用户信息
    2.参数：identity -> 可为手机、用户名、邮箱
**/

func GetUserByIdentity(identity string) ([]*usermodel.User, error) {
	var (
		users []*usermodel.User
		err error
	)

	//若客户端传入参数为手机号,则调用手机号查询语句.否则调用电话、邮箱查询
	if lib.ValidateMobile(identity) {
		mobile, err := lib.Str2int64(identity)
		if err != nil {
			log.Error(3, "convert indentity form string to int64 error: %v", err)
			return nil, err
		}
		err = db.Engine.Sql(QUERY_USER_BY_MOBILE, mobile).Find(&users)
	}else {
		err = db.Engine.Sql(QUERY_USER_BY_USERNAME_OR_EMAIL, identity).Find(&users)
	}

	if err != nil {
		log.Error(3, "get user info by identity error: %v", err)
		return nil, err
	}

	return users, nil
}

func GetUserBriefEntityById(uid int64) (*usermodel.User, error){
	var user usermodel.User

	has, err := db.Engine.Sql("select username, nickname, email, mobile from app_user where status = 1 and id= $1", uid).Get(&user)

	if err != nil {
		log.Error(3, "select user brief info error: %v", err)
		return nil, err
	}

	if !has {
		return nil, nil
	}

	return &user, nil
}

func FindUsersWithNickname(uids []int64) ([]*usermodel.UserNickname, error){
	var users []*usermodel.UserNickname = []*usermodel.UserNickname{}
	user := &usermodel.User{}

	//ids := _domain.Int64Array(uids)

	//err := db.Engine.Sql("select id, nickname, status from app_user where id in $1", ids).Find(&users)

	err := db.Engine.Table(user.TableName()).In("id", uids).Cols("id", "nickname", "status").Find(&users)
	if err != nil {
		log.Error("failed to select user with nickname ", err)
	}

	return users, err
}

func UpdatePassword(password string, userId int64) error {
	sql := "UPDATE app_user SET password = ? where id = ?"
	_, err := db.Engine.Exec(sql, password, userId)
	return err
}

func UpdateNickname(nickname string, userId int64) error {
	sql := "UPDATE app_user SET nickname = ? where id = ?"
	_, err := db.Engine.Exec(sql, nickname, userId)
	return err
}

func UpdateUsername(username string, userId int64) error {
	sql := "UPDATE app_user SET username = ? where id = ?"
	_, err := db.Engine.Exec(sql, username, userId)
	return err
}

const _UPDATE_MOBILE_SQL = `UPDATE app_user SET mobile = ? WHERE id = ?`
func UpdateMobile(mobile int64, userId int64) error {
	_, err := db.Engine.Exec(_UPDATE_MOBILE_SQL, mobile, userId)
	return err
}

