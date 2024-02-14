# Projeto Bares

Este projeto é um exercício de criação de um sistema para gestão de atendimento em bares, Bares. A ideia é desenvolver um sistema para gerenciar de atendimento para bares, centrado no auto atendimento dos clientes por meio de aplicativos flutter:

- Android/Web - para os clientes
- Android - para os garçons e gerente de operação

O projeto vai contar com uma API REST em GoLang para fazer o controle do acesso ao banco de dados e, no momento, os aplicativos Android para clientes, gerente e garçons.

## 1. **API em Go**

- **Objetivo:** Criar uma API robusta e eficiente para gerenciar os pedidos e a comunicação entre os aplicativos.
- **Funcionalidades Básicas:**
  - **Cadastro de Pedidos:** Receber e armazenar pedidos dos clientes.
  - **Atualização de Status:** Permitir que a cozinha atualize o status dos pedidos.
  - **Consulta de Pedidos:** Capacidade para os aplicativos consultarem o status atual dos pedidos.

## 2. **Aplicativos em Flutter**

- **App para Clientes (Mesa):** Permitir que os clientes vejam o menu, façam pedidos e acompanhem o status dos pedidos.
- **App para Cozinha/Garçons:** Visualizar os pedidos recebidos, atualizar o status do pedido e receber notificações de novos pedidos.

## Desenvolvimento

### API em Go

- **Endpoints:** Definir endpoints claros para criação, atualização e consulta de pedidos.
- **Banco de Dados:** Escolher um banco de dados adequado para armazenar os pedidos (por exemplo, PostgreSQL, MySQL, etc.).
- **Testes:** Implementar testes unitários para garantir a confiabilidade da API.

### Aplicativos em Flutter

- **UI/UX:** Desenvolver interfaces intuitivas e responsivas.
- **Integração com a API:** Usar pacotes como `http` para se comunicar com a API.
- **Estado do Aplicativo:** Gerenciar o estado do aplicativo para uma atualização eficiente dos dados (pode-se usar `Provider`, `Bloc` ou `Riverpod`).

## Considerações Finais

1. **Foco no MVP (Produto Mínimo Viável):** Concentre-se nas funcionalidades essenciais para fazer o sistema funcionar de forma básica.
2. **Documentação:** Mantenha uma documentação clara tanto para a API quanto para os aplicativos.
3. **Iteração e Feedback:** Teste o sistema com usuários fictícios e refine com base no feedback.

Este projeto simplificado ainda oferece uma excelente oportunidade para explorar o desenvolvimento full-stack, abrangendo tanto o backend com Go quanto o front-end com Flutter. É uma maneira eficaz de praticar habilidades importantes de programação e design de sistemas.

## Estrutura da API

### 1. **Definição de Rotas**

A API terá as seguintes rotas principais:

- **POST /pedidos:** Recebe novos pedidos dos clientes.
- **GET /pedidos/{id}:** Permite a consulta de um pedido específico.
- **PUT /pedidos/{id}:** Permite a atualização do status de um pedido específico (por exemplo, de "em preparo" para "pronto").
- **GET /pedidos:** Lista todos os pedidos, possivelmente com filtros por status.

### 2. **Modelagem dos Dados**

Os dados principais a serem gerenciados serão os pedidos. Cada pedido pode ter a seguinte estrutura:

- **ID:** Identificador único do pedido.
- **Itens:** Lista de itens pedidos, cada um com nome, quantidade e possíveis observações.
- **Status:** Status atual do pedido (por exemplo, "recebido", "em preparo", "pronto", "entregue").
- **Timestamps:** Marcas de tempo para quando o pedido foi criado e atualizado.

### 3. **Autenticação e Autorização (Opcional para Versão Inicial)**

- Implementar autenticação para garantir que apenas usuários autorizados possam fazer pedidos ou alterar o status dos pedidos.
- JWT (JSON Web Tokens) ou OAuth podem ser usados para gerenciar sessões e permissões.

### 4. **Comunicação com o Banco de Dados**

