package main

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	// 指定默认语言和配置文件
	bundle := i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFile("active.en.toml")
	bundle.LoadMessageFile("active.zh.toml")

	// 创建一个默认的路由引擎
	r := gin.Default()
	// 配置路由
	r.GET("/", func(c *gin.Context) {
		// 使用Query获取请求参数
		lang := c.Query("lang")
		worldName := c.Query("worldName")
		name := c.Query("name")
		acceptHeader := c.GetHeader("Accept-Language")
		localizer := i18n.NewLocalizer(bundle, lang, acceptHeader)
		// go-i18n会根据此处代码自动生成配置文件
		helloWorld := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "HelloWorld",
				Other: "Hello {{.WorldName}}",
			},
			TemplateData: map[string]string{
				"WorldName": worldName,
			},
		})
		helloPerson := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "HelloPerson",
				Other: "Hello {{.Name}}",
			},
			TemplateData: map[string]string{
				"Name": name,
			},
		})
		c.JSON(200, gin.H{
			"message_1": helloWorld,
			"message_2": helloPerson,
		})
	})
	r.Run() // 启动 HTTP 服务
}

