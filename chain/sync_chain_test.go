package chain

import (
	"context"
	"fmt"
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/vm"
	"github.com/filecoin-project/lotus/extern/sector-storage/ffiwrapper"
	"github.com/filecoin-project/lotus/journal"
	"github.com/filecoin-project/lotus/node/repo"
	"github.com/ipfs/go-cid"
	"testing"
)

func TestSync(t *testing.T) {
	r, err := repo.NewFS("/Users/didi/.lotus")
	if err != nil {
		panic(err)
	}

	lr, err := r.Lock(repo.FullNode)
	if err != nil {
		panic(err)
	}

	bs, err := lr.Blockstore(context.TODO(), repo.UniversalBlockstore)
	if err != nil {
		panic(err)
	}

	mds, err := lr.Datastore(context.TODO(), "/metadata")
	if err != nil {
		panic(err)
	}

	j, err := journal.OpenFSJournal(lr, journal.EnvDisabledEvents())
	if err != nil {
		panic(err)
	}

	cst := store.NewChainStore(bs, bs, mds, vm.Syscalls(ffiwrapper.ProofVerifier), j)
	stmt := stmgr.NewStateManager(cst)

	genesis, err := stmt.ChainStore().GetGenesis()
	if err != nil {
		panic(err)
	}
	fmt.Println("genesis", genesis.Height)
	c, err := cid.Parse("bafy2bzacebl3zdyqdxpl4pq2ntxzeklqo3y3ph7dsnyfqjf2xmmznmks35ti6")
	if err != nil {
		panic(err)
	}
	cids := []cid.Cid{c}

	blocks := make([]*types.BlockHeader, 0, 1024)
	LOOP:
	for {
		var p []cid.Cid
		for _, cc := range cids {
			b, err := stmt.ChainStore().GetBlock(cc)
			if err != nil {
				panic(err)
			}
			blocks = append(blocks, b)
			p = b.Parents
			if b.Height == 0 {
				break LOOP
			}
		}
		cids = p
	}

	//for _, block := range blocks {
	//	fmt.Println("block ", block.Height)
	//}
	var w int = 0

	for i := len(blocks) - 2; i >= 0 ; i -- {
		block := blocks[i]
		if block.Height < 100 {
			continue;
		}
		baseTs, err := stmt.ChainStore().LoadTipSet(types.NewTipSetKey(block.Parents...))
		fmt.Println("ready loading block", block.Height, baseTs.Cids())
		if err != nil {
			panic(err)
		}
		stateroot, precp, err := stmt.TipSetState(context.TODO(), baseTs)
		if err != nil {
			panic(err)
		}
		fmt.Println("stateroot", stateroot)
		fmt.Println("block.parentstateroot", block.ParentStateRoot)

		if stateroot != block.ParentStateRoot {
			panic(fmt.Sprintf("parent state root did not match computed state (%s != %s)", stateroot, block.ParentStateRoot))
		}

		if precp != block.ParentMessageReceipts {
			panic(fmt.Sprintf("parent receipts root did not match computed value (%s != %s)", precp, block.ParentMessageReceipts))
		}
		w = w + 1
		if w >= 2 {
			return;
		}
		break;
	}
}

