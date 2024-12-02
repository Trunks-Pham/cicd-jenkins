# CICD Jenkins Pipeline for Golang Application

## Giới thiệu

Dự án này cung cấp một pipeline Jenkins tự động hóa quá trình build, test và deploy một ứng dụng Golang sử dụng Docker. Pipeline này có các bước chính như clone mã nguồn, build Docker image, chạy test, push image lên Docker Hub, và deploy ứng dụng lên môi trường phát triển.

## Ứng dụng Golang

Ứng dụng là một API đơn giản sử dụng framework [Gin](https://github.com/gin-gonic/gin), cung cấp 2 endpoint chính:

- **`GET /products`**: Lấy danh sách sản phẩm.
- **`POST /products`**: Thêm sản phẩm mới.

## Jenkins Pipeline

Pipeline Jenkins có các bước chính sau:

1. **Clone Repository**: Lấy mã nguồn từ GitHub.
2. **Build Docker Image**: Xây dựng Docker image cho ứng dụng Golang.
3. **Run Tests**: Chạy các bài kiểm tra với `go test`.
4. **Push to Docker Hub**: Đẩy Docker image lên Docker Hub.
5. **Deploy to DEV**: Deploy ứng dụng lên môi trường phát triển sử dụng Docker.

### Các phần chính của `Jenkinsfile`:

- **Git Clone**: Clone repository từ GitHub.
- **Build Docker Image**: Sử dụng Docker để xây dựng image ứng dụng.
- **Test**: Chạy kiểm tra Go.
- **Push Image**: Đẩy Docker image lên Docker Hub.
- **Deploy**: Deploy ứng dụng lên môi trường phát triển, sử dụng Docker để chạy container.

### Phần gửi thông báo Telegram:

Gửi thông báo qua Telegram khi build thành công hoặc thất bại, giúp người phát triển theo dõi trạng thái build.

## Dockerfile

Dockerfile cung cấp các bước để xây dựng Docker image cho ứng dụng Golang:

- **Cài đặt Go**: Sử dụng image Golang chính thức.
- **Cài đặt dependencies**: Tải xuống và cài đặt các dependencies từ `go.mod`.
- **Build**: Biên dịch ứng dụng Golang thành file thực thi.
- **Expose Port**: Mở port 3000 để ứng dụng có thể giao tiếp qua HTTP.

## Hướng Dẫn Chạy Ứng Dụng

### Chạy Ứng Dụng Cục Bộ

1. **Cài Đặt Go và Docker**: Cài đặt Go và Docker trên máy của bạn.
2. **Clone Repository**:
   ```bash
   git clone https://github.com/Trunks-Pham/cicd-jenkins.git
   cd cicd-jenkins
   ```
3. **Chạy Ứng Dụng**:
   - **Không sử dụng Docker**:
     ```bash
     go run main.go
     ```
   - **Sử dụng Docker**:
     ```bash
     docker build -t golang-app .
     docker run -p 3000:3000 golang-app
     ```

### Kiểm Tra Các Endpoint

- **`GET /products`**: Lấy danh sách các sản phẩm.
- **`POST /products`**: Thêm một sản phẩm mới.
