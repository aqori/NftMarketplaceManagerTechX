// cmd/nftmarketplacemanagertechx/main.go
package main

import (
"flag"
"log"
"os"

"nftmarketplacemanagertechx/internal/nftmarketplacemanagertechx"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmarketplacemanagertechx.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
