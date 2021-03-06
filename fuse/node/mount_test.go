// +build !nofuse

package node

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"context"

	core "mbfs/go-mbfs/core"
	ipns "mbfs/go-mbfs/fuse/ipns"
	mount "mbfs/go-mbfs/fuse/mount"
	namesys "mbfs/go-mbfs/namesys"

	offroute "mbfs/go-mbfs/gx/QmNuVissmH2ftUd4ADvhm9WER3351wTYduY1EeDDGtP1tM/go-ipfs-routing/offline"
	ci "mbfs/go-mbfs/gx/QmZXjR5X1p4KrQ967cTsy4MymMzUM8mZECF3PV8UcN4o3g/go-testutil/ci"
)

func maybeSkipFuseTests(t *testing.T) {
	if ci.NoFuse() {
		t.Skip("Skipping FUSE tests")
	}
}

func mkdir(t *testing.T, path string) {
	err := os.Mkdir(path, os.ModeDir|os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

// Test externally unmounting, then trying to unmount in code
func TestExternalUnmount(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	// TODO: needed?
	maybeSkipFuseTests(t)

	node, err := core.NewNode(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	err = node.LoadPrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	node.Routing = offroute.NewOfflineRouter(node.Repo.Datastore(), node.RecordValidator)
	node.Namesys = namesys.NewNameSystem(node.Routing, node.Repo.Datastore(), 0)

	err = ipns.InitializeKeyspace(node, node.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	// get the test dir paths (/tmp/fusetestXXXX)
	dir, err := ioutil.TempDir("", "fusetest")
	if err != nil {
		t.Fatal(err)
	}

	ipfsDir := dir + "/ipfs"
	ipnsDir := dir + "/ipns"
	mkdir(t, ipfsDir)
	mkdir(t, ipnsDir)

	err = Mount(node, ipfsDir, ipnsDir)
	if err != nil {
		t.Fatal(err)
	}

	// Run shell command to externally unmount the directory
	cmd, err := mount.UnmountCmd(ipfsDir)
	if err != nil {
		t.Fatal(err)
	}

	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	// TODO(noffle): it takes a moment for the goroutine that's running fs.Serve to be notified and do its cleanup.
	time.Sleep(time.Millisecond * 100)

	// Attempt to unmount IPFS; it should unmount successfully.
	err = node.Mounts.Ipfs.Unmount()
	if err != mount.ErrNotMounted {
		t.Fatal("Unmount should have failed")
	}
}
