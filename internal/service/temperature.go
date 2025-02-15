package service

import (
	"context"
	"fmt"

	pb "github.com/marcosnasp/temperature-converter/proto/gen/go"
)

// TemperatureService implementa a interface do gRPC
type TemperatureService struct {
	pb.UnimplementedTemperatureConverterServer
}

// Convert realiza a conversão entre diferentes unidades de temperatura
func (s *TemperatureService) Convert(
	ctx context.Context,
	req *pb.ConversionRequest,
) (*pb.ConversionResponse, error) {
	temperature := req.Temperature
	from := req.FromUnit
	to := req.ToUnit

	// Verifica se a unidade de origem e destino são válidas
	if from == to {
		return &pb.ConversionResponse{
			Temperature: temperature,
			Unit:        to,
		}, nil
	}

	var convertedTemp float64

	// Conversão para Celsius como unidade intermediária
	switch from {
	case pb.Unit_CELSIUS:
		convertedTemp = temperature
	case pb.Unit_FAHRENHEIT:
		convertedTemp = (temperature - 32) * 5 / 9
	case pb.Unit_KELVIN:
		convertedTemp = temperature - 273.15
	default:
		return nil, fmt.Errorf("unidade de origem desconhecida: %v", from)
	}

	// Conversão de Celsius para a unidade desejada
	switch to {
	case pb.Unit_CELSIUS:
		// Já está em Celsius
	case pb.Unit_FAHRENHEIT:
		convertedTemp = (convertedTemp * 9 / 5) + 32
	case pb.Unit_KELVIN:
		convertedTemp = convertedTemp + 273.15
	default:
		return nil, fmt.Errorf("unidade de destino desconhecida: %v", to)
	}

	// Retorna a resposta
	return &pb.ConversionResponse{
		Temperature: convertedTemp,
		Unit:        to,
	}, nil
}
