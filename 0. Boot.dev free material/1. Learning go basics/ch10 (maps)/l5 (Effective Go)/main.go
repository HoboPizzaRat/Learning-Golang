package main

// If you pass a map to a function that changes the contents of the map,
// the changes will be visible in the caller.

// maps can be easily built using literals during initialization
/*
var timeZone = map[string]int{
    "UTC":  0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,
}
*/

// attempt to fetch a map value with a key that is not present in the map
// will return the zero value of the given type
/*
if attended[person] { // will be false if person is not in the map
    fmt.Println(person, "was at the meeting")
}
*/

// though sometimes you want to be sure that you are specifying a missing value on the map
// rather than the zero value (the key may still exist in the map, but with zero value)
/*
seconds, ok = timeZone[tz]
or
if seconds, ok := timeZone[tz]; ok {
    return seconds
}
*/

// deleting map entires can be done with delete function
/*
delete(timeZone, "PDT")
*/
