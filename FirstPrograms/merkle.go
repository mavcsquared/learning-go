package main

import (
	"fmt"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark/frontend"
)

type merkleCircuit struct {
	rootHash    frontend.Variable `gnark:",public"`
	path        []int
	leaf        string `gnark:",public"`
	index, size int    `gnark:",public"`
	leaves      []byte
}

func (circuit *merkleCircuit) Define(api frontend.API) error {

	circuit.path = getPath(circuit.index, circuit.size)

	mimc := mimc.NewMiMC()
	mimc.Write([]byte(circuit.leaf))
	actual := mimc.Sum(nil)

	newRootHash := circuit.generateProof(circuit.index, actual)

	api.AssertIsEqual(circuit.rootHash, newRootHash)

	return nil
}

func (circuit *merkleCircuit) generateProof(index int, actual []byte) []byte {
	mimc := mimc.NewMiMC()

	if index == len(circuit.path) {
		return actual
	}

	var otherByte []byte

	if circuit.path[index] > 0 {
		otherByte = getHash(index, circuit.size, circuit.leaves)
		mimc.Write(append(actual, otherByte...))
	} else {
		otherByte = getHash(-index, circuit.size, circuit.leaves)
		mimc.Write(append(otherByte, actual...))
	}

	actual = mimc.Sum(nil)

	return circuit.generateProof(index+1, actual)
}

func getPath(index int, size int) []int {
	var path []int

	index = index + size

	for {
		if index == 1 {
			break
		}
		if index%2 == 0 {
			path = append(path, index+1)

		} else {
			path = append(path, -index+1)
		}

		index /= 2
	}

	return path
}

func getHash(index int, size int, leaves []byte) []byte {
	mimc := mimc.NewMiMC()

	if index >= size {
		fmt.Println([]byte{leaves[index-size]})
		mimc.Write([]byte{leaves[index-size]})
		println("LLEGA")
	} else {
		byteLeft := getHash(index*2, size, leaves)
		byteRight := getHash(index*2+1, size, leaves)

		thisByte := append(byteLeft, byteRight...)
		mimc.Write(thisByte)
	}

	return mimc.Sum(nil)
}

// /*
func main() {
	fmt.Println("BEGIN")

	str := "abcdefgh"
	byteFromStr := []byte(str)

	mimc.NewMiMC().Write(byteFromStr[1:2])
}

// */
