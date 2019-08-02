/*
Copyright 2019 The Kubernetes Authors.

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

package status

import (
	"sort"
)

// NodeList defines a list of Node
type NodeList []*Node

// Sort the list of Node wrapper by node provisioning order and by name
func (l NodeList) Sort() {
	sort.Slice(l, func(i, j int) bool {
		return l[i].provisioningOrder() < l[j].provisioningOrder() ||
			(l[i].provisioningOrder() == l[j].provisioningOrder() && l[i].Name() < l[j].Name())
	})
}

// EligibleForActions returns the list of nodes without nodes marked as SkipAction
func (l NodeList) EligibleForActions() NodeList {
	var res NodeList
	for _, n := range l {
		if !n.skip {
			res = append(res, n)
		}
	}
	res.Sort()
	return res
}
