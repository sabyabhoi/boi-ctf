package main

import (
	"log"
	"os/exec"
)

func StartSocat(port string, program string) {
	err := exec.Command("socat", "TCP4-LISTEN:"+port+",reuseaddr,fork", "EXEC:\""+program+"\"").Run()
	if err != nil {
		log.Fatalln(err)
	}
}
