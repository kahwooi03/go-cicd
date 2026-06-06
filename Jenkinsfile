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
                sh 'go test ./...'
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
