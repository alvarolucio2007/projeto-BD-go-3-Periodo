package main

import (
	"log"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/api"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/database"
	grpcclient "github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/client"
	grpcserver "github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/server"
)

func main() {
	log.Print("Tentando conectar banco de dados...")
	_, err := database.ConectarPostgres()
	if err != nil {
		log.Fatalf("\x1b[31m[ERRO CRÍTICO]\x1b[0m Falha ao conectar na Database Postgres: %v", err)
		panic(err)
	}
	log.Print("Banco de dados conectado, tentando migrar base de dados...")
	err = database.MigrarPostgres()
	if err != nil {
		log.Fatalf("\x1b[31m[ERRO CRÍTICO]\x1b[0m Falha ao conectar o servidor gRPC: %v", err)
		panic(err)
	}
	log.Print("Banco de dados migrado, tentando iniciar servidor gRPC...")
	go func() {
		grpcserver.StartServerGeralGRPC()
	}()
	time.Sleep(2 * time.Second)
	log.Print("Servidor gRPC iniciado com sucesso, tentando iniciar client gRPC...")
	grpcclient.ConnectAll("localhost:50051")
	log.Print("Client gRPC iniciado com sucesso, tentando iniciar servidor API REST...")
	api.SetupExtRoutes()
	log.Print("Servidor REST iniciado com sucesso. Aplicação iniciada com sucesso.")
}
