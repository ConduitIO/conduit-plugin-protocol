// Copyright © 2022 Meroxa, Inc.
//
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

package fromproto

import (
	"github.com/conduitio/conduit-commons/opencdc"
	"github.com/conduitio/conduit-connector-protocol/cplugin"
	connectorv1 "github.com/conduitio/conduit-connector-protocol/proto/connector/v1"
)

// -- Request Conversions -----------------------------------------------------

func SourceConfigureRequest(in *connectorv1.Source_Configure_Request) cplugin.SourceConfigureRequest {
	return cplugin.SourceConfigureRequest{
		Config: in.Config,
	}
}

func SourceStartRequest(in *connectorv1.Source_Start_Request) cplugin.SourceOpenRequest {
	return cplugin.SourceOpenRequest{
		Position: in.Position,
	}
}

func SourceRunRequest(in *connectorv1.Source_Run_Request) cplugin.SourceRunRequest {
	return cplugin.SourceRunRequest{
		AckPositions: []opencdc.Position{in.AckPosition},
	}
}

func SourceStopRequest(_ *connectorv1.Source_Stop_Request) cplugin.SourceStopRequest {
	return cplugin.SourceStopRequest{}
}

func SourceTeardownRequest(_ *connectorv1.Source_Teardown_Request) cplugin.SourceTeardownRequest {
	return cplugin.SourceTeardownRequest{}
}

func SourceLifecycleOnCreatedRequest(in *connectorv1.Source_Lifecycle_OnCreated_Request) cplugin.SourceLifecycleOnCreatedRequest {
	return cplugin.SourceLifecycleOnCreatedRequest{
		Config: in.Config,
	}
}
func SourceLifecycleOnUpdatedRequest(in *connectorv1.Source_Lifecycle_OnUpdated_Request) cplugin.SourceLifecycleOnUpdatedRequest {
	return cplugin.SourceLifecycleOnUpdatedRequest{
		ConfigBefore: in.ConfigBefore,
		ConfigAfter:  in.ConfigAfter,
	}
}
func SourceLifecycleOnDeletedRequest(in *connectorv1.Source_Lifecycle_OnDeleted_Request) cplugin.SourceLifecycleOnDeletedRequest {
	return cplugin.SourceLifecycleOnDeletedRequest{
		Config: in.Config,
	}
}

// -- Response Conversions ----------------------------------------------------

func SourceConfigureResponse(_ *connectorv1.Source_Configure_Response) cplugin.SourceConfigureResponse {
	return cplugin.SourceConfigureResponse{}
}

func SourceStartResponse(_ *connectorv1.Source_Start_Response) cplugin.SourceOpenResponse {
	return cplugin.SourceOpenResponse{}
}

func SourceRunResponse(in *connectorv1.Source_Run_Response) (cplugin.SourceRunResponse, error) {
	var rec opencdc.Record
	err := rec.FromProto(in.Record)
	if err != nil {
		return cplugin.SourceRunResponse{}, err
	}
	return cplugin.SourceRunResponse{
		Records: []opencdc.Record{rec},
	}, nil
}

func SourceStopResponse(in *connectorv1.Source_Stop_Response) cplugin.SourceStopResponse {
	return cplugin.SourceStopResponse{
		LastPosition: in.LastPosition,
	}
}

func SourceTeardownResponse(_ *connectorv1.Source_Teardown_Response) cplugin.SourceTeardownResponse {
	return cplugin.SourceTeardownResponse{}
}

func SourceLifecycleOnCreatedResponse(_ *connectorv1.Source_Lifecycle_OnCreated_Response) cplugin.SourceLifecycleOnCreatedResponse {
	return cplugin.SourceLifecycleOnCreatedResponse{}
}
func SourceLifecycleOnUpdatedResponse(_ *connectorv1.Source_Lifecycle_OnUpdated_Response) cplugin.SourceLifecycleOnUpdatedResponse {
	return cplugin.SourceLifecycleOnUpdatedResponse{}
}
func SourceLifecycleOnDeletedResponse(_ *connectorv1.Source_Lifecycle_OnDeleted_Response) cplugin.SourceLifecycleOnDeletedResponse {
	return cplugin.SourceLifecycleOnDeletedResponse{}
}
