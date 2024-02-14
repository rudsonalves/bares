import 'package:bares_app/config/app_config.dart';
import 'package:bares_app/routes.dart';
import 'package:flutter/material.dart';
import 'package:routefly/routefly.dart';

class AppPage extends StatefulWidget {
  const AppPage({super.key});

  @override
  State<AppPage> createState() => _AppPageState();
}

class _AppPageState extends State<AppPage> {
  @override
  void initState() {
    super.initState();

    WidgetsBinding.instance.addPostFrameCallback((timeStamp) async {
      await Future.delayed(const Duration(seconds: 1));
      // Check if an user is logged
      if (AppConfig.instance.isLogged) {
        Routefly.navigate(routePaths.dashboard);
      } else {
        Routefly.navigate(routePaths.login);
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    final onPrimary = Theme.of(context).colorScheme.onPrimary;

    return Scaffold(
      body: Center(
        child: Stack(
          alignment: Alignment.center,
          children: [
            Image.asset(
              'assets/images/chopp.png',
              scale: 2,
            ),
            CircularProgressIndicator(
              color: onPrimary,
            ),
          ],
        ),
      ),
    );
  }
}
