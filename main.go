package main

import (
	"fmt"
	"math"
	"sort"
)

func gessit(stat []int) (nb int, marge int) {
	moyen := 0.0
	Median := 0.0
	Variance := 0.0
	for _, v := range stat {
		moyen = moyen + float64(v)
	}
	sort.Ints(stat)
	if len(stat)%2 == 0 {
		if len(stat) > 2 {
			Median = float64((stat[0] + stat[1]) / 2)
		} else {
			Median = float64(stat[0])
		}
	}
	if len(stat)%2 == 1 {
		if len(stat) > 1 {
			Median = float64(stat[(len(stat)/2)+1])
		} else {
			Median = float64(stat[0])
		}
	}
	moyen = moyen / float64(len(stat))
	for _, v := range stat {
		Variance = Variance + ((float64(v) - moyen) * (float64(v) - moyen))
	}
	Variance = Variance / float64(len(stat))
	stdev := math.Sqrt(Variance)
	nb = stat[len(stat)-1] + int(Median) - int(moyen)
	if stdev < 25 {
		marge = int(stdev)
	} else {
		marge = 25
	}
	return nb, marge
}

func main() {
	var a []int
	for i := 0; i < 12500; i++ {
		var b int
		var nb int
		var marge int
		fmt.Scanf("%d", &b)
		a = append(a, b)
		nb, marge = gessit(a)
		fmt.Println(nb-marge, nb+marge)
	}
}
