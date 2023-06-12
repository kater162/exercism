package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var welcomemsg, tablemsg, neighbormsg string
	welcomemsg = fmt.Sprintf("Welcome to my party, %s!\n", name)
	tablemsg = fmt.Sprintf("You have been assigned to table %03d. Your table is %v, exactly %.1f meters from here.\n", table, direction, distance)
	neighbormsg = fmt.Sprintf("You will be sitting next to %v.", neighbor)
	return welcomemsg + tablemsg + neighbormsg
}
