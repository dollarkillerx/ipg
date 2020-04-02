package main

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

func main() {
	region, err := ip2region.New("ip2region.db")
	defer region.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	ip, err := region.MemorySearch("49.254.13.194")
	fmt.Println(ip, err)
	ip, err = region.BinarySearch("1.224.216.46")
	fmt.Println(ip, err)
	ip, err = region.BtreeSearch("1.201.0.1")
	fmt.Println(ip, err)
}
