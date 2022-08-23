assert := groth16.NewAssert(t)

var mimcCircuit Circuit

{
    assert.ProverFailed(&mimcCircuit, &Circuit{
        Hash: 42,
        PreImage: 42,
    })
}

{
     assert.ProverSucceeded(&mimcCircuit, &Circuit{
        Hash: b.Hash,
        PreImage: 35,
    })
}

