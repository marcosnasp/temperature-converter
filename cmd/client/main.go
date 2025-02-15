package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	pb "github.com/marcosnasp/temperature-converter/proto/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Conexão falhou: %v", err)
	}
	defer conn.Close()
	client := pb.NewTemperatureConverterClient(conn)

	if len(os.Args) != 4 {
		log.Fatalf("Uso: %s <temperatura> <de_unidade> <para_unidade>", os.Args[0])
	}

	temp, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatalf("Temperatura inválida: %v", err)
	}

	fromUnit := strings.ToLower(os.Args[2])
	toUnit := strings.ToLower(os.Args[3])

	unitMap := map[string]pb.Unit{
		"celsius":    pb.Unit_CELSIUS,
		"fahrenheit": pb.Unit_FAHRENHEIT,
		"kelvin":     pb.Unit_KELVIN,
	}

	from, ok := unitMap[fromUnit]
	if !ok {
		log.Fatalf("Unidade de origem inválida: %s", fromUnit)
	}
	to, ok := unitMap[toUnit]
	if !ok {
		log.Fatalf("Unidade de destino inválida: %s", toUnit)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Convert(ctx, &pb.ConversionRequest{
		Temperature: temp,
		FromUnit:    from,
		ToUnit:      to,
	})
	if err != nil {
		log.Fatalf("Erro na conversão: %v", err)
	}

	log.Printf("Resultado: %.2f %s", res.Temperature, toUnit)
}