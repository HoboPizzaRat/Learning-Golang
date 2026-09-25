package main

// With generics, we also got type-set (type-list) interfaces.
//

// Ordered matches any type that supports <, <=, >, and >=.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Because T is constrained by Ordered, the compiler knows
// that < is valid for any T used with this function.
func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
