package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPermissions_IsBlacklisted(t *testing.T) {
	perms := Permissions{
		Blacklist: []uint32{
			uint32(PermClaimValidator),
		},
		Whitelist: nil,
	}

	require.True(t, perms.IsBlacklisted(PermClaimValidator))
	require.False(t, perms.IsBlacklisted(PermClaimGovernanceSeat))
}

func TestPermissions_IsWhitelisted(t *testing.T) {
	perms := NewPermissions([]PermValue{PermClaimGovernanceSeat}, nil)

	require.True(t, perms.IsWhitelisted(PermClaimGovernanceSeat))
	require.False(t, perms.IsWhitelisted(PermClaimValidator))
}