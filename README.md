# Upstart Configuration Generator
Create an upstart configuration file for the executing Go program.

## Usage
First, get this package:
```
go get github.com/pdalinis/upstartConfig
```

Open your main routine, and import this package:
```
import github.com/pdalinis/upstartConfig
```

Add a cli option to your program:
```
var initialize = flag.Bool("init", false, "Intialize service and create the upstart configuration file.")
```

In your main function, remember to parse flags:
```
flag.Parse()
```

Then add logic to call this package if initialize is true:
```
if initialize {
  upstart.Write()
```

## License

This SDK is distributed under the
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0),
see LICENSE.txt and NOTICE.txt for more information.

