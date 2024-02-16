import '../models/user_model.dart';

class AppConst {
  static String serverIp = '192.168.0.22';
  static String apiUrl = 'http://$serverIp:8080';
  static String apiUrlApi = '$apiUrl/api';
  static String jwtToken = 'jwt_token';
}

const Map<Role, String> roleImage = {
  Role.admin: 'assets/images/admin.png',
  Role.gerente: 'assets/images/gerente.png',
  Role.cliente: 'assets/images/cliente.png',
  Role.cozinha: 'assets/images/cozinha.png',
  Role.garcom: 'assets/images/garcom.png',
  Role.caixa: 'assets/images/caixa.png',
};
