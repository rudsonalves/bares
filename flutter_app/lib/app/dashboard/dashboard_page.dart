import 'package:bares_app/routes.dart';
import 'package:flutter/material.dart';
import 'package:routefly/routefly.dart';
import 'package:signals/signals_flutter.dart';

import '../../config/app_config.dart';

class DashboardPage extends StatefulWidget {
  const DashboardPage({super.key});

  @override
  State<DashboardPage> createState() => _DashboardPageState();
}

class _DashboardPageState extends State<DashboardPage> {
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
          children: [
            Expanded(
              child: GridView.count(
                crossAxisCount: 2,
                children: [
                  InkWell(
                    onTap: () => Routefly.push(routePaths.users.path),
                    child: Card(
                      child: Center(
                        child: Column(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: [
                            Image.asset(
                              'assets/images/users_manager.png',
                              scale: 2.5,
                            ),
                            const Text("Gerenciar Usuários"),
                          ],
                        ),
                      ),
                    ),
                  ),
                  InkWell(
                    onTap: () => Routefly.push(routePaths.menu),
                    child: Card(
                      child: Center(
                        child: Column(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: [
                            Image.asset(
                              'assets/images/menu_manager.png',
                              scale: 2,
                            ),
                            const Text("Gerenciar Pratos"),
                          ],
                        ),
                      ),
                    ),
                  ),
                  const Card(
                    child: Center(child: Text("Métricas de Desempenho")),
                  ),
                  InkWell(
                    onTap: () async {
                      await appConfig.logout();
                      Routefly.navigate(routePaths.login);
                    },
                    child: Card(
                      child: Center(
                        child: Column(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: [
                            Image.asset(
                              'assets/images/exit_manager.png',
                              scale: 2,
                            ),
                            const Text("Sair"),
                          ],
                        ),
                      ),
                    ),
                  ),
                  // Adicione mais widgets conforme necessário...
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
