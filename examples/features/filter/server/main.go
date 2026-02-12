//
//
// Tencent is pleased to support the open source community by making tRPC available.
//
// Copyright (C) 2023 Tencent.
// All rights reserved.
//
// If you have downloaded a copy of the tRPC source code from Tencent,
// please note that tRPC source code is licensed under the  Apache 2.0 License,
// A copy of the Apache 2.0 License is included in this file.
//
//

// Package main entry
package main

import (
	"context"

	"github.com/NHYCRaymond/trpc-go/examples/features/filter/shared"

	trpc "github.com/NHYCRaymond/trpc-go"
	"github.com/NHYCRaymond/trpc-go/errs"
	"github.com/NHYCRaymond/trpc-go/examples/features/common"
	"github.com/NHYCRaymond/trpc-go/filter"
	"github.com/NHYCRaymond/trpc-go/log"
	"github.com/NHYCRaymond/trpc-go/server"
	pb "github.com/NHYCRaymond/trpc-go/testdata/trpc/helloworld"
)

func main() {
	// Create a server with filter
	s := trpc.NewServer(server.WithFilter(serverFilter))
	pb.RegisterGreeterService(s, &common.GreeterServerImpl{})
	// Start serving.
	s.Serve()
}

func serverFilter(ctx context.Context, req interface{}, next filter.ServerHandleFunc) (rsp interface{}, err error) {
	log.InfoContext(ctx, "server filter start %s", trpc.GetMetaData(ctx, shared.AuthKey))
	// check token from context metadata
	if !valid(trpc.GetMetaData(ctx, shared.AuthKey)) {
		return nil, errs.Newf(errs.RetServerAuthFail, "auth fail")
	}
	// run business logic
	rsp, err = next(ctx, req)

	log.InfoContext(ctx, "server filter end")
	return rsp, err
}

// valid validates the authorization
func valid(authorization []byte) bool {
	return string(authorization) == shared.Token
}
