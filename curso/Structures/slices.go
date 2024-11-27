package Structures

import "fmt"

func slice() {
	weekDays := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	from := 4
	// The "to" index is exclusive
	to := 7
	lastThreeDays := weekDays[from:to]

	fmt.Println(lastThreeDays)

	// Capacidad
	fmt.Println(cap(lastThreeDays))
	// Logitud
	fmt.Println(len(lastThreeDays))

	// Add element
	lastThreeDays = append(lastThreeDays, "New day")

	// From 0 to 2
	firstTwoDays := weekDays[:2]

	fmt.Println(firstTwoDays)

	// From 2 to last
	days := weekDays[2:]

	fmt.Println(days)

	// Crear un slice vacío
	names := make([]string, 8)

	// Copy names
	copyOfNames := make([]string, len(names))
	countOfCopiedElements := copy(copyOfNames, names)

	fmt.Println(countOfCopiedElements)
}
