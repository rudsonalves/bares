|                                   |      |                         |        | Role    |        |         |       |         |       |
| --------------------------------- |:----:| ----------------------- | ------ | ------- | ------ | ------- | ----- | ------- | ----- |
| Operação                          | Done | Path                    | Método | cliente | garcom | gerente | admin | cozinha | caixa |
| userHandler.CreateUser            | x    | "/users"                | POST   |         | x      | x       | x     |         | x     |
| itemMenuHandler.CreateMenuItem    | x    | "/menuitem"             | POST   |         |        | x       | x     |         |       |
| orderHandler.CreateOrder          | x    | "/order"                | POST   | x       | x      | x       | x     |         |       |
| itemOrderHandler.CreateItemOrder  | x    | "/itemorder"            | POST   | x       | x      | x       | x     |         |       |
| userHandler.GetUser               | x    | "/users/{id}"           | GET    |         |        | x       | x     |         | x     |
| userHandler.GetAllUsers           | x    | "/users"                | GET    |         |        | x       | x     |         | x     |
| itemMenuHandler.GetAllMenuItem    | x    | "/menuitem"             | GET    |         |        | x       | x     | x       |       |
| itemMenuHandler.GetMenuItem       | x    | "/menuitem/{id}"        | GET    |         |        | x       | x     | x       |       |
| itemMenuHandler.GetMenuItemByName | x    | "/menuitem/name/{name}" | GET    | x       | x      | x       | x     | x       |       |
| orderHandler.GetPendingOrder      | x    | "/order"                | GET    | x       | x      | x       | x     | x       | x     |
| orderHandler.GetOrder             | x    | "/order/{id}"           | GET    | x       | x      | x       | x     | x       | x     |
| orderHandler.GetOrderByUser       | x    | "/order/user/{id}"      | GET    | x       | x      | x       | x     | x       | x     |
| itemOrderHandler.GetIItemOrder    | x    | "/itemorder/{id}"       | GET    | x       | x      | x       | x     | x       | x     |
| userHandler.UpdateUser            | x    | "/users/{id}"           | PUT    |         |        | x       | x     |         |       |
| userHandler.UpdateUserPass        | x    | "/users/password/{id}"  | PUT    |         |        | x       | x     |         |       |
| itemMenuHandler.UpdateMenuItem    | x    | "/menuitem/{id}"        | PUT    |         |        | x       | x     | x       |       |
| orderHandler.UpdateOrder          | x    | "/order/{id}"           | PUT    |         | x      | x       | x     | x       | x     |
| itemOrderHandler.UpdateItemOrder  | x    | "/itemorder/{id}"       | PUT    |         | x      | x       | x     | x       | x     |
| userHandler.DeleteUser            | x    | "/users/{id}"           | DELETE |         |        | x       | x     |         |       |
| itemMenuHandler.DeleteMenuItem    | x    | "/menuitem/{id}"        | DELETE |         |        | x       | x     |         |       |
| orderHandler.DeleteOrder          | x    | "/order/{id}"           | DELETE |         | x      | x       | x     |         |       |
| itemOrderHandler.DeleteItemOrder  | x    | "/itemorder/{id}"       | DELETE |         | x      | x       | x     |         |       |
