package websupport

import (
	"fmt"
	"log"
	"os/exec"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Toggles web support for the project
func ToggleWebIntegration(x bool) {

	if x {
		err := exec.Command("flutter", "config", "--enable-web").Run()
		check(err)
		fmt.Println("Web support has been enabled..")
	} else {
		err := exec.Command("flutter", "config", "--no-enable-web").Run()
		check(err)
		fmt.Println("Web support has been disabled..")
	}

}
