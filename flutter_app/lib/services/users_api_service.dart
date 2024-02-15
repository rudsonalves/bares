// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'dart:convert';
import 'dart:developer';

import 'package:http/http.dart' as http;

import '../common/constants/app_const.dart';
import '../common/models/user_model.dart';

class UserApiCreatedException implements Exception {
  final String message;
  UserApiCreatedException(this.message);

  @override
  String toString() => "UserApiCreatedException: $message";
}

class UserApiUpdateException implements Exception {
  final String message;

  UserApiUpdateException(this.message);

  @override
  String toString() => "UserApiUpdateException: $message";
}

class UsersApiService {
  Future<UserModel?> createUser(UserModel user, String token) async {
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
        // If the server returned a 201 CREATED response code,
        // then parse the JSON and return the user model.
        return UserModel.fromJson(response.body);
      } else {
        // If the server did not return a 201 CREATED response code,
        // throw an exception or handle the error as needed.
        log('Error: Failed to create user: ${response.statusCode}');
        throw UserApiCreatedException(
            'Failed to create user: ${response.statusCode}');
      }
    } on UserApiCreatedException {
      return null;
    } catch (err) {
      // Handle connection errors, timeouts, etc.
      log('Error: Handle connection errors, timeouts, ...: ${err.toString()}');
      return null;
    }
  }

  // getAllUsers return all users in sistem
  Future<List<UserModel>> getAllUsers(String token) async {
    final url = Uri.parse('${AppConst.apiUrlApi}/users');

    try {
      final response = await http.get(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token'
        },
      );

      if (response.statusCode == 200) {
        final List<dynamic> usersJson = json.decode(response.body);
        return usersJson
            .map((userJson) => UserModel.fromMap(userJson))
            .toList();
      } else {
        throw Exception('UsersAPI.getAllUser: Falha ao carregar usuários');
      }
    } catch (err) {
      throw Exception('UsersAPI.getAllUser: Falha ao carregar usuários: $err');
    }
  }

  // updateUser update user informations, less password
  Future<UserModel?> updateUser(UserModel user, String token) async {
    if (user.id == null) {
      log('Error: User.id is null');
      throw UserApiUpdateException('Failed: User.id is null');
    }
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
        // If the server returned a 200 UPDATE response code,
        // then parse the JSON and return the user model.
        return UserModel.fromJson(response.body);
      } else {
        // If the server did not return a 201 CREATED response code,
        // throw an exception or handle the error as needed.
        log('Error: Failed to update user: ${response.statusCode}');
        throw UserApiUpdateException(
            'Failed to update user: ${response.statusCode}');
      }
    } on UserApiUpdateException {
      return null;
    } catch (err) {
      // Handle connection errors, timeouts, etc.
      log('Error: Handle connection errors, timeouts, ...: ${err.toString()}');
      return null;
    }
  }

  Future<void> updateUserPass(UserModel user, String token) async {
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
            'Failed to update user: ${response.statusCode}');
      }
    } on UserApiUpdateException {
      return;
    } catch (err) {
      // Handle connection errors, timeouts, etc.
      log('Error: Handle connection errors, timeouts, ...: ${err.toString()}');
      return;
    }
  }
}
