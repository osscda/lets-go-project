package stringlist

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
}

// NewList creates a new list. It returns the first element (Item) in the list
func NewList(firstElement string) Item {
	itemToReturn := Item{
		value: firstElement,
		// we don't actually need this, because it defaults to nil
		// if we don't pass it
		// NOTE: if we ever need to use this pointer in any method,
		// we _always_ need to check if it's nil beforehand
		next: nil, // we need the comma at the end
	}
	// TODO: next week, continue our pointer journey here

	return itemToReturn
}

// GetValue returns the string at the given index. It returns
// an empty string ("") if there is no element at that index
// (e.g. the list is not as long as the index or you passed a
// negative number). The list starts at index 0, so if you passed
// in 5 for example, you are telling GetValue to get the sixth item
// of the list
func (item Item) GetValue(idx int) string {
	// loop all the way
	currentItem := item
	for currentIndex := 0; currentIndex < idx; currentIndex++ {

	}
	return currentItem.value
}
