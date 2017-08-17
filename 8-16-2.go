package main
import "fmt"
import "time"
// recover from a run-time panic

func main() {
    defer func() {
        str := recover()
        fmt.Println(str)
    }
    panic("PANIC")
}
