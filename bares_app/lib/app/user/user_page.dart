import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:signals/signals_flutter.dart';

import '../../common/models/user_model.dart';
import '../../services/users_api_service.dart';
import 'user_page_controller.dart';

class UserPage extends StatefulWidget {
  const UserPage({super.key});

  @override
  State<UserPage> createState() => _UserPageState();
}

class _UserPageState extends State<UserPage> {
  final controller = UserPageController();
  final visibility = signal<bool>(false);

  void addNewUser() async {
    if (controller.isValid()) {
      final user = UserModel(
        name: controller.name(),
        email: controller.email(),
        password: controller.password(),
        role: controller.role(),
      );
      final userApiService = UsersApiService();

      final createUser = await userApiService.createUser(user);
      if (createUser == null) {
        log('Error: Create user return null.');
        throw Exception('Error: Create user return null.');
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('User Page'),
      ),
      body: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 12),
        child: Center(
          child: Column(
            children: [
              TextField(
                onChanged: controller.name.set,
                decoration: InputDecoration(
                  label: const Text('Name'),
                  hintText: 'Your Name',
                  errorText: controller.nameError.watch(context),
                ),
              ),
              TextField(
                onChanged: controller.email.set,
                decoration: InputDecoration(
                  label: const Text('Email'),
                  hintText: 'your@email.address',
                  errorText: controller.emailError.watch(context),
                ),
              ),
              TextField(
                onChanged: controller.password.set,
                obscureText: !visibility(),
                decoration: InputDecoration(
                  label: const Text('Password'),
                  hintText: '****',
                  errorText: controller.passwordError.watch(context),
                  suffixIcon: IconButton(
                    icon: Icon(visibility.watch(context)
                        ? Icons.visibility
                        : Icons.visibility_off),
                    onPressed: () {
                      visibility.value = !visibility.value;
                    },
                  ),
                ),
              ),
              const SizedBox(height: 20),
              const Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  'User Ruler',
                  style: TextStyle(
                    fontSize: 12,
                  ),
                ),
              ),
              DropdownButtonHideUnderline(
                child: ButtonTheme(
                  alignedDropdown: true,
                  child: DropdownButton<Role>(
                    isExpanded: true,
                    value: controller.role.watch(context),
                    items: Role.values
                        .map(
                          (role) => DropdownMenuItem<Role>(
                            value: role,
                            child: Text(role.name),
                          ),
                        )
                        .toList(),
                    onChanged: (value) {
                      if (value != null) {
                        controller.role.set(value);
                      }
                    },
                  ),
                ),
              ),
              const SizedBox(height: 20),
              FilledButton(
                onPressed: addNewUser,
                child: const Text('Create User'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
