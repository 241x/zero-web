// Package server 提供基于 gin 的 HTTP 服务启动、信号处理与优雅关闭能力。
package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/241x/zero-kit/logger"
	"github.com/gin-gonic/gin"
)

// Config 服务配置
type Config struct {
	Host string
	Port int
}

// Option Server 可选配置。
type Option func(*Server)

// WithShutdownTimeout 设置优雅关闭超时时间，默认 5s。
func WithShutdownTimeout(d time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = d
	}
}

// Server HTTP 服务。
type Server struct {
	httpServ        *http.Server
	logger          logger.Logger
	shutdownTimeout time.Duration
}

// New 创建 HTTP 服务。
func New(cfg Config, handler *gin.Engine, log logger.Logger, opts ...Option) *Server {
	s := &Server{
		httpServ: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler: handler,
		},
		logger:          log,
		shutdownTimeout: 5 * time.Second,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Run 启动服务并阻塞，直到收到 SIGINT 或 SIGTERM 后优雅关闭。
func (s *Server) Run() {
	ln, err := net.Listen("tcp", s.httpServ.Addr)
	if err != nil {
		s.logger.Err(err, "Listen failed")
		return
	}

	s.logger.Info("Server Start", "listen", s.httpServ.Addr)

	go func() {
		if err := s.httpServ.Serve(ln); err != nil && err != http.ErrServerClosed {
			s.logger.Err(err, "Serve")
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	if err := s.httpServ.Shutdown(ctx); err != nil {
		s.logger.Err(err, "Server Shutdown")
	}
}
