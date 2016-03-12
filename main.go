package main

import (
	"fmt"
)

func main() {
	if GenerateReg {
		OutputRegFile()
	} else {
		ProcessUrl()
	}

}

func ProcessUrl() {
	var err error
	fmt.Print("Fake Browser: ")

	if (App.UseFiles && App.UseHttp) || (!App.UseFiles && !App.UseHttp) {
		fmt.Println("Sou you want files or http?")
		WaitForUser()
		return
	}
	url := GetUrl()
	if url == "" {
		fmt.Println("No url to open!")
		WaitForUser()
		return
	}
	if App.UseFiles {
		err = WriteFile(url)
	}
	if App.UseHttp {
		err = SendUrl(url)
	}
	if err != nil {
		fmt.Println(err)
		WaitForUser()
	} else {
		fmt.Println(url)
	}
}
