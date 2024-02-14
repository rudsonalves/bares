import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:jwt_decode/jwt_decode.dart';
import 'package:signals/signals_flutter.dart';

import '../common/models/user_model.dart';
import '../services/secure_storage_manager.dart';

class StorageException implements Exception {
  final String message;
  StorageException(this.message);

  @override
  String toString() => "StorageException: $message";
}

class AppConfig {
  AppConfig._();
  static final AppConfig _instance = AppConfig._();
  static AppConfig get instance => _instance;

  final _storage = SecureStorageManager.instance;

  final UserModel _currentUser = UserModel(
    name: '',
    email: '',
    role: Role.cliente,
  );

  final themeMode = signal<ThemeMode>(ThemeMode.dark);

  bool _isLogged = false;
  bool get isLogged => _isLogged;

  void dispose() {
    themeMode.dispose();
  }

  UserModel get user => _currentUser;
  Future<void> setUser(UserModel newUser) async {
    _currentUser.copyUser(newUser);
    await _saveUser();
  }

  Future<void> toogleThemeMode() async {
    switch (themeMode()) {
      case ThemeMode.dark:
        themeMode.value = ThemeMode.light;
        break;
      case ThemeMode.light:
        themeMode.value = ThemeMode.system;
        break;
      case ThemeMode.system:
        themeMode.value = ThemeMode.dark;
        break;
    }
    await _saveThemeMode();
  }

  void _setThemeModeByString(String? theme) {
    switch (theme) {
      case 'dark':
        themeMode.value = ThemeMode.dark;
        break;
      case 'light':
        themeMode.value = ThemeMode.light;
        break;
      case 'system':
        themeMode.value = ThemeMode.system;
        break;
      default:
        break;
    }
  }

  Future<void> saveConfiguration() async {
    await _saveThemeMode();
    await _saveUser();
  }

  Future<void> _saveUser() async {
    try {
      await _storage.write('user', _currentUser.toJson());
    } catch (err) {
      final message = 'Falha ao salvar usuário: $err';
      log(message);
      throw StorageException(message);
    }
  }

  Future<void> _saveThemeMode() async {
    try {
      await _storage.write('themeMode', themeMode().name);
    } catch (err) {
      final message = 'Falha ao salvar themeMode: $err';
      log(message);
      throw StorageException(message);
    }
  }

  Future<void> readConfiguration() async {
    try {
      String? theme = await _storage.read('themeMode');
      String? userJson = await _storage.read('user');

      _setThemeModeByString(theme);

      if (userJson != null) {
        _currentUser.copyUser(UserModel.fromJson(userJson));
      }

      if (_currentUser.id != null) {
        _isLogged = true;
        if (await isTokenExpired()) {
          _currentUser.clearUser();
          _isLogged = false;
          await _storage.deleteToken();
        }
      }
    } catch (err) {
      final message = 'Falha ao ler a configuração: $err';
      log(message);
      throw StorageException(message);
    }
  }

  Future<bool> isTokenExpired() async {
    try {
      String? token = await _storage.getToken();
      if (token == null) return true;
      return Jwt.isExpired(token);
    } catch (err) {
      return false;
    }
  }

  Future<void> logout() async {
    await _storage.deleteToken();
    _currentUser.clearUser();
  }
}
