// Copyright (c) 2016 Tigera, Inc. All rights reserved.

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

package model

import (
	"fmt"

	"regexp"

	"reflect"

	log "github.com/Sirupsen/logrus"
	"github.com/projectcalico/libcalico-go/lib/errors"
)

var (
	matchWorkloadEndpointStatus = regexp.MustCompile("^/?calico/felix/v1/host/([^/]+)/workload/([^/]+)/([^/]+)/endpoint/([^/]+)$")
)

type WorkloadEndpointStatusKey struct {
	Hostname       string `json:"-"`
	OrchestratorID string `json:"-"`
	WorkloadID     string `json:"-"`
	EndpointID     string `json:"-"`
}

func (key WorkloadEndpointStatusKey) defaultPath() (string, error) {
	if key.Hostname == "" {
		return "", errors.ErrorInsufficientIdentifiers{Name: "hostname"}
	}
	if key.OrchestratorID == "" {
		return "", errors.ErrorInsufficientIdentifiers{Name: "orchestrator"}
	}
	if key.WorkloadID == "" {
		return "", errors.ErrorInsufficientIdentifiers{Name: "workload"}
	}
	if key.EndpointID == "" {
		return "", errors.ErrorInsufficientIdentifiers{Name: "endpoint"}
	}
	return fmt.Sprintf("/calico/felix/v1/host/%s/workload/%s/%s/endpoint/%s",
		key.Hostname, escapeName(key.OrchestratorID), escapeName(key.WorkloadID), escapeName(key.EndpointID)), nil
}

func (key WorkloadEndpointStatusKey) defaultDeletePath() (string, error) {
	return key.defaultPath()
}

func (key WorkloadEndpointStatusKey) defaultDeleteParentPaths() ([]string, error) {
	if key.Hostname == "" {
		return nil, errors.ErrorInsufficientIdentifiers{Name: "hostname"}
	}
	if key.OrchestratorID == "" {
		return nil, errors.ErrorInsufficientIdentifiers{Name: "orchestrator"}
	}
	if key.WorkloadID == "" {
		return nil, errors.ErrorInsufficientIdentifiers{Name: "workload"}
	}
	workload := fmt.Sprintf("/calico/felix/v1/host/%s/workload/%s/%s",
		key.Hostname, escapeName(key.OrchestratorID), escapeName(key.WorkloadID))
	endpoints := workload + "/endpoint"
	return []string{endpoints, workload}, nil
}

func (key WorkloadEndpointStatusKey) valueType() reflect.Type {
	return reflect.TypeOf(WorkloadEndpointStatus{})
}

func (key WorkloadEndpointStatusKey) String() string {
	return fmt.Sprintf("WorkloadEndpointStatus(hostname=%s, orchestrator=%s, workload=%s, name=%s)",
		key.Hostname, key.OrchestratorID, key.WorkloadID, key.EndpointID)
}

type WorkloadEndpointStatusListOptions struct {
	Hostname       string
	OrchestratorID string
	WorkloadID     string
	EndpointID     string
}

func (options WorkloadEndpointStatusListOptions) defaultPathRoot() string {
	k := "/calico/felix/v1/host"
	if options.Hostname == "" {
		return k
	}
	k = k + fmt.Sprintf("/%s/workload", options.Hostname)
	if options.OrchestratorID == "" {
		return k
	}
	k = k + fmt.Sprintf("/%s", escapeName(options.OrchestratorID))
	if options.WorkloadID == "" {
		return k
	}
	k = k + fmt.Sprintf("/%s/endpoint", escapeName(options.WorkloadID))
	if options.EndpointID == "" {
		return k
	}
	k = k + fmt.Sprintf("/%s", escapeName(options.EndpointID))
	return k
}

func (options WorkloadEndpointStatusListOptions) KeyFromDefaultPath(ekey string) Key {
	log.Debugf("Get WorkloadEndpoint key from %s", ekey)
	r := matchWorkloadEndpointStatus.FindAllStringSubmatch(ekey, -1)
	if len(r) != 1 {
		log.Debugf("Didn't match regex")
		return nil
	}
	hostname := r[0][1]
	orchID := unescapeName(r[0][2])
	workloadID := unescapeName(r[0][3])
	endpointID := unescapeName(r[0][4])
	if options.Hostname != "" && hostname != options.Hostname {
		log.Debugf("Didn't match hostname %s != %s", options.Hostname, hostname)
		return nil
	}
	if options.OrchestratorID != "" && orchID != options.OrchestratorID {
		log.Debugf("Didn't match orchestrator %s != %s", options.OrchestratorID, orchID)
		return nil
	}
	if options.WorkloadID != "" && workloadID != options.WorkloadID {
		log.Debugf("Didn't match workload %s != %s", options.WorkloadID, workloadID)
		return nil
	}
	if options.EndpointID != "" && endpointID != options.EndpointID {
		log.Debugf("Didn't match endpoint ID %s != %s", options.EndpointID, endpointID)
		return nil
	}
	return WorkloadEndpointStatusKey{
		Hostname:       hostname,
		OrchestratorID: orchID,
		WorkloadID:     workloadID,
		EndpointID:     endpointID,
	}
}

type WorkloadEndpointStatus struct {
	Status string `json:"status"`
}
