// Package main is a exersise solution written by me for quiz 1, URL: https://gophercises.com/exercises/quiz
package main;

import (
	"fmt";
	"os";
	"encoding/csv";
	"log";
	"bufio";  // For reading input from reader, in the end it was better than others, lol
	"math/rand";
	"strconv";
	"time";
	"flag";
)

// myFatal() just calls log.Fatalln(), i was just tierd of writing it

func myFatal(e error) {
	log.Fatalln("An unrecovarable error, ",e);
}

// fileReading reads from the file, and returns the csv.Reader + os.File object for the_game function to launch the quiz
func fileReading(name string) (*csv.Reader, *os.File){

	// Opening the file as requested by the user
	file, err := os.Open(name);
	if err != nil {
		myFatal(err);
	}
	

	// Using encoding csv pakcage to decode the CSV data
	reader := csv.NewReader(file);

	
	return reader,file;
}

// the_game is the actual quiz game, takes csv.Reader, and os.File as arguments
func the_game(reader *csv.Reader,file *os.File) {
	// Ensuring the file is closed by using defer
	defer func(){
		file.Close();
		}();

	// Now actually reading the csv data
	data, err := reader.ReadAll();
	if err != nil {
		fmt.Println("Error while reading CSV data")
		myFatal(err);
	}

	// Holding question, users input and correct answer
	var question string;
	var correct_answer string;
	var user_answer []byte;

	// The bufio reader that reads user input
	input_reader := bufio.NewReader(os.Stdin);

	// The following structure will hold the users result, how many correct and how many incorrect
	type result struct {
		correct []string
		incorrect []string
	};

	// The actual struct
	var r1 result;
	answer := make(chan string);
	timer := time.NewTimer(5 * time.Second);
	// Loop to get all the result
	for _,value := range data {
			question = value[0];
			correct_answer = value[1];
			
			fmt.Printf("Guess the answer to %s :",question);
			go func() {
				user_answer,_,_ = input_reader.ReadLine();
				answer <- string(user_answer);
			}();

			select {
				case _ = <- answer:
					if string(user_answer) == correct_answer {
						fmt.Printf("Congratulations,you are correct nibba")
						r1.correct = append(r1.correct,value[0]);
					} else {
						r1.incorrect = append(r1.incorrect,value[0]);
					}
				case <- timer.C:
					fmt.Println("Sorry time over");
						fmt.Println("\nThe correct answers you got: ",r1.correct);
						fmt.Println("the incorrect answer you got: ",r1.incorrect);	
					os.Exit(-1);

			}
		}
	
	fmt.Println("\nThe correct answers you got: ",r1.correct);
	fmt.Println("the incorrect answer you got: ",r1.incorrect);		
}
// argsHandling() handles the flag from user and do stuff
func argsHandling() (*csv.Reader,*os.File){

	// Getting the list of arguments
	
		
	if (len(os.Args) > 2) {
		if os.Args[1] == "-f" { // If flags were provided

		// Check if file exists or not
			_, err := os.Stat(os.Args[2]);
			if os.IsNotExist(err) {
				fmt.Println("The file you mentioned does not exist, using the default file");
				reader,file := fileReading("sample.csv")
				return reader,file
			} else {
				reader,file := fileReading(os.Args[2]);
				return reader,file
			}
		} else if os.Args[1] == "-r" {
			fmt.Println("You chose to use a randomly generated file, fine by me");
			random();

		} else {
		os.Exit(-1);
		}
	}
	reader,file := fileReading("sample.csv");
	return reader,file
}

// Random function creates random values and writes to csv file
func random() (*csv.Reader,*os.File) {

	file, err := os.OpenFile("random.txt",os.O_WRONLY | os.O_CREATE,0666)
	if err != nil {
		fmt.Println("An error while trying to create a random file")
		myFatal(err);
	}

	defer func(){
		file.Close();
		}();

	var r [][]string;
	var k1,k2 int;
	n := rand.Intn(5);
	for i := 0; i < n; i++ {
		k1 = rand.Intn(100);
		k2 = rand.Intn(100);
	
		r = append(r,[]string{strconv.Itoa(k1)+"+"+strconv.Itoa(k2),strconv.Itoa(k1 + k2)});
	}
	fmt.Println(r);

	csv_writer := csv.NewWriter(file);
	err2 := csv_writer.WriteAll(r);
	if err2 != nil {
		fmt.Println("Error while writing to a random txt")
		myFatal(err);
	} else {
		fmt.Println("Successfully created file")
	}
	reader,file2 := fileReading("random.txt");
	return reader,file2;
}

// main is the main function of the program
func main() {

	reader,file := argsHandling();
	fmt.Println("------------------------WELCOME TO QUIZ CHALLANGE----------------------------");
	the_game(reader,file);

	/* SHOULD HAVE USED FLAG FOr flAG PArsing, made some mistake
	file_name := flag.String("csv","random.txt","a csv file")
	flag.Parse();
	fmt.Println(*file_name);
	*/
}
