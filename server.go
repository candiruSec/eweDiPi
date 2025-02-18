package main
import (
    "fmt" 
    "net"  
)

// Replace me
const IP = "127.0.0.1"
const PORT = 9999

// jklop is "key" to start command
func sendResponse(conn *net.UDPConn, addr *net.UDPAddr, response string) {
    _,err := conn.WriteToUDP([]byte("jklop" + response), addr)
    if err != nil {
        fmt.Printf("Couldn't send response %v", err)
    }
}


func main() {
    p := make([]byte, 2048)

    addr := net.UDPAddr{
        Port: PORT,
        IP: net.ParseIP(IP),
    }

    ser, err := net.ListenUDP("udp", &addr)

    if err != nil {
        fmt.Printf("Some error %v\n", err)
        return
    }

    for {
        var input string
        fmt.Printf("> ")
        fmt.Scanln(&input)

        _, remoteaddr, err := ser.ReadFromUDP(p)
        go sendResponse(ser, remoteaddr, input)
        fmt.Printf("%s\n", p)
        if err !=  nil {
            fmt.Printf("Some error  %v", err)
            continue
        }
    }
}
