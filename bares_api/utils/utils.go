package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz[]_,.;:!@#$%&*-*/+"

// GenerateRandomPassword gera uma senha aleatória com o comprimento especificado.
// A senha pode conter caracteres alfanuméricos e símbolos especiais.
// O parâmetro 'length' define o comprimento da senha desejada.
// Retorna a senha gerada e um erro, se ocorrer.
func GenerateRandomPassword(length int) (string, error) {
  if length <= 0 {
    return "", fmt.Errorf("o comprimento deve ser maior que zero")
  }

  ret := make([]byte, length)
  for i := 0; i < length; i++ {
    num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
    if err != nil {
      return "", err
    }
    ret[i] = letters[num.Int64()]
  }
  return string(ret), nil
}
