# devsecops-reference
> Reference example for a DevSecOps Monorepo

## Getting Started

Each project is meant to be independent. There are shared libs but the repo is not a monolith. All aps can be ran independently and it is done mostly with native tools.

### Service A

```sh
cd apps/service-a
go run main.go
```

You may then query the service: `curl localhost:3000/a`

### Service B

```sh
cd apps/service-b
go run main.go
```

You may then query the service: `curl localhost:3000/b`
