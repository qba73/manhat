package manhat

import (
	"flag"
	"fmt"
	"io"
)

var version, vcsref, buildtime string

// Cli knows hwo to run the app with provided arguments and output
// results to a given io.Writer.
func Cli(args []string, output io.Writer) error {
	flagset := flag.NewFlagSet("manhat", flag.ExitOnError)

	printVersion := flagset.Bool("version", false, "show the version of the manhat app: manhat -version")
	location := flagset.Int("location", 0, "calculate Manhattan-Distance from given location to the center: manhat -location 12")

	flagset.Parse(args)
	if *printVersion {
		fmt.Fprintf(output, "Version: %s\nGitRef: %s\nBuild Time: %s\n", version, vcsref, buildtime)
		return nil
	}

	// bail in case the flag is not provided or lt 0
	if *location == 0 || *location < 0 {
		flagset.Usage()
		return nil
	}

	distance, err := CalculateDistance(*location)
	if err != nil {
		return err
	}

	fmt.Fprintf(output, "%d\n", distance)
	return nil
}
