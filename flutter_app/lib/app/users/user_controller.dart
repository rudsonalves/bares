import 'dart:developer';

import 'package:signals/signals_flutter.dart';

import '../../common/models/user_model.dart';
import '../../services/users_api_service.dart';

enum PageStatus {
  pageStateInitial,
  pageStateLoading,
  pageStateSuccess,
  pageStateError,
}

class UserController {
  final _usersList = <UserModel>[];
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
      final users = await UsersApiService.getAllUsers();
      _usersList.addAll(users);
      state.value = PageStatus.pageStateSuccess;
    } catch (err) {
      log('Erro: $err');
      state.value = PageStatus.pageStateError;
    }
  }

  Future<void> addUser() async {
    _getUsersList();
  }

  Future<void> updateUser() async {
    _getUsersList();
  }
}
