package handler

import (
	"context"
	"github.com/Qalifah/grey-challenge/wallet/database/postgres"
	"github.com/Qalifah/grey-challenge/wallet/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"

	log "github.com/sirupsen/logrus"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	ctx := context.Background()
	conn, err := postgres.CreateTestDB(ctx)
	if err != nil {
		log.Fatalf("Unable to create test database %v", err)
	}
	balanceRepo := postgres.NewBalanceRepository(conn)
	proto.RegisterWalletServer(s, New(balanceRepo))
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestHandler_GetBalance(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer conn.Close()

	client := proto.NewWalletClient(conn)
	resp, err := client.GetBalance(ctx, &proto.GetBalanceRequest{UserId: 1})
	require.NoError(t, err)
	require.Equal(t, 300, int(resp.Amount))
}

func TestHandler_UpdateBalance(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer conn.Close()

	client := proto.NewWalletClient(conn)
	resp, err := client.UpdateBalance(ctx, &proto.UpdateBalanceRequest{UserId: 7, NewBalance: 500})
	require.NoError(t, err)
	require.Equal(t, 500, int(resp.Amount))
}
