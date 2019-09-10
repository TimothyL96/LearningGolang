package key

import (
	"errors"
	"math"
	"strconv"
)

const keyMaxLimit = math.MaxInt32

// Key is the wrapper struct for InterfaceKey.
// Key contains the real key struct BaseKey which can be retrieved with Key()
type Key struct {
	baseKey *BaseKey
}

// BaseKey is the underlying struct for the Key. Its fields are accessible by methods
type BaseKey struct {
	siteKey  int
	majorKey int
	minorKey int
}

// Store the last generated Key
var lastKey *BaseKey

// Key returns the key of the called as the BaseKey type.
// This is used for the Key interface to get the BaseKey before calling any BaseKey's method
func (key *Key) Key() *BaseKey {
	if key == nil {
		panic(errors.New("key is nil").Error())
	}

	return key.baseKey
}

// String converts the key to string with the format "[siteKey.majorKey.minorKey]"
func (key *BaseKey) String() string {
	return "[" + strconv.Itoa(key.siteKey) + "." + strconv.Itoa(key.majorKey) + "." + strconv.Itoa(key.minorKey) + "]"
}

// SiteKey returns the site key of the caller
func (key *BaseKey) SiteKey() *int {
	return &key.siteKey
}

// MajorKey returns the major key of the caller
func (key *BaseKey) MajorKey() *int {
	return &key.majorKey
}

// MinorKey returns the minor key of the caller
func (key *BaseKey) MinorKey() *int {
	return &key.minorKey
}

// NewKey creates and then returns a new key of type Key with it's key value incremented by 1 from the last created key.
//
// Minor key will be incremented first, but if its value reached keyMaxLimit,
// it will be reset to 0 and then major key will be incremented instead,
// or else site key will be incremented and both minor and major key will be reset to 0.
//
// The method panics if the key reaches its limit, which is site, major and minor key values all reached keyMaxLimit
func NewKey() *Key {
	if lastKey == nil {
		// If this is the first key to be generated,
		// assign last key with a BaseKey and all the keys are auto initialized to zero
		lastKey = &BaseKey{}

		// Return the base key wrapped in Key struct
		return &Key{lastKey}
	}

	// Determine the increment for the new key
	siteKey, majorKey, minorKey := lastKey.siteKey, lastKey.majorKey, lastKey.minorKey

	if minorKey != keyMaxLimit {
		minorKey++
	} else if majorKey != keyMaxLimit {
		minorKey = 0
		majorKey++
	} else if siteKey != keyMaxLimit {
		majorKey, minorKey = 0, 0
		siteKey++
	} else {
		panic(errors.New("maxed out key").Error())
	}

	// Create and assign new key to last key
	lastKey = &BaseKey{
		siteKey:  siteKey,
		majorKey: majorKey,
		minorKey: minorKey,
	}

	// Return the base key wrapped in Key struct
	return &Key{lastKey}
}
