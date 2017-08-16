package main
import "fmt"
func sum( x []int ) int {
    sum := 0
    for _ , value := range x {
        sum += value
    }
    return sum
}
func ifodd ( x int ) {
    tmp := []int{1,0}
    fmt.Println("(",tmp[x%2], "," , x%2 == 0 , ")")
}
func find_greatest( x []int ) int {
    max := x[0]
    for _ , value := range x {
        if max < value  {
            max = value
        }
    }
    return max
}
func fib( x int ) int {
    if x == 0 || x == 1 {
        return  1
    }
    return fib(x-1) + fib(x-2)
}
func main () {
    x := []int{1,2,3,4,5,6,7}
    fmt.Println(sum(x))
    ifodd(sum(x))
    fmt.Println("The greatest in the slice is", find_greatest(x))
    fmt.Println("The fib of the slice is" ,fib((sum(x))))
}
