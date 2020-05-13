package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePost{}

type MsgCreatePost struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Body    string         `json:"body" yaml:"body"`
}

// NewMsgCreatePost creates a new MsgCreatePost instance
func NewMsgCreatePost(creator sdk.AccAddress, body string) MsgCreatePost {
	return MsgCreatePost{
		Creator: creator,
		Body:    body,
	}
}

const CreatePostConst = "CreatePost"

// nolint
func (msg MsgCreatePost) Route() string { return RouterKey }
func (msg MsgCreatePost) Type() string  { return CreatePostConst }
func (msg MsgCreatePost) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgCreatePost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCreatePost) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
