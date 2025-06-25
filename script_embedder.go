package simplebashscript

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

// //go:embed scripts/*
// var Script embed.FS

func CallScript() {
	fmt.Println("calling script")
	cmd := exec.Command("./process.sh")
	stdOutPipe, _ := cmd.StdoutPipe()
	stdErrPipe, _ := cmd.StderrPipe()
	stdOutScanner := bufio.NewScanner(stdOutPipe)
	stdErrScanner := bufio.NewScanner(stdErrPipe)

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	//Print the lines returned by the script to stdout as they are received
	go func() {
		for stdOutScanner.Scan() {
			fmt.Println(stdOutScanner.Text())
		}
	}()

	//Print the lines returned by the script to stderr as they are received
	go func() {
		for stdErrScanner.Scan() {
			fmt.Println(stdErrScanner.Text())
		}
	}()

	//this captures exit errors returned by the bash script
	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
	}

	//this captures an error in the Scanner object
	fmt.Printf("Encountered the following errors: %v\n", stdOutScanner.Err())
}
