// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
)

// GetRunsOrErrorResponse is returned by GetRunsOrError on success.
type GetRunsOrErrorResponse struct {
	// Retrieve runs after applying a filter, cursor, and limit.
	RunsOrError GetRunsOrErrorRunsOrError `json:"-"`
}

// GetRunsOrError returns GetRunsOrErrorResponse.RunsOrError, and is useful for accessing the field via an interface.
func (v *GetRunsOrErrorResponse) GetRunsOrError() GetRunsOrErrorRunsOrError { return v.RunsOrError }

func (v *GetRunsOrErrorResponse) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*GetRunsOrErrorResponse
		RunsOrError json.RawMessage `json:"runsOrError"`
		graphql.NoUnmarshalJSON
	}
	firstPass.GetRunsOrErrorResponse = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.RunsOrError
		src := firstPass.RunsOrError
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalGetRunsOrErrorRunsOrError(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"unable to unmarshal GetRunsOrErrorResponse.RunsOrError: %w", err)
			}
		}
	}
	return nil
}

type __premarshalGetRunsOrErrorResponse struct {
	RunsOrError json.RawMessage `json:"runsOrError"`
}

func (v *GetRunsOrErrorResponse) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *GetRunsOrErrorResponse) __premarshalJSON() (*__premarshalGetRunsOrErrorResponse, error) {
	var retval __premarshalGetRunsOrErrorResponse

	{

		dst := &retval.RunsOrError
		src := v.RunsOrError
		var err error
		*dst, err = __marshalGetRunsOrErrorRunsOrError(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"unable to marshal GetRunsOrErrorResponse.RunsOrError: %w", err)
		}
	}
	return &retval, nil
}

// GetRunsOrErrorRunsOrError includes the requested fields of the GraphQL interface RunsOrError.
//
// GetRunsOrErrorRunsOrError is implemented by the following types:
// GetRunsOrErrorRunsOrErrorInvalidPipelineRunsFilterError
// GetRunsOrErrorRunsOrErrorPythonError
// GetRunsOrErrorRunsOrErrorRuns
type GetRunsOrErrorRunsOrError interface {
	implementsGraphQLInterfaceGetRunsOrErrorRunsOrError()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
}

func (v *GetRunsOrErrorRunsOrErrorInvalidPipelineRunsFilterError) implementsGraphQLInterfaceGetRunsOrErrorRunsOrError() {
}
func (v *GetRunsOrErrorRunsOrErrorPythonError) implementsGraphQLInterfaceGetRunsOrErrorRunsOrError() {
}
func (v *GetRunsOrErrorRunsOrErrorRuns) implementsGraphQLInterfaceGetRunsOrErrorRunsOrError() {}

func __unmarshalGetRunsOrErrorRunsOrError(b []byte, v *GetRunsOrErrorRunsOrError) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "InvalidPipelineRunsFilterError":
		*v = new(GetRunsOrErrorRunsOrErrorInvalidPipelineRunsFilterError)
		return json.Unmarshal(b, *v)
	case "PythonError":
		*v = new(GetRunsOrErrorRunsOrErrorPythonError)
		return json.Unmarshal(b, *v)
	case "Runs":
		*v = new(GetRunsOrErrorRunsOrErrorRuns)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing RunsOrError.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for GetRunsOrErrorRunsOrError: "%v"`, tn.TypeName)
	}
}

func __marshalGetRunsOrErrorRunsOrError(v *GetRunsOrErrorRunsOrError) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *GetRunsOrErrorRunsOrErrorInvalidPipelineRunsFilterError:
		typename = "InvalidPipelineRunsFilterError"

		result := struct {
			TypeName string `json:"__typename"`
			*GetRunsOrErrorRunsOrErrorInvalidPipelineRunsFilterError
		}{typename, v}
		return json.Marshal(result)
	case *GetRunsOrErrorRunsOrErrorPythonError:
		typename = "PythonError"

		result := struct {
			TypeName string `json:"__typename"`
			*GetRunsOrErrorRunsOrErrorPythonError
		}{typename, v}
		return json.Marshal(result)
	case *GetRunsOrErrorRunsOrErrorRuns:
		typename = "Runs"

		result := struct {
			TypeName string `json:"__typename"`
			*GetRunsOrErrorRunsOrErrorRuns
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for GetRunsOrErrorRunsOrError: "%T"`, v)
	}
}

