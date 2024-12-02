Sure, I can help with that! Here's a draft for your README file:

```markdown
# cicd-jenkins

This repository contains a sample project demonstrating the use of Jenkins for Continuous Integration and Continuous Deployment (CI/CD) with a Go application.

## Table of Contents

- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This project showcases how to set up a CI/CD pipeline using Jenkins for a Go application. The pipeline includes steps for building, testing, and deploying the application.

## Prerequisites

Before you begin, ensure you have the following installed:

- Docker
- Jenkins
- Go (version 1.16 or later)

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/Trunks-Pham/cicd-jenkins.git
    cd cicd-jenkins
    ```

2. Set up Jenkins:

    - Install Jenkins on your local machine or server.
    - Install the necessary Jenkins plugins (e.g., Docker, Git, Go).
    - Create a new Jenkins pipeline and configure it to use the `Jenkinsfile` in this repository.

## Usage

1. Build the Docker image:

    ```bash
    docker build -t cicd-jenkins .
    ```

2. Run the application:

    ```bash
    docker run -p 8080:8080 cicd-jenkins
    ```

3. Access the application at `http://localhost:8080`.

## Contributing

Contributions are welcome! Please fork this repository and submit a pull request with your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
```

Feel free to customize this README file to better suit your project's needs!
