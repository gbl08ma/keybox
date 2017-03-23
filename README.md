# keybox
This is an extremely simple Go library that is merely a read-only key-value store for strings. Not much more than a convenience wrapper around the JSON decoder.

# Installation

`go get -u github.com/gbl08ma/keybox`

# Usage
The keybox reads the keys and values from a JSON file like this:

```
{
    "secretName": "th15i5mySeCrE7",
    "otherSecret": "passw0rd"
}
```

Use the keybox like this:

```
// ...

import "github.com/gbl08ma/keybox

// ...

var (
	secrets *keybox.Keybox
)

func main() {
    var err error
    secrets, err = keybox.Open("pathTo/FileWithSecrets.json")

    // now you can retrieve secrets like this:
    mySecret, present := secrets.Get("secretName")
	if !present {
		log.Fatal("Missing secret!")
	}
    // use mySecret
    AccessSecretStuff(mySecret)

    // ...
}
```

# License
MIT