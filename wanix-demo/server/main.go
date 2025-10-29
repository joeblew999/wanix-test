package main

import (
	"log"
	"net"
	"net/http"
	"path/filepath"
	"sort"
)

func main() {
	staticDir := filepath.Join("..", "static")
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(staticDir)))

	addr := ":8080"
	log.Printf("serving static assets from %s", staticDir)
	log.Printf("local demo URL:    http://localhost%s", addr)
	for _, ip := range lanIPs() {
		log.Printf("LAN demo URL:      http://%s%s", ip, addr)
	}

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func lanIPs() []string {
	ipSet := make(map[string]struct{})
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			ipSet[ip.String()] = struct{}{}
		}
	}
	if len(ipSet) == 0 {
		return nil
	}
	ips := make([]string, 0, len(ipSet))
	for ip := range ipSet {
		ips = append(ips, ip)
	}
	sort.Strings(ips)
	return ips
}
