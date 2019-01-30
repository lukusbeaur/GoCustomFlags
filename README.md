# GoCustomFlags
Interfaced the flag.Value and made custom flags, sub-commands for each flag. 



Two different commands
systemCommands is called with the System flag [ex. $System -dir=true]
  The only sub-command is -dir = which is a bool. True means it will return your dir, false(default) returns $root
  
  
stringCommands is called with the String flag [ex. $String -far=Argument the rest of the arguments.]
    The only sub command is -far = which pulls the first arg after the flag and echos it.
    the "tail" is everything after the first arg
    
    
  
