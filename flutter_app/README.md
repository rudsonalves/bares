# bares_app

A new Flutter project.

# Changelog

## 2024/02/07 - version: 0.2.0

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