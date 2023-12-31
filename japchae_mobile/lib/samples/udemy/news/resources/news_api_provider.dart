import 'package:http/http.dart' show Client;
import 'dart:convert';
import '../models/item_model.dart';

const _root = "https://hacker-news.firebaseio.com/v0";

class NewsApiProvider {
  Client client = Client();

  // https://github.com/hackernews/api

  fetchTopIds() async {
    final uri = Uri.parse("$_root/topstories.json");
    final response = await client.get(uri);

    final ids = json.decode(response.body);

    return ids;
  }

  fetchItem(int id) async {
    final uri = Uri.parse("$_root/item/$id.json");
    final response = await client.get(uri);
    final parsedJson = json.decode(response.body);

    return Item.fromJson(parsedJson);
  }
}
