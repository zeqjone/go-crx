package command

import (
	"errors"
	"fmt"

	"github/zeqjone/crx"

	"github.com/spf13/cobra"
)

func newIDCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "id [infile]",
		Short: "Generate extension id",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("infile is required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			infile := args[0]
			id, err := crx.ID(infile)
			if err != nil {
				return err
			}
			fmt.Println(id)
			return nil
		},
	}
	return cmd
}
