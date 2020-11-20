// Copyright (C) 2020 Red Hat, Inc.
//
// This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public
// License as published by the Free Software Foundation; either version 2 of the License, or (at your option) any later
// version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
// warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with this program; if not, write to the Free
// Software Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301, USA.

//
// Code generated by `test-network-function-claim/cmd/generate/generate.go` on: 2020-11-12 10:16:09.344265 -0500 EST m=+0.001642744
//
// `https://github.com/a-h/generate` provides a generic set of interfaces to convert JSON schema into
// workable GoLang struct implementations.  However, the code generator is limited and does not allow
// type remapping.  By default, JSON Schema "object" types are remapped to custom struct definitions.
// This becomes a problem in our case, as we do not define certain facets such as "Hosts" or
// "LshwOutput".  This CLI driven generator augments the stock generator to allow overrides to generic
// "map[string]interface{}", which is capable of representing arbitrary JSON data.
//
// Warning:  Do not edit this file by hand.  Instead, use Makefile targets.
//

// Code generated by schema-generate. DO NOT EDIT.

package claim

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// Claim
type Claim struct {

	// Key/Value where the key is the configuration name.
	Configurations map[string]interface{} `json:"configurations"`

	// Hostname/payload key/value information for all hosts in the cluster.
	Hosts    map[string]interface{} `json:"hosts"`
	Metadata *Metadata              `json:"metadata,omitempty"`

	// The test results.
	Results  map[string]interface{} `json:"results"`
	Versions *Versions              `json:"versions"`
}

// Metadata
type Metadata struct {

	// The claim evaluation UTC end time.
	EndTime string `json:"endTime"`

	// The claim evaluation UTC start time.
	StartTime string `json:"startTime"`
}

// Root Represents a test-network-function claim.
type Root struct {
	Claim *Claim `json:"claim"`
}

// Versions
type Versions struct {

	// The test-network-function (tnf) version.
	Openshift string `json:"openshift"`

	// The test-network-function (tnf) version.
	Tnf string `json:"tnf"`
}

func (strct *Claim) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Configurations" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "configurations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"configurations\": ")
	if tmp, err := json.Marshal(strct.Configurations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Hosts" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "hosts" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"hosts\": ")
	if tmp, err := json.Marshal(strct.Hosts); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "metadata" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"metadata\": ")
	if tmp, err := json.Marshal(strct.Metadata); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Results" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "results" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"results\": ")
	if tmp, err := json.Marshal(strct.Results); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Versions" field is required
	if strct.Versions == nil {
		return nil, errors.New("versions is a required field")
	}
	// Marshal the "versions" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"versions\": ")
	if tmp, err := json.Marshal(strct.Versions); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Claim) UnmarshalJSON(b []byte) error {
	configurationsReceived := false
	hostsReceived := false
	resultsReceived := false
	versionsReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "configurations":
			if err := json.Unmarshal([]byte(v), &strct.Configurations); err != nil {
				return err
			}
			configurationsReceived = true
		case "hosts":
			if err := json.Unmarshal([]byte(v), &strct.Hosts); err != nil {
				return err
			}
			hostsReceived = true
		case "metadata":
			if err := json.Unmarshal([]byte(v), &strct.Metadata); err != nil {
				return err
			}
		case "results":
			if err := json.Unmarshal([]byte(v), &strct.Results); err != nil {
				return err
			}
			resultsReceived = true
		case "versions":
			if err := json.Unmarshal([]byte(v), &strct.Versions); err != nil {
				return err
			}
			versionsReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if configurations (a required property) was received
	if !configurationsReceived {
		return errors.New("\"configurations\" is required but was not present")
	}
	// check if hosts (a required property) was received
	if !hostsReceived {
		return errors.New("\"hosts\" is required but was not present")
	}
	// check if results (a required property) was received
	if !resultsReceived {
		return errors.New("\"results\" is required but was not present")
	}
	// check if versions (a required property) was received
	if !versionsReceived {
		return errors.New("\"versions\" is required but was not present")
	}
	return nil
}

func (strct *Metadata) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "EndTime" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "endTime" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"endTime\": ")
	if tmp, err := json.Marshal(strct.EndTime); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "StartTime" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "startTime" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"startTime\": ")
	if tmp, err := json.Marshal(strct.StartTime); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Metadata) UnmarshalJSON(b []byte) error {
	endTimeReceived := false
	startTimeReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "endTime":
			if err := json.Unmarshal([]byte(v), &strct.EndTime); err != nil {
				return err
			}
			endTimeReceived = true
		case "startTime":
			if err := json.Unmarshal([]byte(v), &strct.StartTime); err != nil {
				return err
			}
			startTimeReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if endTime (a required property) was received
	if !endTimeReceived {
		return errors.New("\"endTime\" is required but was not present")
	}
	// check if startTime (a required property) was received
	if !startTimeReceived {
		return errors.New("\"startTime\" is required but was not present")
	}
	return nil
}

func (strct *Root) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Claim" field is required
	if strct.Claim == nil {
		return nil, errors.New("claim is a required field")
	}
	// Marshal the "claim" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"claim\": ")
	if tmp, err := json.Marshal(strct.Claim); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Root) UnmarshalJSON(b []byte) error {
	claimReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "claim":
			if err := json.Unmarshal([]byte(v), &strct.Claim); err != nil {
				return err
			}
			claimReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if claim (a required property) was received
	if !claimReceived {
		return errors.New("\"claim\" is required but was not present")
	}
	return nil
}

func (strct *Versions) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Openshift" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "openshift" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"openshift\": ")
	if tmp, err := json.Marshal(strct.Openshift); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Tnf" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "tnf" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"tnf\": ")
	if tmp, err := json.Marshal(strct.Tnf); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Versions) UnmarshalJSON(b []byte) error {
	openshiftReceived := false
	tnfReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "openshift":
			if err := json.Unmarshal([]byte(v), &strct.Openshift); err != nil {
				return err
			}
			openshiftReceived = true
		case "tnf":
			if err := json.Unmarshal([]byte(v), &strct.Tnf); err != nil {
				return err
			}
			tnfReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if openshift (a required property) was received
	if !openshiftReceived {
		return errors.New("\"openshift\" is required but was not present")
	}
	// check if tnf (a required property) was received
	if !tnfReceived {
		return errors.New("\"tnf\" is required but was not present")
	}
	return nil
}