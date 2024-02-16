import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:routefly/routefly.dart';
import 'package:signals/signals_flutter.dart';

import '../../common/models/user_model.dart';
import '../../routes.dart';
import '../../services/users_api_service.dart';
import 'user_controller.dart';
import 'widgets/dismissible_list_tile.dart';

class UsersPage extends StatefulWidget {
  const UsersPage({super.key});

  @override
  State<UsersPage> createState() => _UsersPageState();
}

class _UsersPageState extends State<UsersPage> {
  final _controller = UserController();

  @override
  void initState() {
    super.initState();

    _controller.init();
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  void callbackFunc() {
    _controller.updateUser();
  }

  Future<bool> editUser(UserModel user) async {
    await Routefly.push(
      routePaths.users.edit,
      arguments: {
        'user': user,
        'callbackFunc': callbackFunc,
      },
    );
    return false;
  }

  Future<bool> deleteUser(UserModel user) async {
    bool? response = await showDialog<bool>(
      context: context,
      builder: (context) => AlertDialog(
        title: const Text('Remover?'),
        content: Text('Confirma a remoção do usuário\n${user.name}?'),
        actions: [
          FilledButton(
            onPressed: () => Navigator.pop(context, true),
            child: const Text('Apagar'),
          ),
          FilledButton(
            onPressed: () => Navigator.pop(context, false),
            child: const Text('Cancelar'),
          ),
        ],
      ),
    );
    if (response != null && response == true) {
      log('Apagar');
      await UsersApiService.deleteUser(user.id!);
      _controller.updateUser();
    }
    return false;
  }

  @override
  Widget build(BuildContext context) {
    final primary = Theme.of(context).colorScheme.primary;

    return Scaffold(
      appBar: AppBar(
        title: const Text('Usuários'),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () async {
          await Routefly.push(
            routePaths.users.edit,
            arguments: {
              'callbackFunc': callbackFunc,
            },
          );
          _controller.addUser();
        },
        child: const Icon(Icons.add),
      ),
      body: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 12),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Expanded(
              child: Watch(
                (context) {
                  final state = _controller.state();

                  if (state == PageStatus.pageStateLoading) {
                    return Center(
                      child: CircularProgressIndicator(
                        color: primary,
                      ),
                    );
                  }

                  if (state == PageStatus.pageStateSuccess) {
                    final userList = _controller.userList;
                    return ListView.builder(
                      itemCount: userList.length,
                      itemBuilder: (context, index) => DismissibleListTile(
                        user: userList[index],
                        onLeftFunc: editUser,
                        onRightFunc: deleteUser,
                      ),
                    );
                  }

                  return const Center(
                    child: Text('Desculpe. Houve um erro.'),
                  );
                },
              ),
            ),
          ],
        ),
      ),
    );
  }
}
