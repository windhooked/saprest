# SAP REST Gateway Service

A simlple nano service for interfacing with SAP NW RFC
The services makes use of https://github.com/SAP/gorfc and echo

It encodes/decodes JSON from from and to map[string]interface{}

Currently with no security

### Env

To compile you should have these set. For env variables to connect look at sap.go

CGO_LDFLAGS_ALLOW=.*
CGO_CFLAGS_ALLOW=.*
CGO_CXXFLAGS_ALLOW=.*
SAPNWRFC_HOME=/home/phiadmin/sap/nwrfcsdk/
CGO_CFLAGS=-I /home/phiadmin/sap/nwrfcsdk/include
CGO_LDFLAGS=-L /home/phiadmin/sap/nwrfcsdk/lib

### Usage

curl -X POST -H "Content-Type: application/json" -d '{}' http://localhost:8088/rfc/RFC_PING

curl -X POST -H "Content-Type: application/json" -d '{}' http://localhost:8088/rfc/STFC_STRUCTURE
