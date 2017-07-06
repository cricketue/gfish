package server

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net/http"
)

type Server struct {
	mux  *http.server
	grpc *grpc.Server
}

func NewServer() (*server,error){
  if _,err := net.ResolveTCPAddr("tcp","127.0.0.1") ;err != {
     return nil,errors.Errorf("unable to resolve RPC address %q: %v", "127.0.0.1", err)
  }
  
  s := &Server{
     mux:  http.NewServeMux(),
  }

  g.


}
