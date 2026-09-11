// Command rp-agent is the Ranger-Pulse agent: it runs on each monitored
// machine, discovers installed plugins, executes their collectors, and
// reports results to the Hub.
//
// This is a placeholder entry point — plugin discovery, collection, and
// reporting arrive with specs/001-mvp-vertical-slice.
package main

import (
	"fmt"

	"github.com/ai4bscale/ranger-pulse/internal/version"
)

func main() {
	fmt.Println("rp-agent", version.String())
}
