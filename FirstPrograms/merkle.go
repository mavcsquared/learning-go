package main

import (
	"fmt"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
)

type merkleCircuit struct {
	rootHash          frontend.Variable `gnark:",public"`
	path, aux, leaves []frontend.Variable
	leaf              frontend.Variable `gnark:",public"`
}

func (circuit *merkleCircuit) isLeaf(api frontend.API, leaf frontend.Variable, index int) {

}

func (circuit *merkleCircuit) Define(api frontend.API, index int) error {
	mimc, _ := mimc.NewMiMC(api)
	mimc.Write(circuit.rootHash)
	//cosas que hacer

	circuit.isLeaf(api, circuit.leaf, index)

	return nil
}

func main() {
	fmt.Println("BEGIN")

}
