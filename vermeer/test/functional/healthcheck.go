//go:build vermeer_test

/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements. See the NOTICE file distributed with this
work for additional information regarding copyright ownership. The ASF
licenses this file to You under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
License for the specific language governing permissions and limitations
under the License.
*/

package functional

import (
	"testing"
	"vermeer/client"

	"github.com/stretchr/testify/require"
)

type HealthCheck struct {
	client []*client.VermeerClient
	t      *testing.T
}

func (h *HealthCheck) Init(t *testing.T, client ...*client.VermeerClient) {
	h.client = client
	h.t = t
}
func (h *HealthCheck) DoHealthCheck() {
	for _, vermeerClient := range h.client {
		ok, err := vermeerClient.HealthCheck()
		require.NoError(h.t, err)
		require.Equal(h.t, true, ok)
	}
}