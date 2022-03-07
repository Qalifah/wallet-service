package handler

import (
	"context"
	core "github.com/Qalifah/grey-challenge/wallet"
	"github.com/Qalifah/grey-challenge/wallet/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type handler struct {
	balanceRepo core.BalanceRepository
	proto.UnimplementedWalletServer
}

func New(balanceRepo core.BalanceRepository) *handler {
	return &handler{balanceRepo: balanceRepo}
}

func (h *handler) GetBalance(ctx context.Context, request *proto.GetBalanceRequest) (*proto.BalanceResponse, error) {
	balance, err := h.balanceRepo.Get(ctx, int(request.UserId))
	if err != nil {
		return nil, err
	}
	return marshalBalanceResponse(balance), nil
}

func (h *handler) UpdateBalance(ctx context.Context, request *proto.UpdateBalanceRequest) (*proto.BalanceResponse, error) {
	balance, err := h.balanceRepo.Update(ctx, int(request.UserId), int(request.NewBalance))
	if err != nil {
		return nil, err
	}
	return marshalBalanceResponse(balance), nil
}

func marshalBalanceResponse(balance *core.Balance) *proto.BalanceResponse {
	return &proto.BalanceResponse{
		UserId:    uint32(balance.UserID),
		Amount:    uint64(balance.Amount),
		UpdatedAt: timestamppb.New(balance.UpdatedAt),
	}
}
