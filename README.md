# go-grpc-boilerplate

#### gRPC microservices with go make easy!

## Prequisites:

#### Options:
1. Manual Hack
- [protobuf](https://developers.google.com/protocol-buffers)
- [CockroachDB](https://www.cockroachlabs.com/)

2. Docker:
- [Docker](https://www.docker.com)
- [docker-compose](https://docs.docker.com/compose)


## Development

### Docker

- Setup .env
```bash
   cp .env.example .env
```
Edit some ENV to satisfy with your setting

- Start docker-compose
```bash
    docker-compose up
```

- After starting docker, you just need to init database name to make server work properly!
```bash
    chmod +x hack/*.sh

    ./hack/crdb-insecure-init-db.sh
```

### Manual Hack

#### Before starting:
- Make sure CockroachDB is started and database initialize correctly
- After init database, use script in `hack` folder which satisfy your settings:
```bash
    chmod +x hack/*.sh

    ./hack/crdb-insecure-init-sb.sh
```

#### Start

- Set up .env
```bash
    cp .env.example .env
```
Edit some ENV to satisfy your settings

- Install dependencies and tools:
```bash
    make install
```

- Generate pb from .proto files:
```bash
    make gen
```

- Start gRPC server:
```bash
   make server
```

- Start gRPC with TLS enable:
```bash
    make server-tls
```

- Start gRPC client (for test):
```bash
    make client
```

- Start gRPC client with TLS enable (for test)
```bash
    make client-tls
```

Please read `Makefile` to get more commands

### Just wanna to get stated with go
- Install go packages:
```bash
    go mod download
```

- Run CLI
```bash
    go run main.go -h
```

- On development mode, you can use [gow](https://github.com/mitranim/gow) (go watch) for hot reload
```bash
    gow run . server
```

# Author
[Nguyen Trung Tin](https://tintrungnguyen.com)

# License
MIT

# Happy Coding! 🚀
