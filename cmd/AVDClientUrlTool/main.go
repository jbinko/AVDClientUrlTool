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
			checkDnsRecord(dnsRecords[i])
		}(i)
	}

	wg.Wait()
}

func checkDnsRecord(dnsRecord dnsRecord) {

	ips, err := net.LookupIP(dnsRecord.name)
	if err != nil {
		dnsRecord.error = err
	} else {

		ipsCount := len(ips)
		ipsString := make([]string, ipsCount)

		for i := 0; i < ipsCount; i++ {
			ipsString[i] = ips[i].String()
		}
	}
}

func main() {

	dnsRecords := []dnsRecord{
		{
			name: "client.wvd.microsoft.com",
		},
		{
			name: "client2.wvd.microsoft.com",
		},
	}

	checkDnsRecords(dnsRecords)

	fmt.Println("\n")
	fmt.Println("AVD")
	fmt.Println("======================================================")
	fmt.Println("\n")
	fmt.Println("NOT Resolvable URLs:")
	fmt.Println("======================================================")
	fmt.Println("\n")
	fmt.Println("Resolvable URLs:")
	fmt.Println("======================================================")
	fmt.Println("client.wvd.microsoft.com")
	fmt.Println("\n")
}
