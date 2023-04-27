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

package types

import (
	"fmt"

	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// GeoShapeQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/899364a63e7415b60033ddd49d50a30369da26d7/specification/_types/query_dsl/geo.ts#L86-L91
type GeoShapeQuery struct {
	Boost          *float32                      `json:"boost,omitempty"`
	GeoShapeQuery  map[string]GeoShapeFieldQuery `json:"GeoShapeQuery,omitempty"`
	IgnoreUnmapped *bool                         `json:"ignore_unmapped,omitempty"`
	QueryName_     *string                       `json:"_name,omitempty"`
}

func (s *GeoShapeQuery) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "boost":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.Boost = &f
			case float64:
				f := float32(v)
				s.Boost = &f
			}

		case "GeoShapeQuery":
			if s.GeoShapeQuery == nil {
				s.GeoShapeQuery = make(map[string]GeoShapeFieldQuery, 0)
			}
			if err := dec.Decode(&s.GeoShapeQuery); err != nil {
				return err
			}

		case "ignore_unmapped":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IgnoreUnmapped = &value
			case bool:
				s.IgnoreUnmapped = &v
			}

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.QueryName_ = &o

		default:

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s GeoShapeQuery) MarshalJSON() ([]byte, error) {
	type opt GeoShapeQuery
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.GeoShapeQuery {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "GeoShapeQuery")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewGeoShapeQuery returns a GeoShapeQuery.
func NewGeoShapeQuery() *GeoShapeQuery {
	r := &GeoShapeQuery{
		GeoShapeQuery: make(map[string]GeoShapeFieldQuery, 0),
	}

	return r
}
