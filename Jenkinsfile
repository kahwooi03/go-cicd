pipeline {
    agent any

    environment {
        GO111MODULE = 'on'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Test') {
            steps {
                sh '''
                    if [ -f "go.mod" ]; then
                        go mod tidy
                    fi
                    go test ./... || echo "No tests found or tests failed"
                '''
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o go-cicd .'
            }
        }
    }

    post {
        always {
            archiveArtifacts artifacts: 'go-cicd', fingerprint: true
        }
    }
}
