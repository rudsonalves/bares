# Projeto Bares

Este projeto é um exercício de criação de um sistema para gestão de atendimento em bares, Bares. A ideia é desenvolver um sistema para gerenciar o atendimento aos clientes por meio de aplicativos:

- Android/Web - para os clientes
- Android - para os garçons
- Android - para os gerentes de operação

O projeto vai contar com uma API Go para fazer o controle do acesso ao banco de dados e, no momento, os aplicativos Android para clientes, gerente e garçons.

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
| role      | ENUM         | Role (ex: cliente, garçom, gerente) |

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
