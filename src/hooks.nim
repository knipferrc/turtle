import system
    
proc ctrlcHook*() {.noconv.} = 
  system.quit(0);