package main

import (
  "fmt"
  "tcpconn"
)

func main() {
  var addrip, port string
  fmt.Printf("Informe o IP a ser analisado: ")
  fmt.Scanf("%s", &addrip)
  fmt.Printf("Informe a porta: ")
  fmt.Scanf("%s", &port)

  tcpconn.NewTcpConn(addrip,port)
}
