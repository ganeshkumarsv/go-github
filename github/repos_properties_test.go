// Copyright 2023 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRepositoriesService_GetAllCustomPropertyValues(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/properties/values", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[
		{
          "property_name": "environment",
          "value": "production"
        },
        {
          "property_name": "service",
          "value": "web"
        }
		]`)
	})

	ctx := context.Background()
	customPropertyValues, _, err := client.Repositories.GetAllCustomPropertyValues(ctx, "o", "r")
	if err != nil {
		t.Errorf("Repositories.GetAllCustomPropertyValues returned error: %v", err)
	}

	want := []*CustomPropertyValue{
		{
			PropertyName: "environment",
			Value:        String("production"),
		},
		{
			PropertyName: "service",
			Value:        String("web"),
		},
	}

	if !cmp.Equal(customPropertyValues, want) {
		t.Errorf("Repositories.GetAllCustomPropertyValues returned %+v, want %+v", customPropertyValues, want)
	}

	const methodName = "GetAllCustomPropertyValues"

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Repositories.GetAllCustomPropertyValues(ctx, "o", "r")
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestCreateOrUpdateRepoCustomPropertyValues(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/properties/values", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		testBody(t, r, `'{"properties":[{"property_name":"environment","value":"production"},{"property_name":"service","value":"web"},{"property_name":"team","value":"octocat"}]}'`+"\n")
	})

	ctx := context.Background()
	_, err := client.Repositories.CreateOrUpdateRepoCustomPropertyValues(ctx, "o", "repo", []*CustomPropertyValue{
		{
			PropertyName: "service",
			Value:        String("string"),
		},
	})
	if err != nil {
		t.Errorf("Repositories.CreateOrUpdateRepoCustomPropertyValues returned error: %v", err)
	}

	const methodName = "CreateOrUpdateRepoCustomPropertyValues"

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		return client.Repositories.CreateOrUpdateRepoCustomPropertyValues(ctx, "o", "", nil)
	})
}
