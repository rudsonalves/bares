import 'package:flutter/material.dart';
import 'package:routefly/routefly.dart';

import '../routes.dart';

class MainApp extends StatelessWidget {
  const MainApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      routerConfig: Routefly.routerConfig(
        routes: routes, // GENERATED
      ),
    );
  }
}
