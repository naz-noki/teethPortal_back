package grpcServer

import (
	"MySotre/internal/delivery"
	"MySotre/internal/service"
	"MySotre/internal/service/ssoService"
	"fmt"
	"net"

	"github.com/naz-noki/teethPortal_proto/gen/go/sso/authApi"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
	"google.golang.org/grpc"
)

type server struct {
	host       string
	port       int
	server     *grpc.Server
	ssoService service.SsoService
}

func (g *server) Start() error {
	l, errListen := net.Listen(
		"tcp",
		fmt.Sprintf("%s:%d", g.host, g.port),
	)

	if errListen != nil {
		return errListen
	}

	return g.server.Serve(l)
}

func New(host string, port int) delivery.Server {
	gs := server{
		host:       host,
		port:       port,
		server:     grpc.NewServer(),
		ssoService: ssoService.New(),
	}

	authApi.RegisterAuthServer(gs.server, gs.ssoService)
	tokensApi.RegisterTokensServer(gs.server, gs.ssoService)

	return &gs
}
