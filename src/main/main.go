///////////////////////////////////////////////////////
//
// Fabian Kirberg
//
//////////////////////////////////////////////////////

package main

import (
	"connect4"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Default board values
var rows = 6
var cols = 7
var win_length = 4

func main() {
	// If there is atleast 1 value given, use it as board size
	if len(os.Args) >= 2 {
		var regex = regexp.MustCompile("^([0-9]+)x([0-9]+)$")
		var match = regex.MatchString(os.Args[1])
		if match {
			var splitString = strings.Split(os.Args[1], "x")
			if rowNum, err := strconv.Atoi(splitString[0]); err == nil {
				rows = rowNum
			}
			if colNum, err := strconv.Atoi(splitString[1]); err == nil {
				cols = colNum
			}
		} else {
			fmt.Println("Board size " + os.Args[1] + " is not formatted properly.")
			os.Exit(0)
		}
	}

	//If there are atleast 2 values given, use them as board size and win length
	if len(os.Args) >= 3 {
		var regex = regexp.MustCompile("^([0-9]+)$")
		var match = regex.MatchString(os.Args[2])
		if match {
			if num, err := strconv.Atoi(os.Args[2]); err == nil {
				win_length = num
			}
		} else {
			fmt.Println("Win length must be an integer.")
			os.Exit(0)
		}
	}

	//Create the game object and play the game
	c4 := connect4.New(rows, cols, win_length)
	c4.PlayGame()
}
