setlocal

set fnName=%1

@rem Creating function...
fission fn create --name %fnName% --env nodejs --code hello.js

@rem Test function...
fission fn test --name %fnName%

@rem Deleting function...
fission fn delete --name %fnName%


