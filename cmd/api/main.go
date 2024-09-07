package main

import (
	"MySotre/internal/config"
	"MySotre/internal/delivery/httpServer"
	"MySotre/pkg/cacheDB"
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
	logsFile, errLoggerInit := logger.Init(config.Config.Logger.Mode, config.Config.Logger.HttpServerLogFile)

	if errLoggerInit != nil {
		log.Fatalln("Error when init logger: ", errLoggerInit)
	}
	defer logsFile.Close()

	// Подключаемся к PostgreSQL
	if errConn := pgDB.Conn(
		config.Config.Postgre.Host, config.Config.Postgre.User, config.Config.Postgre.Password,
		config.Config.Postgre.Name, config.Config.Postgre.Port, config.Config.Postgre.SSLmode,
	); errConn != nil {
		log.Fatalln("Error when connect to PostgreSQL: ", errConn)
	}
	defer pgDB.DB.Close()

	// Подключаемся к Redis
	if errCacheDB := cacheDB.New(config.Config.Redis.Host, config.Config.Redis.Port); errCacheDB != nil {
		log.Fatalln("Error when connect to Redis: ", errCacheDB)
	}
	defer cacheDB.DB.Close()

	// Создаём HTTP сервер
	server := httpServer.New(config.Config.HttpServer.Host, config.Config.GrpcServer.Host, config.Config.HttpServer.Port, config.Config.GrpcServer.Port)
	// Запускаем HTTP сервер
	if errStart := server.Start(); errStart != nil {
		log.Fatalln("Error when start HTTP server: ", errStart)
	}
}
