package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Instance struct {
	Instance []Product `json:"Instances"`
}

type Product struct {
	Type   string `json:"type"`
	VCPU   int    `json:"vCPU"`
	VRam   int    `json:"vRam"`
	Counts int    `json:"counts"`
}

func isJsonFile(fileName string) bool {
	temp := fileName[len(fileName)-7:]
	temp = strings.TrimRight(temp, "\r\n")
	if strings.Compare(temp, ".json") == 0 {
		return true
	} else {
		return false
	}
}

func isExit(fileName string) bool {
	temp := fileName[len(fileName)-5:]
	if strings.Compare(temp, "Exit") == 0 {
		return true
	} else {
		return false
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readData(fileName string) *Instance {
	fileName = strings.TrimRight(fileName, "\r\n")
	jsonFile, err := os.Open(fileName)
	if err != nil {
		//fmt.Println("Error!")
		fmt.Println(err)
	}
	defer jsonFile.Close()
	//instance := make([]Instance, 3)
	var instance Instance
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &instance)
	return &instance
}

// func getData(fileName string) []Instance {
// 	fileName = strings.TrimRight(fileName, "\r\n")

// 	instance := make([]Instance, 3)

// 	dat, err := ioutil.ReadFile(fileName)
// 	check(err)
// 	fmt.Print(string(dat))

// 	json.Unmarshal(dat, instance)
// 	return instance

// }

// func Abs(numb int) int {
// 	if numb > 0 {
// 		return numb
// 	} else {
// 		return -numb
// 	}
// }

// func processData(path string, queue []Instance) []Instance {
// 	product := getData(path)

// 	fmt.Println(product)

// 	if len(product) == 0 {
// 		for i := range product {
// 			queue = append(queue, product.Instance[i].Type)
// 			queue[i].Action = ""
// 		}
// 	} else {
// 		for i := range product {
// 			temp1 := product[i]

// 			temp := queue[0]

// 			queue = queue[1:]

// 			if temp.Counts < temp1.Counts {
// 				temp.Action = "provision"
// 			} else {
// 				temp.Action = "delete"
// 			}
// 			temp = temp1
// 			temp.Amount = Abs(temp1.Counts - temp.Counts)
// 			queue = append(queue, temp)

// 		}
// 	}
// }

func main() {

	var check bool

	check = false
	var oldInstance *Instance
	var newInstance *Instance

	// var quere []Instance
	// var quere2 []Instance

	reader := bufio.NewReader(os.Stdin)
	for exit := 1; exit != 2; {
		fmt.Print("Enter path: ")
		path, _ := reader.ReadString('\n')
		if isExit(path) {
			exit = 2
			break
		} else {
			if isJsonFile(path) {
				//readData(path)
				if check {
					oldInstance = newInstance
					newInstance = readData(path)

					for i := 0; i < len(newInstance.Instance); i++ {
						temp := newInstance.Instance[i].Counts - oldInstance.Instance[i].Counts
						if temp >= 0 {
							fmt.Println("[\""+oldInstance.Instance[i].Type+"\"]"+" [provision] [", temp, "]")
						} else {
							fmt.Println("["+oldInstance.Instance[i].Type+"]"+" [delete] [", -temp, "]")
						}
					}

					// quere = processData(path, quere)
					// for len(queue) > 0 {
					// 	temp := queue[0]
					// 	queue = queue[1:] // Dequeue
					// 	fmt.Println(temp.Type, " ", temp.Action, " ", temp.Counts)
					// 	fmt.Println("===============")
					// 	queue2 = append(queue2, temp)
					// }

					// fmt.Println()
					// queue = queue2
					// for len(queue2) > 0 {
					// 	queue2 = queue2[1:] // Dequeue
					// }
				} else {
					newInstance = readData(path)
					check = true
				}

			} else {
				fmt.Println("Invalid file! Enter path again: ")
			}
		}
	}

}
