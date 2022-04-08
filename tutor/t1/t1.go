package main

import (
	"fmt"
	"net"
	"sync"
)

var x int64
var wg sync.WaitGroup

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}
func main1() {
	ip, _ := GetIP()
	var ips []string
	for _, pp := range ip {
		ips = append(ips, pp)
	}
	print(ips)
}
func GetIP() ([]string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("net.Interfaces failed, err: %s", err.Error())
	}
	var ip = make([]string, 0)
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil && ipnet.IP.String() != "" {
						ip = append(ip, ipnet.IP.String())
					}
				}
			}
		}
	}
	return ip, nil
}
func GetIP2() (map[int]string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("net.Interfaces failed, err: %s", err.Error())
	}
	var ip = make(map[int]string)
	var j = 0
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil && ipnet.IP.String() != "" {
						ip[j] = ipnet.IP.String()
						j++
					}
				}
			}
		}
	}
	return ip, nil
}
