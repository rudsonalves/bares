import 'dart:convert';
import 'dart:developer';

import 'package:http/http.dart' as http;

import '../common/constants/app_const.dart';
import '../common/models/menu_item_model.dart';
import 'functions.dart';

class MenuItemApiCreatedException implements Exception {
  final String message;
  MenuItemApiCreatedException(this.message);

  @override
  String toString() => "MenuItemApiCreatedException: $message";
}

class MenuItemApiGettedException implements Exception {
  final String message;
  MenuItemApiGettedException(this.message);

  @override
  String toString() => "MenuItemApiGettedException: $message";
}

class MenuItemApiUpdateException implements Exception {
  final String message;

  MenuItemApiUpdateException(this.message);

  @override
  String toString() => "MenuItemApiUpdateException: $message";
}

class MenuItemApiDeleteException implements Exception {
  final String message;

  MenuItemApiDeleteException(this.message);

  @override
  String toString() => "MenuItemApiDeleteException: $message";
}

class MenuItemApiService {
  MenuItemApiService._();

  // createMenuItem create a new menuItem
  static Future<MenuItemModel?> createMenuItem(MenuItemModel menuItem) async {
    String token = await getToken();
    final url = Uri.parse('${AppConst.apiUrlApi}/menuitem');

    try {
      final response = await http.post(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
        body: menuItem.toJson(),
      );

      if (response.statusCode == 201) {
        return MenuItemModel.fromJson(response.body);
      } else {
        final message =
            'Error: Failed to create menuItem: ${response.statusCode}';
        log(message);
        throw MenuItemApiCreatedException(message);
      }
    } on MenuItemApiCreatedException {
      rethrow;
    } catch (err) {
      final message =
          'Error: Handle connection errors, timeouts, ...: ${err.toString()}';
      log(message);
      throw Exception(message);
    }
  }

  // getAllMenuItems return all menuItems in system
  static Future<List<MenuItemModel>> getAllMenuItems() async {
    String token = await getToken();
    final url = Uri.parse('${AppConst.apiUrlApi}/menuitem');

    try {
      final response = await http.get(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );

      if (response.statusCode == 200) {
        final List<dynamic> menuItemsJson = json.decode(response.body);
        return menuItemsJson
            .map((menuItem) => MenuItemModel.fromMap(menuItem))
            .toList();
      } else {
        final message = 'Failed to get menu items: ${response.statusCode}';
        log(message);
        throw MenuItemApiGettedException(message);
      }
    } on MenuItemApiGettedException {
      rethrow;
    } catch (err) {
      final message =
          'Error: Handle connection errors, timeouts, ...: ${err.toString()}';
      log(message);
      throw Exception(message);
    }
  }

  // updateMenuItem update user informations, less password
  static Future<MenuItemModel?> updateMenuItem(MenuItemModel menuItem) async {
    if (menuItem.id == null) {
      const message = 'Error: User.id is null';
      log(message);
      throw MenuItemApiUpdateException(message);
    }

    String token = await getToken();
    final url = Uri.parse('${AppConst.apiUrlApi}/menuitem/${menuItem.id}');

    try {
      final response = await http.put(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
        body: menuItem.toJson(),
      );

      if (response.statusCode == 200) {
        return MenuItemModel.fromJson(response.body);
      } else {
        final message =
            'Error: Failed to update menuItem: ${response.statusCode}';
        log(message);
        throw MenuItemApiUpdateException(message);
      }
    } on MenuItemApiUpdateException {
      rethrow;
    } catch (err) {
      final message =
          'Error: Handle connection errors, timeouts, ...: ${err.toString()}';
      log(message);
      throw Exception(message);
    }
  }

  // deleteMenuItem delete an menuItem by id
  static Future<void> deleteMenuItem(int menuItemId) async {
    String token = await getToken();
    final url = Uri.parse('${AppConst.apiUrlApi}/menuitem/$menuItemId');

    try {
      final response = await http.delete(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );

      if (response.statusCode == 200) {
        return;
      } else {
        final message =
            'Error: Failed to delete menuItem: ${response.statusCode}';
        log(message);
        throw MenuItemApiUpdateException(message);
      }
    } on MenuItemApiUpdateException {
      rethrow;
    } catch (err) {
      final message =
          'Error: Handle connection errors, timeouts, ...: ${err.toString()}';
      log(message);
      throw Exception(message);
    }
  }
}
