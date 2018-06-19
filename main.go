package main

import (
	"fmt"
	"time"
	"os"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/uptime"
)

func main() {
	fmt.Println("-----------------------------")
	fmt.Println("LOGTIME", time.Now())
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	fmt.Printf("\nmemory total : %d bytes\n\n", memory.Total)
	fmt.Printf("memory used  : %d bytes      |      %d%%\n\n", memory.Used, memory.Used * 100 / memory.Total)

	uptime, err := uptime.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	fmt.Println("total uptime :", uptime, "\n")
}
