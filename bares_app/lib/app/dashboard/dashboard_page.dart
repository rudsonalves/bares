import 'package:flutter/material.dart';
import 'package:routefly/routefly.dart';

class DashboardPage extends StatefulWidget {
  const DashboardPage({super.key});

  @override
  State<DashboardPage> createState() => _DashboardPageState();
}

class _DashboardPageState extends State<DashboardPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Dashboard Page'),
      ),
      body: Center(
        child: ElevatedButton(
          onPressed: () => Routefly.pop(context),
          child: const Text('Back'),
        ),
      ),
    );
  }
}
