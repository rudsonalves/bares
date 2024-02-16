import 'dart:developer';

import 'secure_storage_manager.dart';

Future<String> getToken() async {
  String? token = await SecureStorageManager.instance.getToken();
  if (token == null) {
    const message = 'Erro: You don\'t have an authentication token';
    log(message);
    throw Exception(message);
  }
  return token;
}
