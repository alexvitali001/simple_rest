pipeline {
    agent any
    tools {
        go 'go1.14'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
    }
    stages {
        stage('Checkout') {
            steps {
                echo 'Running checkout steps'
                git branch: 'master', credentialsId: 'github_token', url: 'https://github.com/wilsontwm/go_simple_rest'
            }
        }
        
        stage('Pre Test') {
            steps {
                echo 'Running pre test steps'
                sh 'go version'
            }
        }
        
        stage('Test') {
            steps {
                echo 'Running test steps'
                sh 'cd test && go test -v'
            }
        }
        
        stage('Build') {
            steps {
                echo 'Running build steps'
                sh 'go build'
            }
        }
    }
}