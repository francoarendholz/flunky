{% autoescape off %}
def jobName = '{{ jobName }}'

def job = Jenkins.instance.getItemByFullName(jobName)
if (job != null) {
    def configXml = job.getConfigFile().asString()
    println(configXml)
} else {
    println("Job not found: $jobName")
}
{% endautoescape %}