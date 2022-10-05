package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func nextfile(lastfile string) string {
	splited := strings.Split(lastfile, "/")
	for i := 1; i < 6; i++ {
		if x, _ := strconv.Atoi(splited[3]); x == 5 {
			splited[3] = "1.txt"
			y, _ := strconv.Atoi(splited[2])
			splited[2] = strconv.Itoa(y + 1)
		} else {
			y, _ := strconv.Atoi(splited[3])
			splited[2] = strconv.Itoa(y+1) + ".txt"
		}
	}
	return splited[0] + "/" + splited[1] + "/" + splited[2] + "/" + splited[3]
}

func gessit(a []int, file string) (nb int, filefind string) {
	text, err := os.ReadFile(file)
	data := string(text)
	table := strings.Split(data, "\n")
	if err != nil {
		gessit(a, nextfile(file))
	} else {
		for i := 0; i < len(a); i++ {
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
	return gessit(a, nextfile(file))
}

func main() {
	var a []int
	file := "./data_sets/1/1.txt"
	for i := 0; i < 12500; i++ {
		var b int
		fmt.Scanf("%d", &b)
		a = append(a, b)
		b, file = gessit(a, file)
		fmt.Println(b, b)
	}
}
