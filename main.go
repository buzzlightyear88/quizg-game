package main

import "time"

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
func parseProblems() {

}

//open a channel and consume the problems
func solveProblems(problems chan Problem) {

}

// startTime is a makeshift timer
func startTimer(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func main() {

}
