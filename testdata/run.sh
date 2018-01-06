# Upload the file to the acceptance server with --insecure. Insecure stops
# certificate validation, but doesn't deactivate the https nature of the
# connection. It can't protect you from a MITM attack, but it does allow you to
# test HTTPS connection parameters.
curl \
  -vvv \
  --insecure \
  --data-ascii 1.csv \
  "https://localhost:9000/v1/camera/upload?deviceId=12345678901234567890"
