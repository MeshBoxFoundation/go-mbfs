package blockstore

import (
	"context"

	cid "mbfs/go-mbfs/gx/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
	blocks "mbfs/go-mbfs/gx/QmWoXtvgC8inqFkAATB7cp2Dax7XBi9VDvSg9RCCZufmRk/go-block-format"
	mh "mbfs/go-mbfs/gx/QmerPMzPk1mJVowm8KgmoknWa4yCYvvugMPsgWmDNUvDLW/go-multihash"
)

// idstore wraps a BlockStore to add support for identity hashes
type idstore struct {
	bs Blockstore
}

func NewIdStore(bs Blockstore) Blockstore {
	return &idstore{bs}
}

func extractContents(k cid.Cid) (bool, []byte) {
	dmh, err := mh.Decode(k.Hash())
	if err != nil || dmh.Code != mh.ID {
		return false, nil
	}
	return true, dmh.Digest
}

func (b *idstore) DeleteBlock(k cid.Cid) error {
	isId, _ := extractContents(k)
	if isId {
		return nil
	}
	return b.bs.DeleteBlock(k)
}

func (b *idstore) Has(k cid.Cid) (bool, error) {
	isId, _ := extractContents(k)
	if isId {
		return true, nil
	}
	return b.bs.Has(k)
}

func (b *idstore) GetSize(k cid.Cid) (int, error) {
	isId, bdata := extractContents(k)
	if isId {
		return len(bdata), nil
	}
	return b.bs.GetSize(k)
}

func (b *idstore) Get(k cid.Cid) (blocks.Block, error) {
	isId, bdata := extractContents(k)
	if isId {
		return blocks.NewBlockWithCid(bdata, k)
	}
	return b.bs.Get(k)
}

func (b *idstore) Put(bl blocks.Block) error {
	isId, _ := extractContents(bl.Cid())
	if isId {
		return nil
	}
	return b.bs.Put(bl)
}

func (b *idstore) PutMany(bs []blocks.Block) error {
	toPut := make([]blocks.Block, 0, len(bs))
	for _, bl := range bs {
		isId, _ := extractContents(bl.Cid())
		if isId {
			continue
		}
		toPut = append(toPut, bl)
	}
	return b.bs.PutMany(toPut)
}

func (b *idstore) HashOnRead(enabled bool) {
	b.bs.HashOnRead(enabled)
}

func (b *idstore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	return b.bs.AllKeysChan(ctx)
}
