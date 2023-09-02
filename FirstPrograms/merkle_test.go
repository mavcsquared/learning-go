package main

import (
	"testing"

	"github.com/consensys/gnark/test"
)

func TestMerkle(t *testing.T) {
	assert := test.NewAssert(t)

	var merkle merkleCircuit

	str := "abcdefgh"

	assert.ProverFailed(&merkle, &merkleCircuit{
		leaf:     "c",
		index:    1,
		size:     len(str),
		rootHash: getHash(1, len(str), []byte(str)),
	})

	assert.ProverSucceeded(&merkle, &merkleCircuit{
		leaf:     "c",
		index:    3,
		size:     len(str),
		rootHash: getHash(3, len(str), []byte(str)),
	})

}
