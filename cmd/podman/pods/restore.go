package pods

import (
	"github.com/containers/podman/v4/cmd/podman/common"
	"github.com/containers/podman/v4/cmd/podman/registry"
	"github.com/spf13/cobra"
)

var (
	restoreDescription = `The pod ID or name can be used.

  Restores each of the specified pods from a checkpoint. The pod name or ID can be used.`

	podRestoreCommand = &cobra.Command{
		Use:               "restore [options] POD",
		Short:             "Restore one or more pods from checkpoint",
		Long:              restoreDescription,
		RunE:              restore,
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: common.AutocompletePods,
		Example:           `podman pod restore podID`,
	}
)

func init() {
	registry.Commands = append(registry.Commands, registry.CliCommand{
		Command: podRestoreCommand,
		Parent:  podCmd,
	})
}

func restore(cmd *cobra.Command, args []string) error {
	return nil
}
