# Exercise: Maps

Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure. 

## WordCount Function

Here’s a description of your solution for the README.md file:

WordCount Function
The `WordCount` function is implemented to count the frequency of each word in a given string `s`. It returns a map where the keys are words and the values are their respective counts.

### How It Works
1. **Splitting the String:**
    - The function first splits the string into individual words using `strings.Fields(s)`. The `strings.Fields` function breaks the input string `s` into a slice of words, separated by spaces, and ignores any leading or trailing whitespace.

2. **Counting the Words:**
    - A map (`m`) of type `map[string]int` is created to store the count of each word. The map's key is the word, and the value is the count of occurrences.

3. **Iterating Over Words:**
   - The function iterates through each word in the slice `words`. For each word, it checks if the word already exists in the map.
     - If the word is found in the map, the count is incremented by 1.
     - If the word is not found, it is added to the map with a count of 1.


4. **Returning the Result:**
    - The function returns the map `m`, which contains the frequency of each word in the input string.

## Example:
```go
    s := "go go golang"
    result := WordCount(s)
    fmt.Println(result)
    // Output: map[go:2 golang:1]
```