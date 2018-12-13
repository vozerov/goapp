pipeline {
  agent any

  stages {
    stage('Build') {
      steps {
        sh "go version"
        sh "go build main.go"
      }
    }
    stage('Test') {
      steps {
        sh "go test"
      }
    }
  }
}
