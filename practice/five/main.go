package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMx, hasSPF, sprRecord,hasDMRAC,dmarRecord")

	for scanner.Scan() {
		checkdomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkdomain(domain string) {
	var hasMx, hasSPF, hasDMRAC bool
	var sprRecord, dmarRecord string

	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Fatal(err)
	}

	if len(mxRecord) > 0 {
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			sprRecord = record
			break
		}
	}

	dmarcRrecod, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range dmarcRrecod {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMRAC = true
			dmarRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMx, hasSPF, sprRecord, hasDMRAC, dmarRecord)
}
