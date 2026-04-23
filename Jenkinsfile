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
            }
        }
        stage('Ginkgo Tests') {
                    steps {
                        sh 'ginkgo -v ./address-api-tests ./ginkgo-tests'
                    }
                }
    }
}