// GetRunsOrErrorRunsOrErrorInvalidPipelineRunsFilterError includes the requested fields of the GraphQL type InvalidPipelineRunsFilterError.
type GetRunsOrErrorRunsOrErrorInvalidPipelineRunsFilterError struct {
	Typename string `json:"__typename"`
}

// GetTypename returns GetRunsOrErrorRunsOrErrorInvalidPipelineRunsFilterError.Typename, and is useful for accessing the field via an interface.
func (v *GetRunsOrErrorRunsOrErrorInvalidPipelineRunsFilterError) GetTypename() string {
	return v.Typename
}

// GetRunsOrErrorRunsOrErrorPythonError includes the requested fields of the GraphQL type PythonError.
type GetRunsOrErrorRunsOrErrorPythonError struct {
	Typename string `json:"__typename"`
}

// GetTypename returns GetRunsOrErrorRunsOrErrorPythonError.Typename, and is useful for accessing the field via an interface.
func (v *GetRunsOrErrorRunsOrErrorPythonError) GetTypename() string { return v.Typename }

// GetRunsOrErrorRunsOrErrorRuns includes the requested fields of the GraphQL type Runs.
type GetRunsOrErrorRunsOrErrorRuns struct {
	Typename string `json:"__typename"`
	Count    int    `json:"count"`
}

// GetTypename returns GetRunsOrErrorRunsOrErrorRuns.Typename, and is useful for accessing the field via an interface.
func (v *GetRunsOrErrorRunsOrErrorRuns) GetTypename() string { return v.Typename }

// GetCount returns GetRunsOrErrorRunsOrErrorRuns.Count, and is useful for accessing the field via an interface.
func (v *GetRunsOrErrorRunsOrErrorRuns) GetCount() int { return v.Count }

// The status of run execution.
type RunStatus string

const (
	// Runs waiting to be launched by the Dagster Daemon.
	RunStatusQueued RunStatus = "QUEUED"
	// Runs that have been created, but not yet submitted for launch.
	RunStatusNotStarted RunStatus = "NOT_STARTED"
	// Runs that are managed outside of the Dagster control plane.
	RunStatusManaged RunStatus = "MANAGED"
	// Runs that have been launched, but execution has not yet started.
	RunStatusStarting RunStatus = "STARTING"
	// Runs that have been launched and execution has started.
	RunStatusStarted RunStatus = "STARTED"
	// Runs that have successfully completed.
	RunStatusSuccess RunStatus = "SUCCESS"
	// Runs that have failed to complete.
	RunStatusFailure RunStatus = "FAILURE"
	// Runs that are in-progress and pending to be canceled.
	RunStatusCanceling RunStatus = "CANCELING"
	// Runs that have been canceled before completion.
	RunStatusCanceled RunStatus = "CANCELED"
)

// __GetRunsOrErrorInput is used internally by genqlient
type __GetRunsOrErrorInput struct {
	Status RunStatus `json:"status"`
}

// GetStatus returns __GetRunsOrErrorInput.Status, and is useful for accessing the field via an interface.
func (v *__GetRunsOrErrorInput) GetStatus() RunStatus { return v.Status }

// The query or mutation executed by GetRunsOrError.
const GetRunsOrError_Operation = `
query GetRunsOrError ($status: RunStatus!) {
	runsOrError(filter: {statuses:[$status]}) {
		__typename
		... on Runs {
			count
		}
	}
}
`

func GetRunsOrError(
	ctx_ context.Context,
	client_ graphql.Client,
	status RunStatus,
) (*GetRunsOrErrorResponse, error) {
	req_ := &graphql.Request{
		OpName: "GetRunsOrError",
		Query:  GetRunsOrError_Operation,
		Variables: &__GetRunsOrErrorInput{
			Status: status,
		},
	}
	var err_ error

	var data_ GetRunsOrErrorResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}
