import 'dart:convert';
import 'dart:developer';

import 'package:http/http.dart' as http;
import 'package:signals/signals.dart';

import '../../common/constants/app_const.dart';
import '../../common/models/user_model.dart';
import '../../config/app_config.dart';
import '../../services/secure_storage_manager.dart';

class LoginException implements Exception {
  final String message;

  LoginException(this.message);

  @override
  String toString() => message;
}

class ServerException implements Exception {
  final String message;

  ServerException(this.message);

  @override
  String toString() => "Erro de servidor: $message";
}

class LoginController {
  final email = signal<String>('');
  final emailError = signal<String?>(null);
  final password = signal<String>('');

  final appConfig = AppConfig.instance;

  void dispose() {
    email.dispose();
    emailError.dispose();
    password.dispose();
  }

  void validateEmail() {
    if (email().contains('@')) {
      emailError.value = null;
    } else {
      emailError.value = 'Enter a valid email.';
    }
  }

  bool isValit() {
    validateEmail();

    return emailError() == null;
  }

  Future<void> login() async {
    try {
      var response = await http.post(
        Uri.parse('${AppConst.apiUrl}/login'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'email': email.value,
          'password': password.value,
        }),
      );

      if (response.statusCode == 200) {
        var data = jsonDecode(response.body);

        var userMap = data['user'];
        final loginUser = UserModel.fromMap(userMap);
        AppConfig.instance.setUser(loginUser);
        log(loginUser.toString());

        var token = data['token'];
        await SecureStorageManager.instance.storeToken(token);
      } else if (response.statusCode == 401) {
        const message = 'Usuário ou senha incorretos.';
        log(message);
        throw LoginException(message);
      } else {
        final message =
            'Falha ao comunicar com o servidor. Código: ${response.statusCode}';
        log(message);
        throw ServerException(message);
      }
    } on LoginException {
      rethrow;
    } on ServerException {
      rethrow;
    } catch (err) {
      final message = 'Erro ao tentar login: $err';
      log(message);
      throw ServerException(message);
    }
  }
}
