package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
)

const TotalRounds = 5
const AcceptableErrorThreshold = 10
const MinRoundsToWin = 4

func Start() {
	var language string
	var repositories SearchResponse
	var statusCode int

	var roundsWon int

	printInstructions()

	for {
		language = askLanguageInput()
		repositories, statusCode = SearchTrendingRepositories(language)
		if statusCode == 422 {
			fmt.Println("Sorry, Github doesn't recognize this language. Try Python, C, Java, Assembly...")
		} else if statusCode == 200 {
			break
		} else {
			log.Fatalf("Github API failed with Status Code - %v", statusCode)
		}
	}

	for round := 1; round <= TotalRounds; round++ {
		roundsWon = roundsWon + playRound(round, TotalRounds, repositories.Items, AcceptableErrorThreshold)
	}

	displayResult(roundsWon, TotalRounds, MinRoundsToWin)
}

func printInstructions() {
	fmt.Println("***** Welcome to Github's Guess the Stars *****")
	fmt.Println("This is a CLI game in which given a trending public repository on Github, you have to guess the number of stars it will have.")
	fmt.Println("Instructions:")
	fmt.Println("\t- You will be optionally asked to choose a language")
	fmt.Println("\t- There will be 5 rounds in total")
	fmt.Println("\t- Each round presents you with a new trending repository")
	fmt.Println("\t- You win the round by guessing the stars within 10% tolerance")
	fmt.Println("\t- You win the game if you win at least 4 of the 5 rounds")
	fmt.Println("")
}

func askLanguageInput() string {
	var language string
	fmt.Println("Enter a language of your choice (Press enter to skip):")
	_, err := fmt.Scanf("%s", &language)
	if (err != nil) && (err.Error() != "unexpected newline") {
		log.Fatal(err)
	}
	return language
}

func playRound(roundNumber, totalRounds int, repositories []RepositoryInfo, acceptableErrorThreshold float64) int {

	displayRoundHeader(roundNumber, totalRounds)

	randRepositoryIndex := randInt(len(repositories))
	repo := repositories[randRepositoryIndex]

	displayRepositoryInfo(repo)

	guessedStars := acceptStarsInput()
	roundWon := computeRoundResult(repo.StargazersCount, guessedStars, acceptableErrorThreshold)

	return roundWon
}

func displayRoundHeader(roundNumber, totalRounds int) {
	fmt.Println("**************")
	fmt.Printf("* Round %v/%v: *\n", roundNumber, totalRounds)
	fmt.Println("**************")
}

func displayRepositoryInfo(repo RepositoryInfo) {
	fmt.Printf("\tRepository name: %v  (%v)\n", repo.Name, repo.Language)
	fmt.Printf("\t%v\n", repo.Description)
	fmt.Printf("\tHint - number of forks: %v\n", repo.Forks)
}

func randInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func acceptStarsInput() int {
	var guessedStars int
	var starsInput string
	for {
		fmt.Printf("Guess the stars: ")
		_, err := fmt.Scan(&starsInput)
		guessedStars, err = strconv.Atoi(starsInput)
		if err != nil {
			fmt.Println("Number of stars should be a valid integer")
		} else {
			break
		}
	}
	return guessedStars
}

func computeRoundResult(actualStars, guessedStars int, acceptableErrorThreshold float64) int {
	var won int
	guessDeviation := math.Abs(float64(actualStars) - float64(guessedStars))
	deviationPercent := (guessDeviation / float64(actualStars)) * 100

	if deviationPercent <= acceptableErrorThreshold {
		won = 1
		fmt.Printf("Awesome guess! You have won this round. Actual stars for this repo - %v\n", actualStars)
	} else {
		won = 0
		fmt.Printf("Uh Oh. Better guess next time. Actual stars for this repo - %v\n", actualStars)
	}
	fmt.Println("")
	return won
}

func displayResult(roundsWon, totalRounds, minRoundsToWin int) {
	if roundsWon >= minRoundsToWin {
		fmt.Println("***** Congratulations! You have won Github's guess the stars *****")
	} else {
		fmt.Println("***** You Lose *****")
		fmt.Printf("You won %v/%v rounds. Win %v+ rounds to win the game\n", roundsWon, totalRounds, minRoundsToWin)
	}
}
