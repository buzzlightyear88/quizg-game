package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
	// make a Problems buffer
	buf := make([]Problem, PROBLEMCOUNT)
	// open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// scan the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// read comma separate lines
		line := strings.Split(scanner.Text(), ",")
		// fist value is the question
		q := line[0]
		// convert string answer to integer
		ans, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		// store problem in buffer
		buf[count] = Problem{q, ans}
		count++
	}
	// shuffle the problems
	if shuffle {
		rand.Seed(time.Now().Unix())
		for i := range buf[:count] {
			j := rand.Intn(i + 1)
			buf[i], buf[j] = buf[j], buf[i]
		}
	}
	// send problems over the channel
	for _, p := range buf[:count] {
		problems <- p
	}
}

//open a channel and consume the problems
func solveProblems(problems chan Problem) {
	// create an IO reader
	scanner := bufio.NewScanner(os.Stdin)
	// start consuming problems from the channel
	for p := range problems {
		// print problem question
		fmt.Printf("%s = ", p.q)
		// scan IO input
		scanner.Scan()
		input := strings.Trim(scanner.Text(), " ") // remove white space
		// convert answer to int
		ans, err := strconv.Atoi(input)
		if err != nil {
			wrong++
			fmt.Printf("'%s' is not a valid asnwer\n", input)
			continue
		} else if ans == p.a {
			correct++
			fmt.Printf("Correct!\n")
		} else {
			wrong++
			fmt.Printf("Wrong!\n")
		}
	}
	return
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
	flag.Parse()
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
	fmt.Printf("Correct answers: %d\n", correct)
	fmt.Printf("Incorrect answers: %d\n", wrong)

}
