package helpers

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"restful-api-kit/config"
	"strconv"
)

var MailConf *MailConfig

type MailConfig struct {
	User     string
	Password string
	Host     string
	Port     string
}

func InitMail() {
	MailConf = &MailConfig{
		User:     config.GlobalConfigSetting.Email.User,
		Password: config.GlobalConfigSetting.Email.Password,
		Host:     config.GlobalConfigSetting.Email.Host,
		Port:     config.GlobalConfigSetting.Email.Port,
	}
}

func SendEMail(subject string, body string, mailTo ...string) error {
	port, err := strconv.Atoi(MailConf.Port)
	if err != nil {
		return err
	}
	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("Seacowtech <%s>", MailConf.User)) //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                                       //发送给多个用户
	m.SetHeader("Subject", fmt.Sprintf("<h1>%s</h1>", subject))        //设置邮件主题
	m.SetBody("text/html", body)                                       //设置邮件正文
	return gomail.NewDialer(MailConf.Host, port, MailConf.User, MailConf.Password).DialAndSend(m)
}
