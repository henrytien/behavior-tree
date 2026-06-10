package behaviortree

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"reflect"
)

// getMd5String returns the hexadecimal MD5 digest of s.
func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// getGuid returns a compact pseudo-random identifier string.
func getGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return getMd5String(base64.URLEncoding.EncodeToString(b))
}

// CreateUUID returns a unique identifier for trees and nodes.
func CreateUUID() string {
	return getGuid()
}

// MinInt returns the smaller of a and b.
func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// RegisterStructMaps stores constructor metadata by name.
type RegisterStructMaps struct {
	maps map[string]reflect.Type
}

// NewRegisterStructMaps creates an empty registry.
func NewRegisterStructMaps() *RegisterStructMaps {
	return &RegisterStructMaps{make(map[string]reflect.Type)}
}

// New creates a new instance registered under name.
func (rsm *RegisterStructMaps) New(name string) (interface{}, error) {
	if v, ok := rsm.maps[name]; ok {
		return reflect.New(v).Interface(), nil
	}

	return nil, fmt.Errorf("not found %s struct", name)
}

// CheckElem reports whether name has been registered.
func (rsm *RegisterStructMaps) CheckElem(name string) bool {
	_, ok := rsm.maps[name]
	return ok
}

// Register stores the concrete type of c under name.
func (rsm *RegisterStructMaps) Register(name string, c interface{}) {
	rsm.maps[name] = reflect.TypeOf(c).Elem()
}
