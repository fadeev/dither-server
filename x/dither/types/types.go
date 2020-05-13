package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Post ...
type Post struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"` // address of the creator
	Body    string         `json:"body" yaml:"body"`       // body
}

// implement fmt.Stringer
func (s Post) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Creator: %s
	Body: %s`,
		s.Creator,
		s.Body,
	))
}
