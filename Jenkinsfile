pipeline {
    agent any
    tools {
    go 'GoTool'
    }

    stages {
        stage('Go Deps') {
            steps {
                sh 'go version'
                sh 'go mod download'
                sh 'go mod tidy'
                sh 'go install github.com/onsi/ginkgo/v2/ginkgo'
            }
        }
        stage('Ginkgo Tests') {
                    steps {
                        sh 'ginkgo -v ./address-api-tests ./ginkgo-tests'
                    }
                }
    }
}
