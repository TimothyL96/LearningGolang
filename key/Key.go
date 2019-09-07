package key

import (
	"errors"
	"strconv"
)

const keyMaxLimit = 1000000

// Key interface allows and forces other objects to use Key()
// before calling Key related methods such as get SiteKey() and String()
type Key interface {
	Key() *BaseKey
}

// BaseKey is the underlying struct for the Key. Its fields are accessible by methods
type BaseKey struct {
	siteKey  int
	majorKey int
	minorKey int
}

// Store the last generated Key
var lastKey *BaseKey

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

// Key returns the key of the called as the BaseKey type.
// This is used for the Key interface to get the BaseKey before calling any BaseKey's method
func (key *BaseKey) Key() *BaseKey {
	return key
}

// NewKey creates and then returns a new key of type BaseKey with it's key value incremented by 1 from the last created key.
// Minor key will be incremented first, but it will be reset to 0 if its value reached keyMaxLimit,
// then major key will be incremented instead, or else site key will be incremented.
// The method panic if the key reaches its limit, which is site, major and minor key values all reached keyMaxLimit
func NewKey() *BaseKey {
	siteKey, majorKey, minorKey := lastKey.siteKey, lastKey.majorKey, lastKey.minorKey

	if lastKey.minorKey != keyMaxLimit {
		minorKey = lastKey.minorKey + 1
	} else if lastKey.majorKey != keyMaxLimit {
		minorKey = 0
		majorKey = lastKey.majorKey + 1
	} else if lastKey.siteKey != keyMaxLimit {
		majorKey, minorKey = 0, 0
		siteKey = lastKey.siteKey + 1
	} else {
		panic(errors.New("maxed out key").Error())
	}

	newKey := &BaseKey{
		siteKey:  siteKey,
		majorKey: majorKey,
		minorKey: minorKey,
	}

	lastKey = newKey

	return newKey
}
