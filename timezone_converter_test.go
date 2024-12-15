package main

import (
	"context"
	"net"
	"testing"
	"time"

	pb "github.com/bernardbaker/time.zone.converter.microservice/proto" // Update with your actual proto package

	"google.golang.org/grpc"
)

// Mock server implementation
type mockTimeZoneConverterServer struct {
	pb.UnimplementedTimeZoneConverterServer
}

func (s *mockTimeZoneConverterServer) ConvertTime(ctx context.Context, req *pb.ConvertTimeRequest) (*pb.ConvertTimeResponse, error) {
	// Mock logic replicating your actual server logic
	loc, err := time.LoadLocation(req.TargetTimezone)
	if err != nil {
		return nil, err
	}

	parsedTime, err := time.Parse(time.RFC3339, req.Timestamp)
	if err != nil {
		return nil, err
	}

	convertedTime := parsedTime.In(loc)
	return &pb.ConvertTimeResponse{
		ConvertedTimestamp: convertedTime.Format(time.RFC3339),
	}, nil
}

func TestTimeZoneConverter(t *testing.T) {
	// Start a test gRPC server
	listener, err := net.Listen("tcp", "localhost:0") // Random free port
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTimeZoneConverterServer(grpcServer, &mockTimeZoneConverterServer{})
	go grpcServer.Serve(listener)
	defer grpcServer.Stop()

	// Setup client connection
	conn, err := grpc.Dial(listener.Addr().String(), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewTimeZoneConverterClient(conn)

	// Test cases
	tests := []struct {
		name             string
		request          *pb.ConvertTimeRequest
		expectedResponse *pb.ConvertTimeResponse
		expectError      bool
	}{
		{
			name: "Valid conversion",
			request: &pb.ConvertTimeRequest{
				Timestamp:      "2024-12-04T15:00:00Z",
				TargetTimezone: "America/New_York",
			},
			expectedResponse: &pb.ConvertTimeResponse{
				ConvertedTimestamp: "2024-12-04T10:00:00-05:00", // Expected conversion
			},
			expectError: false,
		},
		{
			name: "Invalid timezone",
			request: &pb.ConvertTimeRequest{
				Timestamp:      "2024-12-04T15:00:00Z",
				TargetTimezone: "Invalid/Timezone",
			},
			expectedResponse: nil,
			expectError:      true,
		},
		{
			name: "Invalid timestamp format",
			request: &pb.ConvertTimeRequest{
				Timestamp:      "invalid-timestamp",
				TargetTimezone: "America/New_York",
			},
			expectedResponse: nil,
			expectError:      true,
		},
	}

	// Run tests
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := client.ConvertTime(context.Background(), tc.request)

			if tc.expectError {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if resp.ConvertedTimestamp != tc.expectedResponse.ConvertedTimestamp {
				t.Errorf("expected %v, got %v", tc.expectedResponse.ConvertedTimestamp, resp.ConvertedTimestamp)
			}
		})
	}
}
