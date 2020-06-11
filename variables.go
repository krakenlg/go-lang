package main

import "fmt"

func main() {
	var batman string = "I am a batman"
	fmt.Println(batman)

	var superman string
	superman = "I am a superman"
	fmt.Println(superman)

	thor := "I am thor"
	fmt.Println(thor)

	// thorRating := 3
	// fmt.Printf("%v, %T", thorRating,thorRating)

	var Ironman, CatAmerica string = "I am Ironman", "I am capt America"
	fmt.Println(Ironman, CatAmerica)

	var defaultValue int
	fmt.Println(defaultValue)

	var (
		spiderman = "I am a spider man"
		age       = 18
		powers    = "web slings, spidy sense"
		antman    = "I am antman"
	)

	fmt.Println(spiderman, age, powers, antman)
}
