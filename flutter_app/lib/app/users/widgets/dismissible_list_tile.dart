import 'dart:developer';

import 'package:flutter/material.dart';

import '../../../common/constants/app_const.dart';
import '../../../common/models/user_model.dart';

Container baseDismissibleContainer(
  BuildContext context, {
  required AlignmentGeometry alignment,
  required Color color,
  required IconData icon,
  String? label,
  bool enable = true,
}) {
  List<Alignment> alignLeft = [
    Alignment.bottomLeft,
    Alignment.topLeft,
    Alignment.centerLeft,
  ];

  final colorScheme = Theme.of(context).colorScheme;
  late Widget rowIcon;

  Color enableColor = colorScheme.shadow;
  Color disableColor = colorScheme.tertiary;

  if (label != null) {
    if (alignLeft.contains(alignment)) {
      rowIcon = Row(
        mainAxisAlignment: MainAxisAlignment.start,
        children: [
          Icon(
            icon,
            color: enable ? enableColor : disableColor,
          ),
          const SizedBox(width: 16),
          Text(
            label,
            style: TextStyle(
              color: enable ? enableColor : disableColor,
            ),
          ),
        ],
      );
    } else {
      rowIcon = Row(
        mainAxisAlignment: MainAxisAlignment.end,
        children: [
          Text(
            label,
            style: TextStyle(
              color: enable ? enableColor : disableColor,
            ),
          ),
          const SizedBox(width: 16),
          Icon(
            icon,
            color: enable ? enableColor : disableColor,
          ),
        ],
      );
    }
  } else {
    rowIcon = Icon(icon);
  }

  return Container(
    padding: const EdgeInsets.symmetric(horizontal: 18),
    alignment: alignment,
    decoration: BoxDecoration(
      color: color,
      borderRadius: BorderRadius.circular(16),
    ),
    child: rowIcon,
  );
}

class DismissibleListTile extends StatelessWidget {
  final UserModel user;
  final Future<bool> Function(UserModel user) onLeftFunc;
  final Future<bool> Function(UserModel user) onRightFunc;
  final void Function()? onTap;

  const DismissibleListTile({
    super.key,
    required this.user,
    required this.onLeftFunc,
    required this.onRightFunc,
    this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: onTap,
      child: Dismissible(
        key: UniqueKey(),
        background: baseDismissibleContainer(
          context,
          alignment: Alignment.centerLeft,
          color: const Color(0xFF93D895),
          icon: Icons.edit,
          label: 'Editar',
          enable: true,
        ),
        secondaryBackground: baseDismissibleContainer(
          context,
          alignment: Alignment.centerRight,
          color: const Color(0xFFEE918B),
          icon: Icons.delete,
          label: 'Remover',
          enable: true,
        ),
        child: Card(
          child: ListTile(
            title: Text(user.name),
            subtitle: Text(user.email),
            trailing: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              mainAxisSize: MainAxisSize.min,
              children: [
                Image.asset(
                  roleImage[user.role]!,
                  width: 32,
                  height: 32,
                ),
                Text(user.role.name),
              ],
            ),
          ),
        ),
        confirmDismiss: (direction) async {
          if (direction == DismissDirection.startToEnd) {
            log('Editar');
            return await onLeftFunc(user);
          } else if (direction == DismissDirection.endToStart) {
            log('Remover');
            return onRightFunc(user);
          }
          return false;
        },
      ),
    );
  }
}
