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
                    # Initialize module if needed
                    if [ ! -f "go.mod" ]; then
                        go mod init go-cicd
                    fi
                    
                    # Tidy dependencies
                    go mod tidy
                    
                    # Run tests with verbose output
                    go test -v ./... -timeout 30s
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
