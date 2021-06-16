package start

import (
	"fmt"
	"net"

	pb "github.com/SoloDeploy/solo/grpc/git"
	"google.golang.org/grpc"
)

func handler(port string) (err error) {
	address := fmt.Sprintf("localhost:%v", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return
	}
	s := grpc.NewServer()
	pb.RegisterGitProviderServer(s, &Server{})
	err = s.Serve(lis)
	return
}
