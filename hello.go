package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
	"rsc.io/quote"
)
func main(){
	fmt.Println(quote.Go())
	// fmt.Println("Hello world")
	//command line flag parsing
	filename := flag.String("csv","problems.csv","a csv file in the format of 'question,answer'")
	//filename will be a pointer to a string
	//timelimit
	timeLimit := flag.Int("limit",30,"the time limit for the quiz")

	flag.Parse()
	file,err := os.Open(*filename) //actual value usedd
	if(err!=nil){
		// fmt.Println("Failed to open a csv file:%s\n",*filename)
		// os.Exit(1)
		exit(fmt.Sprintf("Failed to open a csv file:%s\n",*filename))
	}
	// _=file
	//using the filname
	// io reader = most common interface
	r:=csv.NewReader(file) //creating a reader
	lines,err := r.ReadAll() //read all lines inside csv
	if(err!=nil){
		exit("Failed to parse the csv file")
	}

	//parse all the lines - ie setting up the entire csv file
	problems:=parseLines(lines)

	//after time expired send some msg to the user
	timer:=time.NewTimer(time.Duration(*timeLimit) * time.Second) //to make same type
	//<- timer.C  //waiting for a message from the channel - it will block the code once the message is received

	correct:=0
	for i,p:=range problems{
		fmt.Printf("Problem #%d: %s = \n",i+1,p.q)
		answerch:=make(chan string)
		go func(){
			var answer string
			fmt.Scanf("%s\n",&answer) 
			answerch<-answer //send answer to answer channel
		}()
		select{
		case <- timer.C:
			fmt.Printf("You scored %d out of %d. \n",correct,len(problems))
			return
		case answer:=<-answerch:
			if(answer==p.a){
				fmt.Println("Correct")
				correct++
			}else{
				fmt.Println("Wrong")
			}
		// default:
			//scanf removes all trailing spaces, it should not be used much /n->new line 
			//answer is what user gives
			//check if correct
			
		}
		
	}
	// fmt.Println(problems)
	//To print score
	fmt.Printf("You scored %d out of %d. \n",correct,len(problems))
}

func parseLines(lines[][] string)[]quiz{
	ret:=make([]quiz,len(lines))
	for i,line:=range lines{
		ret[i] = quiz{
			q:line[0], //ques
			a:strings.TrimSpace(line[1]), //answer
		} 
	}
	return ret
}
type quiz struct{
	q string
	a string
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}