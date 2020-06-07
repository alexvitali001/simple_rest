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
                git branch: 'master', credentialsId: '979d477d-fb5f-432b-a75b-f9916559e105', url: 'https://github.com/wilsontwm/go_simple_rest'
            }
        }
        
        stage('Pre Test') {
            steps {
                echo 'Running pre test steps'
                sh 'go version'
                sh 'go get -u github.com/golang/lint/golint'
            }
        }
        
        stage('Test') {
            steps {
                echo 'Running test steps'
                echo 'Running vetting...'
                sh 'go vet .'
                echo 'Running linting...'
                sh 'golint .'
                echo 'Running test...'
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