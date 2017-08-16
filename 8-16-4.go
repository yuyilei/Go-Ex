package main
import "fmt"
func swap( x *int , y *int ) {
    fmt.Println("x=",*x,"y=",*y)
    tmp := *x
    *x = *y
    *y = tmp
    fmt.Println("x=",*x,"y=",*y)
}
func main() {
    x := 1
    y := 2
    swap(&x,&y)
}
