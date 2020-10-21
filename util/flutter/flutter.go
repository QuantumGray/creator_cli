package flutter

import (
	"creator/util/contexts"
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
	err := os.Chdir(ctx.GetValue["APPNAME"])
	check(err)
	if ctx.GetValue["WEB"] == "enabled" {
		cmd := exec.Command("flutter", "run", "-d", "chrome")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		err := cmd.Run()
		check(err)
	} else {
		cmd := exec.Command("flutter", "run")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		err = cmd.Run()
		check(err)

	}
	err = os.Chdir("/../")
	check(err)
}
