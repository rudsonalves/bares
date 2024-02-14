package utils

import (
	"bares_api/models"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net"
	"regexp"
	"strings"
)

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz[]_,.;:!@#$%&*-*/+"

const minPasswordSize = 6

var (
	emailRegex     = regexp.MustCompile(`^[a-zA-Z0-9.%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	ropeRegex      = regexp.MustCompile(`^mesa[0-9]{2,}@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	numberRegex    = regexp.MustCompile(`[0-9]+`)
	lowercaseRegex = regexp.MustCompile(`[a-z]+`)
	uppercaseRegex = regexp.MustCompile(`[A-Z]+`)
	symbolRegex    = regexp.MustCompile(`[^a-zA-Z0-9]+`)
)

// GenerateRandomPassword generates a random password of the specified length.
// The password can contain alphanumeric characters and special symbols.
// The 'length' parameter defines the length of the desired password.
// Returns the generated password and an error if it occurs.
func GenerateRandomPassword(length int) (string, error) {
	if length < minPasswordSize {
		return "", fmt.Errorf("length must be greater or equal than %d", minPasswordSize)
	}

	passwordBytes := make([]byte, length)
	password := ""
	result := EvaluatePasswordStrength(password)
	for result.Score < 6 {
		for i := 0; i < length; i++ {
			num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
			if err != nil {
				return "", err
			}
			passwordBytes[i] = letters[num.Int64()]
		}

		password = string(passwordBytes)
		result = EvaluatePasswordStrength(password)
	}
	return password, nil
}

func TrimSpaceLB(answer string) string {
	return strings.TrimSpace(strings.Trim(strings.TrimSpace(answer), "\n"))
}

// IPCheck returns the list of IPs connected to the machine.
func IPCheck() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalf("\nIPCheck: %s", err)
	}

	var ipList []string
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ipList = append(ipList, ipNet.IP.String())
				// fmt.Println("IPv4:", ipNet.IP.String())
			} else if ipNet.IP.To16() != nil {
				ipList = append(ipList, ipNet.IP.String())
				// fmt.Println("IPv6:", ipNet.IP.String())
			}
		}
	}

	if len(ipList) > 0 {
		return ipList[0], nil
	}

	return "", fmt.Errorf("error: No connected network device found")
}

// validarEmail validate the email
func ValidateEmail(email string) error {

	if ok := emailRegex.MatchString(email); !ok {
		return errors.New("invalid email address")
	}
	return nil
}

// validateRope checks if the email starts with mesa[0-9]{2,}, 'mesa' + table number.
// In this case the role can only be 'cliente'.
func ValidateRope(email string, role string) error {

	if ok := ropeRegex.MatchString(email); ok {
		if role != string(models.Cliente) {
			return errors.New("e-mail mesaXX@... must have role == cliente")
		}
	}

	return nil
}

type PasswordStrength struct {
	Score        int
	Feedback     string
	HasNumber    bool
	HasLowercase bool
	HasUppercase bool
	HasSymbols   bool
	IsLongEnough bool
}

// NewPasswordStrength craeates a new instance of PasswordStrength with default values.
func NewPasswordStrength() PasswordStrength {
	return PasswordStrength{
		Score:        0,
		Feedback:     "Very weak password",
		HasNumber:    false,
		HasLowercase: false,
		HasUppercase: false,
		HasSymbols:   false,
		IsLongEnough: false,
	}
}

// EvaluatePasswordStrength returns true if the password strength check passes.
func EvaluatePasswordStrength(password string) PasswordStrength {
	result := NewPasswordStrength()

	hasNumber := numberRegex.MatchString(password)
	hasLowecase := lowercaseRegex.MatchString(password)
	hasUppercase := uppercaseRegex.MatchString(password)
	hasSymbols := symbolRegex.MatchString(password)

	lenPassword := len(password)
	result.Score = lenPassword / 2
	if lenPassword <= minPasswordSize {
		result.IsLongEnough = true
		result.Score++
	}

	if hasNumber {
		result.Score++
		result.HasNumber = true
	}

	if hasLowecase {
		result.Score++
		result.HasLowercase = true
	}

	if hasUppercase {
		result.Score++
		result.HasUppercase = true
	}

	if hasSymbols {
		result.Score++
		result.HasSymbols = true
	}

	switch {
	case result.Score < 6:
		result.Feedback = "Weak password"
	case result.Score < 9:
		result.Feedback = "Moderate password"
	default:
		result.Feedback = "Strong password"
	}

	return result
}
