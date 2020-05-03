package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Hostname: %s\n", name)
	percent, _ := cpu.Percent(time.Second, true)
	fmt.Fprintf(w, "CPU %s\n", percent)
	fmt.Fprintf(w, "Current TimeStamp: %s\n", time.Now().Format(time.RFC850))
	addrs, _ := net.LookupIP(name)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			fmt.Fprintf(w, "IPv4: %s\n", ipv4)
		}
	}
	v, _ := mem.VirtualMemory()

	fmt.Fprintf(w, "Memory Total: %d\n", v.Total)
	fmt.Fprintf(w, "Memory Free: %d\n", v.Free)
	fmt.Fprintf(w, "Memory Used: %g\n", v.UsedPercent)

}
