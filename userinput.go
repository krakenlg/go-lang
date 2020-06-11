package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your full name: ")
	// myname, _ := reader.ReadString('\n')
	// fmt.Println(myname)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your rating: ")
	myrating, _ := reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(myrating), 64)
	fmt.Print(mynumRating + 2)
}
