grpcurl -plaintext \
  -import-path ./proto \
  -proto api.proto \
  -d '{"timestamp": "2024-12-04T15:00:00Z", "target_timezone": "America/New_York"}' \
  localhost:8080 \
  timezone.TimeZoneConverter/ConvertTime