package controllers

import (
	"code.isstream.com/stream/lib"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"code.isstream.com/stream/partner-service/service/customerservice"
	log "github.com/Sirupsen/logrus"
	"code.isstream.com/stream/partner-service/model"
	"code.isstream.com/stream/idg"
	"code.isstream.com/stream/s-user/usermodel"
	"code.isstream.com/stream/partner-service/service/feedservice"
)

type CustomerController struct{
	lib.BaseController
}

type CustomerForm struct {
	Mobile      string        	`form:"mobile" json:"mobile"`
	FirstName   string        	`form:"firstname" json:"firstname"`
	LastName    string        	`form:"lastname" json:"lastname"`
	Gender	    int			`form:"gender" json:"gender,string"`
	Email       string        	`form:"email" json:"email"`
	Birthday    string		`form:"birthday" json:"birthday"`
	IdType      int16         	`form:"id_type" json:"id_type"`
	IdNumber    string        	`form:"id_number" json:"id_number"`
	MemberGrade int16         	`form:"member_grade" json:"member_grade"`
}

//checking for the existence of customer's mobile
func (ctr *CustomerController) CheckMobile(c *gin.Context) {
	_, err := ctr.GetUid(c); if err != nil {
		ctr.Unauthorized(c)
		return
	}

	var form CustomerForm
	err = c.Bind(&form); if err != nil {
		log.Error("bad request ", err)
		ctr.BadRequest(c)
		return
	}

	mobile := strings.TrimSpace(form.Mobile)
	//log.Debug("validate mobile:", customer.Mobile)
	if !lib.ValidateMobile(mobile) {
		ctr.Fail(c, model.MOBILE_INVALID)
		return
	}

	customers, err := customerservice.GetCustomerByMobile(mobile); if err != nil {
		log.WithField("get customer by mobile:", mobile).Error("failed to get customer", err)
		panic(err)
	}

	if customers != nil {
		ctr.Fail(c, model.MOBILE_IS_EXIST)
	} else {
		ctr.Success(c)
	}
}

//create customer by owner_id
func (ctr *CustomerController) CreateMyCustomers(c *gin.Context) {
	uid, err := ctr.GetUid(c); if err != nil {
		ctr.Unauthorized(c)
		return
	}

	var form CustomerForm
	err = c.Bind(&form); if err != nil {
		log.Error("bad request ", err)
		ctr.BadRequest(c)
		return
	}

	customer := &model.Customer{}

	customer.Mobile =  strings.TrimSpace(form.Mobile)
	//log.Debug("validate mobile:", customer.Mobile)
	if !lib.ValidateMobile(customer.Mobile) {
		ctr.Fail(c, model.MOBILE_INVALID)
		return
	}

	customers, err := customerservice.GetCustomerByMobile(customer.Mobile); if err != nil {
		log.WithField("get customer by mobile:", customer.Mobile).Error("failed to get customer", err)
		panic(err)
	}

	if customers != nil {
		ctr.Fail(c, model.MOBILE_IS_EXIST)
		return
	}

	customerId, err := idg.Id(); if err != nil {
		log.Error("failed to generate id error:", err)
		return
	}
	log.Debug("customerId", customerId)
	customer.Id = customerId
	customer.FirstName = strings.TrimSpace(form.FirstName)
	//log.Debug("customer.FirstName:", customer.FirstName)
	customer.LastName = strings.TrimSpace(form.LastName)
	customer.Name = customer.FirstName + customer.LastName
	customer.Gender = form.Gender
	customer.Email = strings.TrimSpace(form.Email)
	customer.Birthday = strings.TrimSpace(form.Birthday)
	customer.IdType = form.IdType
	customer.IdNumber = strings.TrimSpace(form.IdNumber)
	customer.MemberGrade = form.MemberGrade
	customer.OwnerId = uid

	err = customerservice.CreateMyCustomers(customer); if err != nil{
		log.WithField("customer", customer).Error("failed to create customer", err)
		panic(err)
	}

	ctr.Success(c)
}

