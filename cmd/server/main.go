package main

import (
	"context"
	"log"
	"net"

	pb "github.com/marcosnasp/temperature-converter/proto/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedTemperatureConverterServer
}

func (s *server) Convert(ctx context.Context, req *pb.ConversionRequest) (*pb.ConversionResponse, error) {
	// Se as unidades forem iguais, retorne o valor original
	if req.FromUnit == req.ToUnit {
		return &pb.ConversionResponse{
			Temperature: req.Temperature,
			Unit:        req.ToUnit,
		}, nil
	}

	// Validação das unidades
	validUnits := map[pb.Unit]bool{
		pb.Unit_CELSIUS:    true,
		pb.Unit_FAHRENHEIT: true,
		pb.Unit_KELVIN:     true,
	}
	if !validUnits[req.FromUnit] || !validUnits[req.ToUnit] {
		return nil, status.Errorf(codes.InvalidArgument, "unidade inválida")
	}

	var result float64

	switch req.FromUnit {
	case pb.Unit_CELSIUS:
		switch req.ToUnit {
		case pb.Unit_FAHRENHEIT:
			result = (req.Temperature * 9 / 5) + 32
		case pb.Unit_KELVIN:
			result = req.Temperature + 273.15
		}
	case pb.Unit_FAHRENHEIT:
		switch req.ToUnit {
		case pb.Unit_CELSIUS:
			result = (req.Temperature - 32) * 5 / 9
		case pb.Unit_KELVIN:
			celsius := (req.Temperature - 32) * 5 / 9
			result = celsius + 273.15
		}
	case pb.Unit_KELVIN:
		switch req.ToUnit {
		case pb.Unit_CELSIUS:
			result = req.Temperature - 273.15
		case pb.Unit_FAHRENHEIT:
			celsius := req.Temperature - 273.15
			result = (celsius * 9 / 5) + 32
		}
	}

	return &pb.ConversionResponse{
		Temperature: result,
		Unit:        req.ToUnit,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Falha ao ouvir: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTemperatureConverterServer(s, &server{})
	log.Printf("Servidor ouvindo em %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falha ao servir: %v", err)
	}
}