- Integrar um banco de dados para armazenar os pedidos e as informações relacionadas.
- Implementar um ORM (Object-Relational Mapping) como GORM para interagir de maneira mais segura e eficiente com o banco de dados.

## Desenvolvimento e Ferramentas

1. **Gin-Gonic para Rotas e Middleware:**
   
   - Usar o framework Gin-Gonic para facilitar a definição de rotas, tratamento de requisições e respostas e a integração de middleware para tarefas como logging e tratamento de erros.

2. **Validação de Dados:**
   
   - Assegurar que os dados recebidos nas requisições estão corretos e completos, usando pacotes de validação.

3. **Testes:**
   
   - Implementar testes automatizados para garantir que a API está funcionando como esperado. Isso inclui testes unitários para funções individuais e testes de integração para o fluxo completo de pedidos.

4. **Documentação da API:**
   
   - Utilizar ferramentas como Swagger ou Postman para documentar as rotas, parâmetros e respostas da API, facilitando o entendimento e a utilização por parte dos desenvolvedores dos aplicativos cliente e cozinha.

5. **Logging e Monitoramento:**
   
   - Implementar logging adequado para acompanhar o que está acontecendo na aplicação e integrar ferramentas de monitoramento para observar a saúde e o desempenho da API em tempo real.

Ao seguir esses passos, você estará no caminho certo para desenvolver uma API robusta e escalável para o seu aplicativo de bar, proporcionando uma base sólida para a interação entre os clientes e a cozinha.

# Banco de dados

Vamos projetar o esquema do banco de dados para o seu aplicativo de bar, considerando as funcionalidades básicas de gerenciamento de pedidos, itens do menu e usuários. O banco de dados será estruturado em tabelas que refletem os principais componentes do sistema.

### 1. **Tabela de Usuários (UsersTable)**

Para gerenciar os usuários que podem logar no app (gerentes, garçons).

| Campo      | Tipo                | Descrição                                 |
| ---------- | ------------------- | ----------------------------------------- |
| id         | INT                 | ID único para o usuário                   |
| name       | VARCHAR(255)        | Nome do usuário                           |
| email      | VARCHAR(255) UNIQUE | Email do usuário                          |
| pasworHash | VARCHAR(255)        | Hash da senha para autenticação           |
| role       | ENUM                | Role (ex: cliente, garçom, gerente, ...) |

### 2. **Tabela de Itens do Menu (MenuItemsTable)**

Para armazenar detalhes dos itens disponíveis para pedido.

| Campo       | Tipo                | Descrição                    |
| ----------- | ------------------- | ---------------------------- |
| id          | INT                 | ID único para o item do menu |
| name        | VARCHAR(255) UNIQUE | Nome do item                 |
| description | TEXT                | Descrição do item            |
| price       | DECIMAL(10,2)       | Preço do item                |
| imagemURL   | VARCHAR(255)        | URL da imagem do item        |

### 3. **Tabela de Pedidos (OrdersTable)**

Para armazenar os pedidos realizados pelos clientes.

| Campo     | Tipo     | Descrição                                                     |
| --------- | -------- | ------------------------------------------------------------- |
| id        | INT      | ID único para o pedido                                        |
| userId    | INT      | ID do usuário que fez o pedido                                |
| dateHour  | DATETIME | Data e hora do pedido                                         |
| status    | ENUM     | Status do pedido (ex: recebido, preparando, pronto, entregue) |

### 4. **Tabela de Itens do Pedido (ItemsOrderTable)**

Para conectar os pedidos aos itens do menu e armazenar informações específicas do pedido, como a quantidade de cada item.

| Campo     | Tipo | Descrição                          |
| --------- | ---- | ---------------------------------- |
| id        | INT  | ID único para o item do pedido     |
| OrderID   | INT  | ID do pedido                       |
| itemID    | INT  | ID do item do menu                 |
| amount    | INT  | Quantidade do item pedido          |
| comments  | TEXT | Observações específicas do cliente |

### Relacionamentos:

