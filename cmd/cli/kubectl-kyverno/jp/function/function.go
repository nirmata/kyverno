package function

import (
	"cmp"
	"fmt"
	"io"
	"strings"

	"github.com/kyverno/kyverno/pkg/config"
	"github.com/kyverno/kyverno/pkg/engine/jmespath"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
	"k8s.io/apimachinery/pkg/util/sets"
)

var description = []string{
	"Provides function informations",
	"For more information visit: https://kyverno.io/docs/writing-policies/jmespath/ ",
}

var examples = []string{
	"  # List functions    \n  kyverno jp function",
	"  # Get function infos\n  kyverno jp function <function name>",
}

func Command() *cobra.Command {
	return &cobra.Command{
		Use:          "function [function_name]...",
		Short:        description[0],
		Long:         strings.Join(description, "\n"),
		Example:      strings.Join(examples, "\n\n"),
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			printFunctions(cmd.OutOrStdout(), args...)
		},
	}
}

func printFunctions(out io.Writer, names ...string) {
	functions := jmespath.GetFunctions(config.NewDefaultConfiguration(false))
	slices.SortFunc(functions, func(a, b jmespath.FunctionEntry) int {
		return cmp.Compare(a.String(), b.String())
	})
	namesSet := sets.New(names...)
	for _, function := range functions {
		if len(namesSet) == 0 || namesSet.Has(function.Name) {
			note := function.Note
			function.Note = ""
			fmt.Fprintln(out, "Name:", function.Name)
			fmt.Fprintln(out, "  Signature:", function.String())
			if note != "" {
				fmt.Fprintln(out, "  Note:     ", note)
			}
			fmt.Fprintln(out)
		}
	}
}
