package main

import (
	"net/http"
	"flag"
	"log"
	"fmt"
	"regexp"
	"os/exec"
)

var Address = flag.String("address", "127.0.0.1", "IP address from which the app listens")
var Port = flag.Uint64("port", 1120, "Port on which the app listens")
var Command = flag.String("cmd", "", "Command to run on detected connection")

func main() {
	checkFlags()

	re := &RemoteExecutor{*Command}

	mux := http.NewServeMux()
	mux.Handle("/", re)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", *Address, *Port), mux); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func checkFlags() {
	flag.Parse()

	ipPattern := "\\A(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\z"
	if matched, err := regexp.MatchString(ipPattern, *Address); !matched || err != nil {
		log.Fatalln("Invalid IP address")
	}

	if *Port > (1<<16)-1 {
		log.Fatalln("Bad port number - max 65535")
	}

	if *Command == "" {
		log.Fatalln("Command must not be empty")
	}
}

type RemoteExecutor struct {
	Command string
}

func (re *RemoteExecutor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	go exec.Command(re.Command).Run()
}
