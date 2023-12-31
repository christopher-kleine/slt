# Slice Tools

[![Go Reference](https://pkg.go.dev/badge/github.com/christopher-kleine/slt.svg)](https://pkg.go.dev/github.com/christopher-kleine/slt)

Some simple functions to work with slices. Functions in the official [slices](https://pkg.go.dev/slices) are not part of this small library.

**NOTE:** These implementations are *not* optimized.

## What does this library bring to the table?

- `All / Every`: Checks if all entries pass the provided function.
- `Any / Some`: Checks if at least one entry passes the provided function.
- `Count`: Counts all entries that pass the provided function.
- `Diff`: Return the elements that are in only one of two slices.
- `Group`: Creates a map of slices, grouped by the provided function.
- `Map`: Creates a new slice based on the provided function.
- `None`: Checks if no entry passes the provided function.
- `Overlap / OverlapFunc`: Checks if two slices have overlapping entries.
- `Reduce`: Loops all entries and returns a new value.
- `Remove`: Remove certain elements from a slice. (slices uses an index, this uses values).
- `Select`: Keeps all entries that pass the provided function.
- `Union`: Merge all slices and remove duplicates.
- `Unique / UniqueFunc`: Removed duplicates from the slice.

Aside from that there are some functions for numbers:

- `Between`: Check if Number is between two values. [Curried Function]
- `Even`: Check if an Integer is dividable by 2.
- `Greater`: Check if a Number is greater than a certain value. [Curried Function]
- `Less`: Check if a Number is below a certain value. [Curried Function]
- `Odd`: Check if an Integer is NOT dividable by 2.
