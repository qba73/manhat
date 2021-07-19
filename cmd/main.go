package main

import (
	"log"
	"os"

	"github.com/qba73/manhat"
)

// This vars represent build metadata generated
// when the binary is built. Values are assigned
// using build flags (ldflags) at the build time.
var (
	Version   = ""
	VcsRef    = ""
	BuildTime = ""
)

func main() {
	meta := manhat.MetaInfo{
		Version:   Version,
		VcsRef:    VcsRef,
		BuildTime: BuildTime,
	}

	if err := manhat.Cli(os.Args[1:], os.Stdout, meta); err != nil {
		log.Fatal(err)
	}
}
