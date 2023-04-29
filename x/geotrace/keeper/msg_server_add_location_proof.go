package keeper

import (
	"context"
	"geotrace/x/geotrace/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddLocationProof(goCtx context.Context, msg *types.MsgAddLocationProof) (*types.MsgAddLocationProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var locationProof = types.LocationProof{
		Creator:   msg.Creator,
		Lat:       msg.Lat,
		Lon:       msg.Lon,
		Timestamp: msg.Timestamp,
	}
	k.CreateLocationProof(ctx, locationProof)
	return &types.MsgAddLocationProofResponse{}, nil
}
