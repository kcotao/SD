// cliente.go
package main

import (
	"context"
	"log"
	"time"

	pb "marina/proto/grpc-server/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	c := pb.NewSaludoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Saludar(ctx, &pb.SaludoRequest{Nombre: "Marina"})
	if err != nil {
		log.Fatalf("Error en la solicitud: %v", err)
	}

	log.Printf("Respuesta: %s", r.Mensaje)
	conn1, err1 := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err1 != nil {
		log.Fatalf("No se pudo conectar: %v", err1)
	}
	defer conn1.Close()

	c1 := pb.NewSaludoServiceClient(conn1)

	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()

	r1, err1 := c1.Saludar(ctx1, &pb.SaludoRequest{Nombre: "Marina"})
	if err1 != nil {
		log.Fatalf("Error en la solicitud: %v", err1)
	}

	log.Printf("Respuesta: %s", r1.Mensaje)

}
