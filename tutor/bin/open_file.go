package bin

import (
	"fmt"
	"os"
	"os/exec"
)

func OpenFile() {
	path, _ := os.Executable()
	dir, _ := os.Getwd()
	fmt.Println("当前路径：", dir)
	fmt.Println(path)
	file, err := os.Open(dir + "\\a.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	fmt.Println(file.Chdir())
	// 关闭文件
	defer file.Close()
}

func OpenCmd() {
	cmd := exec.Command("notepad")
	err := cmd.Run()
	if err != nil {
		fmt.Println("出错了")
	}
	fmt.Print("aaaaaaaaa")
}
