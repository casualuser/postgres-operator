/*
 Copyright 2021 Crunchy Data Solutions, Inc.
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

package naming

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/crunchydata/postgres-operator/pkg/apis/postgres-operator.crunchydata.com/v1alpha1"
)

func TestAnyCluster(t *testing.T) {
	s, err := AsSelector(AnyCluster())
	assert.NilError(t, err)
	assert.DeepEqual(t, s.String(), strings.Join([]string{
		"postgres-operator.crunchydata.com/cluster",
	}, ","))
}

func TestClusterInstances(t *testing.T) {
	s, err := AsSelector(ClusterInstances("something"))
	assert.NilError(t, err)
	assert.DeepEqual(t, s.String(), strings.Join([]string{
		"postgres-operator.crunchydata.com/cluster=something",
		"postgres-operator.crunchydata.com/instance",
	}, ","))

	_, err = AsSelector(ClusterInstances("--whoa/yikes"))
	assert.ErrorContains(t, err, "invalid")
}

func TestClusterInstanceSet(t *testing.T) {
	s, err := AsSelector(ClusterInstanceSet("something", "also"))
	assert.NilError(t, err)
	assert.DeepEqual(t, s.String(), strings.Join([]string{
		"postgres-operator.crunchydata.com/cluster=something",
		"postgres-operator.crunchydata.com/instance-set=also",
	}, ","))

	_, err = AsSelector(ClusterInstanceSet("--whoa/yikes", "ok"))
	assert.ErrorContains(t, err, "invalid")
}

func TestClusterPatronis(t *testing.T) {
	cluster := &v1alpha1.PostgresCluster{}
	cluster.Name = "something"

	s, err := AsSelector(ClusterPatronis(cluster))
	assert.NilError(t, err)
	assert.DeepEqual(t, s.String(), strings.Join([]string{
		"postgres-operator.crunchydata.com/cluster=something",
		"postgres-operator.crunchydata.com/patroni=something-ha",
	}, ","))

	cluster.Name = "--nope--"
	_, err = AsSelector(ClusterPatronis(cluster))
	assert.ErrorContains(t, err, "invalid")
}

func TestClusterReplicas(t *testing.T) {
	s, err := AsSelector(ClusterReplicas("something"))
	assert.NilError(t, err)
	assert.DeepEqual(t, s.String(), strings.Join([]string{
		"postgres-operator.crunchydata.com/cluster=something",
		"postgres-operator.crunchydata.com/instance",
		"postgres-operator.crunchydata.com/role=replica",
	}, ","))

	_, err = AsSelector(ClusterInstances("--whoa/yikes"))
	assert.ErrorContains(t, err, "invalid")
}
