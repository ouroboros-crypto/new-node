package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ouroboros-crypto/node/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgReinvest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgReinvest
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgReinvest{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgReinvest{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
