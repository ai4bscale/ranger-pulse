// Package version holds build-time version info shared by every
// Ranger-Pulse binary (agent, hub). It lives in internal/ so nothing
// outside this module can import it — Go enforces that at compile time.
package version

// These are overridden at build time via -ldflags, e.g.:
//   go build -ldflags "-X github.com/ai4bscale/ranger-pulse/internal/version.Version=v0.1.0"
var (
	Version = "dev"
	Commit  = "unknown"
)

// String returns a human-readable version string, e.g. "dev (unknown)".
func String() string {
	return Version + " (" + Commit + ")"
}
