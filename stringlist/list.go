package stringlist

import "fmt"

// Item is a single element in the list. It is a struct that holds the value (string) in
// that spot in the list, and it holds a _pointer_ to the next node in the list
//
//
//
//
//
// node1 ("abcd", next -> node2), node2 ("defg", next -> node3), node3 ("hijk", next -> nil)
//
// ["open flour", "sift flour", "add baking soda"]
//    ^^ node1       ^^ node2      ^^ node3
type Item struct {
	value string
	next  *Item
	// instead of the endless loop in Append, we could have
	// kept track of the total length of the list, and
	// then looped from zero to the end of the loop,
	// just "traversing" to the next element each time.
	// that would have found us the end of the list also
	// if would look like this:
	//
	// totalLength int
}

// NewList creates a new list. It returns the first element (Item) in the list
//
// TODO: allow the caller to create an empty list.
// we should do this without breaking backward compatibility
func NewList(firstElement string) *Item {
	itemToReturn := Item{
		value: firstElement,
		// we don't actually need this, because it defaults to nil
		// if we don't pass it
		// NOTE: if we ever need to use this pointer in any method,
		// we _always_ need to check if it's nil beforehand
		next: nil, // we need the comma at the end
	}
	// TODO: next week, continue our pointer journey here

	return &itemToReturn
}

// GetValue returns the string at the given index. It returns
// an empty string ("") if there is no element at that index
// (e.g. the list is not as long as the index or you passed a
// negative number). The list starts at index 0, so if you passed
// in 5 for example, you are telling GetValue to get the sixth item
// of the list
func (item *Item) GetValue(idx int) string {
	// loop all the way
	currentItem := item
	currentIndex := 0

	for {
		if currentIndex == idx {
			// we've reached the element that we want,
			// so return the currentItem right here
			return currentItem.value
		}
		currentItem = currentItem.next
		if currentItem == nil {
			// if we get here, it means that we are at the end
			// of the list and we haven't found the item
			// we want yet, so return "not found"
			return ""
		}
		currentIndex++
	}
}

// Append adds the given string to the end of the list
func (item *Item) Append(str string) {
	// start at the first item
	currentItem := item
	// instead of the endless loop, we could have
	// kept track of the total length of the list, and
	// then looped from zero to the end of the loop,
	// just "traversing" to the next element each time
	for {
		// "peek" ahead and see if there is a next item.
		// if not, bail out of the loop
		if currentItem.next == nil {
			// we hit the end of the list, but we don't
			// to bail out of the _function_. we just want
			// to exit the loop
			break
		}
		currentItem = currentItem.next
	}
	// we now know that currentItem is the last
	// item in the list, _and_ that currentItem.next is nil
	// this is the 💖
	newItem := Item{
		value: str,
		next:  nil,
	}
	// we're about to add our 💖 to the existing list 🎇
	// 💖 + 🎇
	currentItem.next = &newItem
}

// if we wanted to keep track of total length, and loop
// to the end of the list using that (instead of an
// endless loop like we do in Append), this is approximately
// how we'd do it.
//
// first item in list
// 	value: "abc",
// 	next:  address of second item in list,
// 	totalLength: 2
// second item in list (after append)
// 	value: "def"
// 	next: nil
// 	totalLength: 0 // always ignored

// IndexNotFoundError is returned by stringlist functions when
// you pass an index that doesn't exist on the list. For example,
// if the list is 3 elements long and you pass _index_ 3 (one
// item off the end of the list, because the list starts
// at index 0), the function will return this error
type IndexNotFoundError struct {
	index int
}

// Error implements the 'error' interface for this type
func (e IndexNotFoundError) Error() string {
	return fmt.Sprintf("Index %d not found", e.index)
}

// Replace overwrites the existing value at the given
// index with newVal
func (item *Item) Replace(index int, newVal string) error {
	currentItem := item
	for i := 0; i < index; i++ {
		// if we hit this, we are at the end of the list
		// but haven't gotten to the index requested yet.
		// that means they passed an invalid index
		if currentItem.next == nil {
			return IndexNotFoundError{index: index}
		}
		currentItem = currentItem.next
	}
	// currentItem is the item at the requested index
	currentItem.value = newVal
	return nil
}
