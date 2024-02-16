import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:routefly/routefly.dart';
import 'package:signals/signals_flutter.dart';

import '../../../common/models/user_model.dart';
import '../../../common/widgets/dialogs.dart';
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

  late final String title;
  late final String buttonLabel;
  late final bool newUserMode;
  late final UserModel? user;
  late final void Function() callbackFunc;

  @override
  void dispose() {
    _controller.dispose();
    _visibility.dispose();
    super.dispose();
  }

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

      final createUser = await UsersApiService.createUser(user);
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

      final updateUser = await UsersApiService.updateUser(user!);
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

  Future<void> _changePassword() async {
    final password = signal<String>('');
    final passwordError = signal<String?>(null);
    final obscureText = signal<bool>(true);

    bool? cancel = await showDialog<bool>(
      context: context,
      builder: (context) => AlertDialog(
        title: const Text('Trocar Senha'),
        content: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Form(
              child: TextField(
                onChanged: password.set,
                obscureText: obscureText.watch(context),
                decoration: InputDecoration(
                  labelText: 'Entre uma nova Senha',
                  errorText: passwordError.watch(context),
                  suffix: IconButton(
                    onPressed: () => obscureText.value = !obscureText(),
                    icon: Icon(
                      !obscureText() ? Icons.visibility : Icons.visibility_off,
                    ),
                  ),
                ),
              ),
            ),
          ],
        ),
        actions: [
          FilledButton(
            onPressed: () => Navigator.pop(context, true),
            child: const Text('Cancela'),
          ),
          FilledButton(
            onPressed: () {
              passwordError.value = validatePassword(password());
              if (passwordError() == null) {
                Navigator.pop(context, false);
              }
            },
            child: const Text('Alterar'),
          ),
        ],
      ),
    );

    if (cancel != null && !cancel) {
      user!.password = password();
      await UsersApiService.updateUserPass(user!);

      if (!mounted) return;
      password.dispose();
      passwordError.dispose();
      obscureText.dispose();
      Navigator.pop(context);
    }

    password.dispose();
    passwordError.dispose();
    obscureText.dispose();
  }

  String? validatePassword(String? value) {
    if (value == null) {
      return "Password não pode ser nulo";
    }

    final regExp = RegExp(r'^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9]).{6,16}$');
    if (!regExp.hasMatch(value)) {
      return "A senha deve possuir 6 a 16 caracteres entre números, letras maiúsculas e minúsculas";
    }

    return null;
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
                        onPressed: _changePassword,
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
