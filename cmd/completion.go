/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"os"

	"github.com/spf13/cobra"
)

func newCompletionCmd() *cobra.Command {
	completionCmd := &cobra.Command{
		Use:   "completion",
		Short: "Generators for shell completions",
		Long: `Output shell completion code for the specified shell (bash or zsh). The shell code must be evaluated to provide interactive completion of kubebuilder commands.  This can be done by sourcing it from the .bash_profile.
 			  Detailed instructions on how to do this are available here:https://github.com/kubernetes-sigs/kubebuilder/tree/master/docs/book/src/reference/completion.md`,
	}
	completionCmd.AddCommand(newZshCmd())
	completionCmd.AddCommand(newBashCmd())
	return completionCmd
}

func newBashCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "bash",
		Short: "Generate bash completions",
		RunE: func(cmd *cobra.Command, cmdArgs []string) error {
			return cmd.Root().GenBashCompletion(os.Stdout)
		},
	}
}

func newZshCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "zsh",
		Short: "Generate zsh completions",
		RunE: func(cmd *cobra.Command, cmdArgs []string) error {
			return cmd.Root().GenZshCompletion(os.Stdout)
		},
	}
}
