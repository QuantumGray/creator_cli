package flutter

import (
	"bufio"
	"creator/util/contexts"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Run function executes "flutter run"
func Run(ctx *contexts.Context) {
	if ctx.GetValue["WEB"] == "enable" {
		cmd := exec.Command("flutter", "run", "-d", "chrome")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		check(err)
	} else {
		err := os.Chdir(ctx.GetValue["APPNAME"])
		check(err)
		cmd := exec.Command("flutter", "run")

		// Stdout + stderr
		out, err := cmd.StderrPipe() // rm writes the prompt to err
		check(err)
		r := bufio.NewReader(out)

		// Stdin
		in, err := cmd.StdinPipe()
		check(err)
		defer in.Close()

		// Start the command!
		err = cmd.Start()
		check(err)

		line, _, err := r.ReadLine()
		for err != nil {
			fmt.Println(line)
			line, _, err = r.ReadLine()
		}
	}
}
