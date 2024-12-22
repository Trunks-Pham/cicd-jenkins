### English Version  

# CICD Jenkins Pipeline for Golang Application  

## Introduction  

This project provides a Jenkins pipeline to automate the process of building, testing, and deploying a Golang application using Docker. The pipeline includes key steps such as source code cloning, building a Docker image, running tests, pushing the image to Docker Hub, and deploying the application to a development environment.  

## Golang Application  

The application is a simple API built using the [Gin](https://github.com/gin-gonic/gin) framework, offering four main endpoints:  

- **`GET /products`**: Retrieve the list of products.  
- **`POST /products`**: Add a new product.  
- **`PUT /products/{id}`**: Update an existing product by ID.  
- **`DELETE /products/{id}`**: Delete a product by ID.  

## Jenkins Pipeline  

The Jenkins pipeline consists of the following main steps:  

1. **Clone Repository**: Fetch the source code from GitHub.  
2. **Build Docker Image**: Build a Docker image for the Golang application.  
3. **Run Tests**: Execute test cases using `go test`.  
4. **Push to Docker Hub**: Push the Docker image to Docker Hub.  
5. **Deploy to DEV**: Deploy the application to the development environment using Docker.  

### Key Sections of the `Jenkinsfile`:  

- **Git Clone**: Clone the repository from GitHub.  
- **Build Docker Image**: Use Docker to build the application image.  
- **Test**: Run tests for the Go application.  
- **Push Image**: Push the Docker image to Docker Hub.  
- **Deploy**: Deploy the application to the development environment by running it in a Docker container.  

### Telegram Notifications  

Send notifications via Telegram when the build succeeds or fails, enabling developers to monitor the build status.  

## Dockerfile  

The Dockerfile outlines the steps to build a Docker image for the Golang application:  

- **Install Go**: Use the official Golang image.  
- **Install Dependencies**: Download and install dependencies from `go.mod`.  
- **Build**: Compile the Golang application into an executable file.  
- **Expose Port**: Open port 3000 for HTTP communication.  

## Application Run Instructions  

### Running the Application Locally  

1. **Install Go and Docker**: Install Go and Docker on your system.  
2. **Clone Repository**:  
   ```bash  
   git clone https://github.com/Trunks-Pham/cicd-jenkins.git  
   cd cicd-jenkins  
   ```  
3. **Run the Application**:  
   - **Without Docker**:  
     ```bash  
     go run main.go  
     ```  
   - **Using Docker**:  
     ```bash  
     docker build -t golang-app .  
     docker run -p 3000:3000 golang-app  
     ```  

### Testing Endpoints  

- **`GET /products`**: Retrieve the list of products.  
- **`POST /products`**: Add a new product.  
- **`PUT /products/{id}`**: Update an existing product by ID.  
- **`DELETE /products/{id}`**: Delete a product by ID.  

---

### Vietnamese Version  

# CICD Jenkins Pipeline cho Ứng dụng Golang  

## Giới thiệu  

Dự án này cung cấp một pipeline Jenkins để tự động hóa quá trình build, test và deploy một ứng dụng Golang sử dụng Docker. Pipeline bao gồm các bước chính như clone mã nguồn, xây dựng Docker image, chạy test, đẩy image lên Docker Hub và triển khai ứng dụng lên môi trường phát triển.  

## Ứng dụng Golang  

Ứng dụng là một API đơn giản sử dụng framework [Gin](https://github.com/gin-gonic/gin), cung cấp 4 endpoint chính:  

- **`GET /products`**: Lấy danh sách sản phẩm.  
- **`POST /products`**: Thêm sản phẩm mới.  
- **`PUT /products/{id}`**: Cập nhật sản phẩm hiện tại theo ID.  
- **`DELETE /products/{id}`**: Xóa sản phẩm theo ID.  

## Pipeline Jenkins  

Pipeline Jenkins bao gồm các bước chính:  

1. **Clone Repository**: Lấy mã nguồn từ GitHub.  
2. **Build Docker Image**: Xây dựng Docker image cho ứng dụng Golang.  
3. **Run Tests**: Chạy các bài kiểm tra bằng `go test`.  
4. **Push to Docker Hub**: Đẩy Docker image lên Docker Hub.  
5. **Deploy to DEV**: Deploy ứng dụng lên môi trường phát triển sử dụng Docker.  

### Các phần chính của `Jenkinsfile`:  

- **Git Clone**: Clone repository từ GitHub.  
- **Build Docker Image**: Sử dụng Docker để xây dựng image ứng dụng.  
- **Test**: Chạy kiểm tra cho ứng dụng Go.  
- **Push Image**: Đẩy Docker image lên Docker Hub.  
- **Deploy**: Deploy ứng dụng lên môi trường phát triển bằng cách chạy container Docker.  

### Thông báo Telegram  

Gửi thông báo qua Telegram khi build thành công hoặc thất bại, giúp các nhà phát triển theo dõi trạng thái build.  

## Dockerfile  

Dockerfile cung cấp các bước để xây dựng Docker image cho ứng dụng Golang:  

- **Cài đặt Go**: Sử dụng image Golang chính thức.  
- **Cài đặt Dependencies**: Tải xuống và cài đặt dependencies từ `go.mod`.  
- **Build**: Biên dịch ứng dụng Golang thành file thực thi.  
- **Expose Port**: Mở port 3000 để ứng dụng giao tiếp qua HTTP.  

## Hướng Dẫn Chạy Ứng Dụng  

### Chạy Ứng Dụng Cục Bộ  

1. **Cài Đặt Go và Docker**: Cài đặt Go và Docker trên máy.  
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
 
### Kiểm Tra Endpoint  

- **`GET /products`**: Lấy danh sách sản phẩm.  
- **`POST /products`**: Thêm sản phẩm mới.  
- **`PUT /products/{id}`**: Cập nhật sản phẩm theo ID.  
- **`DELETE /products/{id}`**: Xóa sản phẩm theo ID.  
 