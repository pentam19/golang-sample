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
	// Middleware1: 実行されない
	// Middleware2: 実行されない
	// Middleware3: 実行されない
	sampleGroup.GET("/path1", func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)
		log.Infof(ctx, "path1 logic")
		c.String(200, "path1 !!!")
	})
	sampleGroup.Use(sampleMiddleware1())
	{
		// path: /test/path2
		// Middleware1: 実行される
		// Middleware2: 実行されない
		// Middleware3: 実行されない
		sampleGroup.GET("/path2", func(c *gin.Context) {
			ctx := appengine.NewContext(c.Request)
			log.Infof(ctx, "path2 logic")
			c.String(200, "path2 !!!")
		})
	}
	// path: /test/path3
	// Middleware1: 実行される
	// Middleware2: 実行されない
	// Middleware3: 実行されない
	sampleGroup.GET("/path3", func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)
		log.Infof(ctx, "path3 logic")
		c.String(200, "path3 !!!")
	})
	// path: /test/path4
	// Middleware1: 実行されない (Use(sampleMiddleware1())の前にGroup作ったため)
	// Middleware2: 実行されない
	// Middleware3: 実行されない
	sampleGroup2.GET("/path4", func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)
		log.Infof(ctx, "path4 logic")
		c.String(200, "path4 !!!")
	})

	sampleGroup3 := sampleGroup2.Group("/")
	sampleGroup3.Use(sampleMiddleware2())
	// path: /test/path5
	// Middleware1: 実行されない (Use(sampleMiddleware1())の前にGroup作ったため)
	// Middleware2: 実行される
	// Middleware3: 実行されない
	sampleGroup3.GET("/path5", func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)
		log.Infof(ctx, "path5 logic")
		c.String(200, "path5 !!!")
	})

	sampleGroup4 := sampleGroup3.Group("/")
	sampleGroup4.Use(sampleMiddleware1())
	sampleGroup4.Use(sampleMiddleware3())
	// path: /test/path6
	// Middleware1: 実行される
	// Middleware2: 実行される (Use(sampleMiddleware2())が引き継がれている)
	// Middleware3: 実行される
	sampleGroup4.GET("/path6", func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)
		log.Infof(ctx, "path6 logic")
		c.String(200, "path6 !!!")
	})

	http.Handle("/", r)
	appengine.Main()
	/*
		result

		[GIN-debug] GET    /test/path1               --> main.main.func1 (1 handlers)
		[GIN-debug] GET    /test/path2               --> main.main.func2 (2 handlers)
		[GIN-debug] GET    /test/path3               --> main.main.func3 (2 handlers)
		[GIN-debug] GET    /test/path4               --> main.main.func4 (1 handlers)
		[GIN-debug] GET    /test/path5               --> main.main.func5 (2 handlers)
		[GIN-debug] GET    /test/path6               --> main.main.func6 (4 handlers)
		INFO     2019-05-12 08:26:19,980 instance.py:294] Instance PID: 52732
		2019/05/12 08:26:27 INFO: path1 logic
		INFO     2019-05-12 08:26:27,499 module.py:861] default: "GET /test/path1 HTTP/1.1" 200 9
		2019/05/12 08:26:29 INFO: [Middleware 1] before logic
		2019/05/12 08:26:29 INFO: path2 logic
		2019/05/12 08:26:29 INFO: [Middleware 1] after logic
		INFO     2019-05-12 08:26:29,688 module.py:861] default: "GET /test/path2 HTTP/1.1" 200 9
		2019/05/12 08:26:31 INFO: [Middleware 1] before logic
		2019/05/12 08:26:31 INFO: path3 logic
		2019/05/12 08:26:31 INFO: [Middleware 1] after logic
		INFO     2019-05-12 08:26:31,817 module.py:861] default: "GET /test/path3 HTTP/1.1" 200 9
		2019/05/12 08:26:34 INFO: path4 logic
		INFO     2019-05-12 08:26:34,061 module.py:861] default: "GET /test/path4 HTTP/1.1" 200 9
		2019/05/12 08:26:36 INFO: [Middleware 2] before logic
		2019/05/12 08:26:36 INFO: path5 logic
		2019/05/12 08:26:36 INFO: [Middleware 2] after logic
		INFO     2019-05-12 08:26:36,492 module.py:861] default: "GET /test/path5 HTTP/1.1" 200 9
		2019/05/12 08:26:38 INFO: [Middleware 2] before logic
		2019/05/12 08:26:38 INFO: [Middleware 1] before logic
		2019/05/12 08:26:38 INFO: [Middleware 3] before logic
		2019/05/12 08:26:38 INFO: path6 logic
		2019/05/12 08:26:38 INFO: [Middleware 3] after logic
		2019/05/12 08:26:38 INFO: [Middleware 1] after logic
		2019/05/12 08:26:38 INFO: [Middleware 2] after logic
		INFO     2019-05-12 08:26:38,790 module.py:861] default: "GET /test/path6 HTTP/1.1" 200 9
	*/
}

func sampleMiddleware1() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)
		log.Infof(ctx, "[Middleware 1] before logic")
		c.Next()
		log.Infof(ctx, "[Middleware 1] after logic")
	}
}

func sampleMiddleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)
		log.Infof(ctx, "[Middleware 2] before logic")
		c.Next()
		log.Infof(ctx, "[Middleware 2] after logic")
	}
}

func sampleMiddleware3() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)
		log.Infof(ctx, "[Middleware 3] before logic")
		c.Next()
		log.Infof(ctx, "[Middleware 3] after logic")
	}
}
