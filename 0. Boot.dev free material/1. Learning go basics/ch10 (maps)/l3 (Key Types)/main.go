package main

// map keys may be of any type that is comparable.

// It's obvious that strings, ints, and other basic types should be available as map keys
// but
// Struct can be used to key data by multiple dimensions

/*
This is a map of string to (map of string to int)
hits := make(map[string]map[string]int)

this can be used like so
n := hits["/doc/"]["au"]
*/

//  for any given outer key you must check if the inner map exists, and create it if needed:
/*
func add(m map[string]map[string]int, path, country string) {
    mm, ok := m[path]
    if !ok {
        mm = make(map[string]int)
        m[path] = mm
    }
    mm[country]++
}
add(hits, "/doc/", "au")
*/
