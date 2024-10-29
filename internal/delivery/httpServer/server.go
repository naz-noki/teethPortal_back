package httpServer

import (
	"MySotre/internal/delivery"
	"MySotre/internal/middlewares"
	"MySotre/internal/routers/authorsRouter"
	"MySotre/internal/routers/tokensRouter"
	"MySotre/internal/routers/usersRouter"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/authApi"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
	"google.golang.org/grpc"
)

type server struct {
	host         string
	port         int
	grpcHost     string
	grpcPort     int
	authClient   authApi.AuthClient
	tokensClient tokensApi.TokensClient
}

func (h *server) Start() error {
	// Устанавливаем соединение с grpc сервером
	grpcAddr := fmt.Sprintf("%s:%d", h.grpcHost, h.grpcPort)
	conn, errDial := grpc.Dial(grpcAddr, grpc.WithInsecure())

	if errDial != nil {
		return errDial
	}
	defer conn.Close()

	// Получаем клиентов для grpc сервера
	h.authClient = authApi.NewAuthClient(conn)
	h.tokensClient = tokensApi.NewTokensClient(conn)

	// Создаём gin сервер
	server := gin.Default()
	addr := fmt.Sprintf("%s:%d", h.host, h.port)

	// Подключаем корсы (CORS)
	server.Use(middlewares.AddCors())

	// Подключаем роуты
	usersRouter.AddUsersRoutes(server, h.authClient)
	tokensRouter.AddTokensRoutes(server, h.tokensClient)
	authorsRouter.AddAuthorsRoutes(server, h.tokensClient)

	// Устанавливаем значения для сервера
	srv := &http.Server{
		Addr:         addr,
		Handler:      server,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Запускаем сервер
	return srv.ListenAndServe()
}

func New(
	host, grpcHost string,
	port, grpcPort int,
) delivery.Server {
	hs := server{
		host:     host,
		port:     port,
		grpcHost: grpcHost,
		grpcPort: grpcPort,
	}

	return &hs
}
