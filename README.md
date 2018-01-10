# Vision Screening Central Server Simulator

This simulator demonstrates the interface I would like to set up to upload
individual vision screening results to a centralized database.

## Instructions

How to use the server.

1.  Follow the Go [instructions](https://golang.org/doc/install#install) to
    install a local development environment.
1.  Follow the instructions in [Test your
    installation](https://golang.org/doc/install#testing) to set your `$GOPATH`
    and compile your first hello world application.
1.  Download this repo's sourcecode with
    `go get github.com/achew22/acceptance-testing-vision-upload-server/cmd`
1.  `cd` into the root of this repo which should be located at
    `$GOPATH/src/github.com/achew22/acceptance-testing-vision-upload-server`
1.  Generate the required keys for testing with `bash generate_cert.sh`.
1.  Run the testing server with `go run ./cmd/main.go`

Success! You now have a working testing server running. You can upload sample
data to it by running `bash testdata/run.sh` and uploading any of the sample
files available in `testdata`. For example, to see a working upload, you can
run `bash testdata/run.sh testdata/good.in`.

## Caveats

In this section I attempt to surface all the places where there are known
differences between the fake implementation of this service and the real one.

### SSL

I will not be issuing any Certificate Authority (CA) signed certificates for
the simulator.  Instead I will be using a leaf 4096 bit RSA public/private
keypair. In production the service will use a certificate (no implementation
details specified) signed by one of the major CA vendors.

I reserve the right to use any algorithm, key size, or signing authority  in
production. However, I will likely use the following TLS configuration:

```
Ciphersuites: ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256
Versions: TLSv1.2
TLS curves: prime256v1, secp384r1, secp521r1
Certificate type: ECDSA
Certificate curve: prime256v1, secp384r1, secp521r1
Certificate signature: sha256WithRSAEncryption, ecdsa-with-SHA256, ecdsa-with-SHA384, ecdsa-with-SHA512
RSA key size: 2048 (if not ecdsa)
DH Parameter size: None (disabled entirely)
ECDH Parameter size: 256
HSTS: max-age=15768000 (See HSTS section below)
Certificate switching: None
```

For more information on how these values were chosen, please see
[Mozilla Server Side TLS - Modern Compabability](
https://wiki.mozilla.org/Security/Server_Side_TLS#Modern_compatibility)

### HSTS

All websites I host use HSTS. However, HSTS can poison connections when used
for local testing. In order to avoid this, the simulator does not use HSTS.

### Hostname

The test service can be accessed by addressing it on the correct port. No HTTP
`host` header will be required to access the site. In production, a host header
will be required.

Additionally TLS SNI (Server Name Indicator) should be allowed but not required
by the client.

### Data storage

The simulator is stateless and will not save anything. It receives the data
from the client, validates it, writes to the console the data that was uploaded
and forgets it. There is no way to retrieve anything from the simulator after
it has been uploaded.
