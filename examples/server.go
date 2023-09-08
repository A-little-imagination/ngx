package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	var s = gin.Default()
	s.GET("/get", get)
	s.GET("/get_body", getWithBody)
	s.POST("/post", post)
	s.POST("/post_body", postWithBody)
	s.PUT("/put", put)
	s.PUT("/put_body", putWithBody)
	s.DELETE("/delete", delete)
	s.DELETE("/delete_body", deleteWithBody)

	s.POST("/redirect", redirect)
	s.POST("/upload", upload)
	s.Run(":9090")
}

func get(c *gin.Context) {
	c.Request.ParseForm()
	for key, values := range c.Request.Form {
		fmt.Println("接收到请求参数:", key, values)
	}
}

func getWithBody(c *gin.Context) {
	c.Request.ParseForm()
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("接收到 Body:", string(body))
	for key, values := range c.Request.Form {
		fmt.Println("接收到请求参数:", key, values)
	}
}

func post(c *gin.Context) {
	c.Request.ParseForm()
	for key, values := range c.Request.Form {
		fmt.Println("接收到请求参数:", key, values)
	}
}

func postWithBody(c *gin.Context) {
	c.Request.ParseForm()
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("接收到 Body:", string(body))
	for key, values := range c.Request.Form {
		fmt.Println("接收到请求参数:", key, values)
	}
}

func put(c *gin.Context) {
	c.Request.ParseForm()
	for key, values := range c.Request.Form {
		fmt.Println("接收到请求参数:", key, values)
	}
}

func putWithBody(c *gin.Context) {
	c.Request.ParseForm()
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("接收到 Body:", string(body))
	for key, values := range c.Request.Form {
		fmt.Println("接收到请求参数:", key, values)
	}
}

func delete(c *gin.Context) {
	c.Request.ParseForm()
	for key, values := range c.Request.Form {
		fmt.Println("接收到请求参数:", key, values)
	}
}

func deleteWithBody(c *gin.Context) {
	c.Request.ParseForm()
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("接收到 Body:", string(body))
	for key, values := range c.Request.Form {
		fmt.Println("接收到请求参数:", key, values)
	}
}

func redirect(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/upload")
}

func upload(c *gin.Context) {
	c.Request.ParseMultipartForm(64 << 20)

	for key, values := range c.Request.Form {
		fmt.Println("接收到请求参数:", key, values)
	}

	for key := range c.Request.MultipartForm.File {
		if err := writeFile(c.Request, key); err != nil {
			fmt.Println("操作文件发生错误:", err)
			return
		}
	}

	c.String(http.StatusOK, "文件上传完成")
}

func writeFile(req *http.Request, name string) error {
	rFile, header, err := req.FormFile(name)
	defer rFile.Close()
	if err != nil {
		return err
	}

	nFile, err := os.Create(fmt.Sprintf("%d-%s", time.Now().UnixNano(), header.Filename))
	if err != nil {
		return err
	}
	defer nFile.Close()
	if _, err = io.Copy(nFile, rFile); err != nil {
		return err
	}

	return nil
}
