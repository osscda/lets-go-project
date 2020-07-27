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
		next: nil, // we need the comma at the end
	}
	// TODO: next week, continue our pointer journey here

	return itemToReturn
}
