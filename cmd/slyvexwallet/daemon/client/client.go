package client

import (
	"context"
	"github.com/slyvex-core/slyvexd/cmd/slyvexwallet/daemon/server"
	"time"

	"github.com/pkg/errors"

	"github.com/slyvex-core/slyvexd/cmd/slyvexwallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the slyvexwalletd server, and returns the client instance
func Connect(address string) (pb.SlyvexwalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("slyvexwallet daemon is not running, start it with `slyvexwallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewSlyvexwalletdClient(conn), func() {
		conn.Close()
	}, nil
}
