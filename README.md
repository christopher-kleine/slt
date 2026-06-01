# Slice Tools

[![Go Reference](https://pkg.go.dev/badge/github.com/christopher-kleine/slt.svg)](https://pkg.go.dev/github.com/christopher-kleine/slt)

Some simple functions to work with slices. Functions in the official
[slices](https://pkg.go.dev/slices) are not part of this small library.

**NOTE:** These implementations are not yet optimized.

## Packages

This library contains the following:

- `github.com/christopher-kleine/slt`: The core of "slice tools". Contains some usefull functions for all types of slices.
- `github.com/christopher-kleine/slt/numbers`: Functions especially for numbers. Mostly closures.
- `github.com/christopher-kleine/slt/chain`: Chainable version of "slice tools" whereever possible. (TODO)

## Functions in "slice tools" and how they compare to the `slices` package

Highlightes functions are optimized to some degree.

| Function      | Description                                                                             | Compared to `slices`                                                 |
| ------------- | --------------------------------------------------------------------------------------- | -------------------------------------------------------------------- |
| All           | Checks if all entries pass the provided function.                                       | `slices.All` is an iterator                                          |
| Any           | Checks if at least one entry passes the provided function.                              |                                                                      |
| Chunk         | Splits a slice in equal sized chunks                                                    | `slices.Chunk` is an iterator                                        |
| ChunkBy       | Splits a slice based on the result of the predicate                                     |                                                                      |
| Count         | Counts all entries that pass the provided function.                                     |                                                                      |
| Diff          | Return the elements that are in only one of two slices.                                 |                                                                      |
| EqualValues   | Returns if two slices have the same content. Regardless of the order.                   | `slices.Equal` requires values and their index to match              |
| Every         | *Alias for All*                                                                         |                                                                      |
| Filter        | *Alias for Select*                                                                      |                                                                      |
| Find          | Find returns the first element (and their index) that match the predicate.              | `slices.Index` only returns the index                                |
| First         | Returns the first element, the rest and an error if the slice is already empty          |                                                                      |
| Flatten       | Takes []S and returns S (S is a slice of any type)                                      |                                                                      |
| Fold          | *Alias for Reduce*                                                                      |                                                                      |
| FoldR         | *Alias for ReduceR*                                                                     |                                                                      |
| GroupBy       | Creates a map of slices, grouped by the provided function.                              |                                                                      |
| GroupCount    | Similiar to GroupBy, but returns a map with int64 instead                               |                                                                      |
| Head          | *Alias for First*                                                                       |                                                                      |
| IndexBy       | Similiar to GroupBy, but the resulting map has only one element                         |                                                                      |
| Intersect     | Takes 2 slices and returns a new slice containing overlapping elements.                 |                                                                      |
| Last          | Returns the last element, the rest and an error in case the slice is already empty      |                                                                      |
| Map           | Creates a new slice based on the provided function. Target type can be different        |                                                                      |
| Some          | *Alias for Any*                                                                         |                                                                      |
| Mode          | Finds the element that appears the most                                                 |                                                                      |
| None          | Checks if no entry passes the provided predicate                                        |                                                                      |
| Overlap       | Checks if two slices overlap                                                            |                                                                      |
| OverlapFunc   | Checks if two slices overlap using a predicate instead of their value                   |                                                                      |
| Pick          | Picks N random elements from the slice and returns a new slice. Duplicates are possible |                                                                      |
| PickUnique    | Similiar to Pick, but makes sure the elements are unique                                |                                                                      |
| Reduce        | Reduces the slice to a single value. Target type can be different from input type       |                                                                      |
| ReduceR       | Similiar to Reduce. But starts at the end instead                                       |                                                                      |
| Reject        | Removes all entries that pass the provided function                                     | `slices.DeleteFunc` modifies the original slice                      |
| Remove        | Removes all entries that match the given values                                         | `slices.Delete` uses an index, not values                            |
| ReplaceValues | Replaces values based on a replace-map                                                  | `slices.Replace` uses a predefined range instead of values           |
| Select        | Creates a new slice containing only the elements that pass the predicate                |                                                                      |
| SplitBy       | Splits a slice in to parts based on the predicate                                       |                                                                      |
| Tail          | *Alias for Last*                                                                        |                                                                      |
| Union         | Merge all slices and remove duplicates from incoming slices                             | `slices.Concat` doesn't remove duplicate values from incoming slices |
| **Unique**    | Removed duplicates from the slice.                                                      | `slices.Compact` only removes duplicates if they're consecutive      |
| UniqueFunc    | Removed duplicates from the slice based on the predicate                                | The same as `Unique` but for `slices.CompactFunc`                    |


## Functions in "numbers"

| Function       | Description                                                             | Closure |
| -------------- | ----------------------------------------------------------------------- | :-----: |
| Above          | Check if a Number is above a certain value                              |   Yes   |
| AboveOrEqual   | Check if a Number is above or equal to a certain value                  |   Yes   |
| Avg            | *Alias for Mean*                                                        |   No    |
| Below          | Check if a Number is below a certain value                              |   Yes   |
| BelowOrEqual   | Check if a Number is below or equal to a certain value                  |   Yes   |
| Between        | Checks if a Number is between two values                                |   Yes   |
| DividableBy    | Check if a Number is dividable by a certain value                       |   Yes   |
| NotDividableBy | Check if a Number is not dividable by a certain value                   |   Yes   |
| Even           | Check if an Integer is dividable by 2 (Alias for DividableBy(2))        |   Yes   |
| Odd            | Check if an Integer is NOT dividable by 2 (Alias for NotDividableBy(2)) |   Yes   |
| Sum            | Calculates the sum of all entries in a numeric slice                    |   No    |
| Mean           | Calculates the mean of all entries in a numeric slice                   |   No    |
| Median         | Calculates the median of all entries in a numeric slice                 |   No    |

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