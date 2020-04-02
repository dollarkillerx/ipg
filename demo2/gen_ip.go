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
	open, err := os.Open("su_ip.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer open.Close()

	out, err := os.Create("out_ip.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer out.Close()
	reader := bufio.NewReader(open)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatalln(err)
		}
		uui := string(line)
		split := strings.Split(uui, "-")
		ips, err := utils.RangeIP(split[0], split[1])
		if err != nil {
			log.Println("S1: ",err)
			continue
		}
		for _,v := range ips {
			//_, err := out.Write([]byte(fmt.Sprintf("%s,%s\n", v, split[2])))
			_, err := out.Write([]byte(fmt.Sprintf("%s\n", v)))
			if err != nil {
				log.Println("S2: ",err)
				continue
			}
		}
	}
}
