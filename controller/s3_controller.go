package controller

import (
	"fmt"
	"gl_s3/initializers"
	"gl_s3/internal/pkg/cloud/aws"
	data_type "gl_s3/type"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	client := c.MustGet("client").(aws.S3)
	if err := client.Create(c, "sample-bucket"); err != nil {
		log.Fatalln(err)
	}
	log.Println("create: ok")
}

func UploadObject(c *gin.Context) {
	client := c.MustGet("client").(aws.S3)

	var body data_type.UploadFile
	if err := c.ShouldBindJSON(&body); err != nil {
		fmt.Println("*******feefefefef***********")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := os.Open(body.FilePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "File not found.",
		})
		return
	}

	filename := path.Base(body.FilePath)
	defer file.Close()

	fmt.Println("******************")
	fmt.Println(file)

	_, err = client.UploadObject(c, body.BucketName, filename, file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "Bucket name: " + body.BucketName + " not found.",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"filepath": initializers.RemoteUrl() + "/api/download/" + body.BucketName + "/" + filename,
	})
}

func DownloadObject(c *gin.Context) {
	client := c.MustGet("client").(aws.S3)
	fileName := c.Param("file_name")
	bucketName := c.Param("bucket_name")
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	if err := client.DownloadObject(c, bucketName, fileName, file); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	//Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.File("./" + fileName)

	return
}

func deleteObject(c *gin.Context) {
	client := c.MustGet("client").(aws.S3)
	if err := client.DeleteObject(c, "aws-test", "id.txt"); err != nil {
		log.Fatalln(err)
	}
	log.Println("delete object: ok")
}

func listObjects(c *gin.Context) {
	client := c.MustGet("client").(aws.S3)
	objects, err := client.ListObjects(c, "aws-test")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("list objects:")
	for _, object := range objects {
		fmt.Printf("%+v\n", object)
	}
}
