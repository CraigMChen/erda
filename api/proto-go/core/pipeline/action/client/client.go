// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: action.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/pipeline/action/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// ActionService action.proto
	ActionService() pb.ActionServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		actionService: pb.NewActionServiceClient(cc),
	}
}

type serviceClients struct {
	actionService pb.ActionServiceClient
}

func (c *serviceClients) ActionService() pb.ActionServiceClient {
	return c.actionService
}

type actionServiceWrapper struct {
	client pb.ActionServiceClient
	opts   []grpc1.CallOption
}

func (s *actionServiceWrapper) List(ctx context.Context, req *pb.PipelineActionListRequest) (*pb.PipelineActionListResponse, error) {
	return s.client.List(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *actionServiceWrapper) Save(ctx context.Context, req *pb.PipelineActionSaveRequest) (*pb.PipelineActionSaveResponse, error) {
	return s.client.Save(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *actionServiceWrapper) Delete(ctx context.Context, req *pb.PipelineActionDeleteRequest) (*pb.PipelineActionDeleteResponse, error) {
	return s.client.Delete(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}