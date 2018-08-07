pipeline {

  agent {
    label "jenkins-docker"
  }

  stages {

    stage('Build Release') {
      when {
        branch 'master'
      }
      steps {
        container('docker') {
          dir ('/home/jenkins/') {
            checkout scm
          }
        }
      }
    }

    stage('Promote to Production') {
      when {
        branch 'master'
      }
      steps {

      }
    }

  }

  post {
    always {
      cleanWs()
    }
  }
}
