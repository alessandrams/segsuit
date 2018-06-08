package port

import (
    "strconv"
  )

var known_ports = map[int]string {
	21:    "FTP",
	22:    "SSH",
	23:    "Telnet",
	25:    "SMTP",
	66:    "Oracle SQL*NET?",
	69:    "TFTP",
	80:    "HTTP",
	88:    "Kerberos",
	109:   "POP",
	110:   "POP3",
  118:   "SQL services",
	123:   "NTP",
	137:   "Netbios",
  138:   "Netbios",
	139:   "Netbios",
  156:   "SQL services",
  194:   "IRC",
	443:   "HTTPS",
  554:   "RTSP",
	1433:  "Microsoft SQL server",
	1434:  "Microsoft SQL monitor",
	3306:  "MySQL",
	3396:  "Novell NDPS Printer Agent",
	3535:  "SMTP (alternate)",
  5432:  "PostgreSQL",
  5433:  "PostgreSQL",
  5800:  "VNC remote desktop",
  7000:  "Cassandra Cluster Com [ http://cassandra.apache.org/ ]",
  9042:  "Cassandra [ http://cassandra.apache.org/ ]",
	9160:  "Cassandra [ http://cassandra.apache.org/ ]",
  27017: "Mongodb [ http://www.mongodb.org/ ]",
	28017: "Mongodb Web Admin [ http://www.mongodb.org/ ]",
}

type Port struct {
   Number string
   IsUp bool
   Description string
}

func NewPort(number string, isup bool) (port Port) {
  port.SetNumber(number)
  port.SetIsup(isup)
  port.DescribePort(number)
  return port
}

func (p *Port) SetNumber(number string) {
  p.Number = number
}

func (p *Port) SetIsup(isup bool) {
  p.IsUp = isup
}

func (p Port) DescribePort(port string) string{
  sport, _ := strconv.Atoi(port)
  if desc, ok := known_ports[sport]; ok {
    return desc
  }
  return "Unknown"
}
