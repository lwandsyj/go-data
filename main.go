package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func get() {
	// get 请求，返回一个response 对象
	res, error := http.Get("http://www.baidu.com")
	if error != nil {
		fmt.Println(res)
		return
	}
	// 输出response 响应状态，200 ok
	fmt.Println(res.Status)
	// 输出response 响应状态 200
	fmt.Println(res.StatusCode)
	// ioutil.ReadAll 读取io.Reader 类型，返回[]byte 切片
	n, error := ioutil.ReadAll(res.Body)
	if error != nil {
		fmt.Println("body read", error)
		return
	}
	// 转成字符串
	fmt.Println("n=", string(n))
}
func postFileAndReturnResponse(filename string) string {
	// 实例化一个Buffer 结构体，buffer 实现了io.Reader 和 io.Writer
	fileDataBuffer := bytes.Buffer{}
	//multipart实现了MIME的multipart解析,Writer类型用于生成multipart信息。
	//cannot use fileDataBuffer (variable of type bytes.Buffer) as io.Writer
	// value in argument to multipart.NewWriter: missing method Write (Write has pointer receiver)
	multipart := multipart.NewWriter(&fileDataBuffer)
	// 返回file 信息
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// CreateFormFile是CreatePart方法的包装， 使用给出的属性名和文件名创建一个新的form-data头。
	formFile, err := multipart.CreateFormFile("myFile", file.Name())
	if err != nil {
		log.Fatal(err)
	}
	// 将reader 中的数据拷贝到writer 中
	_, err = io.Copy(formFile, file)
	multipart.Close()
	// 创建post 文件请求
	req, err := http.NewRequest("POST",
		"http://localhost:8080", &fileDataBuffer)
	if err != nil {
		log.Fatal(err)
	}
	// 设置http 请求header ,文件上传为multipart/form-data
	req.Header.Set("Content-Type",
		//方法返回w对应的HTTP multipart请求的Content-Type的值，多以multipart/form-data起
		multipart.FormDataContentType())
	// response body为io.Reader 类型
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	// ioutil.ReadAll body为io.Reader 类型
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)

}
func main() {
	// 文件内容路径
	path := "./test/hello1.txt"
	// WriteFile 如果文件存在，则会清空文件的内容，然后写入新的内容
	// 如果文件不存在，则按照file.Mode(0666) 创建文件，并写入内容
	err := ioutil.WriteFile(path, []byte("中难过"), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

}
