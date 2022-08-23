package blockchain

import (
  "github.com/consensys/gnark"
)

// Circuit defines a pre-image knowledge proof
// mimc(secret preImage) = public hash
type Circuit struct {
    PreImage frontend.Variable
    Hash     frontend.Variable `gnark:",public"`
}

// Define declares the circuit's constraints
func (circuit *Circuit) Define(api frontend.API) error {
    // hash function
    mimc, err := mimc.NewMiMC(api.Curve())

    // specify constraints
    // mimc(preImage) == hash
    api.AssertIsEqual(circuit.Hash, mimc.Hash(cs, circuit.PreImage))

    return nil
}

var mimcCircuit Circuit
r1cs, err := frontend.Compile(ecc.BN254, r1cs.NewBuilder, &mimcCircuit)

assignment := &Circuit{
    Hash: b.Hash,
    PreImage: 35,
}
witness, _ := frontend.NewWitness(assignment, ecc.BN254)
publicWitness, _ := witness.Public()
pk, vk, err := groth16.Setup(r1cs)
proof, err := groth16.Prove(r1cs, pk, witness)
err := groth16.Verify(proof, vk, publicWitness)
