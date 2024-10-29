package order

import (
	"errors"
	"fmt"
)

// Orderable is an interface that items must implement to be orderable.
type Orderable interface {
	GetID() string
	GetPosition() int
	SetPosition(position int)
}

var (
	ErrItemNotFound    = errors.New("item not found")
	ErrInvalidPosition = errors.New("invalid position")
)

// OrderManager provides methods to manage the order of items.
type OrderManager[T Orderable] struct{}

// NewOrderManager creates a new instance of OrderManager.
func NewOrderManager[T Orderable]() *OrderManager[T] {
	return &OrderManager[T]{}
}

// NormalizePositions ensures that the positions of items are sequential starting from 1.
func (os *OrderManager[T]) NormalizePositions(items []T) {
	for i, item := range items {
		item.SetPosition(i + 1)
	}
}

// GetItemIndexByID returns the index of an item by its ID.
func (os *OrderManager[T]) GetItemIndexByID(items []T, itemID string) (int, error) {
	for index, item := range items {
		if item.GetID() == itemID {
			return index, nil
		}
	}
	return -1, fmt.Errorf("GetItemIndexByID: %w", ErrItemNotFound)
}

// Up moves an item up by one position.
func (os *OrderManager[T]) Up(items []T, itemID string) error {
	index, err := os.GetItemIndexByID(items, itemID)
	if err != nil {
		return err
	}
	if index == 0 {
		// Item is already at the top
		return nil
	}
	// Swap with the item above
	items[index], items[index-1] = items[index-1], items[index]
	// Normalize positions
	os.NormalizePositions(items)
	return nil
}

// Down moves an item down by one position.
func (os *OrderManager[T]) Down(items []T, itemID string) error {
	index, err := os.GetItemIndexByID(items, itemID)
	if err != nil {
		return err
	}
	if index == len(items)-1 {
		// Item is already at the bottom
		return nil
	}
	// Swap with the item below
	items[index], items[index+1] = items[index+1], items[index]
	// Normalize positions
	os.NormalizePositions(items)
	return nil
}

// To moves an item to a specific position.
func (os *OrderManager[T]) To(items []T, itemID string, newPosition int) error {
	if newPosition < 1 || newPosition > len(items) {
		return fmt.Errorf("To: %w", ErrInvalidPosition)
	}

	currentIndex, err := os.GetItemIndexByID(items, itemID)
	if err != nil {
		return err
	}

	// Remove the item from its current position
	itemToMove := items[currentIndex]
	items = append(items[:currentIndex], items[currentIndex+1:]...)

	// Adjust for zero-based index
	insertIndex := newPosition - 1

	// Insert the item at the new position
	items = append(items[:insertIndex], append([]T{itemToMove}, items[insertIndex:]...)...)

	// Normalize positions
	os.NormalizePositions(items)

	return nil
}

// Top moves an item to the first position.
func (os *OrderManager[T]) Top(items []T, itemID string) error {
	return os.To(items, itemID, 1)
}

// Bottom moves an item to the last position.
func (os *OrderManager[T]) Bottom(items []T, itemID string) error {
	return os.To(items, itemID, len(items))
}

// Above moves an item to be directly above the target item.
func (os *OrderManager[T]) Above(items []T, itemID string, targetID string) error {
	targetIndex, err := os.GetItemIndexByID(items, targetID)
	if err != nil {
		return err
	}
	return os.To(items, itemID, targetIndex+1)
}

// Below moves an item to be directly below the target item.
func (os *OrderManager[T]) Below(items []T, itemID string, targetID string) error {
	targetIndex, err := os.GetItemIndexByID(items, targetID)
	if err != nil {
		return err
	}
	return os.To(items, itemID, targetIndex+2)
}
