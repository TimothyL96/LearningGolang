package key

import (
	"errors"
	"math"
	"strconv"
)

// Set the intended key type.
// Note the key type should always be unsigned
type keyType uint32

// Set the limit per key.
// Should not be more than max of the keyType above
const keyMaxLimit = math.MaxUint32

// Key is the wrapper struct for InterfaceKey.
// Key contains the real key struct BaseKey which can be retrieved with Key()
type Key interface {
	Key() *BaseKey
	SiteKey() keyType
	MajorKey() keyType
	MinorKey() keyType
	String() string
}

// BaseKey is the underlying struct for the Key. Its fields are accessible by methods
type BaseKey struct {
	siteKey  keyType
	majorKey keyType
	minorKey keyType
}

// Store the last generated Key
var lastKey *BaseKey

// Convert int to string easily
func (kt keyType) toString() string {
	return strconv.FormatUint(uint64(kt), 10)
}

// Key returns the key of the called as the BaseKey type.
// This is used for the Key interface to get the BaseKey before calling any BaseKey's method
func (key *BaseKey) Key() *BaseKey {
	if key == nil {
		panic(errors.New("key is nil").Error())
	}

	return key
}

// String converts the key to string with the format "[siteKey.majorKey.minorKey]"
func (key *BaseKey) String() string {
	return "[" + key.siteKey.toString() +
		"." + key.majorKey.toString() +
		"." + key.minorKey.toString() + "]"
}

// SiteKey returns the site key of the caller
func (key *BaseKey) SiteKey() keyType {
	return key.siteKey
}

// MajorKey returns the major key of the caller
func (key *BaseKey) MajorKey() keyType {
	return key.majorKey
}

// MinorKey returns the minor key of the caller
func (key *BaseKey) MinorKey() keyType {
	return key.minorKey
}

// NewKey creates and then returns a new key of type Key with it's key value incremented by 1 from the last created key.
func NewKey() Key {
	if lastKey == nil {
		// If this is the first key to be generated,
		// assign last key with a BaseKey and all the keys are auto initialized to zero
		lastKey = &BaseKey{}

		// Return the base key wrapped in Key struct
		return lastKey
	}

	// Determine the increment for the new key
	siteKey, majorKey, minorKey := incrementKey(lastKey.siteKey, lastKey.majorKey, lastKey.minorKey, keyMaxLimit)

	// Create and assign new key to last key
	lastKey = &BaseKey{
		siteKey:  siteKey,
		majorKey: majorKey,
		minorKey: minorKey,
	}

	// Return the base key wrapped in Key struct
	return lastKey
}

// Increment the key appropriately
// Minor key will be incremented first, but if its value reached keyMaxLimit,
// it will be reset to 0 and then major key will be incremented instead,
// or else site key will be incremented and both minor and major key will be reset to 0.
//
// The method panics if the key reaches its limit, which is site, major and minor key values all reached keyMaxLimit
func incrementKey(siteKey, majorKey, minorKey keyType, keyMaxLimit keyType) (keyType, keyType, keyType) {
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

	return siteKey, majorKey, minorKey
}
