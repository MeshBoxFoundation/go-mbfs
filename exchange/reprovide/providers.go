package reprovide

import (
	"context"

	pin "mbfs/go-mbfs/pin"

	cid "gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
	blocks "gx/ipfs/QmSNLNnL3kq3A1NGdQA9AtgxM9CWKiiSEup3W435jCkRQS/go-ipfs-blockstore"
	merkledag "gx/ipfs/QmaDBne4KeY3UepeqSVKYpSmQGa3q9zP6x3LfVF2UjF3Hc/go-merkledag"
	cidutil "gx/ipfs/QmbfKu17LbMWyGUxHEUns9Wf5Dkm8PT6be4uPhTkk4YvaV/go-cidutil"
	ipld "gx/ipfs/QmcKKBwfz6FyQdHR2jsXrrF6XeSBXYL86anmWNewpFpoF5/go-ipld-format"
)

// NewBlockstoreProvider returns key provider using bstore.AllKeysChan
func NewBlockstoreProvider(bstore blocks.Blockstore) KeyChanFunc {
	return func(ctx context.Context) (<-chan cid.Cid, error) {
		return bstore.AllKeysChan(ctx)
	}
}

// NewPinnedProvider returns provider supplying pinned keys
func NewPinnedProvider(pinning pin.Pinner, dag ipld.DAGService, onlyRoots bool) KeyChanFunc {
	return func(ctx context.Context) (<-chan cid.Cid, error) {
		set, err := pinSet(ctx, pinning, dag, onlyRoots)
		if err != nil {
			return nil, err
		}

		outCh := make(chan cid.Cid)
		go func() {
			defer close(outCh)
			for c := range set.New {
				select {
				case <-ctx.Done():
					return
				case outCh <- c:
				}
			}

		}()

		return outCh, nil
	}
}

func pinSet(ctx context.Context, pinning pin.Pinner, dag ipld.DAGService, onlyRoots bool) (*cidutil.StreamingSet, error) {
	set := cidutil.NewStreamingSet()

	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		defer close(set.New)

		for _, key := range pinning.DirectKeys() {
			set.Visitor(ctx)(key)
		}

		for _, key := range pinning.RecursiveKeys() {
			set.Visitor(ctx)(key)

			if !onlyRoots {
				err := merkledag.EnumerateChildren(ctx, merkledag.GetLinksWithDAG(dag), key, set.Visitor(ctx))
				if err != nil {
					log.Errorf("reprovide indirect pins: %s", err)
					return
				}
			}
		}
	}()

	return set, nil
}
