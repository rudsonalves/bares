import 'dart:convert';

class MenuItemModel {
  int? id;
  String name;
  String description;
  double price;
  String imageURL;

  MenuItemModel({
    this.id,
    required this.name,
    required this.description,
    required this.price,
    required this.imageURL,
  });

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'id': id,
      'name': name,
      'description': description,
      'price': price,
      'imageURL': imageURL,
    };
  }

  factory MenuItemModel.fromMap(Map<String, dynamic> map) {
    return MenuItemModel(
      id: map['id'] != null ? map['id'] as int : null,
      name: map['name'] as String,
      description: map['description'] as String,
      price: map['price'] as double,
      imageURL: map['imageURL'] as String,
    );
  }

  String toJson() => json.encode(toMap());

  factory MenuItemModel.fromJson(String source) =>
      MenuItemModel.fromMap(json.decode(source) as Map<String, dynamic>);
}
