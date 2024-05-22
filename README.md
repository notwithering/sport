# Sport

The Sport Go package provides many constants for interacting between sport team names.

## Examples

### NBA

```go
package main

import (
	"fmt"

	nba "github.com/notwithering/sport/NBA"
)

func main() {
	fmt.Println(nba.ClevelandFullTeam)
}
```

> Cleveland Cavaliers

### NFL

```go
package main

import (
	"fmt"

	nfl "github.com/notwithering/sport/NFL"
)

func main() {
	fmt.Println(nfl.CincinnatiFullTeam)
}
```

> Cincinnati Bengals