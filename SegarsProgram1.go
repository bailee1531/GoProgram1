// Bailee Segars
// SegarsProgram1.go
// CS 424-01 Fall 2024
// This program reads a file with player stats and generates a report
// Used VSCode 1.93.1 on Windows 11
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

type Player struct {
	firstName  string
	lastName   string
	batAvg     float64
	sluggPerc  float64
	onBasePerc float64
}

type ErrorLine struct {
	firstName    string
	lastName     string
	errorType    string
	errorLineNum int
}

func CalcBattingAvg(hits, atBats float64) float64 {
	batAvg := hits / atBats
	return batAvg
}

func CalcSlugg(singles, doubles, triples, homeRuns, atBats float64) float64 {
	sluggPerc := ((singles) + (2 * doubles) + (3 * triples) + (4 * homeRuns)) / atBats
	return sluggPerc
}

func CalcOBP(hits, walks, hitByPitch, plateAppear float64) float64 {
	obp := (hits + walks + hitByPitch) / plateAppear
	return obp
}

func CheckData(stats []string) bool {
	var checkData = regexp.MustCompile("^[0-9]+$")
	var isValid bool
	for i := 0; i < len(stats); i++ {
		if checkData.MatchString(stats[i]) {
			isValid = true
			continue
		} else {
			isValid = false
			break
		}
	}
	return isValid
}

func OutputPlayers(playerList Player) {
	fmt.Printf("%s,\t", playerList.lastName)
	fmt.Printf("%s\t\t", playerList.firstName)
	fmt.Printf("%.3f\t\t", playerList.batAvg)
	fmt.Printf("%.3f\t\t", playerList.sluggPerc)
	fmt.Printf("%.3f\t", playerList.onBasePerc)
	fmt.Printf("\n")
}

func OutputErrors(errorList ErrorLine) {
	fmt.Printf("Line %d:\t", errorList.errorLineNum)
	fmt.Printf("%s, ", errorList.lastName)
	fmt.Printf("%s: ", errorList.firstName)
	fmt.Printf("%s\t", errorList.errorType)
	fmt.Printf("\n")
}

func main() {
	var keyboard *bufio.Scanner
	var fileLines *bufio.Scanner
	var numPlayers int
	var lineNum int
	keyboard = bufio.NewScanner(os.Stdin)
	numPlayers = 0
	lineNum = 0
	var playerList []*Player
	var errorList []*ErrorLine

	fmt.Println("Enter your file name")
	keyboard.Scan()
	fileName := keyboard.Text()

	inFile, estat := os.Open(fileName)
	if estat != nil {
		fmt.Println("\nCannot open the file named " + fileName)
		fmt.Println("Exiting program.")
		os.Exit(1)
	} else {
		fmt.Println("\n" + fileName + " was opened successfully\n")
	}

	fileLines = bufio.NewScanner(inFile)

	for fileLines.Scan() {
		stats := strings.Split(fileLines.Text(), " ")
		lineNum++
		var currPlayer Player
		var error ErrorLine
		if len(stats) == 10 && CheckData(stats[2:]) {
			numPlayers++

			currPlayer.firstName = stats[0]
			currPlayer.lastName = stats[1]

			plateAppear, _ := strconv.ParseFloat(strings.TrimSpace(stats[2]), 64)
			atBats, _ := strconv.ParseFloat(strings.TrimSpace(stats[3]), 64)
			singles, _ := strconv.ParseFloat(strings.TrimSpace(stats[4]), 64)
			doubles, _ := strconv.ParseFloat(strings.TrimSpace(stats[5]), 64)
			triples, _ := strconv.ParseFloat(strings.TrimSpace(stats[6]), 64)
			homeRuns, _ := strconv.ParseFloat(strings.TrimSpace(stats[7]), 64)
			walks, _ := strconv.ParseFloat(strings.TrimSpace(stats[8]), 64)
			hitByPitch, _ := strconv.ParseFloat(strings.TrimSpace(stats[9]), 64)

			hits := singles + doubles + triples + homeRuns
			batAvg := CalcBattingAvg(hits, atBats)
			sluggPerc := CalcSlugg(singles, doubles, triples, homeRuns, atBats)
			obp := CalcOBP(hits, walks, hitByPitch, plateAppear)

			currPlayer.batAvg = batAvg
			currPlayer.sluggPerc = sluggPerc
			currPlayer.onBasePerc = obp
			playerList = append(playerList, &currPlayer)

		} else if len(stats) < 10 {
			error.firstName = stats[0]
			error.lastName = stats[1]
			error.errorType = "Line contains not enough data"
			error.errorLineNum = lineNum
			errorList = append(errorList, &error)

		} else if !CheckData(stats) {
			error.firstName = stats[0]
			error.lastName = stats[1]
			error.errorType = "Line contains invalid data"
			error.errorLineNum = lineNum
			errorList = append(errorList, &error)
		}
	}

	sort.Slice(playerList, func(i, j int) bool {
		return playerList[i].sluggPerc > playerList[j].sluggPerc
	})

	sort.Slice(errorList, func(i, j int) bool {
		return errorList[i].errorLineNum < errorList[j].errorLineNum
	})

	fmt.Printf("BASEBALL STATS REPORT -------- %d PLAYERS FOUND IN FILE\n\n", numPlayers)
	fmt.Printf("PLAYER NAME:\t\tAVERAGE\t\tSLUGGING\tONBASE%%\n")
	for i := 0; i < len(playerList); i++ {
		OutputPlayers(*playerList[i])
	}

	fmt.Printf("\n")

	fmt.Printf("ERROR LINES FOUND IN INPUT DATA\n")
	fmt.Println("-------------------------------")
	for i := 0; i < len(errorList); i++ {
		OutputErrors(*errorList[i])
	}

	defer inFile.Close()
}
