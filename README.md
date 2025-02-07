# GOB

An easy way to bootstrap your application with batteries included.

## Status

***Consider this project to be in alpha. The API is not stable and may change at any time.***

## What is included?
- Pairs well with [cobra](github.com/spf13/cobra) and [viper](github.com/spf13/viper) via [gfi](github.com/khulnasoft/gfi), covering CLI arg parsing and config file + env var loading.
- Provides an event bus via [pulsebus](github.com/khulnasoft/go-pulsebus), enabling visibility deep in your execution stack as to what is happening.
- Provides a logger via the [logger interface](github.com/khulnasoft-lab/go-logger), allowing you to swap out for any concrete logger you'd like.
- Supplies a redactor object that can be used to remove sensitive output before it's exposed (in the log or elsewhere).
- Defines a generic UI interface that adapts well to TUI frameworks such as [bubbletea](github.com/charmbracelet/bubbletea).

## Example

Here's a basic example of how to use gob + cobra to get a fully functional CLI application going:

```go
package main

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/khulnasoft/go-pulsebus"
	"github.com/khulnasoft/gob"
	"github.com/khulnasoft/gfi"
)

// Define your per-command or entire application config as a struct
type MyCommandConfig struct {
	TimestampServer string `mapstructure:"timestamp-server"`
	// ...
}

// ... add cobra flags just as you are used to doing in any other cobra application
func (c *MyCommandConfig) AddFlags(flags gfi.FlagSet) {
	flags.StringVarP(
		&c.TimestampServer, "timestamp-server", "",
		"URL to a timestamp server to use for timestamping the signature",
	)
	// ...
}

func MyCommand(app gob.Application) *cobra.Command {
	cfg := &MyCommandConfig{
		TimestampServer: "https://somewhere.out/there", // a default value
	}

	return app.SetupCommand(&cobra.Command{
		Use:     "my-command",
		PreRunE: app.Setup(cfg),
		RunE: func(cmd *cobra.Command, args []string) error {
			// perform command functions here
			return nil
		},
	}, cfg)
}

func main() {
	cfg := gob.NewSetupConfig(gob.Identification{
		Name: "awesome",
		Version: "v1.0.0",
    })

	app := gob.New(*cfg)
	
	root := app.SetupRootCommand(&cobra.Command{})

	root.AddCommand(MyCommand(app))

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
```