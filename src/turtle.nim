import hooks  
import commands  
import config

proc main() = 
  config.loadConfig()
  
  while true:
    stdout.write "$:> "
    var command: string = readLine(stdin)

    commands.executeCommand(command)

  setControlCHook(hooks.ctrlcHook)

when isMainModule:
  main()