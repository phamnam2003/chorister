# Copilot Instructions cho Chorister

## Ngôn ngữ sử dụng

- **Phản hồi, giải thích**: Sử dụng tiếng Việt
- **Code và Comment**: Sử dụng tiếng Anh
- **Viết tắt**: Được phép sử dụng nhưng phải giải thích đầy đủ ý nghĩa và mục đích

## Chuẩn viết code Go

- Tuân thủ các quy ước và idioms chính thống của Go
- Sử dụng chuẩn định dạng `gofmt` (Go formatter - công cụ tự động format code)
- Áp dụng các pattern hiệu quả của Go
- Đặt tên biến và hàm có ý nghĩa, dễ hiểu theo phong cách Go
- Giữ các hàm ngắn gọn, tập trung vào một trách nhiệm duy nhất (SRP - Single Responsibility Principle)

## Cấu trúc Package

- Tổ chức code theo standard Go project layout (bố cục dự án chuẩn)
- Đặt tên package rõ ràng (lowercase, tốt nhất là một từ duy nhất)
- Chỉ export những types và functions cần thiết (viết hoa chữ cái đầu cho public API)
- Giữ dependencies giữa các package tối thiểu và không tuần hoàn (acyclic - tránh circular dependency)
- Document đầy đủ cho tất cả exported functions, types và packages

## Dependency Injection (DI)

- Sử dụng constructor functions để khởi tạo service (pattern `NewXxx()`)
- Truyền dependencies một cách tường minh qua parameters hoặc constructors
- Tránh global state và singleton patterns (tránh biến toàn cục và pattern đơn thể)
- Implement DI cho tất cả services để cải thiện khả năng test
- Đăng ký services và configuration parameters sử dụng DI patterns

## Yêu cầu kiến trúc

### Storage Layer (Lớp lưu trữ)
- Implement interface `IStorage` để trừu tượng hóa database
- Các methods bắt buộc: `SaveTask`, `UpdateTask`, `DeleteTask`, `GetOneTask`, `GetTaskFailed`, `GetTaskInProgress`
- Hỗ trợ nhiều loại database backend thông qua interface implementation

### Transport Layer (Lớp truyền tải)
- Implement interface `Transport` cho message delivery (gửi/nhận tin nhắn)
- Hỗ trợ Redis Pub/Sub, Kafka, RabbitMQ và các messaging services khác
- Cung cấp `UnimplementedTransport` base type để đảm bảo backward compatibility (tương thích ngược - giống như gRPC code generation)
- Cho phép struct embedding để tự động thỏa mãn interface requirements

### Worker Pool Management (Quản lý nhóm workers)
- Implement `CreateWorkerPool()` cho workloads tiêu chuẩn (từ nhỏ đến tương đối lớn)
- Implement `CreateWorkerMultiPool()` cho high-concurrency scenarios (giảm thiểu lock contention - tranh chấp khóa với các goroutine)
- Xử lý graceful shutdown và resource cleanup (đóng dịch vụ an toàn, dọn dẹp tài nguyên)

### Error Handling & Resilience (Xử lý lỗi và khả năng phục hồi)
- Implement panic recovery handlers để tránh service crashes (chặn panic ngăn ứng dụng bị crash)
- Thêm cơ chế retry có thể cấu hình cho failed tasks (thử lại tác vụ thất bại)
- Log errors với đầy đủ context (ngữ cảnh lỗi)

### Observability (Khả năng quan sát)
- Tích hợp OpenTelemetry cho distributed tracing (theo dõi luồng xử lý trong hệ thống phân tán)
- Implement metrics collection cho task execution (thu thập số liệu thực thi task)
- Cung cấp structured logging trong toàn hệ thống
- Đảm bảo đủ 3 trụ cột: traces (dấu vết), metrics (số liệu), logs (nhật ký)

### Testing (Kiểm thử)
- Viết unit tests toàn diện với table-driven test patterns (pattern test theo bảng dữ liệu)
- Sử dụng mock implementations cho interfaces (triển khai giả lập để test)
- Cung cấp testing utilities và mock factories
- Hướng tới test coverage cao trên các critical paths (độ bao phủ test cao ở các luồng quan trọng)

## Tổ chức code

- Giữ interfaces nhỏ gọn và tập trung (Interface Segregation Principle)
- Sử dụng `context.Context` cho cancellation và timeouts (hủy bỏ và giới hạn thời gian)
- Xử lý errors một cách tường minh, không bao giờ bỏ qua
- Sử dụng channels và goroutines theo đúng idioms của Go
- Implement proper resource cleanup với defer (dọn dẹp tài nguyên đúng cách)

## Interface & Duck Typing

- Tận dụng tối đa duck typing của Go (implicit interface implementation - không cần khai báo tường minh)
- Định nghĩa interfaces nhỏ, tập trung vào hành vi cụ thể (behavior-focused)
- Ưu tiên "accept interfaces, return structs" pattern (nhận đầu vào là interface, trả về struct)
- Sử dụng interface để tách biệt dependencies và tăng khả năng test
- Thiết kế interfaces dựa trên những gì consumer cần, không phải producer cung cấp
- Tận dụng struct embedding để compose behaviors (kết hợp các hành vi)
- Áp dụng interface segregation: nhiều interfaces nhỏ tốt hơn một interface lớn
