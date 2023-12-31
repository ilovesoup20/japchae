import 'package:flutter/material.dart';

Widget basicApp() {
  return const MaterialApp(
    title: 'My app',
    home: SafeArea(child: MyScaffold()),
  );
}

class MyAppBar extends StatelessWidget {
  const MyAppBar({required this.title, super.key});

  final Widget title;

  @override
  Widget build(BuildContext ctx) {
    return Container(
        height: 56,
        padding: const EdgeInsets.symmetric(horizontal: 8),
        decoration: BoxDecoration(color: Colors.blue[500]),
        // Row is a horizontal, linear layout
        child: Row(children: [
          const IconButton(
            icon: Icon(Icons.menu),
            tooltip: 'Navigation menu',
            onPressed: null,
          ),
          // Expanded expands its child to fill the available space
          Expanded(
            child: title,
          ),
          const IconButton(
              icon: Icon(Icons.search), tooltip: 'Search', onPressed: null),
        ]));
  }
}

class MyScaffold extends StatelessWidget {
  const MyScaffold({super.key});

  @override
  Widget build(BuildContext ctx) {
    return Material(
      child: Column(
        children: [
          MyAppBar(
            title: Text('Example Title',
                style: Theme.of(ctx).primaryTextTheme.titleLarge),
          ),
          const Expanded(
            child: Center(
              child: Text('Hello, world!!!'),
            ),
          ),
        ],
      ),
    );
  }
}
