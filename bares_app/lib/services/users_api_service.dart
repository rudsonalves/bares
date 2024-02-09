import 'dart:developer';

import 'package:http/http.dart' as http;

import '../common/constants/app_const.dart';
import '../common/models/user_model.dart';

class UsersApiService {
  Future<UserModel?> createUser(UserModel user) async {
    final url = Uri.parse('${AppConst.apiURL}/users');

    try {
      final response = await http.post(
        url,
        headers: {'Content-Type': 'application/json'},
        body: user.toJson(),
      );

      if (response.statusCode == 201) {
        // If the server returned a 201 CREATED response code,
        // then parse the JSON and return the user model.
        // log('response.body: ${response.body}');
        return UserModel.fromJson(response.body);
      } else {
        // If the server did not return a 201 CREATED response code,
        // throw an exception or handle the error as needed.
        log('Error: Failed to create user: ${response.statusCode}');
        throw Exception('Failed to create user: ${response.statusCode}');
      }
    } catch (err) {
      // Handle connection errors, timeouts, etc.
      log('Error: Handle connection errors, timeouts, ...: ${err.toString()}');
      return null;
    }
  }
}
