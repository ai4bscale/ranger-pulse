// Command rp-hub is the Ranger-Pulse Hub: it authenticates agents and
// users, stores reports, and exposes the API the UI is built on.
//
// This is a placeholder entry point — the API, storage, and auth arrive
// with specs/001-mvp-vertical-slice.
package main

import (
	"fmt"

	"github.com/ai4bscale/ranger-pulse/internal/version"
)

func main() {
	fmt.Println("rp-hub", version.String())
}
