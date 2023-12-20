// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: task.proto

package grpc

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TaskManager_CreateTask_FullMethodName = "/TaskManager/CreateTask"
	TaskManager_GetTask_FullMethodName    = "/TaskManager/GetTask"
)

// TaskManagerClient is the client API for TaskManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskManagerClient interface {
	CreateTask(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Task, error)
	GetTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error)
}

type taskManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskManagerClient(cc grpc.ClientConnInterface) TaskManagerClient {
	return &taskManagerClient{cc}
}

func (c *taskManagerClient) CreateTask(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, TaskManager_CreateTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) GetTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, TaskManager_GetTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskManagerServer is the server API for TaskManager service.
// All implementations must embed UnimplementedTaskManagerServer
// for forward compatibility
type TaskManagerServer interface {
	CreateTask(context.Context, *empty.Empty) (*Task, error)
	GetTask(context.Context, *Task) (*Task, error)
	mustEmbedUnimplementedTaskManagerServer()
}

// UnimplementedTaskManagerServer must be embedded to have forward compatible implementations.
type UnimplementedTaskManagerServer struct {
}

func (UnimplementedTaskManagerServer) CreateTask(context.Context, *empty.Empty) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTaskManagerServer) GetTask(context.Context, *Task) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedTaskManagerServer) mustEmbedUnimplementedTaskManagerServer() {}

// UnsafeTaskManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskManagerServer will
// result in compilation errors.
type UnsafeTaskManagerServer interface {
	mustEmbedUnimplementedTaskManagerServer()
}

func RegisterTaskManagerServer(s grpc.ServiceRegistrar, srv TaskManagerServer) {
	s.RegisterService(&TaskManager_ServiceDesc, srv)
}

func _TaskManager_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskManager_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).CreateTask(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskManager_GetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).GetTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskManager_ServiceDesc is the grpc.ServiceDesc for TaskManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TaskManager",
	HandlerType: (*TaskManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskManager_CreateTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _TaskManager_GetTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}