package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/fadeev/dither-server/x/dither/types"
)

// NewQuerier creates a new querier for dither clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryListPosts:
			return listPosts(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown dither query endpoint")
		}
	}
}

func listPosts(ctx sdk.Context, k Keeper) ([]byte, error) {
	var postList types.QueryResPosts

	iterator := k.GetPostsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		postList = append(postList, string(iterator.Key()))
	}

	res, err := codec.MarshalJSONIndent(k.cdc, postList)
	if err != nil {
		return res, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// TODO: Add the modules query functions
// They will be similar to the above one: queryParams()
