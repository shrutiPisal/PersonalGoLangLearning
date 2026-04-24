pipeline {
    agent any
    tools {
    go 'GoTool'
    }

    environment {
                GOPATH = "${WORKSPACE}/go"
                // Explicitly define GOBIN
                GOBIN  = "${WORKSPACE}/go/bin"
                // Add GOBIN to the system PATH
                PATH   = "${GOBIN}:${env.PATH}"
            }

    stages {
        stage('Go Deps') {
            steps {
                sh 'go version'
                sh 'go mod download'
                sh 'go mod tidy'
                sh 'go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo'
            }
        }
        stage('Ginkgo Tests') {
                    steps {
                        sh 'ginkgo -v ./address-api-tests ./ginkgo-tests'
                    }
                }
    }
}
