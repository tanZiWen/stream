package feedservice

import (
	"code.isstream.com/stream/db"
	log "github.com/Sirupsen/logrus"
	"code.isstream.com/stream/setting"
	"code.isstream.com/stream/partner-service/model"
)

const (
	_QUERY_FEEDS_BY_CUSTOMER = `SELECT f.*, au.name as name, au.avatar_url as avatar_url from feed f, app_user au
		WHERE f.customer_id = $1 and f.status = 1
		AND f.creator_id = au.id
		ORDER BY crt DESC OFFSET $2 LIMIT $3;`

	_QUERY_FEEDSRS_BY_IDS = `SELECT f.*, au.name as name, au.avatar_url as avatar_url from feed f, app_user au
		WHERE f.creator_id = au.id
		AND `
)

func GetFeedsByCustomer(customerId int64, page int) (feeds []*model.FeedRS, err error) {
	offset := (page - 1) * setting.Page.Size

	err = db.Engine.SQL(_QUERY_FEEDS_BY_CUSTOMER, customerId, offset, setting.Page.Size).Find(&feeds); if err != nil {
		log.Error("fail to query feeds by customer:", customerId, err)
		return nil, err
	}

	return
}

func GetFeedsByIds(feedIds []int64) (feeds []*model.Feed, err error) {

	err = db.Engine.In("id", feedIds).Find(&feeds); if err != nil {
		log.Error("fail to query feeds by ids:", feedIds, err)
		return nil, err
	}

	return
}

func GetFeedsRSByIds(feedIds []int64) (feeds []*model.FeedRS, err error) {
	sql := _QUERY_FEEDSRS_BY_IDS

	sql = sql + db.InOrEqual("f.id", feedIds) + ";"

	//log.Debug("query feeds sql: ", sql)

	err = db.Engine.SQL(sql).Find(&feeds); if err != nil {
		log.Error("fail to query feeds result set by ids:", feedIds, err)
		return nil, err
	}

	return
}