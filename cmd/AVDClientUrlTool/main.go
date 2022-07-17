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

	dnsRecords := []dnsRecord{
		// These URLs only correspond to client sites and resources
		// https://docs.microsoft.com/en-us/azure/virtual-desktop/safe-url-list?tabs=azure#remote-desktop-clients
		{
			name: "client.wvd.microsoft.com", // TODO - *.wvd.microsoft.com
		},
		{
			name: "watchdog.servicebus.windows.net", // TODO - *.servicebus.windows.net
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
		// Azure Active Directory URLs can be found under IDs 56, 59 and 125
		// https://docs.microsoft.com/en-us/microsoft-365/enterprise/urls-and-ip-address-ranges?view=o365-worldwide#microsoft-365-common-and-office-online
		// Section 56
		{
			name: "credentials.auth.microsoft.com", // TODO - *.auth.microsoft.com
		},
		{
			name: "msftidentity.com", // TODO - *.msftidentity.com
		},
		{
			name: "msidentity.com", // TODO - *.msidentity.com
		},
		{
			name: "account.activedirectory.windowsazure.com",
		},
		{
			name: "accounts.accesscontrol.windows.net",
		},
		{
			name: "adminwebservice.microsoftonline.com",
		},
		{
			name: "api.passwordreset.microsoftonline.com",
		},
		{
			name: "autologon.microsoftazuread-sso.com",
		},
		{
			name: "becws.microsoftonline.com",
		},
		{
			name: "ccs.login.microsoftonline.com",
		},
		{
			name: "clientconfig.microsoftonline-p.net",
		},
		{
			name: "companymanager.microsoftonline.com",
		},
		{
			name: "device.login.microsoftonline.com",
		},
		{
			name: "graph.microsoft.com",
		},
		{
			name: "graph.windows.net",
		},
		{
			name: "login.microsoft.com",
		},
		{
			name: "login.microsoftonline.com",
		},
		{
			name: "login.microsoftonline-p.com",
		},
		{
			name: "login.windows.net",
		},
		{
			name: "logincert.microsoftonline.com",
		},
		{
			name: "loginex.microsoftonline.com",
		},
		{
			name: "login-us.microsoftonline.com",
		},
		{
			name: "nexus.microsoftonline-p.com",
		},
		{
			name: "nexus.microsoftonline-p.com",
		},
		{
			name: "passwordreset.microsoftonline.com",
		},
		{
			name: "provisioningapi.microsoftonline.com",
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

// TODO
// - Check socket port
// - Resolve IPs - E.g. 2603:1006:2000::/48
// - Section 59 and 125
// - webassembly
