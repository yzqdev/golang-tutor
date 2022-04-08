package mail

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendEmail(c *gin.Context) {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "yzq <3131681417@qq.com>"
	// 设置接收方的邮箱
	e.To = []string{"1928315021@qq.com"}
	//设置主题
	e.Subject = "这是主题"
	//设置文件发送的内容
	//设置文件发送的内容
	e.HTML = []byte(`
    <h1><a href="http://www.baidu.com/">go语言中文网站</a></h1>    
    `)
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("1", "3131681417@qq.com", "dvajoafeqazkdhdi", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
}
