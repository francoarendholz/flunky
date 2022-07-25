message='{{ message|capfirst }}'
Jenkins jenkins = Jenkins.get()
jenkins.setSystemMessage(message)
jenkins.save()