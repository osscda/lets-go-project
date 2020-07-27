# String List in Go

This is a library that lets you store and manipulate a list of strings.

(This list is otherwise known as a _data structure_)

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
- Insert into the list - you need the index and the value. This is also called "set" or "replace" on the list, because you're not adding anything new, just setting a new value for an existing item
- Delete from the list - you also just need the index of the item you want to delete




TODOs:

- Introduce initial structure of the list (AKA "nodes")
- Build "constructor" for nodes (AKA `func NewList()`)
- Build `getNode(index int)` - function to get a specific node at a specific index
- Build `Get(index int)`, `Set(...)`, and `Append(...)`, based on `getNode` that we did previously