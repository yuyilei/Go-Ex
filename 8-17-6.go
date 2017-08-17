package main
import "fmt"
type testInt func(int) bool //   声明函数模型
func isOdd(x int) bool{
    if x % 2 == 0 {
        return false
    }
    return true
}
func isEven(x int) bool {
    if x % 2 == 0 {
        return true
    }
    return false
}
func filter(slice []int , f testInt ) []int {
    var res []int
    for _, value := range(slice) {
        if  f(value) {
            res = append(res,value)
        }
    }
    return res
}
func main() {
    slice := []int {1,2,3,4,5,6,7}
    fmt.Println("slice = ",slice)
    odd := filter(slice,isOdd)
    even := filter(slice,isEven)
    fmt.Println("The odd of the elemnts are",odd)
    fmt.Println("THe even of the elemts are",even)
}
