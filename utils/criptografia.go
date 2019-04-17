package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// CriptografaSenha recebe uma string e retorna um hash para ela, útil para inserção no banco de dados
func CriptografaSenha(senha string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.MinCost)
	if err != nil {
		panic(err.Error())
	}

	return string(hash)
}

// ValidaSenha verifica se a senha que foi retornada da query do banco de dados
// bate com a senha digitada pelo usuário.
func ValidaSenha(senhaCriptografada, senhaDigitada string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash, otherPwd := []byte(senhaCriptografada), []byte(senhaDigitada)
	err := bcrypt.CompareHashAndPassword(byteHash, otherPwd)
	if err != nil {
		return false
	}

	return true
}
