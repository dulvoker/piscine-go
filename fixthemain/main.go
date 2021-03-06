package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...\n")
	ptrDoor.state = false
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...\n")
	ptrDoor.state = true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?\n")
	return ptrDoor.state
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	if ptrDoor.state == true {
		return false
	}
	return true
}

func main() {
	door := Door{}
	OpenDoor(&door)
	if IsDoorClose(&door) {
		OpenDoor(&door)
	}
	if IsDoorOpen(&door) {
		CloseDoor(&door)
	}
	if door.state == true {
		CloseDoor(&door)
	}
}
