pipeline{
    agent{
         node {
            label 'master'
            customWorkspace "workspace/${env.BRANCH_NAME}/src/github.com/aditya37/file-service/"
        }
    }
    environment {
        SERVICE  = "file-service"
        NOTIFDEPLOY = -522638644
    }
    options {
        buildDiscarder(logRotator(daysToKeepStr: env.BRANCH_NAME == 'main' ? '90' : '30'))
    }
    stages{
        stage("Checkout"){
            when {
                anyOf { branch 'main'; branch 'develop'; branch 'staging' }
            }
            // Do clone
            steps {
                echo 'Checking out from git'
                checkout scm
            }
        }
        // prepare get credential file
        stage('Prepare') {
            steps {
                // get credential file
                withCredentials([file(credentialsId: 'b067d4c8-d147-4732-8725-cb84c520759b', variable: 'sa')]) {
                    sh "cp $sa firebase-admin-key.json"
                    sh "chmod 644 firebase-admin-key.json"
                }
            }
        }
        stage('Build and deploy') {
            environment {
                GOPATH = "${env.JENKINS_HOME}/workspace/${env.BRANCH_NAME}"
                PATH = "${env.GOPATH}/bin:${env.PATH}"
            }
            stages {
                // build to dev
                stage('Deploy to env development') {
                    when {
                        branch 'develop'
                    }
                    environment {
                        NAMESPACE = 'core-development'
                    }
                    steps {
                        // get credential file
                        withCredentials([file(credentialsId: 'b067d4c8-d147-4732-8725-cb84c520759b', variable: 'sa')]) {
                            echo 'Build image'
                            sh "cp $sa firebase-admin-key.json"
                            sh "chmod 644 firebase-admin-key.json"
                            sh 'chmod +x build.sh'
                            sh './build.sh default'
                            sh 'rm firebase-admin-key.json'
                        }
                    }
                }
            }
        }
    }
    post{
        success{
            telegramSend(message:"Application $SERVICE has been [deployed]",chatId:"$NOTIFDEPLOY")
        }
        failure{
            telegramSend(message:"Application $SERVICE has been [Failed]",chatId:"$NOTIFDEPLOY")
        }
    }
}