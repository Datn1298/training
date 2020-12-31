package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
// 	"bufio"
// 	"os"
// 	"strings"
)

// const path = "./config.json"
// const path1 = "./config_1.json"
// const path2 = "./config_2.json"
// const path3 = "./config_3.json"
// const path4 = "./config_4.json"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Data struct {
	Type   string `json:"type"`
	VCPU   int    `json:"vCPU"`
	VRam   int    `json:"vRAM"`
	Counts int    `json:"counts"`
	Action string

}

type Output struct {
	Type string
	Action string
	Counts int
}

func getProduct(path string) []Data {
	
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}

	data2 := data[16: len(data)-1]

	product := make([]Data, 3)
	json.Unmarshal(data2, &product)
	return product
}
 
func abc(path string, queue []Data) []Data {

	products := getProduct(path)
		
	if len(queue) == 0 {
		for i := range products {
			queue = append(queue, products[i])
			queue[i].Action = "a"
		}	
	} else {
		for i := range products {
			temp1 := products[i]

			temp := queue[0]

			queue = queue[1:]

			if(temp.Counts < temp1.Counts){
				temp = temp1
				temp.Action = "provision"
				queue = append(queue, temp)
			} else {
				temp = temp1
				temp.Action = "delete"
				queue = append(queue, temp)
			}
		}	
	}
	return queue
}


func main()  {
	var queue []Data
	var queue2 []Data
	var i int
	var path string

	for i=0; i<=8; i++{
		
		if i== 0 {
			path = "./config.json"
		} else {
			path = "./config_" + string(i+48) + ".json"
		}

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

	}

	// reader := bufio.NewReader(os.Stdin)
	// for ok := 1; ok != 2; {
	// 	fmt.Print("Enter path: ")
	// 	path, _ := reader.ReadString('\n')

	// 	exe := path[len(path)-7: len(path)]
	// 	exe = strings.TrimRight(exe, "\r\n")

	// 	if exe == ".json" {
	// 		queue = abc(path, queue)		
	// 		for len(queue) > 0 {
	// 			temp := queue[0]
	// 			queue = queue[1:]   // Dequeue
	// 			fmt.Println(temp.Type , " " , temp.Action , " " , temp.Counts)
	// 			fmt.Println("===============")
	// 			queue2 = append(queue2, temp)
	// 		}

	// 		queue = queue2
	// 	} else {
	// 		fmt.Println("Invalid file! Enter path again: ")
	// 	}
	// }

	// queue = abc(path, queue)

	// for len(queue) > 0 {
	// 	temp := queue[0]
	// 	queue = queue[1:]   // Dequeue
	// 	fmt.Println(temp.Type , " " , temp.Action , " " , temp.Counts)
	// 	fmt.Println("===============")
	// 	queue2 = append(queue2, temp)
	// }
	// fmt.Println()

	// queue = queue2

	// for len(queue2) > 0 {
	// 	queue2 = queue2[1:]   // Dequeue
	// }

	// queue = abc(path1, queue)

	// for len(queue) > 0 {
	// 	temp := queue[0]
	// 	queue = queue[1:]   // Dequeue
	// 	fmt.Println(temp.Type , " " , temp.Action , " " , temp.Counts)
	// 	fmt.Println("===============")
	// 	queue2 = append(queue2, temp)
	// }
	// fmt.Println()

	// queue = queue2

	// for len(queue2) > 0 {
	// 	queue2 = queue2[1:]   // Dequeue
	// }

	// queue = abc(path2, queue)

	// for len(queue) > 0 {
	// 	temp := queue[0]
	// 	queue = queue[1:]   // Dequeue
	// 	fmt.Println(temp.Type , " " , temp.Action , " " , temp.Counts)
	// 	fmt.Println("===============")
	// 	queue2 = append(queue2, temp)
	// }
	// fmt.Println()

	// queue = queue2

	// for len(queue2) > 0 {
	// 	queue2 = queue2[1:]   // Dequeue
	// }

	// queue = abc(path3, queue)

	// for len(queue) > 0 {
	// 	temp := queue[0]
	// 	queue = queue[1:]   // Dequeue
	// 	fmt.Println(temp.Type , " " , temp.Action , " " , temp.Counts)
	// 	fmt.Println("===============")
	// 	queue2 = append(queue2, temp)
	// }
	// fmt.Println()

	// queue = queue2

	// for len(queue2) > 0 {
	// 	queue2 = queue2[1:]   // Dequeue
	// }

	// queue = abc(path4, queue)

	// for len(queue) > 0 {
	// 	temp := queue[0]
	// 	queue = queue[1:]   // Dequeue
	// 	fmt.Println(temp.Type , " " , temp.Action , " " , temp.Counts)
	// 	fmt.Println("===============")
	// 	queue2 = append(queue2, temp)
	// }
	// fmt.Println()

	// queue = queue2

	// for len(queue2) > 0 {
	// 	queue2 = queue2[1:]   // Dequeue
	// }
}


