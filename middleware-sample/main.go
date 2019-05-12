package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// Preparation
//  export GO111MODULE=on
//  go mod init project-name
// Local
//  dev_appserver.py app.yaml
// Deploy
//  gcloud app deploy --project [projectid] -v testapiv001
func main() {
	r := gin.New()
	sampleGroup := r.Group("/test")
	sampleGroup2 := sampleGroup.Group("/")

	// path: /test/path1
	// Middleware実行: されない
	sampleGroup.GET("/path1", func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)
		log.Infof(ctx, "path1 logic")
		c.String(200, "path1 !!!")
	})
	sampleGroup.Use(sampleMiddleware())
	{
		// path: /test/path2
		// Middleware実行: される
		sampleGroup.GET("/path2", func(c *gin.Context) {
			ctx := appengine.NewContext(c.Request)
			log.Infof(ctx, "path2 logic")
			c.String(200, "path2 !!!")
		})
	}
	// path: /test/path3
	// Middleware実行: される
	sampleGroup.GET("/path3", func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)
		log.Infof(ctx, "path3 logic")
		c.String(200, "path3 !!!")
	})
	// path: /test/path4
	// Middleware実行: されない
	sampleGroup2.GET("/path4", func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)
		log.Infof(ctx, "path4 logic")
		c.String(200, "path4 !!!")
	})

	http.Handle("/", r)
	appengine.Main()
	/*
		result

		[GIN-debug] GET    /test/path1               --> main.main.func1 (1 handlers)
		[GIN-debug] GET    /test/path2               --> main.main.func2 (2 handlers)
		[GIN-debug] GET    /test/path3               --> main.main.func3 (2 handlers)
		[GIN-debug] GET    /test/path4               --> main.main.func4 (1 handlers)
		INFO     2019-05-12 07:39:15,351 instance.py:294] Instance PID: 52004
		2019/05/12 07:40:35 INFO: path1 logic
		INFO     2019-05-12 07:40:35,539 module.py:861] default: "GET /test/path1 HTTP/1.1" 200 9
		2019/05/12 07:40:42 INFO: before logic
		2019/05/12 07:40:42 INFO: path2 logic
		2019/05/12 07:40:42 INFO: after logic
		INFO     2019-05-12 07:40:42,022 module.py:861] default: "GET /test/path2 HTTP/1.1" 200 9
		2019/05/12 07:40:49 INFO: before logic
		2019/05/12 07:40:49 INFO: path3 logic
		2019/05/12 07:40:49 INFO: after logic
		INFO     2019-05-12 07:40:49,377 module.py:861] default: "GET /test/path3 HTTP/1.1" 200 9
		2019/05/12 07:41:00 INFO: path4 logic
		INFO     2019-05-12 07:41:00,952 module.py:861] default: "GET /test/path4 HTTP/1.1" 200 9
	*/
}

func sampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)
		log.Infof(ctx, "before logic")
		c.Next()
		log.Infof(ctx, "after logic")
	}
}
