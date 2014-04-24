package main

import (
	"fmt"
	"os"
	"simplemath"
	"strconv"
)

var Usage = func() {
	fmt.Println("USAGE:calc command [arguments] ...")
	fmt.Println("\nThe Commands are:\n\tadd\tAddition of two values.\n\tsqrt\tSquare root of a non-negetive value.")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	args = args[1:]
	//os.Args provides access to raw command-line arguments.
	//Note that the first value in this slice is the path to the program,
	//and os.Args[1:] holds the arguments to the program.
	//more info: https://gobyexample.com/command-line-arguments

	switch args[0] {
	case "add":
		if len(args) != 3 {
			fmt.Println("USAGE:calc add <integer1><integer2")
			return
		}
		v1, err1 := strconv.Atoi(args[1])
		v2, err2 := strconv.Atoi(args[2])

		if err1 != nil || err2 != nil {
			fmt.Println("USAGE:calc add <integer1><integer2")
			return
		}

		ret := simplemath.Add(v1, v2)
		fmt.Println("Result:", ret)
	case "sqrt":
		if len(args) != 2 {
			fmt.Println("USAGE:sqrt <integer1>")
			return
		}
		v1, err1 := strconv.Atoi(args[1])
		if err1 != nil {
			fmt.Println("USAGE:sqrt <integer1>")
			return
		}
		ret := simplemath.Sqrt(v1)
		fmt.Println("Result:", ret)
	default:
		Usage()
	}
}
