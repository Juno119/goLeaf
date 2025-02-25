// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             (unknown)
// source: leaf-grpc/v1/segment.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationLeafSegmentServiceGenSegmentCache = "/leafgrpc.v1.LeafSegmentService/GenSegmentCache"
const OperationLeafSegmentServiceGenSegmentDb = "/leafgrpc.v1.LeafSegmentService/GenSegmentDb"
const OperationLeafSegmentServiceGenSegmentId = "/leafgrpc.v1.LeafSegmentService/GenSegmentId"

type LeafSegmentServiceHTTPServer interface {
	GenSegmentCache(context.Context, *IdRequest) (*SegmentBufferCacheViews, error)
	GenSegmentDb(context.Context, *IdRequest) (*LeafAllocDbs, error)
	GenSegmentId(context.Context, *IdRequest) (*IdReply, error)
}

func RegisterLeafSegmentServiceHTTPServer(s *http.Server, srv LeafSegmentServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/api/segment/get/{tag}", _LeafSegmentService_GenSegmentId0_HTTP_Handler(srv))
	r.GET("/monitor/cache", _LeafSegmentService_GenSegmentCache0_HTTP_Handler(srv))
	r.GET("/monitor/db", _LeafSegmentService_GenSegmentDb0_HTTP_Handler(srv))
}

func _LeafSegmentService_GenSegmentId0_HTTP_Handler(srv LeafSegmentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLeafSegmentServiceGenSegmentId)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GenSegmentId(ctx, req.(*IdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IdReply)
		return ctx.Result(200, reply)
	}
}

func _LeafSegmentService_GenSegmentCache0_HTTP_Handler(srv LeafSegmentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLeafSegmentServiceGenSegmentCache)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GenSegmentCache(ctx, req.(*IdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SegmentBufferCacheViews)
		return ctx.Result(200, reply)
	}
}

func _LeafSegmentService_GenSegmentDb0_HTTP_Handler(srv LeafSegmentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLeafSegmentServiceGenSegmentDb)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GenSegmentDb(ctx, req.(*IdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LeafAllocDbs)
		return ctx.Result(200, reply)
	}
}

type LeafSegmentServiceHTTPClient interface {
	GenSegmentCache(ctx context.Context, req *IdRequest, opts ...http.CallOption) (rsp *SegmentBufferCacheViews, err error)
	GenSegmentDb(ctx context.Context, req *IdRequest, opts ...http.CallOption) (rsp *LeafAllocDbs, err error)
	GenSegmentId(ctx context.Context, req *IdRequest, opts ...http.CallOption) (rsp *IdReply, err error)
}

type LeafSegmentServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewLeafSegmentServiceHTTPClient(client *http.Client) LeafSegmentServiceHTTPClient {
	return &LeafSegmentServiceHTTPClientImpl{client}
}

func (c *LeafSegmentServiceHTTPClientImpl) GenSegmentCache(ctx context.Context, in *IdRequest, opts ...http.CallOption) (*SegmentBufferCacheViews, error) {
	var out SegmentBufferCacheViews
	pattern := "/monitor/cache"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLeafSegmentServiceGenSegmentCache))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LeafSegmentServiceHTTPClientImpl) GenSegmentDb(ctx context.Context, in *IdRequest, opts ...http.CallOption) (*LeafAllocDbs, error) {
	var out LeafAllocDbs
	pattern := "/monitor/db"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLeafSegmentServiceGenSegmentDb))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LeafSegmentServiceHTTPClientImpl) GenSegmentId(ctx context.Context, in *IdRequest, opts ...http.CallOption) (*IdReply, error) {
	var out IdReply
	pattern := "/api/segment/get/{tag}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLeafSegmentServiceGenSegmentId))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
