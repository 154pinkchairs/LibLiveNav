package blockchain_test

import (
	"fmt"

	"github.com/154pinkchairs/LibLiveNav/blockchain"
)

func main() {
	if err := blockchain.ValidateProof([]byte("foo"), []byte("bar")); err != nil {
		fmt.Println(err)
	}
}
