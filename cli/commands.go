package cli

import (
	"github.com/urfave/cli/v2"
)

var tctlCommands = []*cli.Command{
	{
		Name:        "namespace",
		Aliases:     []string{"n"},
		Usage:       "Operate Temporal namespace",
		Subcommands: newNamespaceCommands(),
	},
	{
		Name:        "workflow",
		Aliases:     []string{"w"},
		Usage:       "Operate Temporal workflow",
		Subcommands: newWorkflowCommands(),
	},
	{
		Name:        "activity",
		Aliases:     []string{"act"},
		Usage:       "operate activities of workflow",
		Subcommands: newActivityCommands(),
	},
	{
		Name:        "taskqueue",
		Aliases:     []string{"tq"},
		Usage:       "Operate Temporal task queue",
		Subcommands: newTaskQueueCommands(),
	},
	{
		Name:        "batch",
		Usage:       "batch operation on a list of workflows from query.",
		Subcommands: newBatchCommands(),
	},
	{
		Name:        "cluster",
		Aliases:     []string{"cl"},
		Usage:       "Operate Temporal cluster",
		Subcommands: newClusterCommands(),
	},
	{
		Name:        "dataconverter",
		Aliases:     []string{"dc"},
		Usage:       "Operate Custom Data Converter",
		Subcommands: newDataConverterCommands(),
	},
	{
		Name:        "config",
		Aliases:     []string{"c"},
		Usage:       "Configure tctl",
		Subcommands: newConfigCommands(),
	},
}
