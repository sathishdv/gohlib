package collection_test

import (
	"testing"

	"github.com/sathishdv/gohlib/collection"
	"github.com/stretchr/testify/assert"
)

func TestInitSetWithoutElements(t *testing.T) {
	set := collection.NewSet()
	assert.Equal(t, 0, set.Len())
	assert.True(t, set.IsEmpty())
}

func TestInitSetWithElements(t *testing.T) {
	set := collection.NewSet(1, 2, 3)
	assert.Equal(t, 3, set.Len())
	assert.False(t, set.IsEmpty())
}

func TestSetWithAdd(t *testing.T) {
	set := collection.NewSet()
	assert.True(t, set.Add(1))
	assert.Equal(t, 1, set.Len())
	assert.False(t, set.IsEmpty())
	assert.True(t, set.Add(2))
	assert.Equal(t, 2, set.Len())
	assert.False(t, set.IsEmpty())
	assert.True(t, set.Add(3))
	assert.Equal(t, 3, set.Len())
	assert.False(t, set.IsEmpty())
	assert.True(t, set.Add(4))
	assert.Equal(t, 4, set.Len())
	assert.False(t, set.IsEmpty())

	assert.True(t, set.Contains(1))
	assert.True(t, set.Contains(2))
	assert.True(t, set.Contains(3))
	assert.True(t, set.Contains(4))
	assert.False(t, set.Contains(5))
}

func TestSetWithDuplicateItems(t *testing.T) {
	set := collection.NewSet()
	assert.True(t, set.Add(1))
	assert.Equal(t, 1, set.Len())
	assert.False(t, set.IsEmpty())
	assert.False(t, set.Add(1))
	assert.Equal(t, 1, set.Len())
	assert.False(t, set.IsEmpty())
}

func TestRemoveItems(t *testing.T) {
	set := collection.NewSet()
	assert.True(t, set.Add(1))
	assert.Equal(t, 1, set.Len())
	assert.False(t, set.IsEmpty())
	assert.True(t, set.Remove(1))
	assert.Equal(t, 0, set.Len())
	assert.True(t, set.IsEmpty())
	assert.False(t, set.Remove(1))
	assert.Equal(t, 0, set.Len())
	assert.True(t, set.IsEmpty())
}
