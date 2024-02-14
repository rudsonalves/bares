import 'package:signals/signals.dart';

import '../../../common/models/user_model.dart';

class EditController {
  final name = signal<String>('');
  final nameError = signal<String?>(null);
  final email = signal<String>('');
  final emailError = signal<String?>(null);
  final password = signal<String>('');
  final passwordError = signal<String?>(null);
  final role = signal<Role>(Role.cliente);

  void init(UserModel user) {
    name.value = user.name;
    email.value = user.email;
    // password.value = user.password ?? '';
    role.value = user.role;
  }

  void validateName() {
    if (name().length < 4) {
      nameError.value = 'Name must have at least 3 characters.';
    } else {
      nameError.value = null;
    }
  }

  void validateEmail() {
    if (!email().contains('@')) {
      emailError.value = 'Enter a valid email.';
    } else {
      emailError.value = null;
    }
  }

  void validatePassword() {
    if (password().length < 4) {
      passwordError.value = 'Password must have at least 4 characters.';
    } else {
      passwordError.value = null;
    }
  }

  bool isValid([bool withPasswordCheck = true]) {
    validateName();
    validateEmail();

    if (withPasswordCheck) {
      validatePassword();
    } else {
      passwordError.value = null;
    }

    return nameError() == null &&
        emailError() == null &&
        passwordError() == null;
  }
}
