// cliente.go
package main

import (
	"context"
	"log"
	"time"

	pb "cazarrecompensas/proto/grpc-server/proto"

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

	r, err := c.Saludar(ctx, &pb.SaludoRequest{Nombre: "Flavio"})
	if err != nil {
		log.Fatalf("Error en la solicitud: %v", err)
	}

	log.Printf("Respuesta: %s", r.Mensaje)
}
