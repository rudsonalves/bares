import 'package:flutter/material.dart';
import 'package:routefly/routefly.dart';

class AppPage extends StatefulWidget {
  const AppPage({super.key});

  @override
  State<AppPage> createState() => _AppPageState();
}

class _AppPageState extends State<AppPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Main Page'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            FilledButton(
              onPressed: () => Routefly.push('/dashboard'),
              child: const Text('DashBoard Page'),
            ),
            FilledButton(
              onPressed: () => Routefly.push('/user'),
              child: const Text('User Page'),
            ),
          ],
        ),
      ),
    );
  }
}
