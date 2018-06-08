package host

import (
    "net"
    "time"
    "sync"
    "tcpconn"
    "strconv"
)

// Fix and complete interval of TCP ports
const firstTCPport int = 1
const lastTCPPort int = 65535


type Host struct {
   AddrIP string
   Timeout time.Duration
}

func NewHost(ip string, t time.Duration) (host Host) {
    host.SetIp(ip)
    host.SetTimeout(t)
    return host
}

func (h *Host) SetIp(ip string) {
    h.AddrIP = ip
}


func (h *Host) SetTimeout(t time.Duration) {
    tempo := t*time.Second
    h.Timeout = tempo
}


// Verify if a port is open
func (h Host) VerifyOpenPort(port string) (isOpen bool){
    tcpAddr, err := net.ResolveTCPAddr("tcp", tcpconn.DefineBindAddr(h.AddrIP,port))
    if err != nil {
        return false
    }
    conn, err := net.DialTimeout("tcp", tcpAddr.String(), h.Timeout)
    if err != nil {
		    return false
    }
    defer conn.Close()
    return true
}

// Get all open ports and save in a list
func (h Host) GetOpenPorts() []string {
    openports := []string{}
    s := sync.Mutex{}
    c := make(chan bool, 5)
    for port := firstTCPport; port<= lastTCPPort; port++ {
        c <- true
        sport := strconv.Itoa(port)
        go func (port string){
          if h.VerifyOpenPort(sport){
              s.Lock()
              openports = append(openports, sport)
              s.Unlock()
          }
          <- c
        }(sport)
    }
    for i := 0; i < cap(c); i++ {
    	 c <- true
    }
    return openports
}
