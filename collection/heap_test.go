package collection_test

import (
	"os"
	"testing"

	"github.com/sathishdv/gohlib/collection"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestMaxHeap(t *testing.T) {
	// init max heap
	maxHeap := collection.NewMaxHeap()
	maxHeap.Push(1)
	maxHeap.Push(4)
	maxHeap.Push(2)
	maxHeap.Push(3)
	maxHeap.Push(5)

	assert.Equal(t, 5, maxHeap.Len())
	assert.Equal(t, 5, *maxHeap.Peek())
	assert.Equal(t, 5, *maxHeap.Pop())
	assert.Equal(t, 4, maxHeap.Len())
	assert.Equal(t, 4, *maxHeap.Peek())
	assert.Equal(t, 4, *maxHeap.Pop())
	assert.Equal(t, 3, maxHeap.Len())
	assert.Equal(t, 3, *maxHeap.Peek())
	assert.Equal(t, 3, *maxHeap.Pop())
	assert.Equal(t, 2, maxHeap.Len())
	assert.Equal(t, 2, *maxHeap.Peek())
	assert.Equal(t, 2, *maxHeap.Pop())
	assert.Equal(t, 1, maxHeap.Len())
	assert.Equal(t, 1, *maxHeap.Peek())
	assert.Equal(t, 1, *maxHeap.Pop())
	assert.Equal(t, 0, maxHeap.Len())
	assert.Nil(t, maxHeap.Peek())
	assert.Nil(t, maxHeap.Pop())
}

func TestMinHeap(t *testing.T) {
	minHeap := collection.NewMinHeap()
	minHeap.Push(1)
	minHeap.Push(4)
	minHeap.Push(2)
	minHeap.Push(3)
	minHeap.Push(5)

	assert.Equal(t, 5, minHeap.Len())
	assert.Equal(t, 1, *minHeap.Peek())
	assert.Equal(t, 1, *minHeap.Pop())
	assert.Equal(t, 4, minHeap.Len())
	assert.Equal(t, 2, *minHeap.Peek())
	assert.Equal(t, 2, *minHeap.Pop())
	assert.Equal(t, 3, minHeap.Len())
	assert.Equal(t, 3, *minHeap.Peek())
	assert.Equal(t, 3, *minHeap.Pop())
	assert.Equal(t, 2, minHeap.Len())
	assert.Equal(t, 4, *minHeap.Peek())
	assert.Equal(t, 4, *minHeap.Pop())
	assert.Equal(t, 1, minHeap.Len())
	assert.Equal(t, 5, *minHeap.Peek())
	assert.Equal(t, 5, *minHeap.Pop())
	assert.Equal(t, 0, minHeap.Len())
	assert.Nil(t, minHeap.Peek())
	assert.Nil(t, minHeap.Pop())
}

func TestPopWithoutPush(t *testing.T) {
	maxHeap := collection.NewMaxHeap()
	assert.Nil(t, maxHeap.Pop())
	assert.Nil(t, maxHeap.Peek())

	minHeap := collection.NewMinHeap()
	assert.Nil(t, minHeap.Pop())
	assert.Nil(t, minHeap.Peek())
}

func TestPushSameNumber(t *testing.T) {
	maxHeap := collection.NewMaxHeap()
	maxHeap.Push(1)
	maxHeap.Push(1)
	assert.Equal(t, 2, maxHeap.Len())
	assert.Equal(t, 1, *maxHeap.Peek())
	assert.Equal(t, 1, *maxHeap.Pop())
	assert.Equal(t, 1, maxHeap.Len())
	assert.Equal(t, 1, *maxHeap.Peek())
	assert.Equal(t, 1, *maxHeap.Pop())
	assert.Equal(t, 0, maxHeap.Len())
	assert.Nil(t, maxHeap.Peek())
	assert.Nil(t, maxHeap.Pop())

	minHeap := collection.NewMinHeap()
	minHeap.Push(1)
	minHeap.Push(1)
	assert.Equal(t, 2, minHeap.Len())
	assert.Equal(t, 1, *minHeap.Peek())
	assert.Equal(t, 1, *minHeap.Pop())
	assert.Equal(t, 1, minHeap.Len())
	assert.Equal(t, 1, *minHeap.Peek())
	assert.Equal(t, 1, *minHeap.Pop())
	assert.Equal(t, 0, minHeap.Len())
	assert.Nil(t, minHeap.Peek())
	assert.Nil(t, minHeap.Pop())
}
