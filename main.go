package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/0xlichi/govenom/input"
	"github.com/0xlichi/govenom/output"
	"github.com/0xlichi/govenom/ping"
	netscanning "github.com/0xlichi/govenom/reconnaissance/network"
	subenum "github.com/0xlichi/govenom/reconnaissance/subdomain_enumeration"
	webvuln "github.com/0xlichi/govenom/reconnaissance/web"
	"github.com/0xlichi/govenom/ui/banner"
)

func init() {
	banner.GetBanner()
}

func main() {
	host := input.GetHost()

	if !ping.CheckHost(host) {
		fmt.Println(output.Error(fmt.Sprintf("Host '%v' is not reachable.", host)))
		os.Exit(1)
	}
	fmt.Println(output.Success(fmt.Sprintf("Host '%v' is reachable.", host)))

	startTime := time.Now()

	var wg sync.WaitGroup

	fmt.Println(output.Info("Initializing subfinder..."))
	fmt.Println(output.Info("Initializing nmap..."))
	fmt.Println(output.Info("Initializing wafw00f"))

	wg.Add(3)
	go func() { defer wg.Done(); subenum.Subfinder(host) }()
	go func() { defer wg.Done(); netscanning.Nmap(host) }()
	go func() { defer wg.Done(); webvuln.Wafw00f(host) }()
	wg.Wait()

	fmt.Println(output.Info(fmt.Sprintf("Execution done in %v", time.Since(startTime))))
}
