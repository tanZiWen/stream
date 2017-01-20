package itemservice

import (
	"code.isstream.com/stream/db"
	log "github.com/Sirupsen/logrus"
	"code.isstream.com/stream/setting"
	"code.isstream.com/stream/partner-service/model"
)

const (
	//TODO should be removed, use category or collection to filter results for practical case
	_QUERY_ALL_ITEMS = `SELECT s.* from item i, sku s
		WHERE i.leading_sku_id = s.id
		ORDER BY i.publish_time DESC OFFSET $1 LIMIT $2;`

	_QUERY_ITEMS_BY_CATEGORY = `SELECT s.* from item i, sku s
		WHERE i.leading_sku_id = s.id
		AND i.cate_ids && $1
		ORDER BY i.publish_time DESC OFFSET 0 LIMIT 10;`
)

func GetItemList(category int64, page int, sort int) (skus []*model.Sku, err error) {
	offset := (page - 1) * setting.Page.Size

	err = db.Engine.SQL(_QUERY_ALL_ITEMS, offset, setting.Page.Size).Find(&skus); if err != nil {
		log.Error("fail to query all items:", err)
		return nil, err
	}

	return
}