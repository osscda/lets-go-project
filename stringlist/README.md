# String List in Go

This is a library that lets you store and manipulate a list of strings.

(This list is otherwise known as a _data structure_)

>We're using a linked list architecture for this. See [Wikipedia](https://en.wikipedia.org/wiki/Linked_list) for that.

## What is a List?

Lists are an ordered sequence of items. You can access any of the items you want by giving the list an _index_, which is a unique identifier for each item in the list.

The index for any specific item in the list is a number. It represents the "distance" from the start of the list that the item is.

This is usually how you write a list:

```
["open flour", "sift flour", "add baking soda"]
```

You can do several operations on lists:

- Access any item if you have the index - you just need the index, and it returns the value of the item at that index
- Add to the list - you need just the value of the new item (sometimes also called `append` or `prepend`, depending on where you want to add it in the list)
- Replace an item in the list - you need the index and the _new_ value. This is also called "set", because you're not adding anything new, just setting a new value for an existing item
- Delete from the list - you also just need the index of the item you want to delete
- Insert into the list - add a _new_ element into the list, and shift everything after that element to the right


- Introduced initial structure of the list (AKA `Item`)
- Built "constructor" for items (AKA `func NewList()`)
- Built `getValue(index int)` - function to get a specific item at a specific index
- Built `Append(...)`, based on `getValue` that we did previously
- Built `Replace(...)`
- Built `Insert(...)`
- Add documentation

TODOs:

- Build `Delete(...)`
- Add Tests
- Update NewList to allow creating an empty list. 

## About this Project

This is a project built in collaboration with the community on the live stream "Let's Go" with @arschles and @iennae. For schedule details check out the [Let's Go repository](https://github.com/osscda/lets-go-stream). If you'd like to contribute an idea for a topic or project, submit [an issue](https://github.com/osscda/lets-go-stream/issues). 
