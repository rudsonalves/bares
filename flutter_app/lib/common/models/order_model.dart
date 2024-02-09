import 'dart:convert';

// Define an enum 'Status' with possible roles for a order.
enum Status {
  recebido,
  preparando,
  pronto,
  entregue,
}

// Extension on 'Status' to add additional functionality.
extension StatusExtension on Status {
  // Converts a string to a corresponding 'Status' enum value.
  static Status fromString(String statusStr) {
    switch (statusStr) {
      case 'recebido':
        return Status.recebido;
      case 'preparando':
        return Status.preparando;
      case 'pronto':
        return Status.pronto;
      case 'entregue':
        return Status.entregue;
      default:
        throw Exception('Status type "$statusStr" not recognized.');
    }
  }

  // Generates a list of strings representing the names of 'Status' enum values.
  static List<String> stringList() {
    return Status.values.map((status) => status.name).toList();
  }
}

// OrderModel class
class OrderModel {
  int? id;
  int userId;
  DateTime dateHour;
  Status status;

  OrderModel({
    this.id,
    required this.userId,
    required this.dateHour,
    this.status = Status.entregue,
  });

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'id': id,
      'userId': userId,
      'dateHour': dateHour.millisecondsSinceEpoch,
      'status': status.name,
    };
  }

  factory OrderModel.fromMap(Map<String, dynamic> map) {
    return OrderModel(
      id: map['id'] != null ? map['id'] as int : null,
      userId: map['userId'] as int,
      dateHour: DateTime.fromMillisecondsSinceEpoch(map['dateHour'] as int),
      status: StatusExtension.fromString(map['status'] as String),
    );
  }

  String toJson() => json.encode(toMap());

  factory OrderModel.fromJson(String source) =>
      OrderModel.fromMap(json.decode(source) as Map<String, dynamic>);
}
