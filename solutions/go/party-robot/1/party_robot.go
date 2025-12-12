package partyrobot
import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, "+name+ "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
    var nameString = fmt.Sprintf("Happy birthday %s!", name)
    var ageString = fmt.Sprintf("You are now %d years old!", age)
	return nameString + " " + ageString
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
    var welcomeString = Welcome(name)
    var tableString = fmt.Sprintf("You have been assigned to table %.3d. Your table is %s, exactly %.1f meters from here.", table, direction, distance)
    var neighborString = fmt.Sprintf("You will be sitting next to %s.", neighbor)
return welcomeString +"\n" + tableString + "\n" + neighborString
}
