import 'package:flutter/material.dart';
	import 'package:pisser/features/presentation/ui/screens/main_page.dart';
	
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
	}