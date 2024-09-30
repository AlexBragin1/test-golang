package groups

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupSetHasOneGroup(t *testing.T) {
	gset := User
	assert.True(t, gset.HasAccessTo(User))
	assert.False(t, gset.HasAccessTo(Admin))
}

func TestGroupSet3(t *testing.T) {
	uGroups := User | Admin
	assert.True(t, uGroups.HasAccessTo(User|Admin))
	assert.True(t, uGroups.HasAccessTo(User))
}
