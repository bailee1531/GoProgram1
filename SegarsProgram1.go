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
	"strconv"
	"strings"
)

type Player struct {
	firstName   string
	lastName    string
	singles     float64
	doubles     float64
	triples     float64
	homeRuns    float64
	atBats      float64
	walks       float64
	hitByPitch  float64
	plateAppear float64
}

func CalcBattingAvg(hits, atBats float64) float64 {
	batAvg := (hits) / atBats
	return batAvg
}

func CalcSlugg(singles, doubles, triples, homeRuns, atBats float64) float64 {
	sluggPerc := ((singles + 2.0) * (doubles + 3) * (triples + 4) * homeRuns) / atBats
	return sluggPerc
}

func CalcOBP(hits, walks, hitByPitch, plateAppear float64) float64 {
	obp := (hits + walks + hitByPitch) / plateAppear
	return obp
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

	fmt.Println("Enter your file name")
	keyboard.Scan()
	fileName := keyboard.Text()

	inFile, estat := os.Open(fileName)
	if estat != nil {
		fmt.Println("\nCannot open the file named " + fileName)
		fmt.Println("Exiting program.")
		os.Exit(1)
	} else {
		fmt.Println("\n" + fileName + " was opened successfully")
	}

	fileLines = bufio.NewScanner(inFile)
	stats := strings.Split(keyboard.Text(), " ")

	for fileLines.Scan() {
		lineNum++
		var currPlayer Player
		if len(stats) == 9 {
			currPlayer.firstName = stats[0]
			currPlayer.lastName = stats[1]
			singles := currPlayer.singles
			singles, _ = strconv.ParseFloat(strings.TrimSpace(stats[2]), 64)
			doubles := currPlayer.doubles
			doubles, _ = strconv.ParseFloat(strings.TrimSpace(stats[3]), 64)
			triples := currPlayer.triples
			triples, _ = strconv.ParseFloat(strings.TrimSpace(stats[4]), 64)
			homeRuns := currPlayer.homeRuns
			homeRuns, _ = strconv.ParseFloat(strings.TrimSpace(stats[5]), 64)
			atBats := currPlayer.atBats
			atBats, _ = strconv.ParseFloat(strings.TrimSpace(stats[6]), 64)
			walks := currPlayer.walks
			walks, _ = strconv.ParseFloat(strings.TrimSpace(stats[7]), 64)
			hitByPitch := currPlayer.hitByPitch
			hitByPitch, _ = strconv.ParseFloat(strings.TrimSpace(stats[8]), 64)
			plateAppear := currPlayer.plateAppear
			plateAppear, _ = strconv.ParseFloat(strings.TrimSpace(stats[9]), 64)

			hits := singles + doubles + triples + homeRuns
			batAvg := CalcBattingAvg(hits, atBats)
			sluggPerc := CalcSlugg(singles, doubles, triples, homeRuns, atBats)
			obp := CalcOBP(hits, walks, hitByPitch, plateAppear)
			playerList = append(playerList, &currPlayer)
			fmt.Printf("%s\t%s:\t%f\t%f\t%f", stats[0], stats[1], batAvg, sluggPerc, obp)
		} else {
			currPlayer.lastName = stats[0]
			currPlayer.firstName = stats[1]
			stats[2] = "Line contains not enough data"
			playerList = append(playerList, &currPlayer)
			fmt.Printf("Line %d: %s,\t%s: %s\n", lineNum, stats[0], stats[1], stats[2])
		}

	}

	fmt.Printf("BASEBALL STATS REPORT -------- %d PLAYERS FOUND", numPlayers)
	fmt.Println("ERROR LINES FOUND IN INPUT DATA")
	fmt.Println("----------------------------")

	defer inFile.Close()
}
