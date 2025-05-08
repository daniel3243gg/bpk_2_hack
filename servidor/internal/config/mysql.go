package config

import (
	"sistema_artigos_bpk/pkg/errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(user, password, host, dbName string, port int) (*gorm.DB, error) {
    
    
    // Construindo o DSN dinamicamente com os valores do arquivo de configuração
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
        user, password, host, port, dbName)
    
    return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func GetDb(user, password, host, dbName string, port int) (*gorm.DB, erros.CustomError) {
	db, err := ConnectDB(user, password, host,  dbName, port)
	if err != nil {
		return nil, erros.CustomError{
			Message:    "Erro ao conectar ao banco de dados",
			Error:      erros.ErrInterno.Error(),
			Orientacao: "Verifique a conexão com o banco de dados",
		}
	}
	return db,erros.CustomError{}
}