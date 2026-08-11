package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


// Iniciais maiusculas nas funções indicam que ela é Public
func ConnectDB() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")	
	}

	// atribuir o valor da variável de ambiente para dbURL
	dbURL := os.Getenv("DATABASE_URL")

	// Abre conexão do banco
	db, err := sql.Open("postgres", dbURL)

	// caso a conexão seja falha vai retornar o erro
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Erro ao conectar ao banco: (Ping): ", err)
	}

	log.Println("Conectado ao banco de dados com Sucesso!")

	return db

}