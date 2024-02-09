import 'package:flutter_secure_storage/flutter_secure_storage.dart';

import '../common/constants/app_const.dart';

class SecureStorageManager {
  SecureStorageManager._();
  static final SecureStorageManager _instance = SecureStorageManager._();
  static SecureStorageManager get instance => _instance;

  final _storage = const FlutterSecureStorage();

  Future<void> storeToken(String token) async {
    await _storage.write(key: AppConst.jwtToken, value: token);
  }

  Future<String?> getToken() async {
    return await _storage.read(key: AppConst.jwtToken);
  }

  Future<void> deleteToken() async {
    await _storage.delete(key: AppConst.jwtToken);
  }

  Future<void> write(String key, String value) async {
    await _storage.write(key: key, value: value);
  }

  Future<String?> read(String key) async {
    return await _storage.read(key: key);
  }

  Future<void> delete(String key) async {
    await _storage.delete(key: key);
  }
}
