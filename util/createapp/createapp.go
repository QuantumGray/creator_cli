package createapp

import (
	"creator/util/contexts"
	"creator/util/copy"
	"creator/util/gettemplate"
	"creator/util/handledartfiles"
	"creator/util/unzip"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func check(e error) error {
	if e != nil {
		return e
	}
	return nil
}

func passCodeToFile(path string, cont []byte) {
	err := ioutil.WriteFile(path, cont, 0644)
	check(err)
}

func executeFlutterCreate(ctx *contexts.Context) {
	appName := ctx.GetValue["APPNAME"]
	err := exec.Command("flutter", "create", appName).Run()
	check(err)
}

func getAppNameAsInput() string {
	fmt.Println("What is the name of your creator project?")
	var inputString string
	fmt.Scanf("%s", &inputString)
	appName := strings.ToLower(inputString)
	return appName
}

func getTemplate(ctx *contexts.Context) {
	arg := ctx.GetValue["SHA"]

	url := "https://github.com/ben-fornefeld/" + arg + "/archive/main.zip"
	gettemplate.DownloadFile(fmt.Sprintf("fc_t_%v.zip", arg), url) //Downloads file from that url

	ex, err := os.Executable()
	check(err)
	exPath := filepath.Dir(ex)
	ctx.GetValue["EXPATH"] = exPath

	unzip.Unzip(fmt.Sprintf("fc_t_%v.zip", arg), exPath+"/../cache") //Unzips the file to the "cache" folder
	os.Remove(fmt.Sprintf("fc_t_%v.zip", arg))
	copyCacheToProject(ctx)
}

func copyCacheToProject(ctx *contexts.Context) {
	arg := ctx.GetValue["SHA"]
	appName := ctx.GetValue["APPNAME"]
	path := ctx.GetValue["EXPATH"]

	err := copy.CopyDir(path+"/../cache/"+arg+"-main/", appName)
	check(err)
	//handledartfiles.ParseFile(appName+"/lib/main.dart", appName)
	handledartfiles.ScanForFiles(ctx, appName+"/lib")
}

// CreateApp : parent function to delegate creator functions
func CreateApp(ctx *contexts.Context) {
	ctx.GetValue["APPNAME"] = getAppNameAsInput()

	executeFlutterCreate(ctx)

	getTemplate(ctx)

	fmt.Println("Flutter project has been created in a clean way!")
}
