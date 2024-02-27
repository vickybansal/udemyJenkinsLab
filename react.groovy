pipeline {
    agent any
    environment {
        NETLIFY_AUTH_TOKEN = credentials("netlify-personal-access-token")
        NETLIFY_SITE_ID = 'YOUR SITE ID'
    }    
    post {
        success {
            // add post actions here
            archiveArtifacts artifacts: 'build/**', fingerprint: true
        }
    }
    stages {
        stage('Checkout') {
            steps {
                echo 'Checkout'
                git branch: 'main', changelog: false, poll: false, url: 'https://github.com/Veverita-Engineering/Customer-Portal'
            }
        }
        stage('Build') {
            steps {
                nodejs(nodeJSInstallationName: 'node-lts') {
                    sh 'node --version'
                    sh 'npm --version'
                    sh 'npm install'
                    sh 'npm run build'
                }
            }
        }
        stage('Parallel Stage') {
            failFast true
            parallel {
                // parallel stages come here
                stage('Test') {
                   steps {
                        echo 'Test'
                    }
                }
                stage('SVN') {
                    steps {
                        snykSecurity(
                            snykInstallation: 'snyk@latest',
                            snykTokenId: 'snyk-api-token',
                        )
                    }   
                }
            }
        }
        stage('Deploy Stage') {
            steps {
                nodejs(nodeJSInstallationName: 'node-lts') {
                    sh 'node_modules/.bin/netlify deploy --dir=build'
                }
            }
        }
        stage('Deploy Prod') {
            steps {
                echo 'Deploy Prod'
                nodejs(nodeJSInstallationName: 'node-lts') {
                    sh 'node_modules/.bin/netlify deploy --dir=build --prod'
                }
            }
        }
    }
}