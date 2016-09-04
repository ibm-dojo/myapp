package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func getData() string {
	name, _ := os.Hostname()
	result := fmt.Sprintf("<pre><b>Host: %s  Date: %s</b>\n", name, time.Now())
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		str := addr.String()
		if str[0] >= '0' && str[0] <= '9' {
			result += fmt.Sprintf("%s\n", strings.Split(str, "/")[0])
		}
	}
	return result
}

func main() {
	addr, port := "0.0.0.0", "80"

	if len(os.Args) > 2 {
		addr, port = os.Args[1], os.Args[2]
	} else if len(os.Args) > 1 {
		port = os.Args[1]
	}

	fmt.Printf("Will show:\n%s\nListening on: %s:%s\n", getData(), addr, port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(getData()))
	})
	if err := http.ListenAndServe(addr+":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
