import 'package:flutter/material.dart';
import 'package:routefly/routefly.dart';
import 'package:signals/signals_flutter.dart';

import 'common/theme/color_schemes.dart';
import 'config/app_config.dart';
import 'routes.dart';

class MainApp extends StatefulWidget {
  const MainApp({super.key});

  @override
  State<MainApp> createState() => _MainAppState();
}

class _MainAppState extends State<MainApp> {
  final appConfig = AppConfig.instance;

  @override
  void initState() {
    super.initState();

    appConfig.readConfiguration();
  }

  @override
  void dispose() {
    appConfig.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      routerConfig: Routefly.routerConfig(
        routes: routes, // GENERATED
      ),
      theme: ThemeData(
        useMaterial3: true,
        colorScheme: lightColorScheme,
      ),
      darkTheme: ThemeData(
        useMaterial3: true,
        colorScheme: darkColorScheme,
      ),
      themeMode: appConfig.themeMode.watch(context),
      debugShowCheckedModeBanner: false,
    );
  }
}
