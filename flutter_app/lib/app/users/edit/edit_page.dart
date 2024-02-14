import 'dart:developer';

import 'package:bares_app/common/widgets/dialogs.dart';
import 'package:flutter/material.dart';
import 'package:routefly/routefly.dart';
import 'package:signals/signals_flutter.dart';

import '../../../common/models/user_model.dart';
import '../../../services/secure_storage_manager.dart';
import '../../../services/users_api_service.dart';
import 'edit_controller.dart';

class EditPage extends StatefulWidget {
  const EditPage({super.key});

  @override
  State<EditPage> createState() => _EditPageState();
}

class _EditPageState extends State<EditPage> {
  final _controller = EditController();
  final _visibility = signal<bool>(false);
  final _userApiService = UsersApiService();
  final _secureStorage = SecureStorageManager.instance;

  late final String title;
  late final String buttonLabel;
  late final bool newUserMode;
  late final UserModel? user;
  late final void Function() callbackFunc;

  @override
  void initState() {
    super.initState();

    user = Routefly.query.arguments['user'] as UserModel?;
    callbackFunc = Routefly.query.arguments['callbackFunc'] as void Function();

    if (user != null) {
      _controller.init(user!);
      title = 'Atualizar Usuário';
      buttonLabel = 'Atualizar';
      newUserMode = false;
    } else {
      title = 'Novo Usuário';
      buttonLabel = 'Criar';
      newUserMode = true;
    }
  }

  Future<void> addNewUser() async {
    if (_controller.isValid()) {
      final user = UserModel(
        name: _controller.name(),
        email: _controller.email(),
        password: _controller.password(),
        role: _controller.role(),
      );

      String token = await _getToken();

      final createUser = await _userApiService.createUser(user, token);
      if (createUser == null) {
        log('Error: Create user return null.');
        if (!mounted) return;
        await showMessageDialog(
          context,
          title: 'Erro!',
          message: 'Erro na criação de usuário',
        );
      }
      if (!mounted) return;
      callbackFunc();
      Routefly.pop(context);
    }
  }

  Future<void> updateUser() async {
    if (_controller.isValid(false)) {
      user!.name = _controller.name();
      user!.email = _controller.email();
      user!.password = _controller.password();
      user!.role = _controller.role();

      String token = await _getToken();

      final updateUser = await _userApiService.updateUser(user!, token);
      if (updateUser == null) {
        log('Error: Update user return null.');
        if (!mounted) return;
        await showMessageDialog(
          context,
          title: 'Erro!',
          message: 'Erro na atualização de usuário',
        );
      }
      if (!mounted) return;
      callbackFunc();
      Navigator.pop(context);
    }
  }

  Future<String> _getToken() async {
    String? token = await _secureStorage.getToken();
    if (token == null) {
      const message = 'Error: You don\'t have an authentication token';
      log(message);
      throw Exception(message);
    }
    return token;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(title),
      ),
      body: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 12),
        child: Center(
          child: Column(
            children: [
              TextFormField(
                initialValue: _controller.name(),
                onChanged: _controller.name.set,
                decoration: InputDecoration(
                  label: const Text('Nome'),
                  hintText: 'Seu Nome',
                  errorText: _controller.nameError.watch(context),
                ),
              ),
              TextFormField(
                initialValue: _controller.email(),
                onChanged: _controller.email.set,
                decoration: InputDecoration(
                  label: const Text('Email'),
                  hintText: 'seu@email.address',
                  errorText: _controller.emailError.watch(context),
                ),
              ),
              newUserMode
                  ? TextFormField(
                      initialValue: _controller.password(),
                      onChanged: _controller.password.set,
                      obscureText: !_visibility(),
                      decoration: InputDecoration(
                        label: const Text('Senha'),
                        hintText: '****',
                        errorText: _controller.passwordError.watch(context),
                        suffixIcon: IconButton(
                          icon: Icon(_visibility.watch(context)
                              ? Icons.visibility
                              : Icons.visibility_off),
                          onPressed: () {
                            _visibility.value = !_visibility.value;
                          },
                        ),
                      ),
                    )
                  : Padding(
                      padding: const EdgeInsets.only(top: 12),
                      child: OutlinedButton(
                        onPressed: () {},
                        child: const Text('Mudar senha?'),
                      ),
                    ),
              const SizedBox(height: 20),
              const Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  'Papel do Usuário',
                  style: TextStyle(
                    fontSize: 12,
                  ),
                ),
              ),
              DropdownButtonHideUnderline(
                child: ButtonTheme(
                  alignedDropdown: true,
                  child: DropdownButton<Role>(
                    isExpanded: true,
                    value: _controller.role.watch(context),
                    items: Role.values
                        .map(
                          (role) => DropdownMenuItem<Role>(
                            value: role,
                            child: Text(role.name),
                          ),
                        )
                        .toList(),
                    onChanged: (value) {
                      if (value != null) {
                        _controller.role.set(value);
                      }
                    },
                  ),
                ),
              ),
              const SizedBox(height: 20),
              ButtonBar(
                alignment: MainAxisAlignment.center,
                children: [
                  FilledButton(
                    onPressed: newUserMode ? addNewUser : updateUser,
                    child: Text(buttonLabel),
                  ),
                  FilledButton(
                    onPressed: () => Routefly.pop(context),
                    child: const Text('Fechar'),
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
}
