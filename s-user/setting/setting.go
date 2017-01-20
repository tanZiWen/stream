package setting

import (
	"time"
	"code.isstream.com/stream/setting"
)

var Config *MUserConfig
var sectionName string = "s-user"
var smsSectionName string = "s-user.sms"

type SMSConfig struct {
	SecondPeriod         int `ini:"SECOND_PERIOD"`
	MaxPerDay            int `ini:"MAX_PER_DAY"`

	Provider             string `ini:"PROVIDER"`
	Api                  string `ini:"API"`
	ApiSecret            string `ini:"API_SECRET"`

	DynamicPwdTemplateId int64 `ini:"DYNAMIC_PASSWORD_TEMPLATE_ID"`
	VerifyCodeTemplateId int64 `ini:"VERIFY_CODE_TEMPLATE_ID"`
}

type MUserConfig struct {
	MaxPwdRetry int `ini:"MAX_PASSWORD_RETRY"`

	EnableEmailSignup bool `ini:"ENABLE_EMAIL_SIGNUP"`
	EnableMobileSignup bool `ini:"ENABLE_MOBILE_SIGNUP"`

	SMS         *SMSConfig
}

func init() {
	smsConfig := &SMSConfig{
		SecondPeriod: 60,
		MaxPerDay: 10,
	}

	Config = &MUserConfig{MaxPwdRetry: 5, SMS: smsConfig}
	setting.AddMapping(setting.SectionMap{SectionName: sectionName, MapTo: Config})
	setting.AddMapping(setting.SectionMap{SectionName: smsSectionName, MapTo: Config.SMS})
}
