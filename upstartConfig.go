// The upstartConfig package creates an upstart configuration file for the executing program.
//
// For many services written in Go, there is almost no system configuration needed beyond setting up your program to restart if it crashes or the machine reboots.
//
// Using this package, you can integrate that configuration step, removing the need for system configuration tools, or a more complex startup script.
package upstartConfig

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Location to write the upstart configuration file to.
var DestinationPath string = "/etc/init"

//Permisions for the upstart configuration file
var Perms string = "0666"

//Upstart configuration
type Config struct {
	ServiceName   string
	ServicePath   string
	StartRunLevel string
	StopRunLevel  string
	EnableEnv     bool
	Script        string
}

func BinaryToIntStr(in bool) string {
	if in {
		return "1"
	}
	return "0"
}

// Gets the name of the currently executing program. This is used for the name of the upstart configuration file, along with the script.
func GetName() string {
	pathParts := strings.Split(os.Args[0], "/")
	return pathParts[len(pathParts)-1]
}

// Gets the name and path of the currently executing program. This is used for determining what program upstart should manage.
func GetPath() string {
	pathParts := strings.Split(os.Args[0], "/")
	return strings.Join(pathParts[0:len(pathParts)-1], "/")
}

// Writes the upstart configuration file.
func Write(options ...func(*Config)) error {
	contents, config := Generate(options...)
	file := fmt.Sprintf("%s/%s.conf", DestinationPath, config.ServiceName)
	return ioutil.WriteFile(file, []byte(contents), 0666)
}

//Generates the upstart configuration file.
func Generate(options ...func(*Config)) (string, *Config) {
	path := GetPath()
	name := GetName()

	config := &Config{
		ServiceName:   name,
		ServicePath:   path,
		StartRunLevel: "2345",
		StopRunLevel:  "!2345",
		EnableEnv:     true,
		Script:        fmt.Sprintf("cd %s; exec ./%s", path, name),
	}

	//populate options into config, overriding defaults
	for _, o := range options {
		o(config)
	}

	lines := []string{
		fmt.Sprintf("description \"%s\"", config.ServiceName),
		fmt.Sprintf("start on runlevel [%s]", config.StartRunLevel),
		fmt.Sprintf("stop on runlevel [%s]", config.StopRunLevel),
		fmt.Sprintf("env enabled=%s", BinaryToIntStr(config.EnableEnv)),
		"",
		"respawn",
		"",
		"script",
		config.Script,
		"end script",
	}

	return strings.Join(lines, "\n"), config
}
