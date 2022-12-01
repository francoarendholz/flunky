{% autoescape off %}
import org.jenkinsci.plugins.workflow.job.WorkflowJob
import org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition
import org.jenkinsci.plugins.scriptsecurity.scripts.ScriptApproval
import org.jenkinsci.plugins.scriptsecurity.scripts.languages.GroovyLanguage

// https://plugins.jenkins.io/script-security/
// https://javadoc.jenkins.io/plugin/script-security/org/jenkinsci/plugins/scriptsecurity/scripts/ScriptApproval.html

def jobName = '{{ jobName|capfirst }}'

// The pipeline script to be run
script = '''
{{ pipelineScript }}
'''

// TODO: Find a better solution for this
// Have to revert the escaping which comes with the templating of Groovy scripts
// otherwise Jenkins has issues with ''' code ''' in script= above
// See: runPipelineScript.go
polishedScript = script.replaceAll("\\'","'")

// TODO: Some Error handling, currently bad scripts are just timing out.
// Pre-Approve the pipeline script
ScriptApproval.get().preapprove(polishedScript, GroovyLanguage.get())

// Create a temporary "Workflow Job" project to run the pipeline script in.
println("Creating temporary project: "+jobName)
project = Jenkins.instance.createProject(WorkflowJob, jobName)
project.setDefinition(new CpsFlowDefinition(polishedScript, false));
project.save()

println("Sheduling build job...")
Jenkins.instance.queue.schedule(project, 0, null, null)

println("Waiting for build job to start...")
while (project.builds.size() < 1) { }

println("Waiting for build job to finish...")
def execution = project.builds[0].getExecutionPromise().get()
while (!execution.isComplete()) { }

// Print log and delete project
println("Build console log:")
println project.builds[0].log
project.delete()
{% endautoescape %}
