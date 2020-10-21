package updatetemp

import (
	"creator/util/gettemplate"
	"creator/util/unzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//UpdateTemplate updates a specific template
func UpdateTemplate(arg string) error {
	ex, err := os.Executable()
	check(err)
	exPath := filepath.Dir(ex)

	if isTempCached(arg, exPath) {
		err := os.RemoveAll(exPath + "/../cache/" + arg + "/")
		check(err)
		err = downzipTemplate(arg, exPath)
		check(err)
	} else {
		input := getInput()
		if input == "y" {
			downzipTemplate(arg, exPath)
		}
	}
	return err
}

func isTempCached(arg, exPath string) bool {

	files, err := ioutil.ReadDir(exPath + "/../cache")
	check(err)

	for _, f := range files {
		if strings.Contains(f.Name(), arg) {
			return true
		}
	}
	return false
}

func downzipTemplate(arg, exPath string) error {
	fmt.Println("Downloading template..")
	url := "https://github.com/ben-fornefeld/" + arg + "/archive/main.zip"
	err := gettemplate.DownloadFile(fmt.Sprintf("fc_t_%v.zip", arg), url) //Downloads file from that url
	fmt.Println("Deploying template..")
	_, err = unzip.Unzip(fmt.Sprintf("fc_t_%v.zip", arg), exPath+"/../cache") //Unzips the file to the "cache" folder
	err = os.Rename(exPath+"/../cache/"+arg+"-main/", exPath+"/../cache/"+arg+"/")
	err = os.Remove(fmt.Sprintf("fc_t_%v.zip", arg))
	fmt.Println("Template is now ready to use!")
	return err
}

func getInput() string {
	fmt.Println("We couldn't find that template on your machine. Should we get it for you? (y/n)")
	var inputString string
	fmt.Scanf("%s", &inputString)
	input := strings.ToLower(inputString)
	return input
}
