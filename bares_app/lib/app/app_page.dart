import 'package:bares_app/config/app_config.dart';
import 'package:bares_app/routes.dart';
import 'package:flutter/material.dart';
import 'package:routefly/routefly.dart';
import 'package:signals/signals_flutter.dart';

import '../services/secure_storage_manager.dart';

class AppPage extends StatefulWidget {
  const AppPage({super.key});

  @override
  State<AppPage> createState() => _AppPageState();
}

class _AppPageState extends State<AppPage> {
  final appConfig = AppConfig.instance;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Main Page'),
        actions: [
          IconButton(
            onPressed: appConfig.toogleThemeMode,
            icon: Icon(
              appConfig.themeMode.watch(context) == ThemeMode.dark
                  ? Icons.dark_mode
                  : appConfig.themeMode() == ThemeMode.light
                      ? Icons.light_mode
                      : Icons.auto_mode,
            ),
          ),
        ],
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.start,
          children: [
            Text('Bem vindo ${appConfig.user.name}'),
            const Spacer(),
            FilledButton(
              onPressed: () => Routefly.push(routePaths.dashboard),
              child: const Text('DashBoard Page'),
            ),
            FilledButton(
              onPressed: () => Routefly.push(routePaths.user),
              child: const Text('User Page'),
            ),
            FilledButton(
              onPressed: () => Routefly.navigate(routePaths.login),
              child: const Text('Login Page'),
            ),
            FilledButton(
              onPressed: () {
                SecureStorageManager.instance.deleteToken();
              },
              child: const Text('Logout'),
            ),
          ],
        ),
      ),
    );
  }
}
