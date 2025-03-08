package email

import (
	"KazePush/internal/config"
	"fmt"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

type EmailRequest struct {
	To    []string `json:"to"`
	Msg   string   `json:"msg"`
	Title string   `json:"title"`
	Name  string   `json:"name"`
}

func SendEmail(c *gin.Context) {
	var req EmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": 400, "message": "请求参数错误"})
		return
	}

	if len(req.To) == 0 {
		c.JSON(400, gin.H{"code": 400, "message": "必须指定至少一个收件人"})
		return
	}

	// 设置默认值
	senderName := config.GlobalCfg.Smtp.SenderName
	if req.Name != "" {
		senderName = req.Name
	}
	if req.Title == "" {
		req.Title = "默认标题"
	}
	if req.Msg == "" {
		req.Msg = "默认正文"
	}

	if err := sendEmail(senderName, req.To, req.Title, req.Msg); err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": fmt.Sprintf("邮件发送失败，详情: %v", err),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "邮件已成功发送",
	})
}

func sendEmail(senderName string, to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", config.GlobalCfg.Smtp.SenderEmail, senderName)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(
		config.GlobalCfg.Smtp.SmtpServer,
		config.GlobalCfg.Smtp.SmtpPort,
		config.GlobalCfg.Smtp.SenderEmail,
		config.GlobalCfg.Smtp.SenderPassword,
	)

	return d.DialAndSend(m)
}
