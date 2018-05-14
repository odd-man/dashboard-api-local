/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package api

import (
	"net/http"
	"time"

	"github.com/fvbock/endless"
	"golang.org/x/sync/errgroup"
)

// Server api server
type Server struct {
	Server *http.Server
	G      *errgroup.Group
	*API
}

// GetServer get the server
func GetServer(api *API) (server *Server) {
	serverConfig := api.config.ServerConfig
	if serverConfig == nil {
		serverConfig = &ServerConfig{
			ReadTimeout:    0,
			WriteTimeout:   0,
			IdleTimeout:    0,
			MaxHeaderBytes: 1 << 20, //1MB
		}
		api.config.ServerConfig = serverConfig
	}

	//TODO: addr http/https
	currentEngineConfig := &EngineConfig{
		middleware:       nil,
		LimitConnections: api.config.LimitConnection,
	}

	return &Server{
		G:   api.ErrorGroup,
		API: api,
		Server: &http.Server{
			Handler:        currentEngineConfig.Init(),
			ReadTimeout:    serverConfig.ReadTimeout,
			WriteTimeout:   serverConfig.WriteTimeout,
			IdleTimeout:    serverConfig.IdleTimeout,
			MaxHeaderBytes: serverConfig.MaxHeaderBytes,
		},
	}
}

// Run start the server
func (server *Server) Run() {
	// run http server
	server.runServer()
	if server.config.EnableHTTPS {
		if server.config.CertFile == "" || server.config.KeyFile == "" {
			return
		}
		//TODO if certFIle or keyFile not exist runServer
		server.runServerTLS()
	}
	if err := server.G.Wait(); err != nil {
		server.API.log.Fatalln(err)
		// log.Fatal(err)
	}
}

// runServer run our server in a goroutine so that it doesn't block.
func (server *Server) runServer() {
	server.G.Go(func() error {
		// log.Printf("running server %v", server.config.ListenAddr)
		server.API.log.Info("running server %v", server.config.ListenAddr)
		if !server.config.EnableGraceful {
			return http.ListenAndServe(server.config.ListenAddr, server.Server.Handler)
		}
		// Graceful restart or stop we use fvbock/endless to replace the default ListenAndServe
		endless.DefaultHammerTime = server.config.DefaultHammerTime * time.Second
		return endless.ListenAndServe(server.config.ListenAddr, server.Server.Handler)
	})
}

// runServerTLS run our server with tls in a goroutine so that it doesn't block.
func (server *Server) runServerTLS() {
	server.G.Go(func() error {
		server.API.log.Info("running server TLS %v", server.config.ListenAddr)
		// log.Printf("running server TLS %v", server.config.HTTPSAddr)
		if !server.config.EnableGraceful {
			return http.ListenAndServeTLS(server.config.HTTPSAddr, server.config.CertFile, server.config.KeyFile, server.Server.Handler)
		}
		// Graceful restart or stop we use fvbock/endless to replace the default ListenAndServeTLS
		endless.DefaultHammerTime = server.config.DefaultHammerTime * time.Second
		return endless.ListenAndServeTLS(server.config.HTTPSAddr, server.config.CertFile, server.config.KeyFile, server.Server.Handler)
	})
}
