// pipeline {
//     agent any

//     environment {
//         DOCKER_IMAGE = 'phamminhthao/jenkins'
//         DOCKER_TAG = 'loving'
//     }

//     stages {
//         stage('Clone Repository') {
//             steps {
//                 git branch: 'main', url: 'https://github.com/Trunks-Pham/cicd-jenkins.git'
//             }
//         }

//         stage('Build Docker Image') {
//             steps {
//                 script {
//                     docker.build("${DOCKER_IMAGE}:${DOCKER_TAG}")
//                 }
//             }
//         }

//         stage('Run Tests') {
//             steps {
//                 echo 'Running tests...'
//             }
//         }

//         stage('Push to Docker Hub') {
//             steps {
//                 script {
//                     docker.withRegistry('https://index.docker.io/v1/', 'docker-hub-credentials') {
//                         docker.image("${DOCKER_IMAGE}:${DOCKER_TAG}").push()
//                     }
//                 }
//             }
//         }

//         stage('Deploy Golang to DEV') {
//             steps {
//                 echo 'Deploying to DEV...'
//                 sh 'docker image pull phamminhthao/jenkins:loving'
//                 sh 'docker container stop golang-jenkins || echo "this container does not exist"'
//                 sh 'docker network create dev || echo "this network exists"'
//                 sh 'echo y | docker container prune '

//                 sh 'docker container run -d --rm --name server-golang -p 4000:3000 --network dev phamminhthao/jenkins:loving'
//             }
//         }
//     }

//     post {
//         always {
//             cleanWs()
//         }

//         success {
//             sendTelegramMessage("✅ Build #${BUILD_NUMBER} was successful! ✅")
//         }

//         failure {
//             sendTelegramMessage("❌ Build #${BUILD_NUMBER} failed. ❌")
//         }
//     }
// }

// def sendTelegramMessage(String message = "") {
//     if (message.isEmpty()) {
//         error "Message cannot be empty"
//     }
//     def apiToken = "7046314210:AAGqYso31LSx8_kzYpIOcjryCMPqfiztxnE"
//     def chatId = "-1002415710063"
//     def curlCmd = "curl -s -X POST https://api.telegram.org/bot${apiToken}/sendMessage -d chat_id=${chatId} -d text=\"${message}\""
//     sh curlCmd
// }

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
                    try {
                        docker.build("${DOCKER_IMAGE}:${DOCKER_TAG}")
                    } catch (Exception e) {
                        error "Failed to build Docker image: ${e.message}"
                    }
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
                        try {
                            docker.image("${DOCKER_IMAGE}:${DOCKER_TAG}").push()
                        } catch (Exception e) {
                            error "Failed to push Docker image: ${e.message}"
                        }
                    }
                }
            }
        }

        stage('Deploy Golang to DEV') {
            steps {
                script {
                    try {
                        sh 'docker image pull phamminhthao/jenkins:loving'
                        sh 'docker container stop golang-jenkins || true'
                        sh 'docker network create dev || true'
                        sh 'echo y | docker container prune'
                        sh 'docker container run -d --rm --name server-golang -p 4000:3000 --network dev phamminhthao/jenkins:loving'
                    } catch (Exception e) {
                        error "Deployment failed: ${e.message}"
                    }
                }
            }
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

def sendTelegramMessage(String message = "") {
    if (message.isEmpty()) {
        error "Message cannot be empty"
    }
    def apiToken = "7046314210:AAGqYso31LSx8_kzYpIOcjryCMPqfiztxnE"
    def chatId = "-1002415710063"
    def curlCmd = "curl -s -X POST https://api.telegram.org/bot${apiToken}/sendMessage -d chat_id=${chatId} -d text=\"${message}\""
    sh curlCmd
}