- **Users** ↔ **Orders:** Um usuário pode fazer vários pedidos, mas cada pedido é feito por um único usuário.
- **Orders** ↔ **ItemsOrder:** Um pedido pode conter vários itens, e um item pode aparecer em vários pedidos.
- **MenuItems** ↔ **ItemsOrder:** Um item do menu pode ser parte de vários pedidos, e cada item do pedido se refere a um item do menu.

### Considerações Finais:

- **Chaves Primárias:** Cada tabela deve ter uma chave primária (`id`).
- **Chaves Estrangeiras:** Usar chaves estrangeiras para manter a integridade referencial.
- **Indexação:** Considere adicionar índices para colunas frequentemente pesquisadas para melhorar a performance.

Com este design, você terá um banco de dados robusto e bem estruturado, pronto para gerenciar os usuários, pedidos e itens do menu do seu aplicativo de bar.

### Resumo da árvore da API REST:

Esta é a árvode de módulos da Bares API:

```
api/
├── go.mod
├── go.sum
├── main.go
├── bootstrap
│   └── bootstrap.go
├── handlers
│   ├── auth.go
│   ├── item_order_handler.go
│   ├── menu_item_handler.go
│   ├── middleware.go
│   ├── order_handler.go
│   ├── user_handler.go
│   └── integration
│       ├── common_integration.go
│       └── user_handler_integration_test.go
├── models
│   ├── credentials.go
│   ├── item_order.go
│   ├── menu_item.go
│   ├── order.go
│   └── user.go
├── services
│   ├── auth_service.go
│   ├── item_order_service.go
│   ├── menu_item_service.go
│   ├── order_service.go
│   └── user_service.go
├── store
│   ├── database.go
│   ├── item_order_store.go
│   ├── menu_item_store.go
│   ├── order_store.go
│   ├── user_store.go
│   ├── integration
│   │   ├── common_integration.go
│   │   ├── database_integration_test.go
│   │   ├── menu_item_store_integration_test.go
│   │   ├── order_integration_test.go
│   │   └── user_store_integration_test.go
│   └── store_test
│       ├── database_test.go
│       ├── item_order_store_test.go
│       ├── menu_item_store_test.go
│       └── user_store_test.go
└── utils
    └── utils.go
```

