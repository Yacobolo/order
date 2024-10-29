// order/order_test.go
package order_test

import (
	"errors"
	"order"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestItem is a sample struct implementing the Orderable interface.
type TestItem struct {
	ID       uuid.UUID
	Position int
}

// Implement Orderable interface for TestItem.
func (ti *TestItem) GetID() string {
	return ti.ID.String()
}

func (ti *TestItem) GetPosition() int {
	return ti.Position
}

func (ti *TestItem) SetPosition(position int) {
	ti.Position = position
}

func createTestItems(n int) []*TestItem {
	items := make([]*TestItem, n)
	for i := 0; i < n; i++ {
		items[i] = &TestItem{
			ID:       uuid.New(),
			Position: i + 1,
		}
	}
	return items
}

func TestUp(t *testing.T) {
	os := order.NewOrderManager[*TestItem]()
	items := createTestItems(3)

	itemID := items[1].GetID() // Middle item
	err := os.Up(items, itemID)
	assert.NoError(t, err)

	index, _ := os.GetItemIndexByID(items, itemID)
	assert.Equal(t, 0, index)
	assert.Equal(t, 1, items[index].GetPosition())
}

func TestDown(t *testing.T) {
	os := order.NewOrderManager[*TestItem]()
	items := createTestItems(3)

	itemID := items[1].GetID() // Middle item
	err := os.Down(items, itemID)
	assert.NoError(t, err)

	index, _ := os.GetItemIndexByID(items, itemID)
	assert.Equal(t, 2, index)
	assert.Equal(t, 3, items[index].GetPosition())
}

func TestTo(t *testing.T) {
	os := order.NewOrderManager[*TestItem]()
	items := createTestItems(5)

	itemID := items[0].GetID() // First item
	err := os.To(items, itemID, 3)
	assert.NoError(t, err)

	index, _ := os.GetItemIndexByID(items, itemID)
	assert.Equal(t, 2, index)
	assert.Equal(t, 3, items[index].GetPosition())
}

func TestTop(t *testing.T) {
	os := order.NewOrderManager[*TestItem]()
	items := createTestItems(5)

	itemID := items[3].GetID() // Item at position 4
	err := os.Top(items, itemID)
	assert.NoError(t, err)

	index, _ := os.GetItemIndexByID(items, itemID)
	assert.Equal(t, 0, index)
	assert.Equal(t, 1, items[index].GetPosition())
}

func TestBottom(t *testing.T) {
	os := order.NewOrderManager[*TestItem]()
	items := createTestItems(5)

	itemID := items[1].GetID() // Item at position 2
	err := os.Bottom(items, itemID)
	assert.NoError(t, err)

	index, _ := os.GetItemIndexByID(items, itemID)
	assert.Equal(t, 4, index)
	assert.Equal(t, 5, items[index].GetPosition())
}

func TestAbove(t *testing.T) {
	os := order.NewOrderManager[*TestItem]()
	items := createTestItems(5)

	itemID := items[4].GetID()   // Last item
	targetID := items[1].GetID() // Target is at position 2
	err := os.Above(items, itemID, targetID)
	assert.NoError(t, err)

	index, _ := os.GetItemIndexByID(items, itemID)
	assert.Equal(t, 1, index)
	assert.Equal(t, 2, items[index].GetPosition())
}

func TestBelow(t *testing.T) {
	os := order.NewOrderManager[*TestItem]()
	items := createTestItems(5)

	itemID := items[0].GetID()   // First item
	targetID := items[2].GetID() // Target is at position 3
	err := os.Below(items, itemID, targetID)
	assert.NoError(t, err)

	index, _ := os.GetItemIndexByID(items, itemID)
	assert.Equal(t, 3, index)
	assert.Equal(t, 4, items[index].GetPosition())
}

func TestNormalizePositions(t *testing.T) {
	os := order.NewOrderManager[*TestItem]()
	items := createTestItems(3)

	// Manually alter positions
	items[0].SetPosition(10)
	items[1].SetPosition(20)
	items[2].SetPosition(30)

	os.NormalizePositions(items)

	for i, item := range items {
		assert.Equal(t, i+1, item.GetPosition())
	}
}

func TestGetItemIndexByID_NotFound(t *testing.T) {
	os := order.NewOrderManager[*TestItem]()
	items := createTestItems(3)

	nonExistentID := uuid.New().String()
	index, err := os.GetItemIndexByID(items, nonExistentID)
	assert.Error(t, err)
	assert.Equal(t, -1, index)
	assert.Equal(t, order.ErrItemNotFound, errors.Unwrap(err))
}

func TestInvalidPosition(t *testing.T) {
	os := order.NewOrderManager[*TestItem]()
	items := createTestItems(3)

	itemID := items[0].GetID()

	err := os.To(items, itemID, 0)
	assert.Error(t, err)
	assert.Equal(t, order.ErrInvalidPosition, errors.Unwrap(err))

	err = os.To(items, itemID, 5)
	assert.Error(t, err)
	assert.Equal(t, order.ErrInvalidPosition, errors.Unwrap(err))
}
