# Utils Package

This package contains utility functions for the CLIsland game, including character generation.

## Character Generation

### `generateCharacters(characters []string, playerSex string) []string`

Generates a list of character names for the Love Island game based on the player's sex.

**Parameters:**
- `characters`: A slice of existing character names (can be empty)
- `playerSex`: The player's sex ("male" or "female")

**Returns:**
- A slice of character names (not including the player)

**Logic:**
- If player is "male": generates 5 boys and 6 girls (11 total characters)
- If player is "female": generates 6 boys and 5 girls (11 total characters)
- The result is shuffled to randomize the order

**Example:**
```go
// For a male player
characters := generateCharacters([]string{}, "male")
// Returns: ["James", "Emma", "Sophie", "Alex", "Mia", "Liam", "Olivia", "Noah", "Ava", "Ethan", "Isabella"]

// For a female player  
characters := generateCharacters([]string{}, "female")
// Returns: ["Liam", "Emma", "Noah", "Olivia", "Ethan", "Ava", "James", "Sophie", "Alex", "Mia", "Isabella"]
```

### Helper Functions

#### `generateRandomNameFromFile(sex string) string`

Reads a random name from the appropriate name file based on sex.

**Parameters:**
- `sex`: "male" or "female"

**Returns:**
- A random name from the corresponding file (`data/names/boys.txt` or `data/names/girls.txt`)

#### `shuffleStrings(slice []string)`

Shuffles a slice of strings in place using the Fisher-Yates algorithm.

**Parameters:**
- `slice`: The slice of strings to shuffle

## Usage

```go
package main

import "github.com/TuringProblem/CLIsland/utils"

func main() {
    // Generate characters for a male player
    characters := utils.generateCharacters([]string{}, "male")
    fmt.Printf("Generated %d characters: %v\n", len(characters), characters)
}
```

## Testing

Run the tests with:
```bash
go test ./utils/ -v
``` 