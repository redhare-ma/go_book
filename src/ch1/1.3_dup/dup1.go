//dup1输出标准输入中出现次数大于1的行，前面是次数
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin) //bufio可以读取标准输入，以行或者单词为单位断开
	for input.Scan() {
		//检测到 end 就退出，如果没有退出循环的判断，使用Ctrl+d也可以退出
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
		//上面这行代码相当于下面两行代码的意思
		//line := input.Text()
		//counts[line] = counts[line] + 1
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
