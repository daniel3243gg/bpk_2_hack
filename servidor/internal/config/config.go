package config

import (
	"log"
	"os"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
}

var configStruct Config

func LoadConfig(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&configStruct); err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}
}

func GetDatabaseConfig() (string, string, string, int, string) {
	return configStruct.Database.User, configStruct.Database.Password, configStruct.Database.Host, configStruct.Database.Port, configStruct.Database.Name
}


func GetEnv(chave string, valorPadrao string) string {
    valor, existe := os.LookupEnv(chave) // Verifica se a variável de ambiente existe
    if !existe || valor == "" {         // Se não existir ou estiver vazia, retorna o valor padrão
        return valorPadrao
    }
    return valor // Retorna o valor da variável de ambiente
}
func GetJWTSecret() string {
	return configStruct.JWT.Secret
}