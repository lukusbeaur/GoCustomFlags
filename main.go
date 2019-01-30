package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

//new string type
//not used
type argsList []string

//not used
var inputArg []string

//interface flag.value
func (s *argsList) String() string {
	return fmt.Sprintf("%v", *s)
}
func (s *argsList) Set(value string) error {
	*s = strings.Split(value, " ")
	return nil
}

//vars
var currentDir = getCurDir()

func main() {
	stringCommands := flag.NewFlagSet("Text", flag.ExitOnError)  //creates custom flags
	systeCommands := flag.NewFlagSet("System", flag.ExitOnError) //crrates custom flags

	parsText := stringCommands.String("far", "", "This string flag will hold your argument (Required)")
	curDirF := systeCommands.Bool("dir", false, "Returns your current DIR (Required)")

	//Choose which custom flag to parse
	switch os.Args[1] {
	case "Text":
		stringCommands.Parse(os.Args[2:])
	case "System":
		systeCommands.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	//checks to see if flag has been parsed.
	if stringCommands.Parsed() {
		if len(os.Args) < 2 {
			fmt.Print("Need one subcommand.")
			os.Exit(1)
		}
		if *parsText == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}
		flag.Parse()
		fmt.Printf("Your argument :  %v\n", *parsText)
		fmt.Printf("Number of Args: %v\n", 1+len(flag.Args()))
		fmt.Printf("Tail : %v", flag.Args())
	}
	//Only one command which requires no arguments, only the flag. If to add more, need to add len(os.ars <2) call error.
	if systeCommands.Parsed() {
		if len(os.Args) < 2 {
			fmt.Print("You need one subcommand.")
			os.Exit(1)
		}
		if *curDirF == false {
			fmt.Printf("$root")
		}
		if *curDirF == true {
			fmt.Printf(currentDir)
		}
	}

}

func getCurDir() string { // gets dir
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
