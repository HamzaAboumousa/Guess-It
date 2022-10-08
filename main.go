package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func nextfile(lastfile string) string {
	splited := strings.Split(lastfile, "/")
	if len(splited) > 3 {
		if x, _ := strconv.Atoi(string(splited[3][0])); x == 5 {
			splited[3] = "1.txt"
			y, _ := strconv.Atoi(splited[2])
			splited[2] = strconv.Itoa(y + 1)
		} else {
			y, _ := strconv.Atoi(string(splited[3][0]))
			splited[3] = strconv.Itoa(y+1) + ".txt"
		}
		return splited[0] + "/" + splited[1] + "/" + splited[2] + "/" + splited[3]
	}
	return lastfile
}

func gessit(a []int, file string) (nb int, filefind string) {
	text, err := os.ReadFile(file)
	data := string(text)
	table := strings.Split(data, "\n")
	if err != nil {
		return a[len(a)-1], file
	} else {
		for i := 0; i < len(a); i++ {
			if len(table) > i {
				if x, _ := strconv.Atoi(table[i]); x == a[i] {
					nb, _ = strconv.Atoi(table[i+1])
					filefind = file
					if i == len(a)-1 {
						return
					}
				} else {
					break
				}
			}
		}
	}
	return gessit(a, file)
}

func main() {
	var a []int
	file := "./data_sets/1/1.txt"
	for i := 0; i < 12500; i++ {
		var b int
		fmt.Scanf("%d", &b)
		a = append(a, b)
		if len(a) < 3 {
			b, file = gessit(a, file)
		} else {
			text, _ := os.ReadFile(file)
			data := string(text)
			table := strings.Split(data, "\n")
			if len(table) > len(a) {
				b, _ = strconv.Atoi(table[len(a)])
			} else {
				b, _ = gessit(a, file)
			}
		}
		fmt.Println(b, b)
	}
}
