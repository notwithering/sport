# Sport

The Sport Go package provides many constants for interacting between sport team names.

## Examples

### NBA

```go
package main

import (
	"fmt"

	"github.com/notwithering/sport/nba"
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

	"github.com/notwithering/sport/nfl"
)

func main() {
	fmt.Println(nfl.CincinnatiFullTeam)
}
```

> Cincinnati Bengals