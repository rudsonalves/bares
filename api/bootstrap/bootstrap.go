package bootstrap

import (
	"bares_api/models"
	"bares_api/services"
	"bares_api/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/term"
)

// CheckAndCreateAdminUser checks if an Admin user exists. If it does not exist
// an admin user must be created.
func CheckAndCreateAdminUser(userService *services.UserService) error {
	exists, err := userService.CheckIfAdminExists()
	if err != nil {
		log.Printf("Error checking for administrators: %v", err)
		return err
	}

	if exists {
		fmt.Println("Admin user verified!")
		return nil
	}

	// Create admin user
	fmt.Printf("\n\n\n----------------------------------------------------------\n")
	fmt.Printf("\n\n\nNo administrators were found on this system.\n")
	fmt.Println("You must create an Administrator account to continue.")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("==================================")
		fmt.Println("=        Enter information       =")
		fmt.Println("=            user Admin          =")
		fmt.Println("==================================")

		var name string
		fmt.Printf("\nAdmin user name: ")
		name, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading the name, try again.")
			continue
		}
		name = utils.TrimSpaceLB(name)

		var email string
		fmt.Printf("Admin email adreess: ")
		email, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading email, please try again.")
			continue
		}
		email = utils.TrimSpaceLB(email)

		var password string
		var checkPassword string
		for {
			fmt.Println("Use a password of at least 8 characters, with letters and numbers")
			fmt.Printf("Enter a password: ")
			passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				fmt.Println("Error reading password, try again.")
				continue
			}
			password = string(passwordBytes)
			fmt.Println()

			fmt.Printf("Password check: ")
			passwordBytes, err = term.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				fmt.Println("Error reading password, try again.")
				continue
			}
			checkPassword = string(passwordBytes)
			fmt.Println()

			if password != checkPassword {
				fmt.Printf("\nMust-have password! Try again.\n\n")
				continue
			}

			passStrength := utils.EvaluatePasswordStrength(password)
			if passStrength.Score < 6 {
				fmt.Printf("\n\nWeak password! Use a strong password.\n")
				continue
			}

			break
		}

		user := models.User{
			Name:         name,
			Email:        email,
			PasswordHash: password,
			Role:         models.Admin,
		}
		fmt.Printf("\n\nCreate Admin User:\n")
		fmt.Println(user)
		fmt.Println("Confirm the creation of Admin (Y/n): ")

		var ans string
		ans, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading confirmation, please try again.")
			continue
		}
		ans = utils.TrimSpaceLB(ans)

		if strings.EqualFold(ans, "y") || len(ans) == 0 {
			err = userService.CreateUser(&user)
			if err != nil {
				log.Printf("Failed to create admin user: %v", err)
				return err
			}
			fmt.Println("Administrator user created successfully!")
			break
		} else if strings.EqualFold(ans, "n") {
			fmt.Println("Administrator user creation cancelled.")
			continue
		} else {
			fmt.Println("Invalid response. Please answer with 'Y' for yes or 'N' for no.")
		}
	}

	return nil
}
