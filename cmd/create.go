package cmd

/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create an app from a template with a SHA",
	Long:  `do it 4real`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args[0]) != 8 {
			return errors.New("this hasn't got 8 chars")
		}
		if len(args) != 1 {
			return errors.New("too many positional args")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if res, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%v", args[0])); err != nil {
			fmt.Println(err)
		} else {
			createApp()
			fmt.Println(res.StatusCode)
		}
	},
}

var (
	appName string
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createTDDStructure() {
	os.MkdirAll("core/errors", os.ModePerm)
	os.MkdirAll("features/data/data_sources", os.ModePerm)
	os.MkdirAll("features/data/models", os.ModePerm)
	os.MkdirAll("features/data/repositories", os.ModePerm)
	os.MkdirAll("features/domain/entities", os.ModePerm)
	os.MkdirAll("features/domain/repositories", os.ModePerm)
	os.MkdirAll("features/domain/services", os.ModePerm)
	os.MkdirAll("features/presentation/cubits", os.ModePerm)
	os.MkdirAll("features/presentation/ui/screens", os.ModePerm)
	os.MkdirAll("features/presentation/ui/organisms", os.ModePerm)
	os.MkdirAll("features/presentation/ui/molecules", os.ModePerm)
	os.MkdirAll("features/presentation/ui/atoms", os.ModePerm)
}

func writeDartFiles() {
	mainDartContent :=
		[]byte(`import 'package:flutter/material.dart';
	import 'package:` + appName + `/features/presentation/ui/screens/main_page.dart';
	
	void main() {
		runApp(App());
	}
	
	class App extends StatelessWidget {
		@override
		Widget build(BuildContext context) {
			return MaterialApp(
				title: 'Flutter App',
				theme: ThemeData(),
				home: MainPage(),                
				debugShowCheckedModeBanner: false,
			);
		}
	}`)

	mainDartPageContent :=
		[]byte(`import 'package:flutter/material.dart';
	class MainPage extends StatelessWidget {
		@override
		Widget build(BuildContext context) {
			return Scaffold(
				body: Center(
					child: Container(
						child: Text('Be Creative.')
					)    
				)
			);
		}
	}`)

	widgetTestDartContent :=
		[]byte(`// This is a basic Flutter widget test.
		//
		// To perform an interaction with a widget in your test, use the WidgetTester
		// utility that Flutter provides. For example, you can send tap and scroll
		// gestures. You can also use WidgetTester to find child widgets in the widget
		// tree, read text, and verify that the values of widget properties are correct.
		
		import 'package:flutter/material.dart';
		import 'package:flutter_test/flutter_test.dart';
		
		import 'package:` + appName + `/main.dart';
		
		void main() {
		  testWidgets('Counter increments smoke test', (WidgetTester tester) async {
			// Build our app and trigger a frame.
			await tester.pumpWidget(App());
		
			// Verify that our counter starts at 0.
			expect(find.text('0'), findsOneWidget);
			expect(find.text('1'), findsNothing);
		
			// Tap the '+' icon and trigger a frame.
			await tester.tap(find.byIcon(Icons.add));
			await tester.pump();
		
			// Verify that our counter has incremented.
			expect(find.text('0'), findsNothing);
			expect(find.text('1'), findsOneWidget);
		  });
		}
		`)

	_, err := os.Create("features/presentation/ui/screens/main_page.dart")
	check(err)

	passCodeToFile("main.dart", mainDartContent)
	passCodeToFile("features/presentation/ui/screens/main_page.dart", mainDartPageContent)
	os.Chdir("..")
	passCodeToFile("test/widget_test.dart", widgetTestDartContent)
}

func passCodeToFile(path string, cont []byte) {
	err := ioutil.WriteFile(path, cont, 0644)
	check(err)
}

func executeFlutterCreate() {
	err := exec.Command("flutter", "create", appName).Run()
	check(err)
}

func getAppNameAsInput() {
	fmt.Println("What is the name of your new Flutter project?")
	var inputString string
	fmt.Scanf("%s", &inputString)
	appName = strings.ToLower(inputString)
}

func createApp() {
	getAppNameAsInput()

	executeFlutterCreate()

	err := os.Chdir(appName + "/lib")
	check(err)

	createTDDStructure()

	writeDartFiles()

	fmt.Println("New Flutter project has been created in a clean way!")
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
