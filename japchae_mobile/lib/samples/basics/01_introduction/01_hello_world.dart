import 'package:flutter/material.dart';

class HelloWorldApp extends StatelessWidget {
  @override
  Widget build(BuildContext ctx) {
    return Center(
        child: Text('Hello world!', textDirection: TextDirection.ltr));
  }
}
