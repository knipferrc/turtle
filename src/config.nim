import parsecfg

proc loadConfig*() = 
  var config = parsecfg.loadConfig("turtle.ini")
  var name = config.getSectionValue("Author", "name")
  echo name