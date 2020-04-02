package main

import (
	"bufio"
	"fmt"
	"ipg/utils"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("??? ./ipg tag.txt")
	}
	filePath := os.Args[1]
	outPath := utils.Md5Encode(filePath) + "_out.txt"
	Run(filePath,outPath)
}

func Run (tagPath,outPath string) {
	open, err := os.Open(tagPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer open.Close()
	file, err := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 00755)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	reader := bufio.NewReader(open)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		scd := string(line)
		switch  {
		case strings.Index(scd,"/") != -1:
			cidr, err := utils.Cidr(scd)
			if err != nil {
				log.Println(err)
				continue
			}
			for _,v := range cidr {
				file.Write([]byte(fmt.Sprintf("%s\n",v)))
			}
			continue
		case strings.Index(scd,"-") != -1:
			split := strings.Split(scd, "-")
			if len(split) != 2 {
				log.Println("Please check the parameters you entered ???")
				continue
			}
			ips, err := utils.RangeIP(split[0], split[1])
			if err != nil {
				log.Println(err)
				continue
			}
			for _,v := range ips {
				file.Write([]byte(fmt.Sprintf("%s\n",v)))
			}
		}
	}

}
