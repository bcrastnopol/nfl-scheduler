package main

import (
	"fmt"
	"nfl-scheduler/league"
)

func main() {
	fmt.Println("hello")
	var nfl league.Linterface = league.SetUpLeague("NFL")
	_ = "breakpoint"
	fmt.Println(nfl)

}
