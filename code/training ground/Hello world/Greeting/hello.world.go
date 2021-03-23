package Greeting

import (
	"fmt"
	"io/ioutil"
	"log"
)

func HelloWorld() {

	content, err := ioutil.ReadFile("Greeting/hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)
	//fmt.Println("/n/n")

}
