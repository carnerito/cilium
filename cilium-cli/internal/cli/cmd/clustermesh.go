// Copyright 2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"os"

	"github.com/cilium/cilium-cli/clustermesh"

	"github.com/spf13/cobra"
)

func newCmdClusterMesh() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clustermesh",
		Short: "Multi Cluster Management",
		Long:  ``,
	}

	cmd.AddCommand(newCmdClusterMeshEnable())
	cmd.AddCommand(newCmdClusterMeshDisable())
	cmd.AddCommand(newCmdClusterMeshGetAccessToken())
	cmd.AddCommand(newCmdClusterMeshConnect())

	return cmd
}

func newCmdClusterMeshEnable() *cobra.Command {
	var params = clustermesh.Parameters{
		Writer: os.Stdout,
	}

	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Enable ClusterMesh ability in a cluster",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			cm := clustermesh.NewK8sClusterMesh(k8sClient, params)
			return cm.Enable(context.Background())
		},
	}

	cmd.Flags().StringVar(&params.Namespace, "namespace", "kube-system", "Namespace Cilium is running in")
	cmd.Flags().StringVar(&params.ServiceType, "service-type", "ClusterIP", "Type of Kubernetes to expose control plane")

	return cmd
}

func newCmdClusterMeshDisable() *cobra.Command {
	var params = clustermesh.Parameters{
		Writer: os.Stdout,
	}

	cmd := &cobra.Command{
		Use:   "disable",
		Short: "Disable ClusterMesh ability in a cluster",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			cm := clustermesh.NewK8sClusterMesh(k8sClient, params)
			return cm.Disable(context.Background())
		},
	}

	cmd.Flags().StringVar(&params.Namespace, "namespace", "kube-system", "Namespace Cilium is running in")

	return cmd
}

func newCmdClusterMeshGetAccessToken() *cobra.Command {
	var params = clustermesh.Parameters{
		Writer: os.Stdout,
	}

	cmd := &cobra.Command{
		Use:   "get-access-token",
		Short: "Extract the access token to allow another cluster to connect",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			cm := clustermesh.NewK8sClusterMesh(k8sClient, params)
			return cm.GetAccessToken(context.Background())
		},
	}

	cmd.Flags().StringVar(&params.Namespace, "namespace", "kube-system", "Namespace Cilium is running in")

	return cmd
}

func newCmdClusterMeshConnect() *cobra.Command {
	var params = clustermesh.Parameters{
		Writer: os.Stdout,
	}

	cmd := &cobra.Command{
		Use:   "connect",
		Short: "Connect to a remote cluster",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.Flags().StringVar(&params.Namespace, "namespace", "kube-system", "Namespace Cilium is running in")

	return cmd
}