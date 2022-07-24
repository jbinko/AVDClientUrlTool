package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
)

type urlRecord struct {
	dnsName            string
	dnsResolvedIPs     []string
	dnsResolutionError error
	urlConnect         string
	urlConnectStatus   string
	urlConnectError    error
}

func checkDnsRecords(urlRecords []urlRecord) {

	urlRecordsCount := len(urlRecords)
	var wg sync.WaitGroup
	wg.Add(urlRecordsCount)

	for i := 0; i < urlRecordsCount; i++ {
		go func(i int) {
			defer wg.Done()
			err := checkDnsRecord(&urlRecords[i])
			if err == nil {
				checkUrlConnectRecord(&urlRecords[i])
			}
		}(i)
	}

	wg.Wait()
}

func checkDnsRecord(urlRecord *urlRecord) error {

	ips, err := net.LookupIP(urlRecord.dnsName)
	if err != nil {
		urlRecord.dnsResolutionError = err
	} else {

		ipsCount := len(ips)
		ipsString := make([]string, ipsCount)

		for i := 0; i < ipsCount; i++ {
			ipsString[i] = ips[i].String()
		}

		urlRecord.dnsResolvedIPs = ipsString
	}

	return urlRecord.dnsResolutionError
}

func checkUrlConnectRecord(urlRecord *urlRecord) {

	resp, err := http.Get(urlRecord.urlConnect)
	if err != nil {
		urlRecord.urlConnectError = err
	} else {
		defer resp.Body.Close()
		_, err := io.ReadAll(resp.Body)
		if err != nil {
			urlRecord.urlConnectError = err
		} else {
			urlRecord.urlConnectStatus = http.StatusText(resp.StatusCode)
		}
	}
}

func printDnsRecords(urlRecords []urlRecord, printFailedOnly bool) {

	for _, urlRecord := range urlRecords {

		if printFailedOnly == true {

			if urlRecord.dnsResolutionError != nil {
				fmt.Printf("%s (%s)\n", urlRecord.dnsName, urlRecord.dnsResolutionError)
			}
		} else {

			if urlRecord.dnsResolutionError == nil {
				for _, ip := range urlRecord.dnsResolvedIPs {
					fmt.Printf("%s (%s)\n", urlRecord.dnsName, ip)
				}
			}
		}
	}
}

func printUrlConnectRecords(urlRecords []urlRecord, printFailedOnly bool) {

	for _, urlRecord := range urlRecords {

		if printFailedOnly == true {

			if urlRecord.urlConnectError != nil {
				fmt.Printf("%s (%s)\n", urlRecord.urlConnect, urlRecord.urlConnectError)
			}
		} else {

			if urlRecord.urlConnectError == nil {
				fmt.Printf("%s (%s)\n", urlRecord.urlConnect, urlRecord.urlConnectStatus)
			}
		}
	}
}

func main() {

	urlRecords := []urlRecord{
		// These URLs only correspond to client sites and resources
		// https://docs.microsoft.com/en-us/azure/virtual-desktop/safe-url-list?tabs=azure#remote-desktop-clients
		{
			dnsName:    "client.wvd.microsoft.com", // TODO - *.wvd.microsoft.com
			urlConnect: "https://client.wvd.microsoft.com/arm/webclient/index.html",
		},
		{
			dnsName: "watchdog.servicebus.windows.net", // TODO - *.servicebus.windows.net
		},
		{
			dnsName: "go.microsoft.com",
		},
		{
			dnsName: "aka.ms",
		},
		{
			dnsName: "docs.microsoft.com",
		},
		{
			dnsName: "privacy.microsoft.com",
		},
		{
			dnsName: "query.prod.cms.rt.microsoft.com",
		},
		// Azure Active Directory URLs can be found under IDs 56, 59 and 125
		// https://docs.microsoft.com/en-us/microsoft-365/enterprise/urls-and-ip-address-ranges?view=o365-worldwide#microsoft-365-common-and-office-online
		// Section 56
		{
			dnsName: "credentials.auth.microsoft.com", // TODO - *.auth.microsoft.com
		},
		{
			dnsName: "msftidentity.com", // TODO - *.msftidentity.com
		},
		{
			dnsName: "msidentity.com", // TODO - *.msidentity.com
		},
		{
			dnsName: "account.activedirectory.windowsazure.com",
		},
		{
			dnsName: "accounts.accesscontrol.windows.net",
		},
		{
			dnsName: "adminwebservice.microsoftonline.com",
		},
		{
			dnsName: "api.passwordreset.microsoftonline.com",
		},
		{
			dnsName: "autologon.microsoftazuread-sso.com",
		},
		{
			dnsName: "becws.microsoftonline.com",
		},
		{
			dnsName: "ccs.login.microsoftonline.com",
		},
		{
			dnsName: "clientconfig.microsoftonline-p.net",
		},
		{
			dnsName: "companymanager.microsoftonline.com",
		},
		{
			dnsName: "device.login.microsoftonline.com",
		},
		{
			dnsName: "graph.microsoft.com",
		},
		{
			dnsName: "graph.windows.net",
		},
		{
			dnsName: "login.microsoft.com",
		},
		{
			dnsName: "login.microsoftonline.com",
		},
		{
			dnsName: "login.microsoftonline-p.com",
		},
		{
			dnsName: "login.windows.net",
		},
		{
			dnsName: "logincert.microsoftonline.com",
		},
		{
			dnsName: "loginex.microsoftonline.com",
		},
		{
			dnsName: "login-us.microsoftonline.com",
		},
		{
			dnsName: "nexus.microsoftonline-p.com",
		},
		{
			dnsName: "nexus.microsoftonline-p.com",
		},
		{
			dnsName: "passwordreset.microsoftonline.com",
		},
		{
			dnsName: "provisioningapi.microsoftonline.com",
		},
	}

	checkDnsRecords(urlRecords)

	fmt.Println("\n")
	fmt.Println("AVD Client Connectivity Check Tool")
	fmt.Println("======================================================")
	fmt.Println("\n")
	fmt.Println("NOT Resolvable URLs:")
	fmt.Println("======================================================")
	printDnsRecords(urlRecords, true)
	fmt.Println("\n")
	fmt.Println("Resolvable URLs:")
	fmt.Println("======================================================")
	printDnsRecords(urlRecords, false)
	fmt.Println("\n")
	fmt.Println("NOT Reachable URLs:")
	fmt.Println("======================================================")
	printUrlConnectRecords(urlRecords, true)
	fmt.Println("\n")
	fmt.Println("Reachable URLs:")
	fmt.Println("======================================================")
	printUrlConnectRecords(urlRecords, false)
	fmt.Println("\n")
}

// TODO
// - Check socket port
// - Resolve IPs - E.g. 2603:1006:2000::/48
// - Section 59 and 125
// - private/public shortpath
