import system
import os
import osproc
import strutils

proc changeDirectory(args: seq) = 
  if args.len < 2:
    echo "Must provide a path"
  else:
    os.setCurrentDir(args[1])

proc executeCommand*(command: string) = 
  let args = split(command)

  if command == "exit":
    system.quit()
  elif args[0] == "cd":
    changeDirectory(args)
  else:
    let error = osproc.execCmd(command)

    if error != 0:
      echo error