// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'dart:convert';

// Define an enum 'Role' with possible roles for a user.
enum Role {
  cliente,
  garcom,
  gerente,
  admin,
  cozinha,
  caixa,
}

// Extension on 'Role' to add additional functionality.
extension RoleExtension on Role {
  // Converts a string to a corresponding 'Role' enum value.
  static Role fromString(String roleStr) {
    switch (roleStr) {
      case 'cliente':
        return Role.cliente;
      case 'garcom':
        return Role.garcom;
      case 'gerente':
        return Role.gerente;
      case 'admin':
        return Role.admin;
      case 'cozinha':
        return Role.cozinha;
      case 'caixa':
        return Role.caixa;
      default:
        // Throws an exception if the input string doesn't match any role.
        throw Exception('Role type "$roleStr" not recognized.');
    }
  }

  // Generates a list of strings representing the names of 'Role' enum values.
  static List<String> stringList() {
    return Role.values.map((role) => role.name).toList();
  }
}

// UserModel class
class UserModel {
  int? id;
  String name;
  String email;
  String password;
  Role role;

  UserModel({
    this.id,
    required this.name,
    required this.email,
    required this.password,
    required this.role,
  });

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'id': id,
      'name': name,
      'email': email,
      'password': password,
      'role': role.name,
    };
  }

  factory UserModel.fromMap(Map<String, dynamic> map) {
    final pwString =
        map.containsKey('passwordHash') ? 'passwordHash' : 'password';

    return UserModel(
      id: map['id'] as int,
      name: map['name'] as String,
      email: map['email'] as String,
      password: map[pwString] as String,
      role: RoleExtension.fromString(map['role'] as String),
    );
  }

  String toJson() => json.encode(toMap());

  factory UserModel.fromJson(String source) =>
      UserModel.fromMap(json.decode(source) as Map<String, dynamic>);

  @override
  String toString() {
    return 'UserModel(id: $id, name: $name, email: $email, password: $password, role: $role)';
  }
}
