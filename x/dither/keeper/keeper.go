package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fadeev/dither-server/x/dither/types"
)

// Keeper of the dither store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	paramspace types.ParamSubspace
}

// NewKeeper creates a dither keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramspace types.ParamSubspace) Keeper {
	keeper := Keeper{
		storeKey:   key,
		cdc:        cdc,
		paramspace: paramspace.WithKeyTable(types.ParamKeyTable()),
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetPost ...
func (k Keeper) GetPost(ctx sdk.Context, key string) (types.Post, error) {
	store := ctx.KVStore(k.storeKey)
	var post types.Post
	byteKey := []byte(key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &post)
	if err != nil {
		return post, err
	}
	return post, nil
}

// SetPost ...
func (k Keeper) SetPost(ctx sdk.Context, key string, post types.Post) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(post)
	store.Set([]byte(key), bz)
}

// GetPostsIterator gets an iterator over all scavnges in which the keys are the solutionHashes and the values are the scavenges
func (k Keeper) GetPostsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.PostPrefix))
}

// func (k Keeper) delete(ctx sdk.Context, key string) {
// 	store := ctx.KVStore(k.storeKey)
// 	store.Delete([]byte(key))
// }
