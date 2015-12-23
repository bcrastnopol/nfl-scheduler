package main

import (
	"fmt"
	"nfl-scheduler/league"
)

func main() {
	fmt.Println("hello")
	// c := Conference{'XFL', }
	// fmt.Println(Eagles)
	nfl := league.SetUpLeague()
	fmt.println(nfl)
}
