package validation

import (
	"net"
	"strconv"
)

func IsIP(s string) bool {
	return net.ParseIP(s) != nil
}

func IsCIDR(s string) bool {
	_,_,err := net.ParseCIDR(s)
	return err==nil
}

func IsPort(s string) bool {
	if i, err := strconv.Atoi(s); err==nil && i>0 && i<65536 {
		return true
	}
	return false
}

func IsMAC(s string) bool {
	_, err := net.ParseMAC(s)
	return err == nil
}
