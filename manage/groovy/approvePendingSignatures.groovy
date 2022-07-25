import org.jenkinsci.plugins.scriptsecurity.scripts.*

boolean force = {{ force }};

ScriptApproval sa = ScriptApproval.get();

println """
Signatures pending Approval:
-------------------------------------------------------------------------------
"""
for (ScriptApproval.PendingSignature pending : sa.getPendingSignatures()) {
        println "Pending approval : " + pending.signature
}
println """
-------------------------------------------------------------------------------
"""

sa = ScriptApproval.get();

for (item in sa.getPendingSignatures()) {
  if (org.jenkinsci.plugins.scriptsecurity.sandbox.whitelists.StaticWhitelist.isBlacklisted(item.signature)) {
    if (force == false){
      println ("[WARNING] Not approving " + item.signature + " because it is blacklisted")
      continue;
    }
  }
  sa.approveSignature(item.signature);
  println "Approved : " + item.signature
}

sa = ScriptApproval.get();
println "\nChecking for still pending approval signatures:\n"
if (sa.getPendingSignatures().size() > 0) {
  println """
  Signatures still pending Approval: (use -f / --force to force approval)
  -------------------------------------------------------------------------------
  """.stripIndent()
  for (ScriptApproval.PendingSignature pending : sa.getPendingSignatures()) {
    println "Still pending approval : " + pending.signature
  }
  println """
  -------------------------------------------------------------------------------
  """.stripIndent()
} else {
    println "No pending signatures!"
}