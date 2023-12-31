import 'package:flutter/material.dart';

class GridViewApp extends StatelessWidget {
  const GridViewApp({super.key});

  @override
  Widget build(BuildContext ctx) {
    const title = 'Grid List';

    return MaterialApp(
      title: title,
      home: Scaffold(
        appBar: AppBar(
          title: const Text(title),
        ),
        body: GridView.count(
          crossAxisCount: 2,
          children: List.generate(100, (index) {
            return Center(
              child: Text('Item $index',
                  style: Theme.of(ctx).textTheme.headlineSmall),
            );
          }),
        ),
      ),
    );
  }
}
