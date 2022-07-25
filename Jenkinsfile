#!/usr/bin/env groovy

branch = ""
if (env.BRANCH_NAME != null) {
    branch = env.BRANCH_NAME.toLowerCase()
}

try {
    if (branch.startsWith('pr-')) {

        stage('Checkout') {
            checkout()
        }
      
        stage('UnitTests') {
            unitTests()
        }

        stage('Build') {
            prBuild()
        }
      
    } else if ( branch.equalsIgnoreCase('master') || branch.equalsIgnoreCase('dev') || branch.contains(('pipeline-dev'))) {
/**
*  Development Build
*/
        stage ('Checkout'){
            checkout()
        }

        stage('UnitTests') {
            unitTests()
        }

        stage('Security Scans') {
            securityScans()
        }
      
        stage ('Build'){
            devBuild()
        }
      
    } else if ( branch.contains(('release'))){
/**
*  Release Build
*/
        stage ('Checkout'){
            checkout()
        }
      
        stage('UnitTests') {
            unitTests()
        }
      
        stage('Security Scans') {
            securityScans()
        }

        stage('Build - Staging') {
            lock(resource: "${env.JOB_NAME}/10", inversePrecedence: true) {
                milestone 10
                node {
                  releaseBuild("stage")
                }
            }
        }
      
        stage('Integration Tests') {
            securityScans()
        }

        stage('Build - Promote') {
          
            input message: 'Shall we promote staged artifact to release?'
          
            lock(resource: "${env.JOB_NAME}/20", inversePrecedence: true) {
                milestone 20
                node {
                  releaseBuild("promote")
                }
            }
        }
    }
    } catch (Throwable err) { // catch all exceptions
  
        throw err
  
    } finally {
  
      // TODO
  
    }

def checkout() {
    node {
        deleteDir()
        checkout scm
    }
}

def unitTests() {
  // Perform unit tests
}

def securityScans() {
    parallel(
        'ScanTypeA': {
            // Perform scan type A
        },
        'ScanTypeB': {
            // Perform scan type B
        }
    )
}

def prBuild() {
    node{
      // Perform pr build
    }
}

def devBuild() {
    node{
      sh """
        make build-local
      """
    }
}

def releaseBuild(type) {
    node{
      // Perform development build
    }
}