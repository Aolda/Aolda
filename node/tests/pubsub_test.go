package tests

import (
	pusbsub "aolda_node/p2p"
	"testing"
)

func TestPub() {
	PubForTx()
	PubForBlock()
	PubForBlocks()
}

func TestSub() {

}
