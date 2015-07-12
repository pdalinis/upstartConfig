# Upstart Configuration Generator
Create an upstart configuration file for a Go program.

For many services written in Go, there is almost no system configuration needed beyond setting up your program to restart if it crashes or the machine reboots.

Using this package, you can integrate that configuration step, removing the need for system configuration tools, or a more complex startup script.

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
var initialize = flag.Bool("init", false, "Initialize service and create the upstart configuration file.")
```

In your main function, remember to parse flags:
```
flag.Parse()
```

Then add logic to call this package if init is true:
```
if initialize {
  upstart.Write()
```

## CloudFormation Template Example
```
{
  "AWSTemplateFormatVersion" : "2010-09-09",
  "Description"              : "Example CFT",
  "Parameters"               : {
    "AppName" : {
      "Description" : "Name of the application",
      "Type": "String"
    },
  },
  "Resources" : {
    "ExampleEc2" : {
      "Type" : "AWS::EC2::Instance",
      "Properties" : {
      .......
        "UserData" : {
          "Fn::Base64" : {
            "Fn::Join" : [ "", [
              "#!/bin/bash\n",
              "aws s3 cp s3://myBucket/", {"Ref" : "AppName"}, " /opt\n",
              "chmod +x /opt/", {"Ref" : "AppName"}, "\n",
              "/opt/", {"Ref" : "AppName"}, " -init\n",
              "start ", {"Ref" : "AppName"}, "\n"
            ]]
          }
        }
      }
    }
  }
}
```

## License

This package is distributed under the
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0),
see LICENSE.txt and NOTICE.txt for more information.
