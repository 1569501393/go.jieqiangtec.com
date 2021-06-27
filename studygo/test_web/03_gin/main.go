package main

import "github.com/gin-gonic/gin"

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()

	//GET: 请求方式; /hello: 请求路径
	//当客户端以GET请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		//c.JSON:返回JSON格式的数据
		c.JSON(200, gin.H{
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
		c.JSON(200, gin.H{
			"msg": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "DELETE",
		})
	})

	//启动HTTP服务,默认在0.0.0.0:8080启动服务
	//将上面的代码保存并编译执行，然后使用浏览器打开127.0.0.1:8080/hello就能看到一串JSON字符串。
	err := r.Run()
	if err != nil {
		return
	}

}
