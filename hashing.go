package bbolt

import "crypto/sha256"

func CalculatingHashing(n *node) {
	// t: true for leaf node, false for non leaf node
	hashingBytes := make([]byte, 0)
	for _, v := range n.inodes {
		_assert(len(v.hash) == 32, "%v hashing size is %v", len(v.hash))
		hashingBytes = append(hashingBytes, v.hash...)
	}
	hash := sha256.Sum256(hashingBytes)
	n.hash = hash[:]
}
