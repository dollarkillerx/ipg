package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	open, err := os.Open("success_ip.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer open.Close()

	out, err := os.Create("su_ip.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer out.Close()
	reader := bufio.NewReader(open)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		uui := string(line)
		if strings.Index(uui,"KT") != -1 {
			out.Write([]byte(fmt.Sprintf("%s\n",uui)))
		}
	}
}
