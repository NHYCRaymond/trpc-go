//go:build !(linux && amd64)
// +build !linux !amd64

package server

import "github.com/NHYCRaymond/trpc-go/transport"

func attemptSwitchingTransport(o *Options) transport.ServerTransport {
	if o.Transport == nil {
		return transport.DefaultServerStreamTransport
	}
	return o.Transport
}
