package keeper

import (
	"context"
	"fmt"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeletePost(goCtx context.Context, msg *types.MsgDeletePost) (*types.MsgDeletePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	post, ok := k.GetPost(ctx, msg.Id)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("post id %d does not exist", post.Id))
	}

	if msg.Creator != post.Creator {
		return nil, sdkerrors.ErrUnauthorized
	}

	k.RemovePost(ctx, post.Id)

	return &types.MsgDeletePostResponse{}, nil
}
