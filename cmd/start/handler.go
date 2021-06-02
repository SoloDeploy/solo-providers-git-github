package start

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/SoloDeploy/solo/grpc/git"
	"google.golang.org/grpc"
)

// server is used to implement github.com/SoloDeploy/solo/grpc/git.GitProviderServer
type server struct {
	pb.UnimplementedGitProviderServer
}

// GetRepositoryNames implements GitProviderServer.GetRepositoryNames to return a list
// of repository names
func (s *server) GetRepositoryNames(ctx context.Context, in *pb.GetRepositoryNamesRequest) (*pb.GetRepositoryNamesResponse, error) {
	log.Print("Received incoming request")
	names := []string{
		"repository1",
		"repository2",
		"repository3",
	}
	return &pb.GetRepositoryNamesResponse{
		Names: names,
	}, nil
}

func handler(port string) (err error) {
	address := fmt.Sprintf("localhost:%v", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return
	}
	s := grpc.NewServer()
	pb.RegisterGitProviderServer(s, &server{})
	err = s.Serve(lis)
	return
}
