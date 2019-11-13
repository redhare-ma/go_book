//dup2打印输入中多次出现的行的个数和文本
//他从stdin或指定的文件列表读取
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			//fmt.Println("f: ", f)   //内存地址
			//fmt.Println("arg: ", arg) //文件路径
			countLines(f, counts)
			for _, n := range counts {
				if n > 1 {
					fmt.Println("file path: ", arg)
					break
				}
			}
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	//注意，忽略input.Err()中可能的错误
}
