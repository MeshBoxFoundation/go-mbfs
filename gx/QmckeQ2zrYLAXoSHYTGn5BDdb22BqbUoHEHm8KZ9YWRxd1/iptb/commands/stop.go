package commands

import (
	"context"
	"fmt"
	"path"

	cli "mbfs/go-mbfs/gx/Qmc1AtgBdoUHP8oYSqU81NRYdzohmF45t5XNwVMvhCxsBA/cli"

	"mbfs/go-mbfs/gx/QmckeQ2zrYLAXoSHYTGn5BDdb22BqbUoHEHm8KZ9YWRxd1/iptb/testbed"
	"mbfs/go-mbfs/gx/QmckeQ2zrYLAXoSHYTGn5BDdb22BqbUoHEHm8KZ9YWRxd1/iptb/testbed/interfaces"
)

var StopCmd = cli.Command{
	Category:  "CORE",
	Name:      "stop",
	Usage:     "stop specified nodes (or all)",
	ArgsUsage: "[nodes]",
	Action: func(c *cli.Context) error {
		flagRoot := c.GlobalString("IPTB_ROOT")
		flagTestbed := c.GlobalString("testbed")
		flagQuiet := c.GlobalBool("quiet")

		tb := testbed.NewTestbed(path.Join(flagRoot, "testbeds", flagTestbed))
		nodes, err := tb.Nodes()
		if err != nil {
			return err
		}

		nodeRange := c.Args().First()

		if nodeRange == "" {
			nodeRange = fmt.Sprintf("[0-%d]", len(nodes)-1)
		}

		list, err := parseRange(nodeRange)
		if err != nil {
			return fmt.Errorf("could not parse node range %s", nodeRange)
		}

		runCmd := func(node testbedi.Core) (testbedi.Output, error) {
			return nil, node.Stop(context.Background())
		}

		results, err := mapWithOutput(list, nodes, runCmd)
		if err != nil {
			return err
		}

		return buildReport(results, flagQuiet)
	},
}
