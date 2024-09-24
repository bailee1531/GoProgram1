/*
Bailee Segars
SegarsProgram1.go
CS 424-01 Fall 2024
This program reads a file with player stats and generates a report
Used VSCode 1.93.1 on Windows 11
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Holds the calculated stats to be printed for each player
type Player struct {
	firstName  string  //Player's first name
	lastName   string  //Player's last name
	batAvg     float64 //Player's calculated batting average
	sluggPerc  float64 //Player's calculated slugging percentage
	onBasePerc float64 //Player's calculated on base percentage
}

// Holds the info about each error to be printed
type ErrorLine struct {
	firstName    string //First name associated with error line
	lastName     string //Last name associated with error line
	errorType    string //Type of error (not enough data, invalid data)
	errorLineNum int    //Line number of error
}

/*
This function calculates the batting average
Parameters: hits, atBats
Return value: Calculated batting average
CalcBattingAvg calculates the batting average by dividing the given parameters
*/
func CalcBattingAvg(hits, atBats float64) float64 {
	batAvg := hits / atBats //Batting average - to be assigned to player
	return batAvg
}

/*
This function calculates the slugging percentage
Parameters: singles, doubles, triples, homeRuns, atBats
Return value: Calculated slugging percentage
CalcSlugg calculates the batting average by averaging hits and atBats
*/
func CalcSlugg(singles, doubles, triples, homeRuns, atBats float64) float64 {
	//Slugging percentage - to be assigned to player
	sluggPerc := ((singles) + (2 * doubles) + (3 * triples) + (4 * homeRuns)) / atBats
	return sluggPerc
}

/*
This function calculates on base percentage
Parameters: hits, walks, hitByPitch, plateAppear
Return value: Calculated on base percentage
CalcOBP calculates the OBP by adding hits, walks, hitByPitch and dividing the sum by plateAppearances
*/
func CalcOBP(hits, walks, hitByPitch, plateAppear float64) float64 {
	obp := (hits + walks + hitByPitch) / plateAppear //On base percentage - to be assigned to player
	return obp
}

/*
This function determines if player data is valid
Parameter: String array of single player stats
Return value: Boolean that indicates if data is valid
CheckData uses a regular expression to check if the given data is only numerical
*/
func CheckData(stats []string) bool {
	var checkData = regexp.MustCompile("^[0-9]+$") //Variable to use when checking if regex matches
	var isValid bool                               //Boolean to return after determining if the numerical data is valid
	for i := 0; i < len(stats); i++ {
		if checkData.MatchString(stats[i]) { //Continue until reaches the end of stats[] or until false
			isValid = true
			continue
		} else {
			isValid = false
			break
		}
	}
	return isValid
}

/*
This function prints the formatted player stats
Parameters: Instance of Player struct named playerList
Return value: None
OutputPlayer prints formatted data stored in the struct variables
*/
func OutputPlayers(playerList Player) {
	fmt.Printf("%s,\t", playerList.lastName)
	fmt.Printf("%s\t\t", playerList.firstName)
	fmt.Printf("%.3f\t\t", playerList.batAvg)
	fmt.Printf("%.3f\t\t", playerList.sluggPerc)
	fmt.Printf("%.3f\t", playerList.onBasePerc)
	fmt.Printf("\n")
}

/*
This function prints the formatted error information
Parameters: Instance of Error struct named errorList
Return value: None
OutputErrors prints formatted data stored in the struct variables
*/
func OutputErrors(errorList ErrorLine) {
	fmt.Printf("Line %d:\t", errorList.errorLineNum)
	fmt.Printf("%s, ", errorList.lastName)
	fmt.Printf("%s: ", errorList.firstName)
	fmt.Printf("%s\t", errorList.errorType)
	fmt.Printf("\n")
}

