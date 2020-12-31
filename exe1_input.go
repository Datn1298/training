package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"bufio"
	"os"
	"strings"
)

type Instances struct {
	Instances []Data `json:"Instances"`
}

type Data struct {
	Type   string `json:"type"`
	VCPU   int    `json:"vCPU"`
	VRam   int    `json:"vRAM"`
	Counts int    `json:"counts"`
	Action string
	Amount int
}

func Abs(numb int) int {
	if numb > 0 {
		return numb
	} else {
		return -numb
	}
}

const path = "./config.json"
const path1 = "./config_1.json"

// func isJsonFile(fileName string) bool {
// 	temp := fileName[len(fileName)-7:]
// 	temp = strings.TrimRight(temp, "\r\n")
// 	if strings.Compare(temp, ".json") == 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func isExit(fileName string) bool {
// 	temp := fileName[len(fileName)-5:]
// 	if strings.Compare(temp, "Exit") == 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }



func getProduct(path string) []Instances {
	path = strings.TrimRight(path, "\r\n")
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	
	product := make([]Instances, 3)

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &product)
	return product
}
 
func abc(path string, queue []Instances) []Instances {

	products = getProduct(path)
	if len(queue) == 0 {
		for i := range products.Instances {
			queue = append(queue, products.Instances[i])
			queue[i].Action = ""
		}	
	} else {
		for i := range products {
			temp1 := products[i]

			temp := queue[0]

			queue = queue[1:]

			if(temp.Counts < temp1.Counts){		
				temp.Action = "provision"
			} else {
				temp.Action = "delete"
			}
			temp = temp1
			temp.Amount = Abs(temp1.Counts - temp.Counts)
			queue = append(queue, temp)
		}	
	}
	return queue
}

func main()  {
	var queue []Instances
	var queue2 []Instances
	
	queue = abc(path, queue)

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]   // Dequeue
		fmt.Println(temp.Type , " " , temp.Action , " " , temp.Counts)
		fmt.Println("===============")
		queue2 = append(queue2, temp)
	}
	fmt.Println()
	queue = queue2
	for len(queue2) > 0 {
		queue2 = queue2[1:]   // Dequeue
	}

	queue = abc(path2, queue)

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]   // Dequeue
		fmt.Println(temp.Type , " " , temp.Action , " " , temp.Counts)
		fmt.Println("===============")
		queue2 = append(queue2, temp)
	}
	fmt.Println()
	queue = queue2
	for len(queue2) > 0 {
		queue2 = queue2[1:]   // Dequeue
	}
}

	// reader := bufio.NewReader(os.Stdin)
	// for exit := 1; exit != 2; {
	// 	fmt.Print("Enter path: ")
	// 	path, _ := reader.ReadString('\n')

	// 	if isExit(path) {
	// 		exit = 2
	// 		break
	// 	} else {
	// 		if isJsonFile(path) {
	// 			queue = abc(path, queue)

	// 			for len(queue) > 0 {
	// 				temp := queue[0]
	// 				queue = queue[1:]   // Dequeue
	// 				fmt.Println(temp.Type , " " , temp.Action , " " , temp.Amount)
	// 				fmt.Println("===============")
	// 				queue2 = append(queue2, temp)

	// 				fmt.Println()

	// 				queue = queue2

	// 				for len(queue2) > 0 {
	// 					queue2 = queue2[1:]   // Dequeue
	// 				}
	// 			}
	// 		} else {
	// 			fmt.Println("Invalid file! Enter path again: ")
	// 		}

	// 	}
	// }


