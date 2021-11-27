package tree

import "errors"

/*
cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
BenchmarkTwoTree-8                   100          15831217 ns/op
BenchmarkTenTree-8                   451           2316224 ns/op
BenchmarkShallowTree-8               663           1774259 ns/op
PASS
ok      tree    4.271s
*/

// Define the Record type
type Record struct {
	ID     int
	Parent int
}

// Define the Node type
type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	var root *Node
	reclen := len(records)
	sortedR := make([]Record, reclen)
	nodes := make([]*Node, reclen)

	// sort records and check errors
	root, err := createNodeSlice(records, sortedR, reclen)
	if err != nil {
		return root, err
	}

	buildTree(root, nodes, sortedR, reclen)

	return root, nil
}

func buildTree(root *Node, nodes []*Node, r []Record, reclen int) {
	if reclen == 1 {
		return
	}
	for i := 1; i < reclen; i++ {
		item := r[i]
		if nodes[item.ID] == nil {
			nodes[item.ID] = &Node{ID: item.ID}
		}
		if item.Parent == 0 {
			root.Children = append(root.Children, nodes[item.ID])
		} else {
			if nodes[item.Parent] == nil {
				nodes[item.Parent] = &Node{ID: item.Parent}
			}
			node := nodes[item.Parent]
			node.Children = append(node.Children, nodes[item.ID])
		}
	}
}

func createNodeSlice(records []Record, sortedR []Record, reclen int) (*Node, error) {
	if reclen == 0 {
		return nil, nil
	}

	maxID := 0
	for _, item := range records {
		if item.ID > maxID {
			maxID = item.ID
		}
		if item.ID > reclen-1 {
			return nil, errors.New("non-continuous.")
		}
		if item.ID != 0 && item.ID == item.Parent {
			return nil, errors.New("cycle directly.")
		}
		if item.ID < item.Parent {
			return nil, errors.New("higher id parent of lower id.")
		}
		sortedR[item.ID] = item
	}
	if maxID < reclen-1 {
		return nil, errors.New("duplicate nodes.")
	}
	if sortedR != nil && sortedR[0].Parent != 0 {
		return nil, errors.New("root has parent node.")
	}
	return &Node{}, nil
}
