package main
import ( "bytes" ; "strings" )
func main() {
    var buf bytes.Buffer
    buf.Write([]byte("test"))
    strings.NewReader("test")
}

