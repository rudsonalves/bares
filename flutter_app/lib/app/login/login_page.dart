import 'package:bares_app/routes.dart';
import 'package:flutter/material.dart';
import 'package:routefly/routefly.dart';
import 'package:signals/signals_flutter.dart';

import '../../common/widgets/dialogs.dart';
import 'login_controller.dart';

class LoginPage extends StatefulWidget {
  const LoginPage({super.key});

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  final controller = LoginController();
  final visibility = signal<bool>(false);

  @override
  void dispose() {
    controller.dispose();
    visibility.dispose();
    super.dispose();
  }

  Future<void> _loginProcess() async {
    if (controller.isValit()) {
      try {
        await controller.login();
        if (!context.mounted) return;
        Routefly.navigate(routePaths.dashboard);
      } on LoginException catch (_) {
        if (!context.mounted) return;
        showMessageDialog(
          context,
          title: 'Erro nas Credenciais',
          message: 'Email ou senha inválida. Tente novamente',
        );
      } on ServerException catch (_) {
        if (!context.mounted) return;
        showMessageDialog(
          context,
          title: 'Erro no Servidor',
          message: 'Aparentemente estamos com um erro no servidor. '
              'Favor tentar mais tarde!',
        );
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Login Page'),
      ),
      body: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 12),
        child: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              TextField(
                onChanged: controller.email.set,
                decoration: InputDecoration(
                    label: const Text('Email'),
                    errorText: controller.emailError.watch(context)),
              ),
              TextField(
                onChanged: controller.password.set,
                obscureText: !visibility.watch(context),
                decoration: InputDecoration(
                  label: const Text('Senha'),
                  suffixIcon: IconButton(
                    onPressed: () {
                      visibility.value = !visibility.value;
                    },
                    icon: Icon(
                      !visibility() ? Icons.visibility_off : Icons.visibility,
                    ),
                  ),
                ),
              ),
              const SizedBox(height: 20),
              FilledButton(
                onPressed: _loginProcess,
                child: const Text('Login'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
