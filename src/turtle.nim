import hooks  
import commands  

proc main() = 
  while true:
    stdout.write "$:> "
    var command: string = readLine(stdin)

    commands.executeCommand(command)

  setControlCHook(hooks.ctrlcHook)

when isMainModule:
  main()