// List all jobs on Jenkins in a tree-like view

import jenkins.model.*

def getAllJobsAndPipelines() {
    def items = Jenkins.instance.allItems
    def jobList = []

    items.each { item ->
        if (item instanceof hudson.model.Job) {
            jobList.add("${item.fullName}")
        }
    }
    return jobList
}

def jobList = getAllJobsAndPipelines()

def buildTree(list) {
    def tree = [:]
    list.each { item ->
        def current = tree
        item.split('/').each { part ->
            if (!current[part]) {
                current[part] = [:]
            }
            current = current[part]
        }
    }
    return tree
}

def printTree(tree, indent = 0) {
    tree.each { key, value ->
        println(' ' * indent + "|___" + key)
        if (value) {
            printTree(value, indent + 5)
        }
    }
}

def tree = buildTree(jobList)
printTree(tree)