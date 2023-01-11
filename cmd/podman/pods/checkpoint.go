package pods

import (
	"context"
	"errors"

	"github.com/containers/podman/v4/cmd/podman/common"
	"github.com/containers/podman/v4/cmd/podman/registry"
	"github.com/containers/podman/v4/cmd/podman/utils"
	"github.com/containers/podman/v4/cmd/podman/validate"
	"github.com/containers/podman/v4/pkg/domain/entities"
	"github.com/containers/podman/v4/pkg/rootless"
	"github.com/spf13/cobra"
)

var podCheckpointOptions entities.PodCheckpointOptions

var (
	podCheckpointCommand = &cobra.Command{
		// FIXME: Add support for checkpointing multiple Pods.
		Use:   "checkpoint [options] POD [POD...]",
		Short: "Checkpoint one or more pods",
		RunE:  checkpoint,
		Args: func(cmd *cobra.Command, args []string) error {
			return validate.CheckAllLatestAndIDFile(cmd, args, false, "")
		},
		ValidArgsFunction: common.AutocompletePods,
		Example: `podman pod checkpoint podID
  podman pod checkpoint --all
  podman pod checkpoint --leave-running --latest`,
	}
)

func init() {
	registry.Commands = append(registry.Commands, registry.CliCommand{
		Command: podCheckpointCommand,
		Parent:  podCmd,
	})
	flags := podCheckpointCommand.Flags()
	flags.BoolVarP(&podCheckpointOptions.LeaveStopped, "leave-stopped", "s", false, "Leave all containers within a pod in stopped state after checkpoint")
	flags.BoolVarP(&podCheckpointOptions.All, "all", "a", false, "Checkpoint all running pods")

	validate.AddLatestFlag(podCheckpointCommand, &podCheckpointOptions.Latest)
}

func checkpoint(cmd *cobra.Command, args []string) error {
	var errs utils.OutputErrors
	args = utils.RemoveSlash(args)

	if rootless.IsRootless() {
		return errors.New("checkpointing a pod requires root")
	}

	_, err := registry.ContainerEngine().PodCheckpoint(context.Background(), args, podCheckpointOptions)
	if err != nil {
		return err
	}

	return errs.PrintErrors()
}
