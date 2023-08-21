# Latency Service

The Latency Service provides accurate Round-Trip-Time (RTT) measurements from the HTTP client in the form of a JSON response.

## Key Features

* **API Endpoint for RTT Calculation**: Utilize an API endpoint to accurately determine Round-Trip-Time (RTT) from the HTTPS server.
* **Flexible Logging Options**: Log the results to a designated log file if logging is enabled, aiding in performance analysis and troubleshooting.
* **Automated Certificate Management**: Seamlessly obtain and refresh Let's Encrypt certificates to ensure secure communications.

## Configuration

When launching the container or application, make sure to set the following environment variables:

```bash
LATENCY_HOST="your_host_here"
LATENCY_DATA_DIRECTORY="your_data_directory_here"
LATENCY_LISTEN_HTTP="your_http_listener_here"
LATENCY_LISTEN_HTTPS="your_https_listener_here"
LATENCY_LETS_ENCRYPT="true_or_false"
LATENCY_LOGGING="true_or_false"
```

## Deployment Instructions

Deploy the service as a Docker container on your system using the following steps. Ensure that you utilize `--net host` for accurate results. Additionally, set up log rotation for the `LATENCY_DATA_DIRECTORY/latency.logs` file or disable logging by configuring `LATENCY_LOGGING`.

### Example Deployment:

```bash
docker run -t -i --name "latency-service" --net host --pull always --rm \
  -v /opt/latency:/data:z \
  -e LATENCY_LETS_ENCRYPT=true \
  -e LATENCY_HOST=test.latency.g-portal.xyz \
  -e LATENCY_LISTEN_HTTP="127.0.0.1:80" \
  -e LATENCY_LISTEN_HTTPS="127.0.0.1:443" \
  "docker.io/gportal/latency-service:latest"
```

## Sample Response

For example, upon making a `GET` request to `http://127.0.0.1:8443/ping`, you will receive a JSON response as follows:

```json
{
  "source_ip": "172.17.0.1",
  "latency_ms": 23,
  "time": "2022-03-10T16:06:38.711462804Z"
}
```

Experience accurate Round-Trip-Time (RTT) measurement with our Latency Service!