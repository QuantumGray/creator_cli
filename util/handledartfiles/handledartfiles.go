package handledartfiles

import (
	"creator/util/contexts"
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

//Parses the dart file and looks for "-APPNAME-"
func parseFile(dest, appName string) error {
	data, err := ioutil.ReadFile(dest)
	if err != nil {
		return err
	}
	s := string(data)

	if strings.Contains(s, "$APPNAME$") {
		bs := []byte(strings.ReplaceAll(s, "$APPNAME$", appName))
		err = ioutil.WriteFile(dest, bs, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

//Detects all .dart or .dummy files  from root tree
func ScanForFiles(ctx *contexts.Context, root string) {
	appName := ctx.GetValue["APPNAME"]

	var files []string
	var dummys []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".dart") {
			files = append(files, path)
		}
		if strings.Contains(path, ".dummy") {
			dummys = append(dummys, path)
		}
		return nil
	})
	check(err)
	for _, file := range files {
		err := parseFile(file, appName)
		check(err)
	}
	for _, file := range dummys {
		err := os.Remove(file)
		check(err)
	}
}
