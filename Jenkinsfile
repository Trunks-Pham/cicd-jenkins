pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'phamminhthao/jenkins'
        DOCKER_TAG = 'main'
    }

    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'main', url: 'https://github.com/Trunks-Pham/cicd-jenkins.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // Không sử dụng sudo, đảm bảo user Jenkins có quyền chạy Docker
                    sh 'docker build -t ${DOCKER_IMAGE}:${DOCKER_TAG} .'
                }
            }
        }

        stage('Run Tests') {
            steps {
                script {
                    // Chạy các bài kiểm tra Go
                    sh 'go test - run Main ./server'
                }
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    // Đăng nhập và đẩy image lên Docker Hub
                    docker.withRegistry('https://index.docker.io/v1/', 'docker-hub-credentials') {
                        docker.image("${DOCKER_IMAGE}:${DOCKER_TAG}").push()
                    }
                }
            }
        }

        stage('Deploy Golang to DEV') {
            steps {
                echo 'Deploying to DEV...'
                // Kéo image từ Docker Hub
                sh 'docker image pull ${DOCKER_IMAGE}:${DOCKER_TAG}'

                // Dừng container cũ nếu tồn tại
                sh 'docker container stop golang-jenkins || echo "this container does not exist"'

                // Tạo network nếu chưa có
                sh 'docker network create dev || echo "this network exists"'

                // Xóa container không dùng
                sh 'echo y | docker container prune'

                // Chạy container mới
                sh 'docker container run -d --rm --name server-golang -p 4000:3000 --network dev ${DOCKER_IMAGE}:${DOCKER_TAG}'
            }
        }
    }

    post {
        always {
            // Dọn dẹp workspace
            cleanWs()
        }

        success {
            // Gửi thông báo Telegram khi build thành công
            sendTelegramMessage("✅ Build #${BUILD_NUMBER} was successful! ✅")
        }

        failure {
            // Gửi thông báo Telegram khi build thất bại
            sendTelegramMessage("❌ Build #${BUILD_NUMBER} failed. ❌")
        }
    }
}

// Hàm gửi thông báo Telegram
def sendTelegramMessage(String message = "") {
    if (message.isEmpty()) {
        error "Message cannot be empty"
    }
    def apiToken = "7046314210:AAGqYso31LSx8_kzYpIOcjryCMPqfiztxnE"
    def chatId = "-1002415710063"
    def curlCmd = "curl -s -X POST https://api.telegram.org/bot${apiToken}/sendMessage -d chat_id=${chatId} -d text=\"${message}\""
    sh curlCmd
}
