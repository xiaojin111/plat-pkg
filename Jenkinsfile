pipeline{
    agent any

    environment {
        GOROOT="/var/jenkins_home/tools/org.jenkinsci.plugins.golang.GolangInstallation/Go1.13"
        GOPATH="/var/go"
        GOBIN="${env.GOPATH}/bin"
        PATH="${env.GOBIN}:${env.GOROOT}/bin:${env.PATH}"
        GOPROXY="https://goproxy.io,direct"
        GOPRIVATE="github.com/jinmukeji/*"
        AWS_ACCESS_KEY_ID     = credentials('jenkins-aws-secret-key-id')
        AWS_SECRET_ACCESS_KEY = credentials('jenkins-aws-secret-access-key')
        AWS_DEFAULT_REGION    = credentials('jenkins-aws-secret-region')
        PROJECT = "$WORKSPACE"
        DOCKER_HOST_IP ="${DOCKER_HOST_IP}"
        BUILD_NUMBER ="${env.BUILD_NUMBER}"
    }

    stages {    
        // stage('代码拉取'){
        //     steps {
        //         checkout([$class: 'GitSCM', branches: [[name: '${sha1}']], browser: [$class: 'GithubWeb', repoUrl: 'https://github.com/jinmukeji/plat-pkg'], doGenerateSubmoduleConfigurations: false, extensions: [[$class: 'RelativeTargetDirectory', relativeTargetDir: 'plat-pkg']], submoduleCfg: [], userRemoteConfigs: [[credentialsId: '90bdcbdb-808a-40b6-b5db-efe3de9e317d',refspec: '+refs/pull/*:refs/remotes/origin/pr/*', url: 'https://github.com/jinmukeji/plat-pkg.git']]])
        //     }
        // }

        stage('初始化环境'){
            steps {
                sh label: '', script: '$PROJECT/jenkins_ci/ci_init.sh'
             }
        }
        stage('代码格式化'){
            steps{
                sh label: '', script: '$PROJECT/jenkins_ci/ci_format.sh'
            }
        }
        stage('静态代码检查'){
            steps{
                sh label: '', script: '$PROJECT/jenkins_ci/ci_lint.sh'
            }
        }
        stage('编译'){
            steps{
                sh label: '', script: '$PROJECT/jenkins_ci/ci_build.sh'
            }
        }
        stage('单元测试'){
            steps{
                sh label: '', script: '$PROJECT/jenkins_ci/ci_unittest.sh'
            }
        }
     }

    post {
        success {
            emailext (
              subject: "SUCCESSFUL: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'",
              to: "cuimin@jinmuhealth.com",
              body: """SUCCESSFUL: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]':</p>
    <p>Check console output at &QUOT;<a href='${env.BUILD_URL}'>${env.JOB_NAME} [${env.BUILD_NUMBER}]</a>&QUOT;</p>""",
              recipientProviders: [[$class: 'DevelopersRecipientProvider']]
            )
        }
        failure {
            emailext (
              subject: "FAILED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'",
              to: "cuimin@jinmuhealth.com",
              body: """<p>FAILED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]':</p>
                <p>Check console output at &QUOT;<a href='${env.BUILD_URL}'>${env.JOB_NAME} [${env.BUILD_NUMBER}]</a>&QUOT;</p>""",
              recipientProviders: [[$class: 'DevelopersRecipientProvider']]
            )
        }
     }
}
