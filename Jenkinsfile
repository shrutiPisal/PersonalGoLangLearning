pipeline {
    agent any
    tools {
    go 'GoTool'
    }

    stages {
        stage('Go Deps') {
            steps {
                sh 'go version'
                sh 'make deps'
            }
        }
        stage('Ginkgo Tests') {
                    steps {
                        sh 'make functional-tests'
                    }
                }
    }
}
