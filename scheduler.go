package main

import (
	"fmt"
	"nfl-scheduler/league"
)

func main() {
	fmt.Println("hello")
	nfl := league.SetUpLeague("NFL")
	_ = "breakpoint"
	fmt.Println(nfl)

}
