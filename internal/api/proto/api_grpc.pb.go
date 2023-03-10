// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	Play(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Message, error)
	Pause(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Message, error)
	AddSong(ctx context.Context, in *DataSongRequest, opts ...grpc.CallOption) (*AddSongResponse, error)
	Next(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Message, error)
	Prev(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Message, error)
	Stop(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Message, error)
	ShowSong(ctx context.Context, in *SongName, opts ...grpc.CallOption) (*ShowSongResponse, error)
	UpdateSong(ctx context.Context, in *UpdateDataSongRequest, opts ...grpc.CallOption) (*Message, error)
	DeleteSong(ctx context.Context, in *SongName, opts ...grpc.CallOption) (*Message, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) Play(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/api.Api/Play", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Pause(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/api.Api/Pause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) AddSong(ctx context.Context, in *DataSongRequest, opts ...grpc.CallOption) (*AddSongResponse, error) {
	out := new(AddSongResponse)
	err := c.cc.Invoke(ctx, "/api.Api/AddSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Next(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/api.Api/Next", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Prev(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/api.Api/Prev", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Stop(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/api.Api/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ShowSong(ctx context.Context, in *SongName, opts ...grpc.CallOption) (*ShowSongResponse, error) {
	out := new(ShowSongResponse)
	err := c.cc.Invoke(ctx, "/api.Api/ShowSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) UpdateSong(ctx context.Context, in *UpdateDataSongRequest, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/api.Api/UpdateSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) DeleteSong(ctx context.Context, in *SongName, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/api.Api/DeleteSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	Play(context.Context, *Request) (*Message, error)
	Pause(context.Context, *Request) (*Message, error)
	AddSong(context.Context, *DataSongRequest) (*AddSongResponse, error)
	Next(context.Context, *Request) (*Message, error)
	Prev(context.Context, *Request) (*Message, error)
	Stop(context.Context, *Request) (*Message, error)
	ShowSong(context.Context, *SongName) (*ShowSongResponse, error)
	UpdateSong(context.Context, *UpdateDataSongRequest) (*Message, error)
	DeleteSong(context.Context, *SongName) (*Message, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) Play(context.Context, *Request) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Play not implemented")
}
func (UnimplementedApiServer) Pause(context.Context, *Request) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedApiServer) AddSong(context.Context, *DataSongRequest) (*AddSongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSong not implemented")
}
func (UnimplementedApiServer) Next(context.Context, *Request) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Next not implemented")
}
func (UnimplementedApiServer) Prev(context.Context, *Request) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prev not implemented")
}
func (UnimplementedApiServer) Stop(context.Context, *Request) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedApiServer) ShowSong(context.Context, *SongName) (*ShowSongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowSong not implemented")
}
func (UnimplementedApiServer) UpdateSong(context.Context, *UpdateDataSongRequest) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSong not implemented")
}
func (UnimplementedApiServer) DeleteSong(context.Context, *SongName) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSong not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_Play_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Play(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/Play",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Play(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/Pause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Pause(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_AddSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).AddSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/AddSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).AddSong(ctx, req.(*DataSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Next_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Next(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/Next",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Next(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Prev_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Prev(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/Prev",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Prev(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Stop(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ShowSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ShowSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/ShowSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ShowSong(ctx, req.(*SongName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_UpdateSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDataSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).UpdateSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/UpdateSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).UpdateSong(ctx, req.(*UpdateDataSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_DeleteSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).DeleteSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/DeleteSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).DeleteSong(ctx, req.(*SongName))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Play",
			Handler:    _Api_Play_Handler,
		},
		{
			MethodName: "Pause",
			Handler:    _Api_Pause_Handler,
		},
		{
			MethodName: "AddSong",
			Handler:    _Api_AddSong_Handler,
		},
		{
			MethodName: "Next",
			Handler:    _Api_Next_Handler,
		},
		{
			MethodName: "Prev",
			Handler:    _Api_Prev_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Api_Stop_Handler,
		},
		{
			MethodName: "ShowSong",
			Handler:    _Api_ShowSong_Handler,
		},
		{
			MethodName: "UpdateSong",
			Handler:    _Api_UpdateSong_Handler,
		},
		{
			MethodName: "DeleteSong",
			Handler:    _Api_DeleteSong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
