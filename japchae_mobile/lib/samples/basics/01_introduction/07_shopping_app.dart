import 'package:flutter/material.dart';

Widget shoppingApp() {
  return MaterialApp(
    home: Scaffold(
      body: Center(
        child: ShoppingListItem(
          product: const Product(name: 'Chips'),
          inCart: true,
          onCartChanged: (product, inCart) {},
        ),
      ),
    ),
  );
}

class Product {
  const Product({required this.name});

  final String name;
}

typedef CartChangedCallback = Function(Product product, bool inCart);

class ShoppingListItem extends StatelessWidget {
  ShoppingListItem({
    required this.product,
    required this.inCart,
    required this.onCartChanged,
  }) : super(key: ObjectKey(product));

  final Product product;
  final bool inCart;
  final CartChangedCallback onCartChanged;

  Color _getColor(BuildContext ctx) {
    return inCart ? Colors.black54 : Theme.of(ctx).primaryColor;
  }

  TextStyle? _getTextStyle(BuildContext ctx) {
    if (!inCart) return null;

    return const TextStyle(
      color: Colors.black54,
      decoration: TextDecoration.lineThrough,
    );
  }

  @override
  Widget build(BuildContext ctx) {
    return ListTile(
      onTap: () {
        onCartChanged(product, inCart);
      },
      leading: CircleAvatar(
        backgroundColor: _getColor(ctx),
        child: Text(product.name[0]),
      ),
      title: Text(product.name, style: _getTextStyle(ctx)),
    );
  }
}
