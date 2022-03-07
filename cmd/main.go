package main

import (
	"context"
	"flag"
	"github.com/Qalifah/grey-challenge/wallet/config"
	"github.com/Qalifah/grey-challenge/wallet/database/postgres"
	"github.com/Qalifah/grey-challenge/wallet/handler"
	"github.com/Qalifah/grey-challenge/wallet/proto"
	"google.golang.org/grpc"
	"net"
	"os"

	"gopkg.in/yaml.v2"

	log "github.com/sirupsen/logrus"
)

var configPath *string

func init() {
	configPath = flag.String("config_path", "", "path to config file")
	flag.Parse()
	if configPath == nil {
		log.Fatal("-config_path flag is required")
	}
}

func main() {
	file, err := os.Open(*configPath)
	if err != nil {
		log.Fatalf("unable to open config file: %v", err)
	}

	cfg := &config.BaseConfig{}
	err = yaml.NewDecoder(file).Decode(cfg)
	if err != nil {
		log.Fatalf("failed to decode config file: %v", err)
	}

	ctx := context.Background()
	conn, err := postgres.New(ctx, cfg.Postgres)
	if err != nil {
		log.Fatalf("failed to create postgres client: %v", err)
	}
	defer conn.Close(ctx)

	balanceRepo := postgres.NewBalanceRepository(conn)
	ctrl := handler.New(balanceRepo)
	lis, err := net.Listen("tcp", cfg.ServeAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterWalletServer(s, ctrl)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
