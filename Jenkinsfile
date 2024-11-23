pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'phamminhthao/jenkins'
        DOCKER_TAG = 'loving'
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
                    docker.build("${DOCKER_IMAGE}:${DOCKER_TAG}")
                }
            }
        }

        stage('Run Tests') {
            steps {
                echo 'Running tests...'
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    docker.withRegistry('https://index.docker.io/v1/', 'docker-hub-credentials') {
                        docker.image("${DOCKER_IMAGE}:${DOCKER_TAG}").push()
                    }
                }
            }
        }

        stage('Deploy Golang to DEV') {
            steps {
                echo 'Deploying to DEV...'
                sh 'docker image pull trongpham99/golang-jenkins:latest'
                sh 'docker container stop golang-jenkins || echo "this container does not exist"'
                sh 'docker network create dev || echo "this network exists"'
                sh 'echo y | docker container prune '

                sh 'docker container run -d --rm --name server-golang -p 4000:3000 --network dev trongpham99/golang-jenkins:latest'
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }


    post {
        always {
            cleanWs()
        }

        success {
            sendTelegramMessage("✅ Build #${BUILD_NUMBER} was successful! ✅")
        }

        failure {
            sendTelegramMessage("❌ Build #${BUILD_NUMBER} failed. ❌")
        }
    }
}

def sendTelegramMessage(String message) {
    sh """
    curl -s -X POST https://api.telegram.org/bot${7046314210:AAGqYso31LSx8_kzYpIOcjryCMPqfiztxnE}/sendMessage \
    -d chat_id=${-1002415710063} \
    -d text="${message}"
    """
}
