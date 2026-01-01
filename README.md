# Chorister

**High-performance, scalable asynchronous task processing SDK for Go**

## Overview

In modern systems, asynchronous processing is essential for maximizing resource utilization, scaling to handle more clients, and delivering superior user experiences. Chorister provides a unified SDK for managing background task lifecycles in backend applications, enabling seamless task delegation to worker pools while decoupling execution from the main request flow.

## Key Features

- **Flexible Worker Pools**: Choose between `CreateWorkerPool()` for general workloads or `CreateWorkerMultiPool()` for extreme concurrency with reduced lock contention
- **Real-time Updates**: Broadcast task status, progress, and results via WebSocket, Kafka, RabbitMQ, or Redis Pub/Sub
- **Task Query APIs**: Monitor execution state, completion progress, and error details in real-time
- **Intelligent Retry**: Configurable retry strategies for failed tasks ensure system resilience
- **Full Observability**: Integrated OpenTelemetry for distributed tracing, metrics, and logs
- **Panic Recovery**: Built-in panic handler prevents service crashes

## Architecture

- **Dependency Injection**: Clean service registration and configuration management
- **Storage Abstraction**: `IStorage` interface supports multiple databases with methods like `SaveTask`, `UpdateTask`, `GetTaskFailed`, `GetTaskInProgress`
- **Transport Layer**: Pluggable `Transport` interface for message delivery across Redis Pub/Sub, Kafka, RabbitMQ, and more
- **Backward Compatibility**: `UnimplementedTransport` provides forward compatibility via struct embedding (gRPC-style)
- **Comprehensive Testing**: Full mock testing suite included

Chorister reduces backend complexity while enhancing scalability and observability in distributed environments.
