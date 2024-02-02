import 'package:routefly/routefly.dart';

import 'app/home/home_page.dart' as a0;
import 'app/splash/splash_page.dart' as a2;
import 'app/user/user_page.dart' as a1;

List<RouteEntity> get routes => [
  RouteEntity(
    key: '/home',
    uri: Uri.parse('/home'),
    routeBuilder: (ctx, settings) => Routefly.defaultRouteBuilder(
      ctx,
      settings,
      const a0.HomePage(),
    ),
  ),
  RouteEntity(
    key: '/user',
    uri: Uri.parse('/user'),
    routeBuilder: (ctx, settings) => Routefly.defaultRouteBuilder(
      ctx,
      settings,
      const a1.UserPage(),
    ),
  ),
  RouteEntity(
    key: '/splash',
    uri: Uri.parse('/splash'),
    routeBuilder: (ctx, settings) => Routefly.defaultRouteBuilder(
      ctx,
      settings,
      const a2.SplashPage(),
    ),
  ),
];

const routePaths = (
  path: '/',
  home: '/home',
  user: '/user',
  splash: '/splash',
);
