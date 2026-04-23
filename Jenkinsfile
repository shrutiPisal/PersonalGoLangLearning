pipeline {
    agent any
    tools {
    go 'Gotool'
    }

    stages {
        stage('Go Deps') {
            steps {
                sh 'go --version'
            }
        }
        stage('Ginkgo Tests') {
                    steps {
                        echo 'Hello from Jenkins World'
                    }
                }
    }
}
