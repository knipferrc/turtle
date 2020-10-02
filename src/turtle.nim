import system
import osproc
import os
import strutils

proc executeCommand(command: string) = 
  let args = split(command)

  if command == "exit":
    system.quit()
  elif args[0] == "cd":
    os.setCurrentDir(args[1])
  else:
    let error = osproc.execCmd(command)

    if error != 0:
      echo error
    

proc ctrlcHook() {.noconv.} = 
  system.quit();

proc main() = 
  while true:
    stdout.write "$:> "
    var command: string = readLine(stdin)

    executeCommand(command)

setControlCHook(ctrlcHook)
  
when isMainModule:
  main()