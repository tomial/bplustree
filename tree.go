package bplustree

type BTree struct {
	root     *interiorNode
	leaf     int
	interior int
	height   int
}

func newBTree() *BTree {
	leaf := newLeafNode(nil)
	r := newInteriorNode(nil, leaf)
	leaf.p = r
	return &BTree{
		root:     r,
		leaf:     1,
		interior: 1,
		height:   2,
	}
}

// func (bt *BTree) search(key int) (int, *leafNode, ok) {
// 	i, ok := n.find(key)
// 	if !ok {
// 		return nil, n.(*leafNode)
// 	}

// 	switch t := n.(type) {
// 	case *leafNode:
// 		return &t.kvs[i], t
// 	case *interiorNode:
// 		return search(t.kcs[i].child, key)
// 	default:
// 		panic("wrong type")
// 	}
// }

func (bt *BTree) insert(key int, value string) {
	_, leaf := search(bt.root, key)
	p := leaf.parent()
	oldIndex, _ := p.find(key)

	mid, bump := leaf.insert(key, value)

	if !bump {
		return
	}

	var midNode node
	midNode = leaf

	p.kcs[oldIndex].child = leaf.next
	leaf.next.setParent(p)

	interior := p
	interiorP := p.parent()

	for {
		var oldIndex int
		var newNode *interiorNode

		isRoot := interiorP == nil

		if !isRoot {
			oldIndex, _ = interiorP.find(key)
		}

		mid, newNode, bump = interior.insert(mid, midNode)
		if !bump {
			return
		}

		if !isRoot {
			interiorP.kcs[oldIndex].child = newNode
			newNode.setParent(interiorP)

			midNode = interior

			interior = interiorP
			interiorP = interior.parent()
		} else {
			bt.root = newInteriorNode(nil, newNode)
			newNode.p = bt.root
			bt.root.insert(mid, interior)
			return
		}
	}
}