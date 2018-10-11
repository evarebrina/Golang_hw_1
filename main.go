package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var filename string
	var array []string
	var nums []int
	fmt.Scanf("%s", &filename)
	dat, err := ioutil.ReadFile(filename)
	check(err)
	var datstr = string(dat)
	array = strings.Split(datstr, " ")
	for _, i := range array {
		tmp, err := strconv.Atoi(i)
		check(err)
		nums = append(nums, tmp)
	}
	nums = sort(nums)
	fmt.Println(nums)
}

func sort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
	return arr
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
