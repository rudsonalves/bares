import 'dart:developer';

import 'package:bares_app/routes.dart';
import 'package:flutter/material.dart';
import 'package:routefly/routefly.dart';
import 'package:signals/signals_flutter.dart';

import 'user_controller.dart';

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
      body: Center(
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
                      itemBuilder: (context, index) => InkWell(
                        onTap: () async {
                          log(userList[index].toString());
                          await Routefly.push(
                            routePaths.users.edit,
                            arguments: {
                              'user': userList[index],
                              'callbackFunc': callbackFunc,
                            },
                          );
                        },
                        child: ListTile(
                          // leading: Text(userList[index].id.toString()),
                          title: Text(userList[index].name),
                          subtitle: Text(userList[index].email),
                          trailing: Text(userList[index].role.name),
                        ),
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