func main() {
	var keyboard *bufio.Scanner  //Variable to hold file name that the user inputs
	var fileLines *bufio.Scanner //Variable to scan data from file
	keyboard = bufio.NewScanner(os.Stdin)
	numPlayers := 0 //Variable to indicate the number of players with no errors
	lineNum := 0    //Variable to indicate the line number. Used when printing error info

	//Does not need to be a pointer because values are not being passed back to struct
	var playerList []Player   //Slice of player structs so it can be dynamically sized
	var errorList []ErrorLine //Slice of error structs so it can be dynamically sized

	fmt.Println("Enter your file name")
	keyboard.Scan()
	fileName := keyboard.Text()

	inFile, estat := os.Open(fileName) //Variables to point to file and error
	if estat != nil {
		fmt.Println("\nCannot open the file named " + fileName)
		fmt.Println("Exiting program.")
		os.Exit(1)
	} else {
		fmt.Println("\n" + fileName + " was opened successfully\n\n")
	}

	fileLines = bufio.NewScanner(inFile) //Prepares the scanner to read the file

	for fileLines.Scan() { //Executes until EOF has been reached
		//Splits each line into an array of strings and stores it in stats array
		stats := strings.Split(fileLines.Text(), " ")
		lineNum++
		var currPlayer Player //Creates an object of Player struct
		var error ErrorLine   //Creates an object of ErrorLine struct

		//Only executes for players that have the correct amount of data that is all valid
		if len(stats) == 10 && CheckData(stats[2:]) { //Only checks data that comes after the names
			numPlayers++

			currPlayer.firstName = stats[0] //Sets firstName variable to first element scanned
			currPlayer.lastName = stats[1]  //Sets lastName variable to second element scanned

			plateAppear, _ := strconv.ParseFloat(strings.TrimSpace(stats[2]), 64) //Converts numerical string to float to be used in calulations
			atBats, _ := strconv.ParseFloat(strings.TrimSpace(stats[3]), 64)
			singles, _ := strconv.ParseFloat(strings.TrimSpace(stats[4]), 64)
			doubles, _ := strconv.ParseFloat(strings.TrimSpace(stats[5]), 64)
			triples, _ := strconv.ParseFloat(strings.TrimSpace(stats[6]), 64)
			homeRuns, _ := strconv.ParseFloat(strings.TrimSpace(stats[7]), 64)
			walks, _ := strconv.ParseFloat(strings.TrimSpace(stats[8]), 64)
			hitByPitch, _ := strconv.ParseFloat(strings.TrimSpace(stats[9]), 64)

			hits := singles + doubles + triples + homeRuns                      //Adds hits up and stores in hits variable to use as a parameter
			batAvg := CalcBattingAvg(hits, atBats)                              //Holds calculated batting average for that player
			sluggPerc := CalcSlugg(singles, doubles, triples, homeRuns, atBats) //Holds calculated slugging percentage
			obp := CalcOBP(hits, walks, hitByPitch, plateAppear)                //Holds calculated on base percentage for that player

			currPlayer.batAvg = batAvg                  //Sets struct batAvg to calculated batAvg
			currPlayer.sluggPerc = sluggPerc            //Sets struct sluggPerc to calculated sluggPerc
			currPlayer.onBasePerc = obp                 //Sets struct onBasePerc to calculated obp
			playerList = append(playerList, currPlayer) //Appends these values to the Player struct in the playerList slice

		} else if len(stats) < 10 { //Only executes if there is not enough data
			error.firstName = stats[0]                        //Sets firstName error variable to first element in file
			error.lastName = stats[1]                         //Sets lastName error variable to second element in file
			error.errorType = "Line contains not enough data" //Sets the error type string
			error.errorLineNum = lineNum                      //Sets lineNum error variable to the current line number
			errorList = append(errorList, error)              //Appends these values to the ErrorLine struct in the errorList slice

		} else if !CheckData(stats) { //Only executes if there is invalid data found
			error.firstName = stats[0]                     //Sets firstName error variable to the first element in file
			error.lastName = stats[1]                      //Sets lastName error variable to the second element in file
			error.errorType = "Line contains invalid data" //Sets the error type string
			error.errorLineNum = lineNum                   //Sets lineNum error variable to the current line number
			errorList = append(errorList, error)           //Appends these values to the ErrorLine struct in the errorList slice
		}
	}

	sort.Slice(playerList, func(i, j int) bool { //Sorts the slice based on slugging percentage
		return playerList[i].sluggPerc > playerList[j].sluggPerc //Best slugging percentage is at the front
	})

	sort.Slice(errorList, func(i, j int) bool { //Sorts the slice based on the error line number
		return errorList[i].errorLineNum < errorList[j].errorLineNum //Lowest (first encountered) line number is at the front
	})

	fmt.Printf("BASEBALL STATS REPORT -------- %d PLAYERS FOUND IN FILE\n\n", numPlayers)
	fmt.Printf("PLAYER NAME:\t\tAVERAGE\t\tSLUGGING\tONBASE%%\n")
	for i := 0; i < len(playerList); i++ { //Executes for the length of the playerList slice
		OutputPlayers(playerList[i]) //Uses struct at slice element i as parameter
	}

	fmt.Printf("\n")

	fmt.Printf("ERROR LINES FOUND IN INPUT DATA\n")
	fmt.Println("-------------------------------")
	for i := 0; i < len(errorList); i++ { //Executes for the length of the errorList slice
		OutputErrors(errorList[i]) //Uses struct at slice element i as parameter
	}

	defer inFile.Close() //Closes the file
}
