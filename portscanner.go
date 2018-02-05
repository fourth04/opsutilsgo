package opsutils

import (
	"net"
	"strconv"
	"sync"
	"time"
)

func GetOpenedPort(ip string, ports []int, timeout time.Duration, parallelism int) []int {
	rv := []int{}
	sem := make(chan bool, parallelism)
	for _, port := range ports {
		sem <- true
		go func(port int) {
			if IsOpen(ip, port, timeout) {
				rv = append(rv, port)
			}
			<-sem
		}(port)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
	return rv
}

func IsPortsOpen(ip string, ports []int, timeout time.Duration, parallelism int) map[int]bool {
	rv := make(map[int]bool)
	sem := make(chan bool, parallelism)
	lock := sync.Mutex{}

	for _, port := range ports {
		sem <- true
		go func(port int) {
			result := IsOpen(ip, port, timeout)
			lock.Lock()
			defer lock.Unlock()
			rv[port] = result
			<-sem
		}(port)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
	return rv
}

func IsOpen(ip string, port int, timeout time.Duration) bool {
	conn, err := OpenConn(ip, port, timeout)
	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}

func OpenConn(ip string, port int, timeout time.Duration) (net.Conn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ip+":"+strconv.Itoa(port))
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTimeout("tcp", tcpAddr.String(), timeout)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
