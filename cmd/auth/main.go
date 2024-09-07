package main

import (
	"MySotre/internal/config"
	"MySotre/internal/delivery/grpcServer"
	"MySotre/pkg/logger"
	"MySotre/pkg/pgDB"
	"log"
)

func main() {
	// Загружаем конфиг
	if err := config.Load(); err != nil {
		log.Fatalln("Error when load config: ", err)
	}

	// Инициализируем логгер
	logsFile, errLoggerInit := logger.Init(config.Config.Logger.Mode, config.Config.Logger.GrpcServerLogFile)

	if errLoggerInit != nil {
		log.Fatalln("Error when init logger: ", errLoggerInit)
	}
	defer logsFile.Close()

	// Подключаемся к БД
	if errConn := pgDB.Conn(
		config.Config.Postgre.Host, config.Config.Postgre.User, config.Config.Postgre.Password,
		config.Config.Postgre.Name, config.Config.Postgre.Port, config.Config.Postgre.SSLmode,
	); errConn != nil {
		log.Fatalln("Error when connect to PostgreSQL: ", errConn)
	}
	defer pgDB.DB.Close()

	// Создаём GRPC сервер
	server := grpcServer.New(config.Config.GrpcServer.Host, config.Config.GrpcServer.Port)
	// Запускаем GRPC сервер
	if errStart := server.Start(); errStart != nil {
		log.Fatalln("Error when start GRPC server: ", errStart)
	}
}
