package hashMap

// https://leetcode.com/problems/design-hashmap/

type KeyValuePair struct {
	Key, Value int
}

type ListNode struct {
	Value *KeyValuePair
	Next  *ListNode
}

type MyHashMap struct {
	array []*ListNode
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	hashMap := MyHashMap{
		make([]*ListNode, 7919),
	}
	return hashMap
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	ind := key % len(this.array)

	curr := this.array[ind]
	for curr != nil {
		if curr.Value.Key == key {
			curr.Value.Value = value
			return
		}
		curr = curr.Next
	}

	newHead := &ListNode{Value: &KeyValuePair{key, value}}
	newHead.Next = this.array[ind]
	this.array[ind] = newHead
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	ind := key % len(this.array)

	curr := this.array[ind]
	for curr != nil {
		if curr.Value.Key == key {
			return curr.Value.Value
		}
		curr = curr.Next
	}

	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	ind := key % len(this.array)

	curr := this.array[ind]
	if curr == nil {
		return
	}

	if curr.Value.Key == key {
		this.array[ind] = curr.Next
	}

	for curr.Next != nil {
		if curr.Next.Value.Key == key {
			curr.Next = curr.Next.Next
			return
		}
		curr = curr.Next
	}
}
