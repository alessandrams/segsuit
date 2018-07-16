package tcpconn

import (
  "fmt"
  "net"
  "time"
  )

func NewTcpConn(ip string, port string) (net.Conn, error) {
  bindaddr := DefineBindAddr(ip, port)
  tcpc, err := CreateConnection(bindaddr)
  if err != nil{
    return nil, err
  }
  return tcpc, nil
}

func DefineBindAddr(ip string, port string) string {
  addr := fmt.Sprintf("%s:%s", ip, port)
  return addr
}

func CreateConnection(addr string) (net.Conn, error){
    tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
    if err != nil {
      return nil, err
    }
    conn, err := net.DialTimeout("tcp", tcpAddr.String(), 1*time.Second)
    if err != nil {
		  return nil, err
    }
    return conn, nil
}
