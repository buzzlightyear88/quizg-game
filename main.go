package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Problem struct {
	q string
	a int
}

// set max number of problems in the buffer
const PROBLEMCOUNT = 100

// default time limit
const TIMELIMIT = 10

// Problem counter
var count int

// Score counter
var correct int
var wrong int

// parseProblems reads the problems from the file, line by line
// sends problems into a channel
func parseProblems(problems chan Problem, filename string, shuffle bool) {

}

//open a channel and consume the problems
func solveProblems(problems chan Problem) {

}

// startTime is a makeshift timer
func startTimer(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func main() {
	// parse command line flags
	filenameFl := flag.String("f", "problems.csv", "name of file")
	secondsFl := flag.Int("t", TIMELIMIT, "time to solve the quiz")
	shuffleFl := flag.Bool("s", false, "shuffle or not the questions")
	debugFl := flag.Bool("debug", false, "show debug information")

	// discard debug data if not wanted
	if !*debugFl {
		log.SetOutput(ioutil.Discard)
	}
	// show logging information if flag is called
	log.Println("debug:", *debugFl)
	log.Println("filename:", *filenameFl)
	log.Println("timer:", *secondsFl)
	log.Println("shuffle: ", *shuffleFl)

	// make a Problems channel
	problems := make(chan Problem, PROBLEMCOUNT)

	// read Problems goroutine
	go parseProblems(problems, *filenameFl, *shuffleFl)

	// prompt to start the quiz
	fmt.Printf("Press any key to start the quiz!")
	bufio.NewScanner(os.Stdin).Scan()

	// solve Problems goroutine
	go solveProblems(problems)

	// start the timer
	startTimer(*secondsFl)

	// show final tally
	fmt.Printf("\nNumber of questions: %d\n", count)
	fmt.Printf("Correct answers: %d", correct)
	fmt.Printf("Incorrect answers: %d", wrong)

}
