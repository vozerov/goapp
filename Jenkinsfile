pipeline {
  agent any

  stages {
    stage('Build') {
      steps {
        sh "go build main.go"
      }
    }
    stage('Test') {
      steps {
        sh "go test"
      }
    }
    stage('Deploy') {
      steps {
        sh "rsync -av main root@goapp.srwx.net:/usr/local/bin/goapp"
        sh "ssh root@goapp.srwx.net -l root '/etc/init.d/goapp stop'"
        sh "ssh root@goapp.srwx.net -l root '/etc/init.d/goapp start'"
      }
    }
  }
}
