import 'package:flutter/material.dart';

Future<void> showMessageDialog(
  BuildContext context, {
  required String title,
  required String message,
  IconData iconData = Icons.warning,
}) async {
  showDialog(
    context: context,
    builder: (context) => AlertDialog(
      icon: Icon(
        iconData,
        size: 32,
      ),
      title: Text(title),
      content: Text(message),
      actions: [
        FilledButton(
          onPressed: () => Navigator.pop(context),
          child: const Text('Fechar'),
        ),
      ],
    ),
  );
}
