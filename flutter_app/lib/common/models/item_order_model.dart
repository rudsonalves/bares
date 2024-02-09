import 'dart:convert';

// ignore_for_file: public_member_api_docs, sort_constructors_first
class ItemOrderModel {
  int? id;
  int orderId;
  int itemId;
  int amount;
  String comments;

  ItemOrderModel({
    this.id,
    required this.orderId,
    required this.itemId,
    this.amount = 1,
    this.comments = '',
  });

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'id': id,
      'orderId': orderId,
      'itemId': itemId,
      'amount': amount,
      'comments': comments,
    };
  }

  factory ItemOrderModel.fromMap(Map<String, dynamic> map) {
    return ItemOrderModel(
      id: map['id'] != null ? map['id'] as int : null,
      orderId: map['orderId'] as int,
      itemId: map['itemId'] as int,
      amount: map['amount'] as int,
      comments: map['comments'] as String,
    );
  }

  String toJson() => json.encode(toMap());

  factory ItemOrderModel.fromJson(String source) =>
      ItemOrderModel.fromMap(json.decode(source) as Map<String, dynamic>);
}
