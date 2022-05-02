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

package peeroob

const (
	// AuthURLFlagName contains the name of auth-url flag.
	AuthURLFlagName = "auth-url"
	// ClusterIDFlagName contains the name of cluster ID flag.
	ClusterIDFlagName = "cluster-id"
	// ClusterTokenFlagName contains the name for the token flag.
	ClusterTokenFlagName = "auth-token"

	// SuccessfulMessage is printed when an add cluster command succeeds.
	SuccessfulMessage = `
Hooray 🎉! You have correctly peered with cluster %q.
You can now:

* Check the status of the peering to see when it is completely established 👓.
Every field of the foreigncluster (but IncomingPeering) should be in "Established":

kubectl get foreignclusters %s

* Check if the virtual node is correctly created (this should take less than ~30s) 📦:

kubectl get nodes %s

* Ready to go! Let's deploy a simple cross-cluster application using Liqo 🚜:

kubectl create ns liqo-demo # Let's create a demo namespace
liqoctl offload ns liqo-demo # Enable Liqo offloading on this namespace (Check out https://doc.liqo.io/usage for more details).
kubectl apply -n liqo-demo -f https://get.liqo.io/app.yaml # Deploy a sample application in the namespace to trigger the offloading.

* For more information about Liqo have a look to: https://doc.liqo.io
`
)
