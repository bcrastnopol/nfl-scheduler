package main
import (
	"fmt"
	"nfl-scheduler/libs"
)
func main() {
	fmt.Println("hello")
	Eagles := teams.NewTeam("Eagles", "East",  "NFC")
	fmt.Println(Eagles)
}