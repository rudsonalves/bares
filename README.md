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

### 1. **Tabela de Usuários (Usuarios)**

Para gerenciar os usuários que podem logar no app (gerentes, garçons).

| Campo     | Tipo         | Descrição                            |
| --------- | ------------ | ------------------------------------ |
| usuarioID | INT          | ID único para o usuário              |
| nome      | VARCHAR(255) | Nome do usuário                      |
| email     | VARCHAR(255) | Email do usuário                     |
| senhaHash | VARCHAR(255) | Hash da senha para autenticação      |
| papel     | ENUM         | Papel (ex: cliente, garçom, gerente) |

### 2. **Tabela de Itens do Menu (ItensMenu)**

Para armazenar detalhes dos itens disponíveis para pedido.

| Campo     | Tipo          | Descrição                    |
| --------- | ------------- | ---------------------------- |
| itemID    | INT           | ID único para o item do menu |
| nome      | VARCHAR(255)  | Nome do item                 |
| descricao | TEXT          | Descrição do item            |
| preco     | DECIMAL(10,2) | Preço do item                |
| imagemURL | VARCHAR(255)  | URL da imagem do item        |

### 3. **Tabela de Pedidos (Pedidos)**

Para armazenar os pedidos realizados pelos clientes.

| Campo     | Tipo     | Descrição                                                     |
| --------- | -------- | ------------------------------------------------------------- |
| pedidoID  | INT      | ID único para o pedido                                        |
| usuarioID | INT      | ID do usuário que fez o pedido                                |
| dataHora  | DATETIME | Data e hora do pedido                                         |
| status    | ENUM     | Status do pedido (ex: recebido, preparando, pronto, entregue) |

### 4. **Tabela de Itens do Pedido (ItensPedido)**

Para conectar os pedidos aos itens do menu e armazenar informações específicas do pedido, como a quantidade de cada item.

| Campo        | Tipo | Descrição                          |
| ------------ | ---- | ---------------------------------- |
| itemPedidoID | INT  | ID único para o item do pedido     |
| pedidoID     | INT  | ID do pedido                       |
| itemID       | INT  | ID do item do menu                 |
| quantidade   | INT  | Quantidade do item pedido          |
| observacoes  | TEXT | Observações específicas do cliente |

### Relacionamentos:

- **Usuarios** ↔ **Pedidos:** Um usuário pode fazer vários pedidos, mas cada pedido é feito por um único usuário.
- **Pedidos** ↔ **ItensPedido:** Um pedido pode conter vários itens, e um item pode aparecer em vários pedidos.
- **ItensMenu** ↔ **ItensPedido:** Um item do menu pode ser parte de vários pedidos, e cada item do pedido se refere a um item do menu.

### Considerações Finais:

- **Chaves Primárias:** Cada tabela deve ter uma chave primária (`usuarioID`, `itemID`, `pedidoID`, `itemPedidoID`).
- **Chaves Estrangeiras:** Usar chaves estrangeiras para manter a integridade referencial (ex: `usuarioID` em `Pedidos` refere-se a `usuarioID` em `Usuarios`; `itemID` em `ItensPedido` refere-se a `itemID` em `ItensMenu`; `pedidoID` em `ItensPedido` refere-se a `pedidoID` em `Pedidos`).
- **Indexação:** Considere adicionar índices para colunas frequentemente pesquisadas para melhorar a performance.

Com este design, você terá um banco de dados robusto e bem estruturado, pronto para gerenciar os usuários, pedidos e itens do menu do seu aplicativo de bar.

### Resumo da árvore da API REST:

Esta é a árvode de módulos da Bares API:

```
bares_api$ tree .
.
├── go.mod
├── go.sum
├── main.go
├── handlers
│   ├── auth.go
│   ├── item_menu_handler.go
│   ├── item_pedido_handler.go
│   ├── middleware.go
│   ├── pedido_handler.go
│   └── usuario_handler.go
├── models
│   ├── credentials.go
│   ├── item_menu.go
│   ├── item_pedido.go
│   ├── pedido.go
│   └── usuario.go
├── services
│   ├── auth_service.go
│   ├── item_menu_service.go
│   ├── item_pedido_service.go
│   ├── pedido_service.go
│   └── usuario_service.go
├── store
│   ├── database.go
│   ├── item_menu_store.go
│   ├── item_pedido_store.go
│   ├── pedido_store.go
│   └── usuario_store.go
└── utils
    └── utils.go
```

1. **handlers**: Este diretório contém os manipuladores HTTP para diferentes endpoints da API. Cada arquivo representa um manipulador relacionado a uma entidade específica, como pedidos, itens do menu e usuários. Também inclui middleware para tratamento de autenticação e autorização.

2. **models**: Aqui, você encontra os modelos de dados que representam as entidades do sistema, como credenciais, itens de menu, itens de pedido, pedidos e usuários. Esses modelos são usados para mapear dados entre a aplicação e o banco de dados.

3. **services**: O diretório de serviços contém lógica de negócios relacionada a cada entidade. Cada serviço corresponde a uma entidade específica, como autenticação, itens do menu, itens do pedido, pedidos e usuários. Essa camada é responsável por fornecer funcionalidades de alto nível para manipulação de dados.

4. **store**: Aqui estão os pacotes relacionados ao armazenamento de dados, incluindo integração com o banco de dados.

5. **utils**: Este diretório contém utilitários gerais que podem ser compartilhados em todo o projeto, como funções auxiliares, structs e constantes úteis.

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

### 2024/02/01 version 0.3.0:

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


