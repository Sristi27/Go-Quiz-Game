package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"rsc.io/quote"
)
func main(){
	fmt.Println(quote.Go())
	// fmt.Println("Hello world")
	//command line flag parsing
	filename := flag.String("csv","problems.csv","a csv file in the format of 'question,answer'")
	//filename will be a pointer to a string
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

	problems:=parseLines(lines)

	for i,p:=range problems{
		fmt.Printf("Problem #%d: %s = \n",i+1,p.q)
		var answer string
		fmt.Scanf("%s\n",&answer) //scanf removes all trailing spaces, it should not be used much /n->new line 
		//answer is what user gives
		//check if correct
		if(answer==p.a){
			fmt.Println("Correct")
		}else{
			fmt.Println("Wrong")
		}
	}
	
	// fmt.Println(problems)
}

func parseLines(lines[][] string)[]quiz{
	ret:=make([]quiz,len(lines))
	for i,line:=range lines{
		ret[i] = quiz{
			q:line[0], //ques
			a:line[1], //answer
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