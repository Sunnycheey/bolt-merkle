// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// StorageServiceClient is the client API for StorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageServiceClient interface {
	Set(ctx context.Context, in *RequestKV, opts ...grpc.CallOption) (*Status, error)
	SetBatch(ctx context.Context, in *RequestKVs, opts ...grpc.CallOption) (*Status, error)
	Get(ctx context.Context, in *RequestKey, opts ...grpc.CallOption) (*ReturnVal, error)
	GetBatch(ctx context.Context, in *RequestKeys, opts ...grpc.CallOption) (*ReturnVals, error)
	InitDatabase(ctx context.Context, in *InitParam, opts ...grpc.CallOption) (*Status, error)
	GetRootHash(ctx context.Context, in *Bucket, opts ...grpc.CallOption) (*RootHash, error)
	CloseDB(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error)
	Del(ctx context.Context, in *RequestKey, opts ...grpc.CallOption) (*Status, error)
	DelBatch(ctx context.Context, in *RequestKeys, opts ...grpc.CallOption) (*Status, error)
	RangeQuery(ctx context.Context, in *ArrayRangeKey, opts ...grpc.CallOption) (*ReturnKVs, error)
}

type storageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageServiceClient(cc grpc.ClientConnInterface) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) Set(ctx context.Context, in *RequestKV, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/BSSE.protocol.StorageService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) SetBatch(ctx context.Context, in *RequestKVs, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/BSSE.protocol.StorageService/SetBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) Get(ctx context.Context, in *RequestKey, opts ...grpc.CallOption) (*ReturnVal, error) {
	out := new(ReturnVal)
	err := c.cc.Invoke(ctx, "/BSSE.protocol.StorageService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetBatch(ctx context.Context, in *RequestKeys, opts ...grpc.CallOption) (*ReturnVals, error) {
	out := new(ReturnVals)
	err := c.cc.Invoke(ctx, "/BSSE.protocol.StorageService/GetBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) InitDatabase(ctx context.Context, in *InitParam, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/BSSE.protocol.StorageService/InitDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetRootHash(ctx context.Context, in *Bucket, opts ...grpc.CallOption) (*RootHash, error) {
	out := new(RootHash)
	err := c.cc.Invoke(ctx, "/BSSE.protocol.StorageService/GetRootHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) CloseDB(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/BSSE.protocol.StorageService/CloseDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) Del(ctx context.Context, in *RequestKey, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/BSSE.protocol.StorageService/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) DelBatch(ctx context.Context, in *RequestKeys, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/BSSE.protocol.StorageService/DelBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) RangeQuery(ctx context.Context, in *ArrayRangeKey, opts ...grpc.CallOption) (*ReturnKVs, error) {
	out := new(ReturnKVs)
	err := c.cc.Invoke(ctx, "/BSSE.protocol.StorageService/RangeQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServiceServer is the server API for StorageService service.
// All implementations should embed UnimplementedStorageServiceServer
// for forward compatibility
type StorageServiceServer interface {
	Set(context.Context, *RequestKV) (*Status, error)
	SetBatch(context.Context, *RequestKVs) (*Status, error)
	Get(context.Context, *RequestKey) (*ReturnVal, error)
	GetBatch(context.Context, *RequestKeys) (*ReturnVals, error)
	InitDatabase(context.Context, *InitParam) (*Status, error)
	GetRootHash(context.Context, *Bucket) (*RootHash, error)
	CloseDB(context.Context, *Empty) (*Status, error)
	Del(context.Context, *RequestKey) (*Status, error)
	DelBatch(context.Context, *RequestKeys) (*Status, error)
	RangeQuery(context.Context, *ArrayRangeKey) (*ReturnKVs, error)
}

// UnimplementedStorageServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStorageServiceServer struct {
}

func (UnimplementedStorageServiceServer) Set(context.Context, *RequestKV) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedStorageServiceServer) SetBatch(context.Context, *RequestKVs) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBatch not implemented")
}
func (UnimplementedStorageServiceServer) Get(context.Context, *RequestKey) (*ReturnVal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStorageServiceServer) GetBatch(context.Context, *RequestKeys) (*ReturnVals, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBatch not implemented")
}
func (UnimplementedStorageServiceServer) InitDatabase(context.Context, *InitParam) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitDatabase not implemented")
}
func (UnimplementedStorageServiceServer) GetRootHash(context.Context, *Bucket) (*RootHash, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRootHash not implemented")
}
func (UnimplementedStorageServiceServer) CloseDB(context.Context, *Empty) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseDB not implemented")
}
func (UnimplementedStorageServiceServer) Del(context.Context, *RequestKey) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (UnimplementedStorageServiceServer) DelBatch(context.Context, *RequestKeys) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelBatch not implemented")
}
func (UnimplementedStorageServiceServer) RangeQuery(context.Context, *ArrayRangeKey) (*ReturnKVs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RangeQuery not implemented")
}

// UnsafeStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServiceServer will
// result in compilation errors.
type UnsafeStorageServiceServer interface {
	mustEmbedUnimplementedStorageServiceServer()
}

func RegisterStorageServiceServer(s grpc.ServiceRegistrar, srv StorageServiceServer) {
	s.RegisterService(&StorageService_ServiceDesc, srv)
}

func _StorageService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestKV)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BSSE.protocol.StorageService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).Set(ctx, req.(*RequestKV))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_SetBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestKVs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).SetBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BSSE.protocol.StorageService/SetBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).SetBatch(ctx, req.(*RequestKVs))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BSSE.protocol.StorageService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).Get(ctx, req.(*RequestKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestKeys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BSSE.protocol.StorageService/GetBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetBatch(ctx, req.(*RequestKeys))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_InitDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).InitDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BSSE.protocol.StorageService/InitDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).InitDatabase(ctx, req.(*InitParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetRootHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetRootHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BSSE.protocol.StorageService/GetRootHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetRootHash(ctx, req.(*Bucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_CloseDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).CloseDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BSSE.protocol.StorageService/CloseDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).CloseDB(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BSSE.protocol.StorageService/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).Del(ctx, req.(*RequestKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_DelBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestKeys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).DelBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BSSE.protocol.StorageService/DelBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).DelBatch(ctx, req.(*RequestKeys))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_RangeQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArrayRangeKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).RangeQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BSSE.protocol.StorageService/RangeQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).RangeQuery(ctx, req.(*ArrayRangeKey))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageService_ServiceDesc is the grpc.ServiceDesc for StorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BSSE.protocol.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _StorageService_Set_Handler,
		},
		{
			MethodName: "SetBatch",
			Handler:    _StorageService_SetBatch_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _StorageService_Get_Handler,
		},
		{
			MethodName: "GetBatch",
			Handler:    _StorageService_GetBatch_Handler,
		},
		{
			MethodName: "InitDatabase",
			Handler:    _StorageService_InitDatabase_Handler,
		},
		{
			MethodName: "GetRootHash",
			Handler:    _StorageService_GetRootHash_Handler,
		},
		{
			MethodName: "CloseDB",
			Handler:    _StorageService_CloseDB_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _StorageService_Del_Handler,
		},
		{
			MethodName: "DelBatch",
			Handler:    _StorageService_DelBatch_Handler,
		},
		{
			MethodName: "RangeQuery",
			Handler:    _StorageService_RangeQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
