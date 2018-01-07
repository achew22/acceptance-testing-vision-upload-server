# Specification
> Version 2018-01-07

## Purpose

This document specifies a way for a vision screening camera to upload results
to a user designated server for aggregation and assimilation.

## Definitions

OD
 : Oculus dexter

OS
 : Oculus sinister

mm
 : A Système international (d'unités) (SI) millimetre

dpt (dioptre)
 : A unit of measurement of the optical power of a lens or curved mirror, which
   is equal to the reciprocal of the focal length measured in metres (that is,
   1/metres)

SSV
 : Semicolon Separated Value

TLS
 : Transport Layer Security

HTTPS
 : HTTP connection sent over a connection using the prescribed connection
   parameters.

HTTP POST
 : A request method supported by the HTTP protocol that can be used for
   uploading files.

Client
 : The software uploading vision screening data files who initiates the HTTPS
   connection and POSTs data

Server
 : The software receiving vision screening data files from listens for
   connections and receives POSTed data.

UTF-8
 : A character set and encoding standard described by RFC3629

RFC3339 compatible timestamp
 : A date denoted by the ABNF definition `date-time` in RFC3339 section 5.6.
   "Internet Date/Time Format"

RFC3339 compatible date
 : A date denoted by the ABNF definition `full-date` in RFC3339 section 5.6.
   "Internet Date/Time Format"

"MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT",
"RECOMMENDED",  "MAY", and "OPTIONAL"
 : to be interpreted as described in [RFC 2119](
   https://www.ietf.org/rfc/rfc2119.txt)

# Detailed specification

## File Specification

### Header

The file MUST begin with a SSV compatible header with precisely the following
text.

```
Date of measurement;Primary key;Family name;First name;Date of birth;ID;Location;Contact;Sphere [dpt] OD;Cylinder [dpt] OD;Axis [°] OD;Pupil size [mm] OD;Sphere [dpt] OS;Cylinder[dpt] OS;Axis [°] OS;Pupil size [mm] OS;Gaze asymmetry [°] OS;Pupil distance [mm];Monocular {1=OD, 2=OS, 3=Binocular};Screening result {0=Pass, 1=Refer, 2=Refer or try again};
```

This header is used to denote version 1 of the protocol. Any divergence from it
SHALL necessitate a change to the protocol and specification of this document.
Future versions of this file may use any header they please as long as it is
not identical to the header provided above.

### Fields

#### Date of measurement

A RFC3339 compatible time stamp indicating the time that the measurement was
taken.

 *  Format: RFC3339 compatable timestamp
 *  Example: "2013-10-02T15:00:00Z" or "2013-10-01T10:00:00-05:00"

#### Primary key

Unique identifier to the uploader. This value MUST be unique to the screening
result per device and will be used to return error messages to the client.

 *  Format: UTF-8 string
 *  Example: "1", "2", "1000", "abc123"

#### Family name

Family name of the person screened.

 *  Format: UTF-8 string
 *  Example: "Anderson", "Åke", "சுப்ரமணிய ", "María-Jose Carreño Quiñones"

#### First name

First (or given) name of the person screened.

 *  Format: UTF-8 string
 *  Example: "Martin", "Aðalfríður", "Andrea", "யார்"

#### Date of birth

Gregorian date of birth of the person screened.

 *  Format: RFC3339 compatible date
 *  Example: 2004-12-31 (December 31, 2004), 2007-07-08 (July 8, 2007),
    2013-02-04 (Febuary 4, 2013)

#### ID

A string based identifier for the current record.

 *  Format: UTF-8 string
 *  Example: `Anonymous-20131011_194642`

#### Location

Description of the location in which the screening took place. Freeform

 *  text, not to be interpreted by computer.
 *  Format: UTF-8 string
 *  Example: "Miss Vicky's 1st grade class", "Room 101", "123 S. Fake St."

#### Contact

Person to contact for follow-up information. Used to record the name of the

 *  parent or guardian of a minor. The provided name has NO relation to the
 *  patient name. You MUST NOT assume a shared last name.
 *  Format: UTF-8 string
 *  Example: "George Marbrook", "Nicholas Zelfor"

#### Sphere OD

 *  Format: Number
 *  Unit: Dioptres
 *  Example: -1.5, 0, 2.9472

#### Cylinder OD

 *  Format: Number
 *  Unit: Dioptres
 *  Example: -1.5, 0, 2.9472

#### Axis OD

 *  Format: Number
 *  Unit: Degrees
 *  Example: -1.5, 0, 2.9472

#### Pupil size OD

 *  Format: Number
 *  Unit: mm
 *  Example: -1.5, 0, 2.9472

#### Sphere OS

 *  Format: Number
 *  Unit: Dioptres
 *  Example: -1.5, 0, 2.9472

#### Cylinder OS

 *  Format: Number
 *  Unit: Dioptres
 *  Example: -1.5, 0, 2.9472

#### Axis OS

 *  Format: Number
 *  Unit: Degrees
 *  Example: -1.5, 0, 2.9472

#### Pupil size OS

 *  Format: Number
 *  Unit: mm
 *  Example: -1.5, 0, 2.9472

#### Gaze asymmetry OS

 *  Format: Number
 *  Unit: Degrees
 *  Example: -1.5, 0, 2.9472

#### Pupil distance [mm]

 *  Format: Number
 *  Unit: mm
 *  Example: -1.5, 0, 2.9472

#### Monocular {1=OD, 2=OS, 3=Binocular}

Rating of the monocular status of the screened patient.

 *  Format: Number 1 to 3
 *  Unit: Enumeration
 *  Example: 1 (indicating OD preference), 2 (indicating OS preference), 3
    (indicating binoclar)

#### Screening result {0=Pass, 1=Refer, 2=Refer or try again}

Camera determined score for the screened patient.

 *  Format: Number 1 to 2
 *  Unit: Enumeration
 *  Example: 0 (Pass), 1 (Refer), 2 (Refer or try again)

### ABNF for file format

This section describes an ABNF ([RFC 2234](
https://www.ietf.org/rfc/rfc2234.txt)) parser for parsing the file format.

```
ENDLINE   = CR / CRLF
SEMICOLON = ";"

HEADER = "Date of measurement;Primary key;Family name;First name;Date of birth;ID;Location;Contact;Sphere [dpt] OD;Cylinder [dpt] OD;Axis [°] OD;Pupil size [mm] OD;Sphere [dpt] OS;Cylinder[dpt] OS;Axis [°] OS;Pupil size [mm] OS;Gaze asymmetry [°] OS;Pupil distance [mm];Monocular {1=OD, 2=OS, 3=Binocular};Screening result {0=Pass, 1=Refer, 2=Refer or try again};"
HEADER_ROW = HEADER ENDLINE

date-time = RFC3339 date-time (defined in section 5.6)
full-date = RFC3339 full-date (defined in section 5.6)

INTEGER = 1*DIGIT
NEGATION_OPERATOR = "-"
NUMBER = [NEGATION_OPERATOR]INTEGER[. INTEGER]

UTF_STRING = 1*UTF8_RUNE


DATE_OF_MEASUREMENT = date-time
PRIMARY_KEY = [INTEGER]
FAMILY_NAME = [UTF_STRING]
FIRST_NAME = [UTF_STRING]
DATE_OF_BIRTH = full-date
ID = [UTF_STRING]
LOCATION = [UTF_STRING]
CONTACT = [UTF_STRING]
SPHERE_OD = NUMBER
CYLINDER_OD = NUMBER
AXIS_OD = NUMBER
PUPIL_SIZE_OD = NUMBER
SPHERE_OS = NUMBER
CYLINDER_OS = NUMBER
AXIS_OS = NUMBER
PUPIL_SIZE_OS = NUMBER
GAZE_ASYMMETRY_OS = NUMBER
PUPIL_DISTANCE = NUMBER
MONOCULAR = 1 / 2 / 3
SCREENING_RESULT = 0 / 1 / 2

DATA_ROW = DATE_OF_MEASUREMENT SEMICOLON PRIMARY_KEY SEMICOLON FAMILY_NAME SEMICOLON FIRST_NAME SEMICOLON DATE_OF_BIRTH SEMICOLON ID SEMICOLON LOCATION SEMICOLON CONTACT SEMICOLON SPHERE_OD SEMICOLON CYLINDER_OD SEMICOLON AXIS_OD SEMICOLON PUPIL_SIZE_OD SEMICOLON SPHERE_OS SEMICOLON CYLINDER_OS SEMICOLON AXIS_OS SEMICOLON PUPIL_SIZE_OS SEMICOLON GAZE_ASYMMETRY_OS SEMICOLON PUPIL_DISTANCE SEMICOLON MONOCULAR SEMICOLON SCREENING_RESULT ENDLINE

DOCUMENT = HEADER_ROW 1*DATA_ROW
```

## Network connection specification

To upload the file for consideration by the server the client MUST connect to a
specified host using TLS (connection parameters below) and use the HTTP POST
verb to post the file.

The post operation should be sent to the URL
`/v1/camera/upload` with a query parameter specifying the device ID (or serial
number) of the device doing the uploading. The device ID, MUST be transmitted
in the query string for the request with the named parameter `deviceId`.

For example, if the server were listing on the hostname "example.com" and the
device id of the client is "ABCD12345", the POST would be sent to

`https://example.com/v1/camera/upload?deviceId=ABCD12345`

The server will reply with a JSON object containing a UTF-8 encoded field that
MUST be displayed to the user, `message`. This field can hold anywhere from 0
to 1000 letters and must be displayed in its entirety to the user.

### Example

Since the connection will be relayed through a TLS connection, it is impossible
to describe the over wire bytes. Instead in this example I will provide the
bytes that are transmitted through TLS as if they were not encryped.

Presuming a camera made by "EyeCorp" model "Iris 2000" would like to upload
patient data for 1 patient

Key:

`>`
 : Indicates data sent from client to server

`<`
 : Indicates data sent from server to client

```
> POST /v1/camera/upload?deviceId=12345678901234567890 HTTP/1.1
> Host: localhost:9000
> User-Agent: curl/7.47.0
> Accept: */*
> Content-Length: 784
> Content-Type: application/x-www-form-urlencoded
>
> Date of measurement;Primary key;Family name;First name;Date of birth;ID;Location;Contact;Sphere [dpt] OD;Cylinder [dpt] OD;Axis [°] OD;Pupil size [mm] OD;Sphere [dpt] OS;Cylinder[dpt] OS;Axis [°] OS;Pupil size [mm] OS;Gaze asymmetry [°] OS;Pupil distance [mm];Monocular {1=OD, 2=OS, 3=Binocular};Screening result {0=Pass, 1=Refer, 2=Refer or try again};
> 2018-01-07-16:24:04Z;1;Smith;John;2013-06-22;Anonymous_2018_01_07_16_24_04;;Marrybeth Doe;0.2;0.3;0.4;3.5;-0.2;-0.1;-0.5;3.2;1.8;4;3;1
>
< HTTP/1.1 200 OK
< Date: Sun, 07 Jan 2018 23:23:48 GMT
< Content-Length: 12
< Content-Type: text/plain; charset=utf-8
<
< {"message":"Successfully uploaded 1 record"}
```


### Required headers

In the request a number of headers MUST be included:

#### Host

Hostname and port (if provided) used in the connection.

 *  Example: "localhost:9000", "upload.mydomain.com"

#### User-Agent

The name of the software and the version that made the request

 *  Example: "Panopticon/4.55.8"

#### Content-Type

The MIME content type being transmitted. This can be one of.

 *  Example: application/x-www-form-urlencoded, application/csv, text/csv

### TLS Connection requirements
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

### Error states

#### Cryptographic error states

In the event that any cryptographic precondition is not met the connection
should immediately be terminated.

#### Unexpected inputs

In the event that the server receives an unexpected input the server should
respond with a valid JSON [RFC7159](https://tools.ietf.org/html/rfc7159)
payload specifying that the error occurred, providing a numeric reference for
the error along with a message to be displayed to the user. The text error
message MUST be displayed to the end user.

The HTTP response MUST use standard HTTP error codes in addition to the JSON
payload that is provided in the error case.

Error structure:

```
{
  error: <string>,
  code: <int>,
  details: <Array<Object>>,
}
```

Example error output:

```
{
  "code": 1,
  "error": "Request canceled by server",
  details: []
}
```

#### Valid error reference codes

Error codes taken from [gRPC error
code](
https://github.com/grpc/grpc/blob/master/doc/statuscodes.md) list.

0
 : OK -- is returned on success.

1
 : Canceled -- indicates the operation was canceled (typically by the caller).

2
 : UnknownError -- An example of where this error may be returned is
   if a Status value received from another address space belongs to
   an error-space that is not known in this address space. Also
   errors raised by APIs that do not return enough error information
   may be converted to this error.

3
 : InvalidArgument -- indicates client specified an invalid argument.
   Note that this differs from FailedPrecondition. It indicates arguments
   that are problematic regardless of the state of the system
   (e.g., a malformed file name).

4
 : DeadlineExceeded -- means operation expired before completion.
   For operations that change the state of the system, this error may be
   returned even if the operation has completed successfully. For
   example, a successful response from a server could have been delayed
   long enough for the deadline to expire.

5
 : NotFound -- means some requested entity (e.g., file or directory) was
   not found.

6
 : AlreadyExists -- means an attempt to create an entity failed because one
   already exists.

7
 : PermissionDenied -- indicates the caller does not have permission to
   execute the specified operation. It must not be used for rejections
   caused by exhausting some resource (use ResourceExhausted
   instead for those errors). It must not be
   used if the caller cannot be identified (use Unauthenticated
   instead for those errors).

8
 : ResourceExhausted -- indicates some resource has been exhausted, perhaps
   a per-user quota, or perhaps the entire file system is out of space.

9
 : FailedPrecondition -- indicates operation was rejected because the
   system is not in a state required for the operation's execution.
   For example, directory to be deleted may be non-empty, an rmdir
   operation is applied to a non-directory, etc.

   A litmus test that may help a service implementor in deciding
   between FailedPrecondition, Aborted, and Unavailable:
    (a) Use Unavailable if the client can retry just the failing call.
    (b) Use Aborted if the client should retry at a higher-level
        (e.g., restarting a read-modify-write sequence).
    (c) Use FailedPrecondition if the client should not retry until
        the system state has been explicitly fixed. E.g., if an "rmdir"
        fails because the directory is non-empty, FailedPrecondition
        should be returned since the client should not retry unless
        they have first fixed up the directory by deleting files from it.
    (d) Use FailedPrecondition if the client performs conditional
        REST Get/Update/Delete on a resource and the resource on the
        server does not match the condition. E.g., conflicting
        read-modify-write on the same resource.

10
 : Aborted -- indicates the operation was aborted, typically due to a
   concurrency issue like sequencer check failures, transaction aborts,
   etc.

   See litmus test above for deciding between FailedPrecondition,
   Aborted, and Unavailable.

11
 : OutOfRange -- means operation was attempted past the valid range.
   E.g., seeking or reading past end of file.

   Unlike InvalidArgument, this error indicates a problem that may
   be fixed if the system state changes. For example, a 32-bit file
   system will generate InvalidArgument if asked to read at an
   offset that is not in the range [0,2^32-1], but it will generate
   OutOfRange if asked to read from an offset past the current
   file size.

   There is a fair bit of overlap between FailedPrecondition and
   OutOfRange. We recommend using OutOfRange (the more specific
   error) when it applies so that callers who are iterating through
   a space can easily look for an OutOfRange error to detect when
   they are done.

12
 : Unimplemented -- indicates operation is not implemented or not
   supported/enabled in this service.

13
 : Internal -- errors. Means some invariants expected by underlying
   system has been broken. If you see one of these errors,
   something is very broken.

14
 : Unavailable -- indicates the service is currently unavailable.
   This is a most likely a transient condition and may be corrected
   by retrying with a backoff.

   See litmus test above for deciding between FailedPrecondition,
   Aborted, and Unavailable.

15
 : DataLoss -- indicates unrecoverable data loss or corruption.

16
 : Unauthenticated -- indicates the request does not have valid
   authentication credentials for the operation.