//get customer by owner_id
func (ctr *CustomerController) GetMyCustomers(c *gin.Context)  {
	uid, err := ctr.GetUid(c); if err != nil {
		ctr.Unauthorized(c)
		return
	}

	var page int
	pageStr := c.Query("page")

	if strings.TrimSpace(pageStr) == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageStr); if err != nil {
			log.Error("failed to get page", err)
			panic(err)
		}
	}

	customers, err := customerservice.GetMyCustomers(page, uid); if err != nil {
		log.Error("failed to get customers ", err)
		panic(err)
	}

	customerMap := map[int64]*model.Customer{}
	var feedIds []int64 = []int64{}

	for _, customer := range customers {
		customerMap[customer.Id] = customer
		customer.Name = lib.MaskName(customer.Name)
		if customer.LastFeedId > 0 {
			feedIds = append(feedIds, customer.LastFeedId)
		}
	}

	if len(feedIds) > 0 {
		feeds, err := feedservice.GetFeedsRSByIds(feedIds); if err != nil {
			panic(err)
		}

		for _, feed := range feeds {
			customer := customerMap[feed.CustomerId];if customer != nil {
				customer.Feeds = []*model.FeedRS{feed}
			}
		}
	}

	ctr.SuccessPair(c, "customers", customers)
}

//update customer by owner_id and customer's id
func (ctr * CustomerController) UpdateMyCustomers(c *gin.Context) {
	_, err := ctr.GetUid(c); if err != nil {
		ctr.Unauthorized(c)
		return
	}

	var form CustomerForm
	err = c.Bind(&form); if err != nil {
		log.Error("bad request ", err)
		ctr.BadRequest(c)
		return
	}

	idStr := c.Param("id")
	customerId, err := strconv.ParseInt(idStr, 10, 64); if err != nil {
		ctr.Fail(c, model.ID_INVALID)
		return
	}

	customers, err := customerservice.GetCustomerById(customerId); if err != nil {
		log.WithField("get customer by id:", customerId).Error("failed to get customer", err)
		panic(err)
	}

	if customers == nil {
		ctr.Fail(c, usermodel.USER_NOTEXIST)
		return
	}

	customer := &model.Customer{}
	customer.Mobile = strings.TrimSpace(form.Mobile)
	customer.FirstName = strings.TrimSpace(form.FirstName)
	customer.LastName = strings.TrimSpace(form.LastName)
	customer.Name = customer.FirstName + customer.LastName
	customer.Email = strings.TrimSpace(form.Email)
	customer.Gender = form.Gender
	customer.Birthday = strings.TrimSpace(form.Birthday)
	customer.IdType = form.IdType
	customer.IdNumber = strings.TrimSpace(form.IdNumber)
	customer.MemberGrade = form.MemberGrade

	err = customerservice.UpdateMyCustomers(customerId, customer); if err != nil{
		log.Error("failed to update customer")
		panic(err)
	}

	ctr.Success(c)
}

func (ctr * CustomerController) GetCustomer(c *gin.Context) {
	//uid, err := ctr.GetUid(c); if err != nil {
	//	ctr.Unauthorized(c)
	//	return
	//}

	log.Debug("customer id ", c.Param("id"))
	customerId, err := lib.Str2int64(c.Param("id")); if err != nil {
		ctr.BadRequest(c)
		return
	}

	customer, err := customerservice.GetCustomerById(customerId); if err != nil {
		panic(err)
	}

	customer.Name = lib.MaskName(customer.Name)

	//TODO Calculate customer's age and return
	//customer.age = age

	//TODO check privilege
	//if customer.OwnerId != uid {
	//	log.Warn("try to create order on behalf of customer illegally")
	//	ctr.Unauthorized(c)
	//	return
	//}


	feeds, err := feedservice.GetFeedsByCustomer(customerId, 1); if err != nil {
		panic(err)
	}

	customer.Feeds = feeds

	ctr.SuccessPair(c, "customer", customer)
}