package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
    storeKey sdk.StoreKey
    cdc      *codec.Codec
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
    return Keeper{
        storeKey: key,
        cdc:      cdc,
    }
}

func (k Keeper) CreateSomething(ctx sdk.Context, data string) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
    store.Set([]byte(data), []byte(data))
}
