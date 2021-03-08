# App_api Service

This is the App_api service

Generated with

```
micro new app_api --namespace=easymall --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: easymall.web.app_api
- Type: web
- Alias: app_api

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./app_api-web
```

Build a docker image
```
make docker
```