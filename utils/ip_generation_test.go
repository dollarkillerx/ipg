package utils

import (
	"fmt"
	"log"
	"testing"
)

func TestCidr(t *testing.T) {
	cidr, err := Cidr("192.168.11.9/27")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cidr)
}

func TestRangeIP(t *testing.T) {
	ip, err := RangeIP("221.177.0.0", "221.177.255.1")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(ip)
}