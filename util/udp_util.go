package util

import (
	"container/list"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func SendUdpPoint(ip string, port int, message string, useGbk bool) {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
	})
	if err != nil {

		return
	}
	defer socket.Close()
	sendData := []byte(message)
	if useGbk {
		gbk, err := Utf8ToGbk(sendData)
		if err == nil {
			sendData = gbk
		} else {
			log.Println(err)
		}
	}
	socket.Write(sendData)
	socket.Close()
}

func getIpList(startIp string, endIp string) *list.List {
	ips := list.New()
	starts := strings.Split(startIp, ".")
	ends := strings.Split(endIp, ".")
	startIp3, _ := strconv.Atoi(starts[2])
	startIp4, _ := strconv.Atoi(starts[3])
	endIp3, _ := strconv.Atoi(ends[2])
	endIp4, _ := strconv.Atoi(ends[3])

	ip := []string{starts[0], starts[1], starts[2], starts[3]}
	for i := startIp3; i <= startIp4; i++ {
		for x := 2; x <= 240; x++ {
			if (i == startIp3 && x >= startIp4) || (i == endIp3 && x <= endIp4) || (i > startIp3 && i < endIp3) {
				ip[2] = strconv.Itoa(i)
				ip[3] = strconv.Itoa(x)
				ips.PushBack(strings.Join(ip, "."))
			}
		}
	}

	return ips
}

func SendUdpMultiPoint(ips *list.List, port int, message string) {
	for e := ips.Front(); e != nil; e = e.Next() {

		switch e.Value.(type) {
		case string:
			SendUdpPoint(e.Value.(string), port, message, false)
		}
	}

}

func SendMultiPointFrom1To3(port int, message string) {
	ips := getIpList("191.168.1.6", "191.168.3.225")
	ips.PushBack("191.168.6.44")
	ips.PushBack("191.168.6.225")
	ips.PushBack(GetIp())
	SendUdpMultiPoint(ips, port, message)
}

func SendMultiPointFrom1To6(port int, message string) {
	ips := getIpList("191.168.1.6", "191.168.6.225")
	SendUdpMultiPoint(ips, port, message)
}

func GetIp() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}
	var ip = "localhost"
	for i := 0; i < len(netInterfaces); i++ {
		// log.Println(netInterfaces[i].Name)
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if inet, ok := address.(*net.IPNet); ok && !inet.IP.IsLoopback() {
					if inet.IP.To4() != nil && strings.HasPrefix(inet.IP.String(), "191.168.") {
						ip = inet.IP.String()
					}
					if inet.IP.To4() != nil && ip == "localhost" && strings.HasPrefix(netInterfaces[i].Name, "eth") {
						ip = inet.IP.String()
					}
				}
			}
		}
	}
	return ip
}
