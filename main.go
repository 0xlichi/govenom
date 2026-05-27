package main

import (
	"fmt"
	"os"
	"time"

	"github.com/0xlichi/govenom/input"
	"github.com/0xlichi/govenom/output"
	"github.com/0xlichi/govenom/ping"
	netscanning "github.com/0xlichi/govenom/reconnaissance/network"
	subenum "github.com/0xlichi/govenom/reconnaissance/subdomain_enumeration"
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
	fmt.Println(output.Success(fmt.Sprintf("Host '%v' is reachable.\n", host)))

	startTime := time.Now()

	subenum.Subfinder(host)
	netscanning.Nmap(host)

	fmt.Println(output.Info(fmt.Sprintf("Execution done in %v", time.Since(startTime))))
}
