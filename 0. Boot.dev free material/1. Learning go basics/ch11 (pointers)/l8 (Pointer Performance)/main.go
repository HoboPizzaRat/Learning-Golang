package main

// if you think that passing pointers is better for performance,
// think again!

// local non-pointer variables are generally faster to pass around than pointers

// Before even thinking about using pointers to optimize your code,
// use pointers when you need a shared reference to a value; otherwise, just use values.
