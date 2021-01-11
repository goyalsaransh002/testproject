//______________________________________________________________________________
//
// Proprietary Software Product
//
// Hewlett-Packard Enterprise.
//
// © Copyright 2020 Hewlett Packard Enterprise Development LP
//     All Rights Reserved
//______________________________________________________________________________
//
// Content: CMS Jenkins Pipeline
//______________________________________________________________________________
//
// Purpose:
//
//    Execute CMS CI tasks. The full pipeline has the following stages
//      - Preparation : prepare build environment (Git, ...)
//      - Build       : execute build tasks for all projects
//      - UnitTests   : execute configured Unit Tests
//      - Analyze     : static analysis of source code
//      - Package     : generate packages (rpm, docker, ...) and sign them
//      - Deploy      : deploy artifacts to configured target systems
//      - SystemTests : launch remote tests on deployed systems
//      - Results     : gather all reports local and remote
//
//______________________________________________________________________________
//

@Library('cms5g-shared-library') _

pipeline {
    agent any

    environment {
        EMAIL_TO = 'goyalsaransh002@gmail.com'
        EMAIL_CC = ''
        PROJ_NAME = 'MP OBSERVABILITY'
        PROJ_SONAR_NAME = 'M5G:MP:OBS'
        IS_SNAPSHOT = readMavenPom().getVersion().endsWith("-SNAPSHOT")
    }

    options {
        buildDiscarder(logRotator(numToKeepStr: '10', artifactNumToKeepStr: '10'))
        timestamps()
    }

    tools {
        maven 'M3'
        jdk 'OpenJDK 11'
    }

    stages {
        // ------------------------------------------------------------
        //   Stage PREPARATION
        //
        //   Tasks
        //     - retrieve code from GitHub
        //     - set some environment variables
        // ------------------------------------------------------------
        stage('Checkout') {
            steps {
                echo "STAGE - PREPARATION"
                sh "git clean -fdx"
                checkout scm
            }
        }

        // ------------------------------------------------------------
        //   Stage BUILD
        //
        //   Tasks
        //     - run Maven build
        // ------------------------------------------------------------
        stage('Build') {
            options {
                timeout(time: 30, unit: 'MINUTES')
            }
            steps {
                echo "STAGE - BUILD"
                sh "'mvn' --batch-mode -U clean test-compile"
            }
        }

        // ------------------------------------------------------------
        //   Stage UNIT TESTS
        //
        //   Tasks
        //     - run Unit Tests through maven test target
        // ------------------------------------------------------------
        stage('UnitTests') {
            steps {
                echo "STAGE - UNIT TESTS"
                echo "STAGE - NO UNIT TESTS CURRENTLY"
                sh "'mvn' --batch-mode -Pkubernetes,-developer verify"
            }
        }

        // ------------------------------------------------------------
        //   Stage PRODUCTION
        //
        //   Tasks
        //     - run sub stages only on production branches (master)
        // ------------------------------------------------------------
        stage('Production') {
            when {
                beforeAgent true
                anyOf {
                    branch 'master'
                }
            }

            stages {
                // ------------------------------------------------------------
                //   Stage ANALYZE
                //
                //   Tasks
                //     - run static analyzers
                // ------------------------------------------------------------
                stage('Analyze') {
                    steps {
                        echo "STAGE - ANALYZE"
                        // OWASP Dependency Check
                        checkDeps()
                        // Sonar Analysis through sonar:sonar goal
                        doSonarMaven("$env.PROJ_SONAR_NAME")
                    }
                }

                // ------------------------------------------------------------
                //   Stage PACKAGE
                //
                //   Tasks
                //     - build rpms
                // ------------------------------------------------------------
                stage('Package') {
                    steps {
                        echo "STAGE - PACKAGE"
                    }
                }

                // ------------------------------------------------------------
                //   Stage DEPLOY
                //
                //   Tasks
                //     - archive java artifacts
                //     - archive rpms
                //     - publish rpms to target system
                // ------------------------------------------------------------
                stage('Deploy') {
                    steps {
                        echo "STAGE - DEPLOY"
                       // withCredentials([usernamePassword(credentialsId: 'nexus-public-user', passwordVariable: 'NEXUS_PASS', usernameVariable: 'NEXUS_USER')]) {
                          //  sh "'mvn' --batch-mode -DskipTests -P docker,kubernetes,sign deploy"
                        //}
                    }
                }

                // ------------------------------------------------------------
                //   Stage RESULTS
                //
                //   Tasks
                //     - collect static analyzers reports
                //     - collect tests reports
                //
                // ------------------------------------------------------------
                stage('Results') {
                    steps {
                        echo "STAGE - RESULTS"
                //        junit allowEmptyResults: true, testResults: '**/target/surefire-reports/TEST*.xml'
                //        publishHTML([allowMissing: false, alwaysLinkToLastBuild: false, keepAll: false, reportDir: 'user-doc/target/out', reportFiles: 'index.html', reportName: 'User doc', reportTitles: ''])
                    }
                }
            }
        }
    }
    post {
        always {
            sendEmail("$env.PROJ_NAME", "$env.EMAIL_TO", "$env.EMAIL_CC")
            script {
                echo "Build duration = " + currentBuild.durationString
            }
        }
        success {
            getBuildNumber()
        }
    }

}