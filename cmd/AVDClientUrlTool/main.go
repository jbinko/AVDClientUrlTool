package main

import (
	"fmt"
	"net"
	"sync"
)

type dnsRecord struct {
	name  string
	ips   []string
	error error
}

func checkDnsRecords(dnsRecords []dnsRecord) {

	dnsRecordsCount := len(dnsRecords)
	var wg sync.WaitGroup
	wg.Add(dnsRecordsCount)

	for i := 0; i < dnsRecordsCount; i++ {
		go func(i int) {
			defer wg.Done()
			checkDnsRecord(&dnsRecords[i])
		}(i)
	}

	wg.Wait()
}

func checkDnsRecord(dnsRecord *dnsRecord) {

	ips, err := net.LookupIP(dnsRecord.name)
	if err != nil {
		dnsRecord.error = err
	} else {

		ipsCount := len(ips)
		ipsString := make([]string, ipsCount)

		for i := 0; i < ipsCount; i++ {
			ipsString[i] = ips[i].String()
		}

		dnsRecord.ips = ipsString
	}
}

func printDnsRecords(dnsRecords []dnsRecord, printFailed bool) {

	for _, dnsRecord := range dnsRecords {

		if printFailed == true {

			if dnsRecord.error != nil {
				fmt.Printf("%s (%s)\n", dnsRecord.name, dnsRecord.error)
			}
		} else {

			if dnsRecord.error == nil {
				for _, ip := range dnsRecord.ips {
					fmt.Printf("%s (%s)\n", dnsRecord.name, ip)
				}
			}
		}
	}
}

func main() {

	// https://docs.microsoft.com/en-us/azure/virtual-desktop/safe-url-list?tabs=azure#remote-desktop-clients
	dnsRecords := []dnsRecord{
		{
			name: "client.wvd.microsoft.com",
		},
		{
			name: "watchdog.servicebus.windows.net",
		},
		{
			name: "go.microsoft.com",
		},
		{
			name: "aka.ms",
		},
		{
			name: "docs.microsoft.com",
		},
		{
			name: "privacy.microsoft.com",
		},
		{
			name: "query.prod.cms.rt.microsoft.com",
		},
	}

	checkDnsRecords(dnsRecords)

	fmt.Println("\n")
	fmt.Println("AVD Client Connectivity Check Tool")
	fmt.Println("======================================================")
	fmt.Println("\n")
	fmt.Println("NOT Resolvable URLs:")
	fmt.Println("======================================================")
	printDnsRecords(dnsRecords, true)
	fmt.Println("\n")
	fmt.Println("Resolvable URLs:")
	fmt.Println("======================================================")
	printDnsRecords(dnsRecords, false)
	fmt.Println("\n")
}
