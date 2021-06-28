package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()

	//GET: 请求方式; /hello: 请求路径
	//当客户端以GET请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		//c.JSON:返回JSON格式的数据
		c.JSON(http.StatusOK, gin.H{
			"msg": "Hello world!",
		})
	})

	/*RESTful API
	REST与技术无关，代表的是一种软件架构风格，REST是Representational State Transfer的简称，中文翻译为“表征状态转移”或“表现层状态转化”。

	推荐阅读阮一峰 理解RESTful架构

	简单来说，REST的含义就是客户端与Web服务器之间进行交互的时候，使用HTTP协议中的4个请求方法代表不同的动作。

	GET用来获取资源
	POST用来新建资源
	PUT用来更新资源
	DELETE用来删除资源。
	只要API程序遵循了REST风格，那就可以称其为RESTful API。目前在前后端分离的架构中，前后端基本都是通过RESTful API来进行交互。

	例如，我们现在要编写一个管理书籍的系统，我们可以查询对一本书进行查询、创建、更新和删除等操作，我们在编写程序的时候就要设计客户端浏览器与我们Web服务端交互的方式和路径。按照经验我们通常会设计成如下模式：

	请求方法	URL	含义
	GET	/book	查询书籍信息
	POST	/create_book	创建书籍记录
	POST	/update_book	更新书籍信息
	POST	/delete_book	删除书籍信息
	同样的需求我们按照RESTful API设计如下：

	请求方法	URL	含义
	GET	/book	查询书籍信息
	POST	/book	创建书籍记录
	PUT	/book	更新书籍信息
	DELETE	/book	删除书籍信息
	Gin框架支持开发RESTful API的开发。*/

	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "DELETE",
		})
	})

	//获取参数

	//获取querystring参数
	//querystring指的是URL中?后面携带的参数，例如：/user/search?username=小王子&address=沙河。 获取请求的querystring参数的方法如下：

	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "jieqiang")
		//username := c.Query("username")

		address := c.Query("address")

		c.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"address":  address,
		})
	})

	//获取form参数
	//前端请求的数据通过form表单提交时，例如向/user/search发送一个POST请求，获取请求数据的方式如下：

	r.POST("/user/search", func(c *gin.Context) {
		username := c.PostForm("username")
		address := c.PostForm("address")

		c.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"address":  address,
		})
	})

	//获取json参数
	//当前端请求的数据通过JSON提交时，例如向/json发送一个POST请求，则获取请求参数的方式如下：
	r.POST("/json", func(c *gin.Context) {
		// 从c.Request.Body读取请求数据
		b, _ := c.GetRawData()

		//定义map或结构体
		var m map[string]interface{}

		//反序列化
		json.Unmarshal(b, &m)
		c.JSON(http.StatusOK, m)
	})

	//获取path参数
	//请求的参数通过URL路径传递，例如：/user/search/小王子/沙河。 获取请求URL路径中的参数的方式如下。
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")

		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"address":  address,
		})
	})

	// 绑定JSON的示例 ({"user": "q1mi", "password": "123456"})
	r.POST("/loginJSON", func(c *gin.Context) {
		var login Login

		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定form表单示例 (user=q1mi&password=123456)
	r.POST("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定QueryString示例 (/loginQuery?user=q1mi&password=123456)
	r.GET("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	//启动HTTP服务,默认在0.0.0.0:8080启动服务
	//将上面的代码保存并编译执行，然后使用浏览器打开127.0.0.1:8080/hello就能看到一串JSON字符串。
	err := r.Run()
	if err != nil {
		return
	}

}

//参数绑定
//为了能够更方便的获取请求相关参数，提高开发效率，我们可以基于请求的Content-Type识别请求数据类型并利用反射机制自动提取请求中QueryString、form表单、JSON、XML等参数到结构体中。 下面的示例代码演示了.ShouldBind()强大的功能，它能够基于请求自动提取JSON、form表单和QueryString类型的数据，并把值绑定到指定的结构体对象。
// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
