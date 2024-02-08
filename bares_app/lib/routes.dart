import 'package:routefly/routefly.dart';

import 'app/app_page.dart' as a4;
import 'app/dashboard/dashboard_page.dart' as a3;
import 'app/login/login_page.dart' as a2;
import 'app/splash/splash_page.dart' as a1;
import 'app/user/user_page.dart' as a0;

List<RouteEntity> get routes => [
  RouteEntity(
    key: '/user',
    uri: Uri.parse('/user'),
    routeBuilder: (ctx, settings) => Routefly.defaultRouteBuilder(
      ctx,
      settings,
      const a0.UserPage(),
    ),
  ),
  RouteEntity(
    key: '/splash',
    uri: Uri.parse('/splash'),
    routeBuilder: (ctx, settings) => Routefly.defaultRouteBuilder(
      ctx,
      settings,
      const a1.SplashPage(),
    ),
  ),
  RouteEntity(
    key: '/login',
    uri: Uri.parse('/login'),
    routeBuilder: (ctx, settings) => Routefly.defaultRouteBuilder(
      ctx,
      settings,
      const a2.LoginPage(),
    ),
  ),
  RouteEntity(
    key: '/dashboard',
    uri: Uri.parse('/dashboard'),
    routeBuilder: (ctx, settings) => Routefly.defaultRouteBuilder(
      ctx,
      settings,
      const a3.DashboardPage(),
    ),
  ),
  RouteEntity(
    key: '/',
    uri: Uri.parse('/'),
    routeBuilder: (ctx, settings) => Routefly.defaultRouteBuilder(
      ctx,
      settings,
      const a4.AppPage(),
    ),
  ),
];

const routePaths = (
  path: '/',
  user: '/user',
  splash: '/splash',
  login: '/login',
  dashboard: '/dashboard',
);
