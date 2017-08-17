package main
import "fmt"

func main() {
    var a = [...]int{1,2,3,4,5,6,7}
    b := a[2:4]
    b[1] = 0
    // slice 是引用类型，改变其元素的值时，所有的引用的值都会改变
    fmt.Println(a,b)
    m := make(map[string]string)
    m["hello"] = "you"
    m1 := m
    m1["hello"] = "me"
    fmt.Println(m["hello"])
    // map是引用类型，改变一个引用，所有引用的值都会改变
}
