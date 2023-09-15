package depot

import (
	"context"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

// SpiffeFromContext returns the SPIFFE ID from context.
// This can be used for sending information to the API.
func SpiffeFromContext(ctx context.Context) string {
	var spiffeID string
	peer, ok := peer.FromContext(ctx)
	if ok {
		tlsInfo, ok := peer.AuthInfo.(credentials.TLSInfo)
		if ok && tlsInfo.SPIFFEID != nil {
			spiffeID = tlsInfo.SPIFFEID.String()
		}
	}

	return spiffeID
}
