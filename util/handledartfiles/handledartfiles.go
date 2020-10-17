package handledartfiles

import (
	"creator/util/contexts"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

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

//Detects all '.dart' files in root
func ScanForFiles(ctx *contexts.Context, root string) {
	appName := ctx.GetValue["APPNAME"]

	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".dart") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		parseFile(file, appName)
	}
}
