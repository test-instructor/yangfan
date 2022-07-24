/*
@Author : 锐杰
@Time : 2022/7/21 20:15
*/
package runTestCase

import (
	"fmt"
	"os"
)

func CreateDebugTalk(path string) string {
	filePath := "examples/" + path
	_, err := os.Stat(filePath)

	//isnotexist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		err := os.Mkdir(filePath, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录异常 -> %v\n", err)
		} else {
			fmt.Println("创建成功!")
		}
	}
	return filePath
}
