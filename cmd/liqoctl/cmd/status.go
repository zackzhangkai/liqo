// Copyright 2019-2022 The Liqo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"

	"github.com/spf13/cobra"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	"github.com/liqotech/liqo/pkg/liqoctl/completion"
	"github.com/liqotech/liqo/pkg/liqoctl/factory"
	"github.com/liqotech/liqo/pkg/liqoctl/status"
)

const liqoctlStatusLongHelp = `Show the status of Liqo.

The command performs a set of checks to verify the status of the Liqo control
plane, its configuration, as well as the characteristics of the currently
active peerings, and reports the outcome in a human-readable format.

Examples:
  $ {{ .Executable }} status
or
  $ {{ .Executable }} status --namespace liqo-system
`

func newStatusCommand(ctx context.Context, f *factory.Factory) *cobra.Command {
	options := status.Options{Factory: f}
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Show the status of Liqo",
		Long:  WithTemplate(liqoctlStatusLongHelp),
		Args:  cobra.NoArgs,

		PersistentPreRun: func(cmd *cobra.Command, args []string) { singleClusterPersistentPreRun(cmd, f) },

		RunE: func(cmd *cobra.Command, args []string) error {
			return options.Run(ctx)
		},
	}

	f.AddLiqoNamespaceFlag(cmd.Flags())
	utilruntime.Must(cmd.RegisterFlagCompletionFunc(factory.FlagNamespace, completion.Namespaces(ctx, f, completion.NoLimit)))
	return cmd
}
