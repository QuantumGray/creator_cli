package createapp

import (
	"fluttercreator/util/gettemplate"
	"fluttercreator/util/unzip"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// CreateContext : context parameter that gets passed arround by creator functions
type CreateContext struct {
	getValue map[string]string
}

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

func writeDartFiles(appName string) {
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
				title: 'Flutterinio App',
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

func executeFlutterCreate(appName string) {
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

func getTemplate(arg string) {
	url := "https://github.com/ben-fornefeld/" + arg + "/archive/main.zip"
	gettemplate.DownloadFile(fmt.Sprintf("fc_t_%v.zip", arg), url) //Downloads file from that url

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	unzip.Unzip(fmt.Sprintf("fc_t_%v.zip", arg), exPath+"/../cache") //Unzips the file to the "cache" folder
	os.Remove(fmt.Sprintf("fc_t_%v.zip", arg))
}

// CreateApp : parent function to delegate creator functions
func CreateApp(arg string) {
	appName := getAppNameAsInput()

	getTemplate(arg)

	executeFlutterCreate(appName)

	err := os.Chdir(appName + "/lib")
	check(err)

	createTDDStructure()

	writeDartFiles(appName)

	fmt.Println("Flutter project has been created in a clean way!")
}
