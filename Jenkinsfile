pipeline {
  stages {
    stage('Build') {
      steps {
        script {
          go build main.go
        }
      }
    }
    stage('Test') {
      steps {
        script {
          go test
        }
      }
    }
    stage('Deploy') {
      steps {
        sh "rsync -av main root@goapp.srwx.net:"
        sh "ssh root@goapp.srwx.net -l root '(pkill -9 goapp || true) && /root/main >/tmp/goapp.log 2>&1 &'"
      }
    }
  }
}
