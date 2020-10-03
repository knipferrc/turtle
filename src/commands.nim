import system
import os
import osproc
import strutils

proc executeCommand*(command: string) = 
  let args = split(command)

  if command == "exit":
    system.quit()
  elif args[0] == "cd":
    os.setCurrentDir(args[1])
  else:
    let error = osproc.execCmd(command)

    if error != 0:
      echo error