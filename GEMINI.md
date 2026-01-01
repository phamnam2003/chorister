# Chorister Project Context

## Project Overview
Chorister is a Go library designed to provide a unified interface for asynchronous task processing and worker pool management. 
Currently, the codebase primarily defines the core interfaces (`CPool`, `CMultiPool`) and their generic counterparts (`CPoolWithFunc`, `CMultiPoolWithFunc`). It also includes "Unimplemented" reference implementations, likely intended for embedding to ensure forward compatibility or for use in mocking/testing scenarios.

The `README.md` describes a full-featured SDK with real-time updates, retry logic, and observability, but the current source files (`pool.go`, `multipool.go`) strictly contain the interface definitions and stubs. This suggests the project is either in an early stage, is a contract-only package, or relies on specific implementations to be injected.

## Key Files
- **`pool.go`**: Defines the `CPool` interface for standard worker pools and `CPoolWithFunc[T]` for generic task execution. Includes `UnimplementCPool` and `UnimplementCPoolWithFunc` structs.
- **`multipool.go`**: Defines the `CMultiPool` interface for multi-pool architectures and `CMultiPoolWithFunc[T]`. Includes `UnimplementCMultiPool` and corresponding stubs.
- **`go.mod`**: Defines the module `github.com/phamnam2003/chorister` and specifies Go version `1.25.5`.

## Architecture & Conventions
- **Interface-First Design**: The project emphasizes defining behaviors via interfaces rather than concrete implementations, promoting loose coupling.
- **Generics**: Utilizes Go 1.18+ generics for type-safe task handling (e.g., `CPoolWithFunc[T]`).
- **Forward Compatibility**: The pattern of providing `Unimplement...` structs (similar to gRPC/Protobuf in Go) allows new methods to be added to interfaces without breaking existing implementations that embed these stubs.
- **Dependency**: The code comments mention `github.com/phamnam2003/ants`, indicating the interfaces are likely modeled after or intended to wrap a fork of the `ants` worker pool library.

## Building & Testing
Since this is a library project:
- **Build**: Run `go build ./...` to verify the code compiles.
- **Test**: Run `go test ./...` (Note: No explicit test files are currently visible in the root directory).
