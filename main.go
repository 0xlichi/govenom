package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/0xlichi/govenom/input"
	"github.com/0xlichi/govenom/output"
	"github.com/0xlichi/govenom/ping"
	dnsrecon "github.com/0xlichi/govenom/recon/dns"
	netscanning "github.com/0xlichi/govenom/recon/network"
	subenum "github.com/0xlichi/govenom/recon/subdomain_enumeration"
	webrecon "github.com/0xlichi/govenom/recon/web"
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
	output.NewLine()

	startTime := time.Now()

	// ── Phase 1: Independent reconnaissance ──────────────────────────────────
	fmt.Println(output.Info("Phase 1: Reconnaissance"))
	fmt.Println(output.Info("Initializing subfinder..."))
	fmt.Println(output.Info("Initializing amass..."))
	fmt.Println(output.Info("Initializing assetfinder..."))
	fmt.Println(output.Info("Initializing nmap..."))
	fmt.Println(output.Info("Initializing wafw00f..."))
	fmt.Println(output.Info("Initializing gospider..."))
	fmt.Println(output.Info("Initializing whois..."))
	fmt.Println(output.Info("Initializing dig..."))
	output.NewLine()

	tracker1 := output.NewTracker()
	var wg1 sync.WaitGroup
	wg1.Add(8)

	go func() {
		tracker1.Add("subfinder")
		defer tracker1.Remove("subfinder")
		defer wg1.Done()
		subenum.Subfinder(host, tracker1)
	}()
	go func() {
		tracker1.Add("amass")
		defer tracker1.Remove("amass")
		defer wg1.Done()
		subenum.Amass(host, tracker1)
	}()
	go func() {
		tracker1.Add("assetfinder")
		defer tracker1.Remove("assetfinder")
		defer wg1.Done()
		subenum.Assetfinder(host, tracker1)
	}()
	go func() {
		tracker1.Add("nmap")
		defer tracker1.Remove("nmap")
		defer wg1.Done()
		netscanning.Nmap(host, tracker1)
	}()
	go func() {
		tracker1.Add("wafw00f")
		defer tracker1.Remove("wafw00f")
		defer wg1.Done()
		webrecon.Wafw00f(host, tracker1)
	}()
	go func() {
		tracker1.Add("gospider")
		defer tracker1.Remove("gospider")
		defer wg1.Done()
		webrecon.Gospider(host, tracker1)
	}()
	go func() {
		tracker1.Add("whois")
		defer tracker1.Remove("whois")
		defer wg1.Done()
		dnsrecon.Whois(host, tracker1)
	}()
	go func() {
		tracker1.Add("dig")
		defer tracker1.Remove("dig")
		defer wg1.Done()
		dnsrecon.Dig(host, tracker1)
	}()

	done1 := make(chan bool)
	go output.Spinner(tracker1, done1)
	wg1.Wait()
	done1 <- true

	output.NewLine()
	fmt.Println(output.Success("Phase 1 complete."))
	output.NewLine()

	// ── Phase 2: Deep scanning using phase 1 output ───────────────────────────
	fmt.Println(output.Info("Phase 2: Deep scanning"))
	fmt.Println(output.Info("Initializing dnsx..."))
	fmt.Println(output.Info("Initializing whatweb..."))
	output.NewLine()

	tracker2 := output.NewTracker()
	var wg2 sync.WaitGroup
	wg2.Add(2)

	go func() {
		tracker2.Add("dnsx")
		defer tracker2.Remove("dnsx")
		defer wg2.Done()
		dnsrecon.Dnsx(host, tracker2)
	}()
	go func() {
		tracker2.Add("whatweb")
		defer tracker2.Remove("whatweb")
		defer wg2.Done()
		webrecon.WhatWeb(host, tracker2)
	}()

	done2 := make(chan bool)
	go output.Spinner(tracker2, done2)
	wg2.Wait()
	done2 <- true

	output.NewLine()
	fmt.Println(output.Success("Phase 2 complete."))

	// ── Summary ───────────────────────────────────────────────────────────────
	output.NewLine()
	fmt.Println(output.Info("Scan summary:"))
	fmt.Println(output.Success("Subdomains (subfinder)   ->  logs/" + host + "/subdomains/subfinder-result.txt"))
	fmt.Println(output.Success("Subdomains (amass)       ->  logs/" + host + "/subdomains/amass-result.txt"))
	fmt.Println(output.Success("Subdomains (assetfinder) ->  logs/" + host + "/subdomains/assetfinder-result.txt"))
	fmt.Println(output.Success("Subdomains (merged)      ->  logs/" + host + "/subdomains/merged-subdomains.txt"))
	fmt.Println(output.Success("Network                  ->  logs/" + host + "/network/nmap-result.txt"))
	fmt.Println(output.Success("WAF                      ->  logs/" + host + "/web/wafw00f-result.txt"))
	fmt.Println(output.Success("Spider                   ->  logs/" + host + "/web/gospider-result.txt"))
	fmt.Println(output.Success("WhatWeb                  ->  logs/" + host + "/web/whatweb-result.txt"))
	fmt.Println(output.Success("WHOIS                    ->  logs/" + host + "/dns/whois-result.txt"))
	fmt.Println(output.Success("Dig                      ->  logs/" + host + "/dns/dig-result.txt"))
	fmt.Println(output.Success("DNS resolution (dnsx)    ->  logs/" + host + "/dns/dnsx-result.txt"))
	output.NewLine()
	fmt.Println(output.Info(fmt.Sprintf("Execution done in %v", time.Since(startTime))))
}
