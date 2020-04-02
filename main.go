package main

import (
	"fmt"
	"ipg/utils"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("???    ./ipg 192.168.11.9/27   or  ./ipg 221.177.0.0-221.177.255.1")
	}
	param := os.Args[1]
	outPath := utils.Md5Encode(param) + "_out.txt"
	switch  {
	case strings.Index(param,"/") != -1:
		generationCidr(outPath,param)
	case strings.Index(param,"-") != -1:
		generationRangIP(outPath,param)
	}
	fmt.Println("Over")
}

func generationCidr(path ,cidr string) {
	i, err := utils.Cidr(cidr)
	if err != nil {
		log.Fatalln(err)
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 00755)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	for _,v := range i {
		file.Write([]byte(fmt.Sprintf("%s\n",v)))
	}
}

func generationRangIP(path,rangs string) {
	split := strings.Split(rangs, "-")
	if len(split) != 2 {
		log.Fatalln("Please check the parameters you entered ???")
	}
	ip, err := utils.RangeIP(split[0], split[1])
	if err != nil {
		log.Fatalln(err)
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 00755)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	for _,v := range ip {
		file.Write([]byte(fmt.Sprintf("%s\n",v)))
	}
}