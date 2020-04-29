package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"strconv"
	controller "./controller"
)

func main(){
	route := gin.Default()
	
	route.LoadHTMLGlob("./views/*.html")
	
	route.GET("/", func(c *gin.Context){
		c.HTML(200, "top.html", gin.H{})
	})

	
	route.GET("/book/:id", func(c *gin.Context){
		id, _ := strconv.Atoi(c.Param("id"))
		book := controller.BookFind(id)

		c.HTML(200, "show.html", gin.H{"book": book})
	})

	route.GET("/books", func(c *gin.Context){
		books := controller.BookAll()
		c.HTML(200, "index.html", gin.H{"books": books})
	})



	route.GET("/book/:id/edit", func(c *gin.Context){
		id, _ := strconv.Atoi(c.Param("id"))
		book := controller.BookFind(id)
		c.HTML(200, "edit.html", gin.H{"book": book})
	})


	route.POST("/books", func(c *gin.Context){
		c.Request.ParseForm()
		title := c.Request.Form["title"][0]
		body := c.Request.Form["body"][0]
		controller.CreateBook(title, body)
		//net/httpを使ったリダイレクト		
		c.Redirect(http.StatusMovedPermanently, "/books")
		// c.Request.URL.Path = "/index"
		// route.HandleContext(c)
	})

	route.POST("/update/:id", func(c *gin.Context){
		id, _ := strconv.Atoi(c.Param("id"))
		c.Request.ParseForm()
		title := c.Request.Form["title"][0]
		body := c.Request.Form["body"][0]
		controller.UpdateBook(id, title, body)
		path := "/book/" + c.Param("id")
		fmt.Println(path)
		c.Redirect(http.StatusMovedPermanently, path)
	})

	route.POST("/delete/:id", func(c *gin.Context){
		id, _ := strconv.Atoi(c.Param("id"))
		controller.DeleteBook(id)
		// message :=  "データの削除に成功しました!"
		c.Redirect(http.StatusMovedPermanently, "/books")
	})
	
	route.Run(":8080")
}