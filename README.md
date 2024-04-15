# Nội Dung Khóa Học Golang

## TOPIC 1: NGÔN NGỮ GOLANG – KEY FEATURE
Go effective, các quy ước để có source go đẹp và chuẩn. </br>
Go channel: Giao tiếp giữa các Goroutines (concurrent). </br>
Buffer Channel trong Golang. </br>
Cơ chế Defer, Recover trong Golang. </br>
Sử dụng Interface trong Golang. </br>
Slice, buffer, json decode / encode trong Golang. </br>

## TOPIC 2: PHÂN TÍCH DỰ ÁN TỪ GIAO DIỆN CHO TRƯỚC
Thiết lập requirement, user story từ giao diện. </br>
Phân tích chức năng, flow và các APIs cần có.</br>
Phân tích modules hoặc theo domain driven.</br>
Bonus: phân chia task hiệu quả để quản lý tiến độ tốt.</br>

## TOPIC 3: THIẾT LẬP DATABASE CHO DỰ ÁN FOOD DELIVERY
Cài đặt và kết nối database service: MySQL / PostgreSQL.</br>
Từ kết quả phân tích, thiết lập các bảng dữ liệu.</br>
Thiết lập các mối quan hệ giữa các bảng dữ liệu.</br>
Kỹ thuật đánh khoá chính và index để có kết quả truy xuất tốt nhất.</br>
Bonus: kinh nghiệm thiết kế database cho app Food Delivery.</br>

## TOPIC 4: VIẾT API (CƠ BẢN) TRONG GOLANG
Tìm hiểu REST API convention.</br>
Các API cơ bản: Create-Read-Update-Delete (CRUD) cơ bản.</br>
Các API CRUD trên nhiều bảng và transaction.</br>
Authen với JWT, cách sử dụng JWT để xác thực người dùng.</br>

## TOPIC 5: VIẾT API (MỞ RỘNG) TRONG GOLANG
Sử dụng middleware: tiền xử lý, xác thực quyền hạn, bắt lỗi crash.</br>
Upload files: xử lý, lưu trữ với các cloud storage (AWS S3) và CDN.</br>
Giao tiếp API giữa các module.</br>
Tổng hợp và link data các module.</br>
Bonus: Cách thiết kế giảm lệ thuộc giữa các module, tăng tốc xử lý, chống leak memory.</br>

## TOPIC 6: ASYNC HANDLERS, XỬ LÝ SIDE EFFECT TRONG GOLANG
Cách xây dựng async job trong Golang.</br>
Giải quyết timeout, retry cho async job.</br>
Đồng bộ dữ liệu với các async job.</br>
Pub / Sub trong Golang.</br>
Xây dựng async job queue & message broker.</br>

## TOPIC 7: TRIỂN KHAI (DEPLOY) & MONITORING
Log system trong Golang.</br>
Cơ chế tự động phục hồi kết nối DB (resilience).</br>
Cách sử dụng environment trong Golang.</br>
Build & Deploy với Docker.</br>
Sử dụng nginx (container Docker) làm reverse proxy.</br>
Bonus: Monitoring & Tracing.</br>

## TOPIC 8: SỬ DỤNG GRPC ĐỂ TĂNG TẢI SERVICE
Giới thiệu gRPC.</br>
Lập trình Protobuf 3.</br>
Tạo các service sử dụng gRPC cơ bản.</br>
gRPC streaming.</br>
Sử dụng gRPC Gateway để hỗ trợ thêm REST API.</br>
Bonus: Các kinh nghiệm xử lý gRPC trong thực tế.</br>

##TOPIC 9: MICROSERVICE CƠ BẢN (KHOÁ LIVESTREAM MỚI)
Hiểu rõ về Stateless service.</br>
Phân tách và deploy nhiều Microservices.</br>
Sử dụng Redis và NATs để tăng tải các services</br>.
Một số kinh nghiệm xử lý các vấn đề trong Microservice.</br>

## TOPIC 10: CÁC KỸ NĂNG KHÁC
Tư duy backend và hệ thống.</br>
Thuật toán & cấu trúc dữ liệu cơ bản.</br>
Xây dựng profile Github.</br>
Quản lý dự án, teamwork.</br>
Kiến trúc ứng dụng và các design pattern thường gặp trong Golang.</br>
Xây dựng CV để ứng tuyển vị trí Golang (có hỗ trợ review trainee, interview thử).</br>
