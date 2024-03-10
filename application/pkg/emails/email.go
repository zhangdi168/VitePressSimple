// Package emails
// @Author: zhangdi
// @File: email
// @Version: 1.0.0
// @Date: 2023/6/25 15:08
package emails

import "gopkg.in/gomail.v2"

//const Email = "1778871523@qq.com"
//const Pwd = "rwehrxsnnvrneaae"

const Email = "ideatools@qq.com"
const Pwd = "rtcxcomjggocbefb"
const Port = 25
const Host = "smtp.qq.com"

//func SendEmail(toEmail string, title string, content string) error {
//
//	// 配置SMTP服务器地址、端口和身份验证信息
//	smtpHost := Host
//	smtpPort := Port
//	username := Email
//	password := Pwd
//
//	// 构造邮件内容
//	from := Email
//	to := []string{toEmail}
//	subject := title
//	body := content
//
//	// 构建消息体
//	message := []byte("To: " + to[0] + "\r\n" +
//		"Subject: " + subject + "\r\n" +
//		"\r\n" +
//		body + "\r\n")
//
//	// 连接到SMTP服务器，进行身份验证
//	auth := smtp.PlainAuth("", username, password, smtpHost)
//	err := smtp.SendMail(smtpHost+":"+string(smtpPort), auth, from, to, message)
//	if err != nil {
//
//		log.Fatal(err)
//	}
//
//	return nil
//}

func SendEmail1(toEmail string, title string, content string) error {
	// 邮件配置
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", Email)
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", title)
	mailer.SetBody("text/html", content)

	// 发送邮件
	dialer := gomail.NewDialer(Host, Port, Email, Pwd)
	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	println("Email sent successfully!")

	return nil
}
