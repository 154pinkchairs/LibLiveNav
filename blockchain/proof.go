package blockchain

import (
	"github.com/consensys/gnark"
)

//declare a public key
func PublicKey() *gnark.PublicKey {
	return gnark.NewPublicKey()
}

// generate a random secret key
func RandomSecretKey() *gnark.SecretKey {
	return gnark.NewSecretKey()
}

//Implement a basic zero knowledge proof for the blockchain
func Proof(pk *gnark.PublicKey, sk *gnark.SecretKey, inputs []*gnark.Variable, outputs []*gnark.Variable) *gnark.Proof {
	// Create a new proof
	proof := gnark.NewProof()
	// Create a new circuit
	circuit := gnark.NewCircuit()
	// Add the circuit to the proof
	proof.AddCircuit(circuit)
	// Add the public key to the circuit
	circuit.AddPublicKey(pk)
	// Add the secret key to the circuit
	circuit.AddSecretKey(sk)
	// Add the inputs to the circuit
	circuit.AddInputs(inputs)
	// Add the outputs to the circuit
	circuit.AddOutputs(outputs)
	// Add the constraints to the circuit
	circuit.AddConstraints()
	// Add the circuit to the proof
	proof.AddCircuit(circuit)
	// Return the proof
	return proof
}

// Verify a proof
func Verify(pk *gnark.PublicKey, proof *gnark.Proof) bool {
	// Verify the proof
	return proof.Verify(pk)
}
