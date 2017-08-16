package main
import "fmt"
func first() {
    fmt.Println("first")
}
func second() {
    fmt.Println("second")
}
func third() {
    fmt.Println("third")
}
func main() {
    defer first()
    second()
    third()
    panic("PANIC")
}

