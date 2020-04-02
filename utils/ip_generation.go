package utils

import (
	"bytes"
	"encoding/binary"
	"net"
	"strconv"
)

func Cidr(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	return ips[1 : len(ips)-1], nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}


//ip到数字
func ip2Long(ip string) (uint32,error) {
	var long uint32
	err := binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &long)
	return long,err
}

//数字到IP
func backtoIP4(ipInt int64) string {
	// need to do two bit shifting and “0xff” masking
	b0 := strconv.FormatInt((ipInt>>24)&0xff, 10)
	b1 := strconv.FormatInt((ipInt>>16)&0xff, 10)
	b2 := strconv.FormatInt((ipInt>>8)&0xff, 10)
	b3 := strconv.FormatInt((ipInt & 0xff), 10)
	return b0 + "." + b1 + "." + b2 + "." + b3
}

func RangeIP(start,end string) ([]string,error) {
	result := make([]string,0)
	ip1,err := ip2Long(start)
	if err != nil {
		return nil,err
	}
	ip2,err := ip2Long(end)
	if err != nil {
		return nil,err
	}
	for i := ip1; i <= ip2; i++ {
		i := int64(i)
		result = append(result,backtoIP4(i))
	}
	return result,nil
}