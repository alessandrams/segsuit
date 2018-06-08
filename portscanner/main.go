package main

import (
    "fmt"
    "strings"
    "portscanner/host"
    "portscanner/port"
    //"bufio"
    //"os"
)

func main() {

    var addrip string
    fmt.Printf("Informe o IP a ser analisado: ")
    fmt.Scanf("%s", &addrip)
    h := host.NewHost(addrip, 2)

    y := strings.Repeat("-", 45)
    fmt.Println("\n[INFO] Starting Port-Scanning... \n\n Port\t\tStatus\t\tDescription\n", y)

    openports := h.GetOpenPorts()

    p := port.NewPort("1", false)

    for i := 0; i < len(openports); i++ {
     	port := openports[i]
      desc := p.DescribePort(port)
     	fmt.Println(" ",port, "\t--->\t[open]\t---->\t", desc)
     }
}
