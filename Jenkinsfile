pipeline {
  agent {
    label any
  }
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
        script {
          rsync -av main root@goapp.srwx.net:
          ssh root@goapp.srwx.net -l root '(pkill -9 goapp || true) && /root/main >/tmp/goapp.log 2>&1 &'
        }
      }
    }
  }
}
