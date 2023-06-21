package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Data struct {
	Message  string `json:"message"`
	Hostname string `json:"hostname"`
}

func main() {
	http.HandleFunc("/get", getData)
	fmt.Print("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getData(w http.ResponseWriter, _ *http.Request) {
	// 创建要返回的数据对象
	go func() {
		fmt.Println("Hello, World!")
	}()
	hostname, _ := os.Hostname()
	data := Data{
		Message:  "Hello, World!",
		Hostname: hostname,
	}

	// 将数据对象转换为 JSON 格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 设置响应头部为 JSON 类型
	w.Header().Set("Content-Type", "application/json")

	// 将 JSON 数据写入响应体
	_, err = w.Write(jsonData)
	if err != nil {
		return
	}
}
