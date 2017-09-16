package merkle

import "testing"
import "strconv"
import "time"

func makeTree(n int) *MerkleTree {
	strs := []string{}
	for i := 2; i < n; i++ {
		x := i * 7777777777
		strs = append(strs, strconv.Itoa(x))
	}
	t, _ := MerkleTreeFromStrings(strs)
	return t
}

func TestMerkleTree(t *testing.T) {
	tree, err := MerkleTreeFromStrings([]string{
		"abc",
		"xyz",
		"tree",
	})
	if err != nil {
		t.Fail()
	}
	s, err := tree.RootHash()
	if err != nil {
		t.Fail()
	}
	t.Log(s)
}

func TestMerkleTreeLarge(t *testing.T) {
	nNodes := 1000000
	start := time.Now()
	tree := makeTree(nNodes)
	h, err := tree.RootHash()
	if err != nil {
		t.Fail()
	}
	t.Log("Hash of", h)
	t.Log("took", time.Now().Sub(start), "to make a tree of", nNodes)
}
