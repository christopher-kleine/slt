# Slice Tools

[![Go Reference](https://pkg.go.dev/badge/github.com/christopher-kleine/slt.svg)](https://pkg.go.dev/github.com/christopher-kleine/slt)

Some simple functions to work with slices. Functions in the official [slices](https://pkg.go.dev/slices) are not part of this small library.

**NOTE:** These implementations are *not* optimized.

## What does this library bring to the table?

- `All / Every`: Checks if all entries pass the provided function.
- `Any / Some`: Checks if at least one entry passes the provided function.
- `Count`: Counts all entries that pass the provided function.
- `Diff`: Return the elements that are in only one of two slices.
- `Find`: Find returns the first element that match the predicate.
- `GroupBy`: Creates a map of slices, grouped by the provided function.
- `Intersect`: Takes 2 slices and returns a new slice containing overlapping elements.
- `Map`: Creates a new slice based on the provided function.
- `Mode`: Calculates the mode of a slice.
- `None`: Checks if no entry passes the provided function.
- `Overlap / OverlapFunc`: Checks if two slices have overlapping entries.
- `Reduce`: Loops all entries and returns a new value.
- `Reject`: Removes all entries that pass the provided function. (This is similiar to [slices.DeleteFunc](https://pkg.go.dev/slices#DeleteFunc), but it doesn't change the original slice.)
- `Remove`: Remove certain elements from a slice. ([slices.Delete](https://pkg.go.dev/slices#Delete) uses an index, this uses values).
- `Select`: Keeps all entries that pass the provided function.
- `SplitBy`: Splits a slice into 2 parts using the provided function.
- `Union`: Merge all slices and remove duplicates.
- `Unique / UniqueFunc`: Removed duplicates from the slice.

Aside from that there are some functions for numbers:

- `Above`: Check if a Number is above a certain value. [Closure]
- `Below`: Check if a Number is below a certain value. [Closure]
- `Between`: Check if Number is between two values. [Closure]
- `DividableBy`: Check if a Number is dividable by a certain value. [Closure]
- `NotDividableBy`: Check if a Number is not dividable by a certain value. [Closure]
- `Even`: Check if an Integer is dividable by 2. [Alias for DividableBy(2)]
- `Odd`: Check if an Integer is NOT dividable by 2. [Alias for NotDividableBy(2)]
- `Sum`: Calculates the sum of all entries in a numeric slice.
- `Mean / Avg`: Calculates the mean of all entries in a numeric slice.
- `Median`: Calculates the median of all entries in a numeric slice.

## Examples

For the sake of keeping things short, the following code is assumed in all cases:

```go
type Person struct {
    Name string
    Age  int
}

type Hero struct {
    Name     string
    Universe string
}

func isMinor(person Person) bool {
    return person.Age < 18
}

var (
    viewers = []Person{
        { Name: "Chris", Age: 38 },
        { Name: "Danny", Age: 21 },
        { Name: "Betty", Age: 30 },
        { Name: "Karen", Age: 55 },
        { Name: "Peter", Age: 17 },
    }

    roomA = []string{
        "Betty", "Chris", "Danny", "Peter",
    }

    roomB = []string{
        "Betty", "Danny", "Karen", "Peter",
    }

    heroes = []Hero{
        { Name: "Batman", Universe: "DC" },
        { Name: "Captain America", Universe: "Marvel" },
        { Name: "Flash", Universe: "DC" },
        { Name: "Hulk", Universe: "Marvel" },
        { Name: "Iron Man", Universe: "Marvel" },
        { Name: "Robin", Universe "DC" },
        { Name: "Superman", Universe: "DC" },
        { Name: "Thor", Universe: "Marvel" },
    }
)
```


### All / Every

The function "All" loops through all entries of a given slice and executes the provided callback on every element. That is, until the callback returns false or there are no more entries to check.

The function "Every" is simply an alias to "All" and works the same way.

```go
func main() {
    // This returns false right at the first entry "Chris"
    if slt.All(viewers, isMinor) == false {
        fmt.Println("Not all viewers are minors!")
    }
}
```

### Any / Some

The function "Any" loops through all entries of a given slice and executes the provided callback on every element. That is, until the callback returns true or there are no more entries to check.

The function "Some" is simply an alias to "Any" and works the same way.

```go
func main() {
    // This returns true at the last entry "Peter".
    if slt.Any(viewers, isMinor) == false {
        fmt.Println("At least one viewer is a minor!")
    }
}
```

### Count

The function "Count" loops through all entries of a given slice and executes the provided callback on every element. It counts all entries that return true.

```go
func main() {
    fmt.Printf("There are %d minor viewer(s).", slt.Count(viewers, isMinor))

    // Output:
    // There are 1 minor viewer(s).
}
```

### Diff

The function "Diff" takes two slices of the same type and returns the elements that appear only in one of the slices.

```go
func main() {
    fmt.Println(slt.Diff(roomA, roomB))

    // Output:
    // - Chris
    // - Karen
}
```

### Find

The function "Find" takes a slice, an offset and a predicate. It returns the first element - and it's index - in the slice that matches the predicate. If there are no elements that match, the default value for that datatype and -1 for the index are returned.

This functions is less useful when using primitive types. For this, the [`slices.Index`](https://pkg.go.dev/slices#Index) would be more useful.

```go
func main() {
    someone, index := slt.Find(room, 0, func(person Person) bool {
        return !isMinor(person)
    })

    fmt.Printf("The first person of age in the room is %s at position %d", someone.Name, index)

    // Output:
    // The first person of age in the room is Chris at position 0
}
```

### GroupBy

The function "GroupBy" creates a map based on the result of the callback. This can be used to group a slice into multiple sub-slices.

```go
func main() {
    byUniverse := func(hero Hero) string {
        return hero.Universe
    }

    grouped := slt.GroupBy(heroes, byUniverse)
	for universe, heroes := range grouped {
		fmt.Printf("Universe: %v\n", universe)
		fmt.Printf("%+v\n", heroes)
	}

    // Output:
    // Universe: DC
    // [{Name:Batman Universe:DC} {Name:Flash Universe:DC} {Name:Robin Universe:DC} {Name:Superman Universe:DC}]
    // Universe: Marvel
    // [{Name:Captain America Universe:Marvel} {Name:Hulk Universe:Marvel} {Name:Iron Man Universe:Marvel} {Name:Thor Universe:Marvel}]
}
```