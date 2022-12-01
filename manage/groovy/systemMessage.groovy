{% autoescape off %}
message='{{ message|capfirst }}'
Jenkins jenkins = Jenkins.get()

try {
    jenkins.setSystemMessage(message)
    jenkins.save()
} catch(Throwable err) {
    println('There was an error setting the new system message!\n'+err.toString)
    throw err
}

println('Successfully set new System Message!')
{% endautoescape %}
