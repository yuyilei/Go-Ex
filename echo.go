// echo函数 
package main 

import (
	"fmt"
	"os"
)

func main(){
	s, sep := "", ""
	// 和 var s, sep string  相同，没有显式地赋初值，就默认为0值（数字为0，字符型为空字符串） 
	for _, arg := range os.Args[1:] {
		s += sep + arg 
		sep = " "
	}
	fmt.Println(s)
}