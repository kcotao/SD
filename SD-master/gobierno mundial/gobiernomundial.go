// servidor.go
package main

import (
	"context"
	"log"
	"net"

	pb "gobierno/proto/grpc-server/proto" // import generado desde .proto

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSaludoServiceServer
}

func (s *server) Saludar(ctx context.Context, req *pb.SaludoRequest) (*pb.SaludoResponse, error) {
	log.Printf("Recibido: %s", req.Nombre)
	return &pb.SaludoResponse{Mensaje: "Hola " + req.Nombre}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051") // escucha en el puerto 50051
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSaludoServiceServer(grpcServer, &server{})
	log.Println("Servidor gRPC en ejecución...")
	grpcServer.Serve(lis)
}
