// servidor.go
package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "cazarecompensas/proto/grpc-server/proto" // import generado desde .proto

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

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	c := pb.NewSaludoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Saludar(ctx, &pb.SaludoRequest{Nombre: "Cazzarecompensas"})
	if err != nil {
		log.Fatalf("Error en la solicitud: %v", err)
	}

	log.Printf("Respuesta: %s", r.Mensaje)

	lis, err := net.Listen("tcp", ":50053") // escucha en el puerto 50051
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSaludoServiceServer(grpcServer, &server{})
	log.Println("Servidor gRPC cazarrecompensa en ejecución...")
	grpcServer.Serve(lis)
}
