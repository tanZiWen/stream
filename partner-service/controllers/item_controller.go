package controllers

import (
	"code.isstream.com/stream/lib"
	"github.com/gin-gonic/gin"
	log "github.com/Sirupsen/logrus"
	"code.isstream.com/stream/partner-service/service/itemservice"
)

type ItemController struct {
	lib.BaseController
}

type ItemListForm struct {
	Page     int            `form:"p"`
	Sort     int            `form:"s"`
	Category int64            `form:"c"`
}

/**
	get item list
 */
func (ctr *ItemController) GetItems(c *gin.Context) {
	//uid, err := ctr.GetUid(c); if err != nil {
	//	ctr.Unauthorized(c)
	//	return
	//}

	//var page int
	var err error
	//pageStr := c.Query("page")
	//
	//if strings.TrimSpace(pageStr) == "" {
	//	page = 1
	//} else {
	//	page, err = strconv.Atoi(pageStr); if err != nil {
	//		log.Info("bad param page", page)
	//		ctr.BadRequest(c)
	//		return
	//	}
	//}

	var form ItemListForm
	err = c.Bind(&form); if err != nil {
		log.Error("bad request ", err)
		ctr.BadRequest(c)
		return
	}

	if form.Page <= 0 {
		form.Page = 1
	}

	items, err := itemservice.GetItemList(form.Category, form.Page, form.Sort); if err != nil {
		log.Error("fail to get items ", err)
		panic(err)
	}

	ctr.SuccessPair(c, "items", items)
}