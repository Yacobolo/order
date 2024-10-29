
# orderedcollection

A Go package that provides ordering functionalities for collections of items implementing the `Orderable` interface. Inspired by `django-ordered-model`, it allows you to easily move items up, down, to specific positions, above or below other items, and to the top or bottom of a collection.

## Features

- **Move Items Up or Down**: Shift items one position up or down.
- **Move Items to Specific Positions**: Place items at any valid position in the collection.
- **Move Items Above or Below Others**: Position items directly above or below another item.
- **Move Items to Top or Bottom**: Quickly move items to the start or end of the collection.
- **Normalize Positions**: Ensure item positions are sequential and consistent.

## Installation

To install the package, run:

```bash
go get github.com/yacobolo/orderedcollection
```


## Usage

First, import the package into your Go project:

```go
import "github.com/yourusername/orderedcollection"
```

### Implement the `Orderable` Interface

Your item type must implement the `Orderable` interface:

```go
type Orderable interface {
    GetID() string
    GetPosition() int
    SetPosition(position int)
}
```

Here's an example of a custom type implementing `Orderable`:

```go
type Item struct {
    ID       uuid.UUID
    Position int
    Name     string
}

func (i *Item) GetID() string {
    return i.ID.String()
}

func (i *Item) GetPosition() int {
    return i.Position
}

func (i *Item) SetPosition(position int) {
    i.Position = position
}
```

### Initialize the Ordering Service

Create an instance of the `OrderingService`:

```go
os := orderedcollection.NewOrderingService[*Item]()
```

### Examples

#### Moving an Item Up

```go
err := os.Up(items, itemID)
if err != nil {
    // Handle error
}
```

#### Moving an Item Down

```go
err := os.Down(items, itemID)
if err != nil {
    // Handle error
}
```

#### Moving an Item to a Specific Position

```go
err := os.To(items, itemID, newPosition)
if err != nil {
    // Handle error
}
```

#### Moving an Item Above Another

```go
err := os.Above(items, itemID, targetID)
if err != nil {
    // Handle error
}
```

#### Moving an Item Below Another

```go
err := os.Below(items, itemID, targetID)
if err != nil {
    // Handle error
}
```

#### Moving an Item to the Top

```go
err := os.Top(items, itemID)
if err != nil {
    // Handle error
}
```

#### Moving an Item to the Bottom

```go
err := os.Bottom(items, itemID)
if err != nil {
    // Handle error
}
```

### Full Example

Here's a full example demonstrating how to use the package:

```go
package main

import (
    "fmt"

    "github.com/google/uuid"
    "github.com/yourusername/orderedcollection"
)

type Item struct {
    ID       uuid.UUID
    Position int
    Name     string
}

func (i *Item) GetID() string {
    return i.ID.String()
}

func (i *Item) GetPosition() int {
    return i.Position
}

func (i *Item) SetPosition(position int) {
    i.Position = position
}

func main() {
    // Initialize the ordering service
    os := orderedcollection.NewOrderingService[*Item]()

    // Create a slice of items
    items := []*Item{
        {ID: uuid.New(), Position: 1, Name: "Item A"},
        {ID: uuid.New(), Position: 2, Name: "Item B"},
        {ID: uuid.New(), Position: 3, Name: "Item C"},
    }

    // Move "Item B" to the top
    err := os.Top(items, items[1].GetID())
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Move "Item C" below "Item A"
    err = os.Below(items, items[2].GetID(), items[0].GetID())
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Print the items after reordering
    fmt.Println("Items after reordering:")
    for _, item := range items {
        fmt.Printf("Position: %d, Name: %s\n", item.GetPosition(), item.Name)
    }
}
```

**Output:**

```
Items after reordering:
Position: 1, Name: Item B
Position: 2, Name: Item A
Position: 3, Name: Item C
```

## Error Handling

All methods return an error if the operation fails. Common errors include:

- `ErrItemNotFound`: The item with the specified ID was not found.
- `ErrInvalidPosition`: The specified position is out of bounds.

Example of error handling:

```go
err := os.To(items, itemID, newPosition)
if err != nil {
    if errors.Is(err, orderedcollection.ErrItemNotFound) {
        // Handle item not found
    } else if errors.Is(err, orderedcollection.ErrInvalidPosition) {
        // Handle invalid position
    } else {
        // Handle other errors
    }
}
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

