//error is there

package main

import (
	"flag"
	"fmt"
	"os/exec"
	"runtime"
)

func execute(command string, args_arr []string) {
	args := args_arr
	// we can store the output of this in our out variable
	// and catch any errors in err
	//args... means many values are there
	out, err := exec.Command(command, args...).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}
	// as the out variable defined above is of type []byte we need to convert
	// this to a string or else we will see garbage printed out in our console
	// this is how we convert it to a string

	output := string(out[:])
	fmt.Println(output)

}

func main() {
	//to get input from user and to get command line interface use flag module
	// iface := flag.String(name:"iface",value:"eth0",usage:"Interface for which you want to cahnge the MAC")
	// newMac := flag.String(name:"newMac",value:"",usage:"Provide the new MAC address")
	//below is in the above format
	iface := flag.String("iface", "eth0", "Interface for which you want to change the MAC")
	newMac := flag.String("newMac", "", "Provide the new MAC address")

	//flag will return the pointer so we have toparse this
	flag.Parse()

	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute("sudo", []string{"ifconfig", *iface, "down"})
		execute("sudo", []string{"ifconfig", *iface, "hw", "ether", *newMac})
		execute("sudo", []string{"ifconfig", *iface, "up"})

	}
}
