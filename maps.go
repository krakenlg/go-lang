package main

import (
	"fmt"
)

func main() {

	score := make(map[string]int)
	score["lgk"] = 90
	fmt.Println(score)

	score["jo"] = 14
	score["jos"] = 24
	score["joh"] = 34
	score["josh"] = 44
	fmt.Println(score)

	getJosScore := score["jos"]
	fmt.Println(getJosScore)

	delete(score, "joh")

	fmt.Println(score)

	for k, v := range score {
		fmt.Printf("Score of %v is %v\n", k, v)
	}
}
