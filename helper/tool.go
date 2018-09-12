package helper

import (
	"os/exec"
	"os"
	"strings"
	"fmt"
)

//获取当前路径
func GetCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

//打印数据
func Dd(d interface{}){
	fmt.Println()
	fmt.Printf("%v", d)
	fmt.Println()
}