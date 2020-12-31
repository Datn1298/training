package main

import (
	"fmt"
	// "bufio"
    // "fmt"
    // "io"
    "io/ioutil"
    // "os"	
)

const path = "./config.json"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// type Data struct {
// 	type string
// 	vCPU int
// 	vRam int
// 	counts int
// }


func main()  {

	exe := path[len(path)-4: len(path)]

	if exe == "json"{
		data, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(string(data))
		
	}

	
}