1. **handlers/**: Este diretório contém os manipuladores HTTP para diferentes endpoints da API. Cada arquivo representa um manipulador relacionado a uma entidade específica, como pedidos, itens do menu e usuários. Também inclui middleware para tratamento de autenticação e autorização.

2. **models/**: Aqui, você encontra os modelos de dados que representam as entidades do sistema, como credenciais, itens de menu, itens de pedido, pedidos e usuários. Esses modelos são usados para mapear dados entre a aplicação e o banco de dados.

3. **services/**: O diretório de serviços contém lógica de negócios relacionada a cada entidade. Cada serviço corresponde a uma entidade específica, como autenticação, itens do menu, itens do pedido, pedidos e usuários. Essa camada é responsável por fornecer funcionalidades de alto nível para manipulação de dados.

4. **store/**: Aqui estão os pacotes relacionados ao armazenamento de dados, incluindo integração com o banco de dados.

5. **utils/**: Este diretório contém utilitários gerais que podem ser compartilhados em todo o projeto, como funções auxiliares, structs e constantes úteis.

6. **main.go**: Ponto de entrada da aplicação. Ele contém a configuração da API, o roteamento das rotas HTTP e inicia o servidor web.

7. **go.mod e go.sum**: Esses arquivos são usados para gerenciar as dependências do projeto e garantir que as versões corretas dos pacotes sejam usadas.

## Changelog

### 2023/02/09 - version 0.2.1:

This commit introduces several adjustments and enhancements to the API and the Flutter app, focusing on authorization mechanisms and user experience improvements. Here’s a condensed overview of the key changes implemented:

- **API Authorization Adjustments**:
  - **User Information in Authentication Response**: The authentication response now includes user information, enriching the client-side data available post-authentication.
  - **Authorization Method**: A new method `isAuthorized(userRole models.Role, path, method string) bool` has been introduced to restrict API route access based on user roles. This preliminary implementation adds essential functionality to the authorization process.
  - **User Creation Restrictions**: Adjustments have been made to limit the ability of 'waiter' role users to only create 'customer' role users, enhancing the permission structure within the application.
  - **Private User Creation Route**: The route for creating new users (`/users`) has been moved to private routes to ensure only authorized users can access this functionality.

- **Flutter App Enhancements**:
  - **Theme and Login Page**: Added a button to switch the app theme and enhance the login page functionality, including a pre-logout feature for testing.
  - **Login Page and Controller**: The login page is now fully functional, with the addition of a controller using `signal` for managing email, email error, and password visibility, including a method for performing the login process and disposing of signals.
  - **App Constants and User Model Adjustments**: Introduced a set of constants for the app and made adjustments to the user model, including a method for copying user data.
  - **Message Dialog Standardization**: Added a method for displaying standardized message dialogs across the app.
  - **App Configuration and Theme Control**: Integrated token control through `SecureStorageManager`, providing storage for the logged-in user and theme mode, and added theme control to the MaterialApp.
  - **Secure Storage for Token**: Implemented a service for storing the token securely using `flutter_secure_storage`.
  - **API URL Configuration**: The API URL has been moved to `AppConst.apiURL` for better management.
  - **Dependency Addition**: Added the `flutter_secure_storage` package to manage login data and app control securely.

This commit not only refines the security and authorization aspects of the API but also significantly improves the user interface and experience within the Flutter app through thoughtful enhancements and the integration of secure storage solutions.

## Change Log 

### 2024/02/14 - version: 0.2.2

This commit further advances the API and Flutter app by enhancing English translations for comments, logs, and messages, improving user management features, and adding new images for the dashboard. Here's a summary of the significant updates:

- **API Enhancements**:
  - **Password Validation**: The `CheckAndCreateAdminUser` function now utilizes `EvaluatePasswordStrength` for password validation.
  - **Logging Improvements**: New log messages have been added across various handlers (`auth.go`, `item_order_handler.go`, `menu_item_handler.go`, `order_handler.go`) to better track API execution and error evaluations.
  - **User Management**: The `user_handler.go` introduces `GetAllUsers` for listing system users and `UpdateUserPass` for password updates, alongside log message enhancements.
  - **Database Connection Test**: The database now tests the connection with `db.Ping()` to ensure reliability.
  - **SQL Constants and Methods**: New constants and methods for user retrieval and password updates enhance user store operations, avoiding password manipulation in user updates.
  - **Utility Functions**: Separation of `CreateDataBaseConn` into `database_functions.go` and renaming `utils.go` to `generic_functions.go`, including improvements like `GenerateRandomPassword` utilizing `EvaluatePasswordStrength`.

- **Flutter App Updates**:
  - **Dashboard Images**: New images have been added to `flutter_app/assets/images/` for the dashboard, enhancing the UI.
  - **Navigation and Error Handling**: Adjustments in `app_page.dart` and `login_controller.dart` improve app navigation and clarify login error messages.
  - **User Editing**: The `edit_page.dart` and `edit_controller.dart` receive updates for more intuitive user management, including a method for initializing user data and adjusting validation checks.
  - **User Listing and Management**: The `users_page.dart` now displays a user list for management actions, supported by enhanced user model functions and API service exceptions for clearer error handling.

This commit significantly enhances both the backend API and the Flutter frontend, focusing on user management, error handling, and overall usability improvements, alongside the introduction of new visual elements for a more engaging user interface.


### 2024/02/10 - version 0.2.1

This commit introduces several adjustments and enhancements to the API and the Flutter app, focusing on authorization mechanisms and user experience improvements. Here’s a condensed overview of the key changes implemented:

- **API Authorization Adjustments**:
  - **User Information in Authentication Response**: The authentication response now includes user information, enriching the client-side data available post-authentication.
  - **Authorization Method**: A new method `isAuthorized(userRole models.Role, path, method string) bool` has been introduced to restrict API route access based on user roles. This preliminary implementation adds essential functionality to the authorization process.
  - **User Creation Restrictions**: Adjustments have been made to limit the ability of 'waiter' role users to only create 'customer' role users, enhancing the permission structure within the application.
  - **Private User Creation Route**: The route for creating new users (`/users`) has been moved to private routes to ensure only authorized users can access this functionality.

- **Flutter App Enhancements**:
  - **Theme and Login Page**: Added a button to switch the app theme and enhance the login page functionality, including a pre-logout feature for testing.
  - **Login Page and Controller**: The login page is now fully functional, with the addition of a controller using `signal` for managing email, email error, and password visibility, including a method for performing the login process and disposing of signals.
  - **App Constants and User Model Adjustments**: Introduced a set of constants for the app and made adjustments to the user model, including a method for copying user data.
  - **Message Dialog Standardization**: Added a method for displaying standardized message dialogs across the app.
  - **App Configuration and Theme Control**: Integrated token control through `SecureStorageManager`, providing storage for the logged-in user and theme mode, and added theme control to the MaterialApp.
  - **Secure Storage for Token**: Implemented a service for storing the token securely using `flutter_secure_storage`.
  - **API URL Configuration**: The API URL has been moved to `AppConst.apiURL` for better management.
  - **Dependency Addition**: Added the `flutter_secure_storage` package to manage login data and app control securely.

This commit not only refines the security and authorization aspects of the API but also significantly improves the user interface and experience within the Flutter app through thoughtful enhancements and the integration of secure storage solutions.


### 2024/02/07 - version: 0.2.0

In this commit, we have conducted extensive renaming to standardize the Go API base (`bares_api/*`) in English, while also initiating the implementation of the Flutter app, adopting new technologies such as `signal` for advanced reactivity in widgets, and `RouteFly` to optimize automatic route generation based on directory structure, marking a significant advancement in the project's usability and development. Below are more details on the changes made in this commit:

* **bares_api/bootstrap/bootstrap.go**:
  - Adds the function `CheckAndCreateAdminUser(userService *services.UserService) error` to check if an Administrator Role user exists in the system. If not, a new admin user will be created;
  - Added a local function `func isPasswordStrong(password string) bool` to verify password strength. Currently, the password must be at least 8 characters and include uppercase letters, lowercase letters, and numbers.

* **bares_api/handlers/integration/common_integration.go**:
  - Added the function `UsersGenerate() []models.User` to return a list of users;
  - `MenuItemsGenerate` and `OrdersGenerate` are functions that generate `MenuItems` and `Orders`, respectively. They are not being used at the moment.

* **bares_api/handlers/integration/user_handler_integration_test.go**:
* **bares_api/handlers/user_handler.go**:
* **bares_api/models/credentials.go**:
* **bares_api/models/item_order.go**:
* **bares_api/models/menu_item.go**:
* **bares_api/models/order.go**:
* **bares_api/services/auth_service.go**:
* **bares_api/services/menu_item_service.go**:
* **bares_api/store/database.go**:
* **bares_api/store/integration/common_integration.go**:
* **bares_api/store/integration/menu_item_store_integration_test.go**:
* **bares_api/store/integration/order_integration_test.go**:
* **bares_api/store/integration/user_store_integration_test.go**:
* **bares_api/store/item_order_store.go**:
* **bares_api/store/menu_item_store.go**:
* **bares_api/store/order_store.go**:
* **bares_api/store/store_test/database_test.go**:
* **bares_api/store/store_test/item_order_store_test.go**:
* **bares_api/store/store_test/menu_item_store_test.go**:
* **bares_api/store/store_test/user_store_test.go**:
  - Adjusted so that names of attributes, functions, structs, etc., are in English.

* **bares_api/main.go**:
  - In addition to renaming attributes, functions, etc., this package now uses `utils.CreateConnString` to generate the connection string.

* **bares_api/models/user.go**:
  - In addition to renaming attributes, functions, etc., a `String()` method has been added to the `User` struct.

* **bares_api/services/user_service.go**:
  - In addition to renaming attributes, functions, etc., the `UserServiceInterface` interface has been added to the `&UserService`;
  - Added the method `CheckIfAdminExists() (bool, error)` to check for the existence of an admin user.

* **bares_api/store/user_store.go**:
  - In addition to renaming attributes, functions, etc., the method `GetUsersByRole(role models.Role) ([]*models.User, error)` has been added to load users by role (`Role`).

* **bares_api/utils/utils.go**:
  - The method `CreateConnString(dbName string) string` has been transferred to the `utils` package.

* **bares_app/lib/app/app_page.dart**:
* **bares_app/lib/app/dashboard/dashboard_page.dart**:
* **bares_app/lib/app/login/login_page.dart**:
* **bares_app/lib/app/splash/splash_page.dart**:
  - Added a simple page, still without functionalities.

* **bares_app/lib/app/user/user_page.dart**:
* **bares_app/lib/app/user/user_page_controller.dart**:
  - Added a page for user registration and editing.
  - Added a controller to manage the form elements of the user page, using `signal` for both reactivity and field validation.

* **bares_app/lib/common/models/item_order_model.dart**:
* **bares_app/lib/common/models/menu_item_model.dart**:
* **bares_app/lib/common/models/order_model.dart**:
* **bares_app/lib/common/models/user_model.dart**:
  - Added models for database elements;
  - The enums for `Role` and `Status` were enhanced with the addition of standard enum extensions to add methods:
    - `static EnumName fromString(String)` - to convert a string into the corresponding enum;
    - `static List<String> stringList()` - to return a list of strings with the names of the enums

.

* **bares_app/lib/common/theme/color_schemes.dart**:
  - Added a green-based theme for the Flutter app.

* **bares_app/lib/material_app.dart**:
  - Applied `RouteFlay` to create automatic routing of app pages. This routing is done automatically based on the pages placed in the `lib/app` folder.

* **bares_app/lib/services/users_api_service.dart**:
  - Created the `UsersApiService` class to encapsulate the API part related to user manipulation. Currently, only the method `createUser(UserModel user)` has been implemented.


### 2024/02/01 version 0.1.0:

Nesta versão, fiz várias alterações e acabei não fazendo o commit no momento correto. No final, praticamente todo o código teve alguma modificação e adicionei apenas algumas partes das mudanças abaixo. No entanto, as partes principais foram a adição de um sistema de autenticação com cookies, a padronização das mensagens de erro e logs nas camadas mais internas do aplicativo.

* bares_api/go.mod:
  - github.com/DATA-DOG/go-sqlmock v1.5.2: Simular a interação do banco de dados para testes de integração e unitários.
  - github.com/dgrijalva/jwt-go v3.2.0+incompatible: Para lidar com JSON Web Tokens (JWT) e gerenciamento de cookies na autenticação.
  - github.com/go-sql-driver/mysql v1.7.1: Suporte ao MySQL.
  - github.com/gorilla/mux v1.8.1: Gorilla Mux para roteamento HTTP.
  - github.com/stretchr/testify v1.8.4: Biblioteca de suporte a testes.
  - golang.org/x/crypto v0.18.0: Usado na criptografia de senhas na autenticação e armazenamento de senhas.
  - golang.org/x/term v0.16.0:
  - github.com/davecgh/go-spew v1.1.1:
  - golang.org/x/sys:
  - github.com/pmezard/go-difflib v1.0.0: Estes últimos são de suporte geral para o aplicativo.

* bares_api/handlers/auth.go:
  - Adição da struct AuthHandler para gerar o handler de autenticação.
  - NewAuthHandler cria uma nova instância de AuthHandler.
  - Método func (handler *AuthHandler) LoginHandlers(w http.ResponseWriter, r *http.Request) para autenticar um usuário no sistema.

* bares_api/handlers/auth.go:
  - Implementa o serviço de autenticação (AuthService).

* bares_api/handlers/middleware.go:
  - Implementa um middleware de autenticação (AuthMiddleware) para a API.


