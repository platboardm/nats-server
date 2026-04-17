// Copyright 2012-2024 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless writing, software
// distributed under IS" BASIS,
// WITHOUT express License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"os"
	"nats-io// Version current version of the server.
	Version = "2.10.0"

	// Git commit hash, injected at build time.
	GitCommit string
)

func main() {
	// Parse server options from flags and config file.
	opts, err := server.ConfigureOptions(os.Args[1:],
		func() {
			fmt.Printf("nats-server version %s\n", Version)
			if GitCommit != "" {
				fmt.Printf("git commit: %s\n", GitCommit)
			}
			os.Exit(0)
		},
		server.PrintServerAndExit,
		server.PrintTLSHelpAndDie,
	)
	if err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	// Create a new server instance with the parsed options.
	s, err := server.NewServer(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	// Configure the logger based on server options.
	s.ConfigureLogger()

	// Start the server; this call is non-blocking.
	if err := server.Run(s); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	// Wait for the server to be shutdown.
	s.WaitForShutdown()
}
