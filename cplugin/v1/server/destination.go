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

package server

import (
	"context"

	"github.com/conduitio/conduit-connector-protocol/cplugin"
	"github.com/conduitio/conduit-connector-protocol/cplugin/v1/fromproto"
	"github.com/conduitio/conduit-connector-protocol/cplugin/v1/toproto"
	connectorv1 "github.com/conduitio/conduit-connector-protocol/proto/connector/v1"
)

func NewDestinationPluginServer(impl cplugin.DestinationPlugin) connectorv1.DestinationPluginServer {
	return &destinationPluginServer{impl: impl}
}

type destinationPluginServer struct {
	connectorv1.UnimplementedDestinationPluginServer
	impl cplugin.DestinationPlugin
}

func (s *destinationPluginServer) Configure(ctx context.Context, protoReq *connectorv1.Destination_Configure_Request) (*connectorv1.Destination_Configure_Response, error) {
	goReq := fromproto.DestinationConfigureRequest(protoReq)
	goResp, err := s.impl.Configure(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.DestinationConfigureResponse(goResp), nil
}
func (s *destinationPluginServer) Start(ctx context.Context, protoReq *connectorv1.Destination_Start_Request) (*connectorv1.Destination_Start_Response, error) {
	goReq := fromproto.DestinationStartRequest(protoReq)
	goResp, err := s.impl.Open(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.DestinationStartResponse(goResp), nil
}
func (s *destinationPluginServer) Run(stream connectorv1.DestinationPlugin_RunServer) error {
	return s.impl.Run(stream.Context(), &DestinationRunStream{stream: stream})
}
func (s *destinationPluginServer) Stop(ctx context.Context, protoReq *connectorv1.Destination_Stop_Request) (*connectorv1.Destination_Stop_Response, error) {
	goReq := fromproto.DestinationStopRequest(protoReq)
	goResp, err := s.impl.Stop(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.DestinationStopResponse(goResp), nil
}
func (s *destinationPluginServer) Teardown(ctx context.Context, protoReq *connectorv1.Destination_Teardown_Request) (*connectorv1.Destination_Teardown_Response, error) {
	goReq := fromproto.DestinationTeardownRequest(protoReq)
	goResp, err := s.impl.Teardown(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.DestinationTeardownResponse(goResp), nil
}
func (s *destinationPluginServer) LifecycleOnCreated(ctx context.Context, protoReq *connectorv1.Destination_Lifecycle_OnCreated_Request) (*connectorv1.Destination_Lifecycle_OnCreated_Response, error) {
	goReq := fromproto.DestinationLifecycleOnCreatedRequest(protoReq)
	goResp, err := s.impl.LifecycleOnCreated(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.DestinationLifecycleOnCreatedResponse(goResp), nil
}
func (s *destinationPluginServer) LifecycleOnUpdated(ctx context.Context, protoReq *connectorv1.Destination_Lifecycle_OnUpdated_Request) (*connectorv1.Destination_Lifecycle_OnUpdated_Response, error) {
	goReq := fromproto.DestinationLifecycleOnUpdatedRequest(protoReq)
	goResp, err := s.impl.LifecycleOnUpdated(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.DestinationLifecycleOnUpdatedResponse(goResp), nil
}
func (s *destinationPluginServer) LifecycleOnDeleted(ctx context.Context, protoReq *connectorv1.Destination_Lifecycle_OnDeleted_Request) (*connectorv1.Destination_Lifecycle_OnDeleted_Response, error) {
	goReq := fromproto.DestinationLifecycleOnDeletedRequest(protoReq)
	goResp, err := s.impl.LifecycleOnDeleted(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.DestinationLifecycleOnDeletedResponse(goResp), nil
}

// DestinationRunStream is the server-side implementation of the
// cplugin.DestinationRunStream interface.
type DestinationRunStream struct {
	stream connectorv1.DestinationPlugin_RunServer
}

func (s *DestinationRunStream) Client() cplugin.DestinationRunStreamClient {
	panic("invalid use of server.DestinationRunStream - it is a server-side type only")
}
func (s *DestinationRunStream) Server() cplugin.DestinationRunStreamServer {
	if s.stream == nil {
		panic("invalid use of server.DestinationRunStream - stream has not been initialized using DestinationPluginServer.Run")
	}
	return s
}

func (s *DestinationRunStream) Send(in cplugin.DestinationRunResponse) error {
	out := toproto.DestinationRunResponse(in)
	for _, out := range out {
		err := s.stream.Send(out)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DestinationRunStream) Recv() (cplugin.DestinationRunRequest, error) {
	in, err := s.stream.Recv()
	if err != nil {
		return cplugin.DestinationRunRequest{}, err
	}
	out, err := fromproto.DestinationRunRequest(in)
	if err != nil {
		return cplugin.DestinationRunRequest{}, err
	}
	return out, nil
}
