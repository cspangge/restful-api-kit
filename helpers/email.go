package helpers

import (
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

func SendMail(mailTo []string, subject string, body string) error {
	port, _ := strconv.Atoi(MailConf.Port) //转换端口类型为int
	m := gomail.NewMessage()
	m.SetHeader("From", "XD Game"+"<"+MailConf.User+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                         //发送给多个用户
	m.SetHeader("Subject", subject)                      //设置邮件主题
	m.SetBody("text/html", body)                         //设置邮件正文

	d := gomail.NewDialer(MailConf.Host, port, MailConf.User, MailConf.Password)
	err := d.DialAndSend(m)
	return err
}
