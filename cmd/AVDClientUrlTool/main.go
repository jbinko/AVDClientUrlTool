package main

import (
	"fmt"
	"net"
)

type dnsRecord struct {
	name  string
	ips   []string
	error error
}

func checkDnsRecords(dnsRecords []dnsRecord) {

	for _, dnsRecord := range dnsRecords {
		checkDnsRecord(dnsRecord)
	}
}

func checkDnsRecord(dnsRecord dnsRecord) {

	ips, err := net.LookupIP(dnsRecord.name)
	if err != nil {
		dnsRecord.error = err
	} else {

		ipsString := make([]string, len(ips))

		for i := 0; i < len(ips); i++ {
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
