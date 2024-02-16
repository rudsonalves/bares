import 'dart:developer';

import 'package:signals/signals_flutter.dart';

import '../../common/models/menu_item_model.dart';
import '../../services/menu_item_api_service.dart';

enum PageStatus {
  pageStateInitial,
  pageStateLoading,
  pageStateSuccess,
  pageStateError,
}

class MenuController {
  final _menuItemList = <MenuItemModel>[];
  final state = signal<PageStatus>(PageStatus.pageStateInitial);

  List<MenuItemModel> get menuItemList => _menuItemList;

  void init() {
    _getMenuItemsList();
  }

  void dispose() {
    state.dispose();
  }

  Future<void> _getMenuItemsList() async {
    try {
      state.value = PageStatus.pageStateLoading;
      _menuItemList.clear();
      await Future.delayed(const Duration(milliseconds: 1000));
      final menuItems = await MenuItemApiService.getAllMenuItems();
      _menuItemList.addAll(menuItems);
      state.value = PageStatus.pageStateSuccess;
    } catch (err) {
      log('Erro: $err');
      state.value = PageStatus.pageStateError;
    }
  }

  Future<void> addUser() async {
    _getMenuItemsList();
  }

  Future<void> updateUser() async {
    _getMenuItemsList();
  }
}
