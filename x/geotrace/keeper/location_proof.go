package keeper

import (
	// "encoding/binary"
	"geotrace/x/geotrace/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CreateLocationProof(ctx sdk.Context, location_proof types.LocationProof) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LocationProofKey))
	appendedValue := k.cdc.MustMarshal(&location_proof)
	proofKey := location_proof.Creator + "T" + location_proof.Timestamp
	store.Set([]byte(proofKey), appendedValue)
}

func (k Keeper) GetLocationProof(ctx sdk.Context, id string) (val types.LocationProof, found bool) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LocationProofKey))
    b := store.Get([]byte(id))
    if b == nil {
        return val, false
    }
    k.cdc.MustUnmarshal(b, &val)
    return val, true
}