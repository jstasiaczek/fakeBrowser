package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"time"
)

func GetUrl() string {
	length := len(os.Args)
	url := ""
	for i := 1; i < length; i++ {
		url += os.Args[i]
	}
	return url
}

func GetUniqueId() string {
	nano := time.Now().Format(time.StampNano)
	md5raw := md5.Sum([]byte(nano))
	return fmt.Sprintf("%x", md5raw)
}

func WaitForUser() {
	fmt.Println("Please press ENTER to continue...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
