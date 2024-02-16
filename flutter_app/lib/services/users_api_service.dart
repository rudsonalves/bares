import 'dart:convert';
import 'dart:developer';

import 'package:http/http.dart' as http;

import '../common/constants/app_const.dart';
import '../common/models/user_model.dart';
import 'functions.dart';

class UserApiCreatedException implements Exception {
  final String message;
  UserApiCreatedException(this.message);

  @override
  String toString() => "UserApiCreatedException: $message";
}

class UserApiGettedException implements Exception {
  final String message;
  UserApiGettedException(this.message);

  @override
  String toString() => "UserApiGettedException: $message";
}

class UserApiUpdateException implements Exception {
  final String message;

  UserApiUpdateException(this.message);

  @override
  String toString() => "UserApiUpdateException: $message";
}

class UserApiDeleteException implements Exception {
  final String message;

  UserApiDeleteException(this.message);

  @override
  String toString() => "UserApiDeleteException: $message";
}

class UsersApiService {
  UsersApiService._();

  // createUser create a new user
  static Future<UserModel?> createUser(UserModel user) async {
    String token = await getToken();
    final url = Uri.parse('${AppConst.apiUrlApi}/users');

    try {
      final response = await http.post(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
        body: user.toJson(),
      );

      if (response.statusCode == 201) {
        return UserModel.fromJson(response.body);
      } else {
        final message = 'Error: Failed to create users: ${response.statusCode}';
        log(message);
        throw UserApiCreatedException(message);
      }
    } on UserApiCreatedException {
      rethrow;
    } catch (err) {
      final message =
          'Error: Handle connection errors, timeouts, ...: ${err.toString()}';
      log(message);
      throw Exception(message);
    }
  }

  // getAllUsers return all users in system
  static Future<List<UserModel>> getAllUsers() async {
    String token = await getToken();
    final url = Uri.parse('${AppConst.apiUrlApi}/users');

    try {
      final response = await http.get(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );

      if (response.statusCode == 200) {
        final List<dynamic> usersJson = json.decode(response.body);
        return usersJson
            .map((userJson) => UserModel.fromMap(userJson))
            .toList();
      } else {
        const message = 'UsersAPI.getAllUser: Falha ao carregar usuários';
        log(message);
        throw UserApiGettedException(message);
      }
    } on UserApiGettedException {
      rethrow;
    } catch (err) {
      // Handle connection errors, timeouts, etc.
      final message =
          'Error: Handle connection errors, timeouts, ...: ${err.toString()}';
      log(message);
      throw Exception(message);
    }
  }

  // updateUser update user informations, less password
  static Future<UserModel?> updateUser(UserModel user) async {
    if (user.id == null) {
      const message = 'Error: User.id is null';
      log(message);
      throw UserApiUpdateException(message);
    }

    String token = await getToken();
    final url = Uri.parse('${AppConst.apiUrlApi}/users/${user.id}');

    try {
      final response = await http.put(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
        body: user.toJson(),
      );

      if (response.statusCode == 200) {
        return UserModel.fromJson(response.body);
      } else {
        final message = 'Error: Failed to update user: ${response.statusCode}';
        log(message);
        throw UserApiUpdateException(message);
      }
    } on UserApiUpdateException {
      rethrow;
    } catch (err) {
      final message =
          'Error: Handle connection errors, timeouts, ...: ${err.toString()}';
      log(message);
      throw Exception(message);
    }
  }

  // updateUserPass update user password
  static Future<void> updateUserPass(UserModel user) async {
    String token = await getToken();
    if (user.id == null) {
      log('Error: User.id is null');
      throw UserApiUpdateException('Failed: User.id is null');
    }
    final url = Uri.parse('${AppConst.apiUrlApi}/users/password/${user.id}');

    try {
      final response = await http.put(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
        body: user.toJson(),
      );
      if (response.statusCode == 200) {
        return;
      } else {
        log('Error: Failed to update user: ${response.statusCode}');
        throw UserApiUpdateException(
          'Failed to update user: ${response.statusCode}',
        );
      }
    } on UserApiUpdateException {
      rethrow;
    } catch (err) {
      // Handle connection errors, timeouts, etc.
      final message =
          'Error: Handle connection errors, timeouts, ...: ${err.toString()}';
      log(message);
      throw Exception(message);
    }
  }

  // deleteUser delete an user by id
  static Future<void> deleteUser(int userId) async {
    String token = await getToken();
    final url = Uri.parse('${AppConst.apiUrlApi}/users/$userId');

    try {
      final response = await http.delete(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );

      if (response.statusCode == 200) {
        return;
      } else {
        final message = 'Error: Failed to delete user: ${response.statusCode}';
        log(message);
        throw UserApiUpdateException(message);
      }
    } on UserApiDeleteException {
      rethrow;
    } catch (err) {
      final message =
          'Error: Handle connection errors, timeouts, ...: ${err.toString()}';
      log(message);
      throw Exception(message);
    }
  }
}
