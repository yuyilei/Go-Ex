package main
import "fmt"
func main() {
    for i := 0 ; i < 10 ; i++ {
        defer fmt.Println(i)
        // 如果多次调用defer，defer 是先进后出的，会输出 9 ，8，7 。。。
    }
}
