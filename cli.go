package manhat

import (
	"flag"
	"fmt"
	"io"
)

// MetaInfo holds information about cli binary.
// Values are populated from ldflags at build time.
type MetaInfo struct {
	Version   string
	VcsRef    string
	BuildTime string
}

// Cli knows hwo to run the app with provided arguments and output
// results to a given io.Writer.
func Cli(args []string, output io.Writer, meta MetaInfo) error {
	flagset := flag.NewFlagSet("manhat", flag.ExitOnError)

	printVersion := flagset.Bool("version", false, "show the version of the manhat app: manhat -version")
	location := flagset.Int("location", 0, "calculate Manhattan-Distance from given location to the center: manhat -location 12")

	flagset.Parse(args)
	if *printVersion {
		fmt.Fprintf(output, "Version: %s\nGitRef: %s\nBuild Time: %s\n", meta.Version, meta.VcsRef, meta.BuildTime)
		return nil
	}

	// bail in case the flag is not provided or lt 0
	if *location == 0 || *location < 0 {
		flagset.Usage()
		return nil
	}

	distance, err := CalculateDistance(float64(*location))
	if err != nil {
		return err
	}

	fmt.Fprintf(output, "%d\n", int(distance))
	return nil
}
