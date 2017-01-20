package customerservice

import (
	"code.isstream.com/stream/partner-service/model"
	"code.isstream.com/stream/db"
	log "github.com/Sirupsen/logrus"
)

const (
	ALL_CUSTOMERS = "select id, name, gender, mobile, email, birthday, last_feed_id from customer where status = 1 " +
		"and owner_id = ? ORDER BY crt DESC limit 20 offset ?"
	QUERY_CUSTOMER_BY_MOBILE = "select * from customer where status = 1 and mobile = ?"
	QUERY_CUSTOMER_BY_ID = "select id, name, gender, mobile, email, birthday from customer where status = 1 and id = ?"
)

func CreateMyCustomers(customer *model.Customer) error {
	_, err := db.Engine.Omit("status").Insert(customer); if err != nil {
		log.Error("create customer error:", err)
		return err
	}
	return nil
}

func GetMyCustomers(page int, id int64) (customers []*model.Customer, err error) {
	offset := (page - 1) * 20
	err = db.Engine.Sql(ALL_CUSTOMERS, id, offset).Find(&customers); if err != nil {
		log.Error("get customers info error:", err)
		return nil, err
	}

	return customers, nil
}

func GetCustomerByMobile(mobile string) (*model.Customer, error) {
	customer := &model.Customer{}
	isExist, err := db.Engine.Sql(QUERY_CUSTOMER_BY_MOBILE, mobile).Get(customer); if err != nil {
		log.Error("query customer by mobile error:", err)
		return nil, err
	}

	if !isExist {
		return nil, nil
	}
	return customer, nil
}

func GetCustomerById(id int64) (*model.Customer, error) {
	customer := &model.Customer{}
	isExist, err := db.Engine.Sql(QUERY_CUSTOMER_BY_ID, id).Get(customer); if err != nil {
		log.Error("query customer by id error:", err)
		return nil, err
	}

	if !isExist {
		return nil, nil
	}
	return customer, nil
}

func UpdateMyCustomers(id int64, customer *model.Customer) error {
	_, err := db.Engine.Id(id).Cols("mobile", "name", "first_name", "last_name", "email", "gender", "id_type", "birthday",
		"member_grade", "id_type", "id_number").Update(customer)
	if err != nil {
		log.Error("failed to update customer", err)
		return err
	}

	return nil
}