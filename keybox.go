package keybox

import (
	"encoding/json"
	"os"
)

// A Keybox is a read-only object that stores secrets like API keys.
type Keybox struct {
	keys map[string]interface{}
}

// Open reads a Keybox from the specified file name.
// The should contain JSON-encoded map of key names to secrets.
func Open(filename string) (*Keybox, error) {
	keybox := Keybox{}
	file, _ := os.Open(filename)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&keybox.keys)
	if err != nil {
		return nil, err
	}
	return &keybox, nil
}

// Get retrieves a key from the Keybox, returning the key and whether it is present in the Keybox.
func (box *Keybox) Get(key string) (string, bool) {
	value, exists := box.keys[key]
	if exists {
		value, casts := value.(string)
		return value, casts
	}
	return "", exists
}

// GetBox retrieves a nested Keybox from the Keybox, returning the Keybox and whether it is present in the Keybox.
func (box *Keybox) GetBox(key string) (*Keybox, bool) {
	value, exists := box.keys[key].(map[string]interface{})
	if exists {
		return &Keybox{
			keys: value,
		}, true
	}
	return nil, exists
}
