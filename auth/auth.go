package auth
import (
    log "github.com/Sirupsen/logrus"
    "github.com/gin-gonic/gin"
    "strings"
    "code.isstream.com/stream/auth/setting"
    globalsetting "code.isstream.com/stream/setting"
    "net/http"
    "code.isstream.com/stream/m-user/userservice"
    "code.isstream.com/stream/lib"
    "strconv"
    "errors"
    "code.isstream.com/stream/domain"
    "time"
)

var Auth *GinJWTMiddleware

func Initialize() {
    if !globalsetting.Initialized {
        err := errors.New("try to initialize auth before global setting initialized")
        log.Error(err)
        panic(err)
    }

    Auth = &GinJWTMiddleware{
        Realm:      setting.Config.Realm,
        Key:        []byte(setting.Config.Secret),
        Timeout:    time.Duration(setting.Config.Timeout) * time.Hour,
        MaxRefresh: time.Duration(setting.Config.MaxRefresh) * time.Hour,
        Authenticator: authenticator,
        Authorizator: func(userId string, c *gin.Context) bool {
            return strings.TrimSpace(userId) != ""
        },
        Unauthorized: unauthorized,
    }
}

func authenticator(userId string, password string, c *gin.Context) (uid string, data map[string]interface{}, success bool) {
    if userId == "" || password == "" {
        return
    }

    identity := strings.ToLower(userId)

    users, err := userservice.GetUserByIdentity(identity); if err != nil {
        log.WithField("id", identity).Error("failed to get user", err)
        return
    }

    log.Debug("users found when login", users)

    if len(users) != 1 {
        return
    }

    user := users[0]

    if user.Status == domain.USER_STATUS_FREEZE || user.Status == domain.DB_COMMON_STATUS_DELETED {
        return
    }

    pwdParts := strings.Split(user.Password, "_")
    log.Debug("saved password, s%", user.Password)

    encryptedPassword, err := lib.EncryptPassword("rsmc")

    log.Debug("saved password, s%", encryptedPassword, err)

    encryptedPassword, err = lib.EncryptPassword(password, pwdParts[2])

    if err != nil {
        log.Error("failed to encrypt password", err)
        return
    }

    log.Debug("encrypted password", encryptedPassword)

    if encryptedPassword != user.Password {
        return
    }

    uid = strconv.FormatInt(user.Id, 10)
    data = map[string]interface{}{
        "user": user,
    }
    success = true

    return
}

func unauthorized(c *gin.Context, code int, message string) {
    log.Debug("unauthorized ", code, message)

    c.JSON(http.StatusUnauthorized, gin.H{
        "code":    code,
        "message": message,
    })
}