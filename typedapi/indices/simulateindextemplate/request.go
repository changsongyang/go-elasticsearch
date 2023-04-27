// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/899364a63e7415b60033ddd49d50a30369da26d7

package simulateindextemplate

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package simulateindextemplate
//
// https://github.com/elastic/elasticsearch-specification/blob/899364a63e7415b60033ddd49d50a30369da26d7/specification/indices/simulate_index_template/IndicesSimulateIndexTemplateRequest.ts#L33-L71
type Request struct {
	AllowAutoCreate *bool                       `json:"allow_auto_create,omitempty"`
	ComposedOf      []string                    `json:"composed_of,omitempty"`
	DataStream      *types.DataStreamVisibility `json:"data_stream,omitempty"`
	IndexPatterns   []string                    `json:"index_patterns,omitempty"`
	Meta_           types.Metadata              `json:"_meta,omitempty"`
	Priority        *int                        `json:"priority,omitempty"`
	Template        *types.IndexTemplateMapping `json:"template,omitempty"`
	Version         *int64                      `json:"version,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Simulateindextemplate request: %w", err)
	}

	return &req, nil
}
