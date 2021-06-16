package start

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/SoloDeploy/solo-providers-git-github/lib"
	pb "github.com/SoloDeploy/solo/grpc/git"
)

// server is used to implement github.com/SoloDeploy/solo/grpc/git.GitProviderServer
type Server struct {
	pb.UnimplementedGitProviderServer
}

// GetRepositoryNames implements GitProviderServer.GetRepositoryNames to return a list
// of repository names
func (s *Server) GetRepositoryNames(ctx context.Context, in *pb.GetRepositoryNamesRequest) (*pb.GetRepositoryNamesResponse, error) {
	log.Println("Received incoming request")
	names, _ := lib.GetRepositoryNames("org", "pat")
	return &pb.GetRepositoryNamesResponse{
		Names: names,
	}, nil
}

// Close implements GitProviderServer.Close to stop the listening gRPC server and kill the process
func (s *Server) Close(ctx context.Context, in *pb.CloseRequest) (*pb.CloseResponse, error) {
	log.Println("Stopping the Solo GitHub Provider gRPC server")
	go func() {
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}()
	return &pb.CloseResponse{}, nil
}
