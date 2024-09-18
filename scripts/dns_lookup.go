package main

import (
	"fmt"
	"net"
)

func dns_addr(domain string) []string {
	// if len(os.Args) < 2 {
	// 	fmt.Println("usage: ./scraper <domain>")
	// 	return
	// }
	// domain := os.Args[1]

	addr, err := net.LookupHost(domain)
	if err != nil {
		fmt.Printf("Encountered Error : %s", err)
		return nil
	} else {
		fmt.Println("Records for :", domain)
		for num, adr := range addr {
			fmt.Println(num, " ", adr)
		}
		return addr
	}
}

func dns_ns(domain string) []string {
	NSrecs, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println("Encountered Error : ", err)
	} else {

		fmt.Println("NS Records :")
		for ns := range NSrecs {
			println(ns)
		}
		return NSrecs
}