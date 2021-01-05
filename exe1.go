package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type Instance struct {
	VCPU   float64 `json:"vCPU"`
	VRAM   float64 `json:"vRAM"`
	Counts float64 `json:"counts"`
}

// ktra input dau vao co phai file json
func isJsonFile(fileName string) bool {
	if strings.HasSuffix(fileName, ".json") {
		return true
	} else {
		return false
	}
}

// ktra input dau vao co phai exit
func isExit(fileName string) bool {
	if strings.HasPrefix(fileName, "Exit") {
		return true
	} else {
		return false
	}
}

// nhap input dau vao
func inputFilePath() string {
	var fileName string
	fmt.Print("Enter config path: ")
	fmt.Scanln(&fileName)

	if !isJsonFile(fileName) {
		fmt.Print("Invalid file, enter path: ")
		fmt.Scanln(&fileName)
	}

	return fileName
}

// doc file
func readFile(path string) ([]byte, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	return dat, err
}

func process() {

	var path string
	var result map[string]interface{}
	rs1 := make(map[string]Instance)

	check := true
	for check {
		path = inputFilePath()
		if isExit(path) {
			check = false
			break
		}
		data, _ := ioutil.ReadFile(path)
		json.Unmarshal(data, &result)
		if len(rs1) == 0 {
			for _, v := range result["Instances"].([]interface{}) {
				instance := v.(map[string]interface{})

				rs1[instance["type"].(string)] = Instance{
					instance["vCPU"].(float64),
					instance["vRam"].(float64),
					instance["counts"].(float64),
				}
			}

		} else {
			for _, v := range result["Instances"].([]interface{}) {
				instance := v.(map[string]interface{})
				type1 := instance["type"].(string)
				elem, ok := rs1[type1]
				if ok {
					if elem.Counts < instance["counts"].(float64) {
						fmt.Println("[", instance["type"], "] \t", " [ Provision ] \t ", "[", instance["counts"].(float64)-elem.Counts, "]")
					} else {
						fmt.Println("[", instance["type"], "] \t", " [ Delete ] \t ", "[", elem.Counts-instance["counts"].(float64), "]")
					}
					elem.Counts = instance["counts"].(float64)
				}
			}
		}
	}
}

func main() {
	process()
}
