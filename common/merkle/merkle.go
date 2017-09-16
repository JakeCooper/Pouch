package merkle

import (
	"io"

	"github.com/erikreppel/pouch/constants"
	"github.com/pkg/errors"
)

type MerkleString string

func (m MerkleString) Bytes() []byte {
	b, _ := constants.Encoder.DecodeString(string(m))
	return b
}

// MerkleTree is used to compose state
type MerkleTree struct {
	RootNode  *Node
	LeafNodes []*Node
}

// RootHash returns the state of the tree
func (m *MerkleTree) RootHash() (string, error) {
	return m.RootNode.ComputeSum()
}

// newLeafNode returns a new leaf node for the tree
func newLeafNode(i []byte) (*Node, error) {
	n := &Node{}
	h := constants.NewHash()
	if _, err := h.Write(i); err != nil {
		return nil, err
	}
	sum := h.Sum(nil)
	n.Sum = constants.Encoder.EncodeToString(sum)
	return n, nil
}

// Node is a node in a merkletree
type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Sum    string
}

// IsLeaf confirms if the node is a lead
func (n *Node) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

// ComputeSum returns the hash of data or computes down tree
func (n *Node) ComputeSum() (string, error) {
	if n.IsLeaf() {
		if n.Sum == "" {
			return "", errors.New("sum not set in leaf")
		}
		return n.Sum, nil
	}

	lSum, err := n.Left.ComputeSum()
	if err != nil {
		return "", nil
	}
	rSum, err := n.Right.ComputeSum()
	if err != nil {
		return "", nil
	}

	h := constants.NewHash()
	childSum := lSum + rSum

	if _, err := io.WriteString(h, childSum); err != nil {
		return "", err
	}
	sum := h.Sum(nil)
	return constants.Encoder.EncodeToString(sum), nil
}

// BuildMerkleTree creates a tree from Leaf
func buildMerkleTreeFromNodes(leafNodes ...*Node) *MerkleTree {
	root := createRoot(leafNodes...)
	return &MerkleTree{
		LeafNodes: leafNodes,
		RootNode:  root,
	}
}

func createRoot(leafNodes ...*Node) *Node {
	var nodes []*Node
	for i := 0; i < len(leafNodes)-1; i = i + 2 {
		parent := &Node{}
		parent.Left = leafNodes[i]
		parent.Right = leafNodes[i+1]

		parent.Right.Parent = parent
		parent.Left.Parent = parent
		nodes = append(nodes, parent)
	}
	if len(nodes) == 1 {
		return nodes[0]
	}
	return createRoot(nodes...)
}

// TreeFromStrings takes strings a returns a tree
func TreeFromStrings(datas []string) (*MerkleTree, error) {
	leafNodes := []*Node{}
	for _, data := range datas {
		b := []byte(data)
		n, err := newLeafNode(b)
		if err != nil {
			return nil, err
		}
		leafNodes = append(leafNodes, n)
	}
	return buildMerkleTreeFromNodes(leafNodes...), nil
}

// CreateMerkleTree returns a merkle tree from an interface
func CreateMerkleTree(items [][]byte) (*MerkleTree, error) {
	leafNodes := []*Node{}
	for _, item := range items {
		n, err := newLeafNode(item)
		if err != nil {
			return nil, err
		}
		leafNodes = append(leafNodes, n)
	}
	return buildMerkleTreeFromNodes(leafNodes...), nil
}
