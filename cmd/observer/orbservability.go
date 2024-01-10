package main

import (
	"context"
	"fmt"
	"orbservability/observer/pkg/config"
	"reflect"

	pb "github.com/orbservability/schemas/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createGrpcStream(ctx context.Context, cfg *config.Config) (*grpc.ClientConn, pb.EventGatewayService_StreamEventsClient, error) {
	conn, err := grpc.Dial(cfg.OrbservabilityURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	client := pb.NewEventGatewayServiceClient(conn)
	stream, err := client.StreamEvents(ctx)
	if err != nil {
		return nil, nil, err
	}

	return conn, stream, nil
}

func streamEvent(stream pb.EventGatewayService_StreamEventsClient, recordMap map[string]interface{}) error {
	msg := &pb.PixieEvent{}
	err := mapToProto(recordMap, msg)
	if err != nil {
		return err
	}

	if err := stream.Send(msg); err != nil {
		return err
	}

	return nil
}

func mapToProto(recordMap map[string]interface{}, msg *pb.PixieEvent) error {
	msgVal := reflect.ValueOf(msg).Elem()

	for key, value := range recordMap {
		field := msgVal.FieldByName(key)
		if !field.IsValid() || !field.CanSet() {
			continue // Skip invalid or unsettable fields
		}

		fieldValue := reflect.ValueOf(value)

		// Check if the field is a nested message
		if field.Kind() == reflect.Struct && fieldValue.Kind() == reflect.Map {
			nestedMap, ok := value.(map[string]interface{})
			if !ok {
				return fmt.Errorf("expected map for nested field %s", key)
			}

			// Recursively call mapToProto for the nested object
			err := mapToProto(nestedMap, field.Addr().Interface().(*pb.PixieEvent))
			if err != nil {
				return fmt.Errorf("error setting nested field %s: %v", key, err)
			}

		} else if field.Type() == fieldValue.Type() {
			field.Set(fieldValue)
		} else {
			// Handle type conversion or return an error
			return fmt.Errorf("type mismatch for field %s", key)
		}
	}

	return nil
}
