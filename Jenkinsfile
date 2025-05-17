pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Go Test') {
            steps {
                sh 'go test .'
            }
        }

        stage('Docker Build') {
            steps {
                sh 'docker build .'
            }
        }
    }
}
