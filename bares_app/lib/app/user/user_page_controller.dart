import 'package:signals/signals.dart';

import '../../common/models/user_model.dart';

class UserPageController {
  final name = signal<String>('');
  final nameError = signal<String?>(null);
  final email = signal<String>('');
  final emailError = signal<String?>(null);
  final password = signal<String>('');
  final passwordError = signal<String?>(null);
  final role = signal<Role>(Role.cliente);

  validateName() {
    if (name().length < 4) {
      nameError.value = 'Name must have at least 3 characters.';
    } else {
      nameError.value = null;
    }
  }

  validateEmail() {
    if (!email().contains('@')) {
      emailError.value = 'Enter a valid email.';
    } else {
      emailError.value = null;
    }
  }

  validatePassword() {
    if (password().length < 4) {
      passwordError.value = 'Password must have at least 4 characters.';
    } else {
      passwordError.value = null;
    }
  }

  bool isValid() {
    validateName();
    validateEmail();
    validatePassword();

    return nameError() == null &&
        emailError() == null &&
        passwordError() == null;
  }
}
