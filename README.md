# Project: Time Zone Converter Microservice

## Overview: A gRPC service where a client sends a request with:

1. A timestamp (in UTC or any other timezone)

2. A target timezone

The service responds with the equivalent **timestamp** in the **target** timezone.

---

# Features:

1. **Proto Definition**: Define a .proto file with:

   - A ConvertTimeRequest message containing the timestamp and target timezone.

   - A ConvertTimeResponse message containing the converted timestamp.

   - A TimeZoneConverter service with one method, ConvertTime.

2. **Server Implementation**: Implement the TimeZoneConverter service in Golang:

   - Use the time and github.com/rickar/cal libraries for timezone conversions.

   - Handle basic errors (e.g., invalid timezones).

3. **Client Implementation**: Write a simple Golang client that:

   - Accepts input for timestamp and timezone via command line or hardcoded values.

   - Calls the gRPC server and prints the converted timestamp.

---

# Explanation

1. Parsing the Time:

   - The input time (2024-12-04T15:00:00Z) is in UTC, so it’s parsed using time.Parse with the time.RFC3339 layout.

   - The result is a time.Time object in the UTC time zone.

2. Loading the Target Time Zone:

   - Use time.LoadLocation with the IANA time zone name (e.g., America/New_York).

3. Converting Time:

   - The In(location) method converts the time to the desired time zone.

4. Output:

   - Print both the original UTC time and the converted time for verification.

---

# Testing

## Run module tests:

```bash
go run main.go
```

```bash
go test -v
```

## Example grpcurl Command with .proto File:

```bash
grpcurl -plaintext \
  -import-path ./proto \
  -proto api.proto \
  -d '{"timestamp": "2024-12-04T15:00:00Z", "target_timezone": "America/New_York"}' \
  localhost:8080 \
  timezone.TimeZoneConverter/ConvertTime
```

## Explanation:

`-import-path ./proto`: Specifies the directory where your .proto file is located (adjust the path accordingly).

`-proto timezone.proto`: Specifies the .proto file that contains the service definition.

`-d`: Specifies the JSON payload, as in the previous example.

`localhost:50051`: Replace with your actual server address.

`timezone.TimeZoneConverter/ConvertTime`: Specifies the full service and method name, as per your .proto file.
