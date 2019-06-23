package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
)

type Person struct {
	ID int `uri:"id" binding:"required,min=1,max=100"` // 値1-100まで
	//ID   string `uri:"id" binding:"required,min=1,max=3"` // 文字数1-3文字まで
	Name string `uri:"name" binding:"required"`
}

// Preparation
//  export GO111MODULE=on
//  go mod init project-name
// Local
//  dev_appserver.py app.yaml
// Deploy
//  gcloud app deploy --project [projectid] -v testapiv001
func main() {
	r := gin.New()

	r.GET("/path/:name/:id", func(c *gin.Context) {
		var person Person
		//if err := c.ShouldBindUri(&person); err != nil {
		if err := c.BindUri(&person); err != nil { // パスが数字だけであればintにバインドしてくれる
			c.JSON(400, gin.H{"ErrorMsg": err})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "id": person.ID})
	})

	http.Handle("/", r)
	appengine.Main()
}
