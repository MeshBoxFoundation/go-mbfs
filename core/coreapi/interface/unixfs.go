package iface

import (
	"context"
	"io"

	options "mbfs/go-mbfs/core/coreapi/interface/options"

	files "mbfs/go-mbfs/gx/QmZMWMvWMVKCbHetJ4RgndbuEF1io2UpUxwQwtNjtYPzSC/go-ipfs-files"
	ipld "mbfs/go-mbfs/gx/QmcKKBwfz6FyQdHR2jsXrrF6XeSBXYL86anmWNewpFpoF5/go-ipld-format"
)

// TODO: ideas on making this more coreapi-ish without breaking the http API?
type AddEvent struct {
	Name  string
	Hash  string `json:",omitempty"`
	Bytes int64  `json:",omitempty"`
	Size  string `json:",omitempty"`
}

type UnixfsFile interface {
	files.SizeFile
	io.Seeker
}

// UnixfsAPI is the basic interface to immutable files in IPFS
// NOTE: This API is heavily WIP, things are guaranteed to break frequently
type UnixfsAPI interface {
	// Add imports the data from the reader into merkledag file
	//
	// TODO: a long useful comment on how to use this for many different scenarios
	Add(context.Context, files.File, ...options.UnixfsAddOption) (ResolvedPath, error)

	// Get returns a read-only handle to a file tree referenced by a path
	//
	// Note that some implementations of this API may apply the specified context
	// to operations performed on the returned file
	//Get(context.Context, Path) (UnixfsFile, error)
	// added by vingo
	Get(context.Context, Path, []byte) (UnixfsFile, error)

	// Ls returns the list of links in a directory
	Ls(context.Context, Path) ([]*ipld.Link, error)
}
