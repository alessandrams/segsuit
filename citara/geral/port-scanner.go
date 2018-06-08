package main

import (
    "net"
    "fmt"
    "bufio"
    "os"
    "time"
)

type Host struct {
   AddrIP string
   ServerAddr string
   timeout time.Duration
   IsUp  bool
}

func (h Host) serverHost(port int, ip string) (serverAddr string){
    serverAddr = fmt.Sprintf("%s:%d", h.host, port)
    return serverAddr
}

func readInputIP() (addrIP string){
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("Informe o IP a ser analisado: ")
    addrIP, _ := reader.ReadString('\n')
    return addrIP
}


func verifyOpenPort() (isOpen bool){
    tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
    if err != nil {
      fmt.Println("Impossible to resolve TCP Addr for %s",servAddr)
    }
    conn, err := net.Dial("tcp", servAddr, time.Duration(sec)*time.Second)
}

func main(){
    var host Host
    host.AddrIP := readIP()


}
