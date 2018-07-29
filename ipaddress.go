package opsutilsgo

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// IP2Num convert IP address to number
func IP2Num(ip string) int {
	canSplit := func(c rune) bool { return c == '.' }
	lisit := strings.FieldsFunc(ip, canSplit) //[58 215 20 30]
	ip1StrInt, _ := strconv.Atoi(lisit[0])
	ip2StrInt, _ := strconv.Atoi(lisit[1])
	ip3StrInt, _ := strconv.Atoi(lisit[2])
	ip4StrInt, _ := strconv.Atoi(lisit[3])
	return ip1StrInt<<24 | ip2StrInt<<16 | ip3StrInt<<8 | ip4StrInt
}

// Num2IP convert Number to IP address
func Num2IP(num int) string {
	ip1Int := (num & 0xff000000) >> 24
	ip2Int := (num & 0x00ff0000) >> 16
	ip3Int := (num & 0x0000ff00) >> 8
	ip4Int := num & 0x000000ff
	data := fmt.Sprintf("%d.%d.%d.%d", ip1Int, ip2Int, ip3Int, ip4Int)
	return data
}

// GenIPs make ips from start ip to stop ip
func GenIPs(Aip1 string, Aip2 string) []string {
	return GenIPsByNum(IP2Num(Aip1), IP2Num(Aip2))
}

// GenIPsByNum make ips from start ip to stop ip, but parameter is int
func GenIPsByNum(Aip1 int, Aip2 int) []string {
	l := []string{}
	index := Aip1
	for index <= Aip2 {
		ipData := Num2IP(index)
		l = append(l, ipData)
		index++
	}
	return l
}

// IP2Int convert ipv4 to uint32, ipv6 to 0
func IP2Int(ip net.IP) uint32 {
	ipv4 := ip.To4()
	if ipv4 != nil {
		return binary.BigEndian.Uint32(ipv4)
	}
	return 0
}

// IPStr2Int convert ip address string to int
func IPStr2Int(ipStr string) (uint32, error) {
	ip := net.ParseIP(ipStr)

	if ip == nil {
		return 0, errors.New("Err:无效的地址")
	}
	return IP2Int(ip), nil
}
