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
        sh "rsync -av main root@goapp.srwx.net:goapp"
        sh "ssh root@goapp.srwx.net -l root '(pkill -9 goapp || true) && nohup /root/main >/tmp/goapp.log'"
      }
    }
  }
}
