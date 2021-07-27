package manhat

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

// These vars represent build metadata generated
// when the binary is built. Values are assigned
// using build flags (ldflags) at the build time.
var version, vcsref, buildtime string

// config is a struct that holds config
// parameters for the cli.
type config struct {
	args   []string
	output io.Writer
}

type option func(*config)

// WithArgs is an option constructor.
func WithArgs(args []string) option {
	return func(c *config) {
		c.args = args
	}
}

// WithOutput is an option constructor.
func WithOutput(out io.Writer) option {
	return func(c *config) {
		c.output = out
	}
}

// Cli knows how to run the app with provided
// arguments and prints result to a given io.Writer.
func Cli(opts ...option) error {

	// Default config settings are applied when
	// no func option is provided.
	c := config{
		args:   os.Args[1:],
		output: os.Stdout,
	}

	for _, o := range opts {
		o(&c)
	}

	flagset := flag.NewFlagSet("manhat", flag.ExitOnError)

	printVersion := flagset.Bool("version", false, "show the version of the manhat app: manhat -version")
	location := flagset.Int("location", 0, "calculate Manhattan-Distance from given location to the center: manhat -location 12")

	flagset.Parse(c.args)
	flagset.SetOutput(c.output)

	if *printVersion {
		fmt.Fprintf(c.output, "Version: %s\nGitRef: %s\nBuild Time: %s\n", version, vcsref, buildtime)
		return nil
	}

	// bail in case the flag is not provided or lt 0
	if *location == 0 || *location < 0 {
		flagset.Usage()
		return errors.New("incorrect input")
	}

	distance, err := CalculateDistance(*location)
	if err != nil {
		return err
	}

	fmt.Fprintf(c.output, "%d\n", distance)
	return nil
}
