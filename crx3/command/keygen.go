package command

import (
	"errors"
	"strings"

	crx3 "crx"

	"github.com/spf13/cobra"
)

func newKeygenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keygen [file]",
		Short: "Create a new private key",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("infile is required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			filename := args[0]
			if !strings.HasSuffix(filename, ".pem") {
				filename = filename + ".pem"
			}
			pk, err := crx3.NewPrivateKey()
			if err != nil {
				return err
			}
			return crx3.SavePrivateKey(filename, pk)
		},
	}
	return cmd
}
