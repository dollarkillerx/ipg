package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	open, err := os.Open("ip.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer open.Close()

	reader := bufio.NewReader(open)

	out, err := os.Create("success_ip.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer out.Close()
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		uui := string(line)
		if strings.Index(uui,"数据中心") != -1 {
			out.Write([]byte(fmt.Sprintf("%s\n",uui)))
		}
	}
}

