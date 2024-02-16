import 'package:routefly/routefly.dart';

import 'app/app_page.dart' as a5;
import 'app/dashboard/dashboard_page.dart' as a1;
import 'app/login/login_page.dart' as a0;
import 'app/menu/menu_page.dart' as a4;
import 'app/users/edit/edit_page.dart' as a3;
import 'app/users/users_page.dart' as a2;

List<RouteEntity> get routes => [
  RouteEntity(
    key: '/login',
    uri: Uri.parse('/login'),
    routeBuilder: (ctx, settings) => Routefly.defaultRouteBuilder(
      ctx,
      settings,
      const a0.LoginPage(),
    ),
  ),
  RouteEntity(
    key: '/dashboard',
    uri: Uri.parse('/dashboard'),
    routeBuilder: (ctx, settings) => Routefly.defaultRouteBuilder(
      ctx,
      settings,
      const a1.DashboardPage(),
    ),
  ),
  RouteEntity(
    key: '/users',
    uri: Uri.parse('/users'),
    routeBuilder: (ctx, settings) => Routefly.defaultRouteBuilder(
      ctx,
      settings,
      const a2.UsersPage(),
    ),
  ),
  RouteEntity(
    key: '/users/edit',
    uri: Uri.parse('/users/edit'),
    routeBuilder: (ctx, settings) => Routefly.defaultRouteBuilder(
      ctx,
      settings,
      const a3.EditPage(),
    ),
  ),
  RouteEntity(
    key: '/menu',
    uri: Uri.parse('/menu'),
    routeBuilder: (ctx, settings) => Routefly.defaultRouteBuilder(
      ctx,
      settings,
      const a4.MenuPage(),
    ),
  ),
  RouteEntity(
    key: '/',
    uri: Uri.parse('/'),
    routeBuilder: (ctx, settings) => Routefly.defaultRouteBuilder(
      ctx,
      settings,
      const a5.AppPage(),
    ),
  ),
];

const routePaths = (
  path: '/',
  login: '/login',
  dashboard: '/dashboard',
  users: (
    path: '/users',
    edit: '/users/edit',
  ),
  menu: '/menu',
);
