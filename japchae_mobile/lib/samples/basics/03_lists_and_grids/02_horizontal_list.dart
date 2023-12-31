import 'package:flutter/material.dart';

class HorizontalList extends StatelessWidget {
  const HorizontalList({super.key});

  @override
  Widget build(BuildContext ctx) {
    const title = 'Horizontal List';
    const width = 160.0;

    return MaterialApp(
      title: title,
      home: Scaffold(
        appBar: AppBar(title: const Text(title)),
        body: Container(
          margin: const EdgeInsets.symmetric(vertical: 20),
          height: 200,
          child: ListView(
            scrollDirection: Axis.horizontal,
            children: <Widget>[
              Container(
                width: width,
                color: Colors.red,
              ),
              Container(
                width: width,
                color: Colors.blue,
              ),
              Container(
                width: width,
                color: Colors.green,
              ),
              Container(
                width: width,
                color: Colors.yellow,
              ),
              Container(
                width: width,
                color: Colors.orange,
              ),
            ],
          ),
        ),
      ),
    );
  }
}
