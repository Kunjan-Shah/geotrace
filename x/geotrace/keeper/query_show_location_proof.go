package keeper

import (
	"context"

	"geotrace/x/geotrace/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowLocationProof(goCtx context.Context, req *types.QueryShowLocationProofRequest) (*types.QueryShowLocationProofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	locationProof, found := k.GetLocationProof(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}
	return &types.QueryShowLocationProofResponse{LocationProof: locationProof}, nil
}
