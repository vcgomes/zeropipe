package main

import (
	"fmt"
	"github.com/godbus/dbus"
	"os"
)

const (
	avahiServiceName = "org.freedesktop.Avahi"
	serverInterface = avahiServiceName + ".Server"
	entryGroupInterface = avahiServiceName + ".EntryGroup"
	serviceBrowserInterface = avahiServiceName + ".ServiceBrowser"
)

func registerLocalService(conn *dbus.Conn, name, stype string, port int, txt string) {
	server := conn.Object(avahiServiceName, "/")

	var entryPath dbus.ObjectPath
	err := server.Call(serverInterface + ".EntryGroupNew", 0).Store(&entryPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to call `EntryGroupNew`:", err)
		os.Exit(1)
	}

	fmt.Printf("Entry %s\n", entryPath)
}

func main() {
	conn, err := dbus.SystemBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to system bus:", err)
		os.Exit(1)
	}

	registerLocalService(conn, "", "", 0, "")
}
