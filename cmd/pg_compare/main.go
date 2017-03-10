package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/jelmersnoeck/pg_compare/compare"
	"github.com/jelmersnoeck/pg_compare/config"
	"github.com/jelmersnoeck/pg_compare/drivers/pg"
	"github.com/spf13/cobra"
)

func main() {
	cfg := &config.Config{}

	cmd := &cobra.Command{
		Use:   "pg_compare [source] [destination]",
		Short: "PG Compare is a tool to compare two PGSQL databases with each other.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("Expected a `source` and a `destination` as arguments.")
			}

			cfg.Source = args[0]
			cfg.Destination = args[1]

			source, err := pg.NewDriver(cfg.Source)
			if err != nil {
				return err
			}
			destination, err := pg.NewDriver(cfg.Destination)
			if err != nil {
				return err
			}

			if cfg.Views {
				fmt.Println("==> Comparing views")
				if err := compare.Views(source, destination); err != nil {
					return err
				}
				fmt.Println("==> Comparing views: OK")
			}

			if cfg.Columns {
				fmt.Println("==> Comparing columns")
				if err := compare.Columns(source, destination); err != nil {
					return err
				}
				fmt.Println("==> Comparing columns: OK")
			}

			if cfg.ForeignKeys {
				fmt.Println("==> Comparing foreign keys")
				if err := compare.ForeignKeys(source, destination); err != nil {
					return err
				}
				fmt.Println("==> Comparing foreign keys: OK")
			}

			if cfg.Rows {
				fmt.Println("==> Comparing row counts")
				if err := compare.Count(source, destination); err != nil {
					return err
				}
				fmt.Println("==> Comparing row counts: OK")
			}

			if cfg.Indexes {
				fmt.Println("==> Comparing indexes")
				if err := compare.Indexes(source, destination); err != nil {
					return err
				}
				fmt.Println("==> Comparing indexes: OK")
			}

			fmt.Println("")
			fmt.Println("Databases match")
			return nil
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&cfg.Rows, "rows", "r", false, "Compare rows. This does a simple count of the number of rows")
	flags.BoolVarP(&cfg.Columns, "columns", "c", false, "Compare table columns")
	flags.BoolVarP(&cfg.Views, "views", "v", false, "Compare views")
	flags.BoolVarP(&cfg.ForeignKeys, "foreign-keys", "k", false, "Compare foreign keys")
	flags.BoolVarP(&cfg.Indexes, "indexes", "i", false, "Compare indexes")

	flags.BoolVar(&config.Verbose, "verbose", false, "Verbose")

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
