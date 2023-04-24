package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddLocationProof = "add_location_proof"

var _ sdk.Msg = &MsgAddLocationProof{}

func NewMsgAddLocationProof(creator string, lat string, lon string, timestamp string) *MsgAddLocationProof {
	return &MsgAddLocationProof{
		Creator:   creator,
		Lat:       lat,
		Lon:       lon,
		Timestamp: timestamp,
	}
}

func (msg *MsgAddLocationProof) Route() string {
	return RouterKey
}

func (msg *MsgAddLocationProof) Type() string {
	return TypeMsgAddLocationProof
}

func (msg *MsgAddLocationProof) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddLocationProof) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddLocationProof) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
