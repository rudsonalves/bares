import 'dart:developer';

import 'package:signals/signals_flutter.dart';

import '../../common/models/user_model.dart';
import '../../services/secure_storage_manager.dart';
import '../../services/users_api_service.dart';

enum PageStatus {
  pageStateInitial,
  pageStateLoading,
  pageStateSuccess,
  pageStateError,
}

class UserController {
  final _usersList = <UserModel>[];
  final _userApi = UsersApiService();
  final state = signal<PageStatus>(PageStatus.pageStateInitial);

  List<UserModel> get userList => _usersList;

  void init() {
    _getUsersList();
  }

  void dispose() {
    state.dispose();
  }

  Future<void> _getUsersList() async {
    try {
      state.value = PageStatus.pageStateLoading;
      _usersList.clear();
      await Future.delayed(const Duration(milliseconds: 1000));
      final users = await _userApi.getAllUsers(await _getToken());
      _usersList.addAll(users);
      state.value = PageStatus.pageStateSuccess;
    } catch (err) {
      log('Erro: $err');
      state.value = PageStatus.pageStateError;
    }
  }

  Future<String> _getToken() async {
    String? token = await SecureStorageManager.instance.getToken();
    if (token == null) {
      const message = 'Erro: You don\'t have an authentication token';
      log(message);
      throw Exception(message);
    }
    return token;
  }

  Future<void> addUser() async {
    _getUsersList();
  }

  Future<void> updateUser() async {
    _getUsersList();
  }
}
