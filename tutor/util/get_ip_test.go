package util

import "testing"

func TestGetIP(t *testing.T) {
	ip, _ := GetIP()
	var ips []string
	for _, pp := range ip {
		ips = append(ips, pp)
	}
	print(ips)

}

func TestLocalIPv4s(t *testing.T) {
	s, err := LocalIPv4s()
	if err != nil {
		return
	}
	for st := range s {
		println(s[st])
	}
}
