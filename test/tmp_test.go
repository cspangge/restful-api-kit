package test

import (
	"fmt"
	"restful-api-kit/helpers"
	"testing"
)

func TestEmail(t *testing.T) {

	user := "yang**@yun*.com"
	password := "***"
	host := "smtp.exmail.qq.com:25"
	to := "397685131@qq.com"

	subject := "使用Golang发送邮件"

	body := `
		<html>
		<body>
		<h3>
		"Test send to email"
		</h3>
		</body>
		</html>
		`
	fmt.Println("send email")
	err := helpers.SendToMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}
