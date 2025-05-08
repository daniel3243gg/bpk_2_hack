package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashString(input string) string {
	hash := sha256.New()                   // Cria um novo hash SHA-256
	hash.Write([]byte(input))              // Escreve os dados da string no hash
	hashedBytes := hash.Sum(nil)           // Calcula o hash
	return hex.EncodeToString(hashedBytes) // Converte o hash para string hexadecimal
}