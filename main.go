package main

import (
	"gl_s3/controller"
	initializers "gl_s3/initializers"
	"gl_s3/internal/pkg/cloud/aws"

	"github.com/gin-gonic/gin"
)

func main() {
	initializers.LoadEnv()
	sess := initializers.InitAWS()
	client := aws.NewS3(sess)
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("client", client)
		c.Next()
	})

	publicRoutes := router.Group("/api")
	publicRoutes.POST("/upload", controller.UploadObject)
	publicRoutes.GET("/download/:bucket_name/:file_name", controller.DownloadObject)

	_ = router.Run(":4003")
}
