package websupport

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

//Enables chrome support for the project
func EnableWeb(ctx *contexts.Context) {
	appName := ctx.GetValue["APPNAME"]

	err := os.Chdir(appName)
	check(err)

	err = exec.Command("flutter", "channel", "beta").Run()
	check(err)

	err = exec.Command("flutter", "upgrade").Run()
	check(err)

	err = exec.Command("flutter", "config", "--enable-web").Run()
	check(err)
}

//Disables web support for the project
func DisableWeb() {
	err := exec.Command("flutter", "config", "--no-enable-web").Run()
	check(err)
}
