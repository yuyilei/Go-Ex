package main
import "fmt"
import "os"
func main() {
    var total float64 = 0
    x := [5]float64{ 98, 93, 77, 82, 83 }
    for _ , value := range x {
        total += value
    }
    total /= float64(len(x))
    fmt.Println(total)
    slice1 := []int{1,2,3}
    slice2 := make([]int,2,10)
    copy(slice2,slice1)
    i , err := fmt.Println(slice1,slice2)
    fmt.Println(i,err)
    for _ , value := range slice2 {
        fmt.Println(value)
    }
    k := make(map[string]int)
    k["s"] = 10
    k["t"] = 20
    name , ok := k["p"]
    fmt.Println(k["p"])
    fmt.Println(name,ok)
    fmt.Println(k)
    fmt.Println(len(k))
    os.Exit(0)
}
