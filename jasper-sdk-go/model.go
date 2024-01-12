package jasper_sdk_go

import (
	"encoding/json"
	"time"
)

type Ref struct {
	Url           string                     `json:"url"`
	Origin        string                     `json:"origin"`
	Title         string                     `json:"title"`
	Comment       string                     `json:"comment"`
	Tags          []string                   `json:"tags"`
	Sources       []string                   `json:"sources"`
	AlternateUrls []string                   `json:"alternateUrls"`
	Plugins       map[string]json.RawMessage `json:"plugins"`
	Metadata      Metadata                   `json:"metadata"`
	Published     time.Time                  `json:"published"`
	Created       time.Time                  `json:"created"`
	Modified      time.Time                  `json:"modified"`
}

type Metadata struct {
	Modified          string          `json:"modified"`
	Responses         uint            `json:"responses"`
	InternalResponses uint            `json:"internalResponses"`
	Plugins           map[string]uint `json:"plugins"`
	UserUrls          []string        `json:"userUrls"`
	Obsolete          bool            `json:"obsolete"`
}

type Ext struct {
	Tag      string          `json:"tag"`
	Origin   string          `json:"origin"`
	Name     string          `json:"name"`
	Config   json.RawMessage `json:"config"`
	Modified time.Time       `json:"modified"`
}

type User struct {
	Tag            string    `json:"tag"`
	Origin         string    `json:"origin"`
	Name           string    `json:"name"`
	Role           string    `json:"role"`
	ReadAccess     []string  `json:"readAccess"`
	WriteAccess    []string  `json:"writeAccess"`
	TagReadAccess  []string  `json:"tagReadAccess"`
	TagWriteAccess []string  `json:"tagWriteAccess"`
	Modified       time.Time `json:"modified"`
	PubKey         []byte    `json:"pubKey"`
}

type PluginDto struct {
	Tag              string                      `json:"tag"`
	Origin           string                      `json:"origin"`
	Name             string                      `json:"name"`
	Config           json.RawMessage             `json:"config"`
	Defaults         *json.RawMessage            `json:"defaults"`
	Schema           *map[string]json.RawMessage `json:"schema"`
	GenerateMetadata bool                        `json:"generateMetadata"`
	UserUrl          bool                        `json:"userUrl"`
	Modified         time.Time                   `json:"modified"`
}

type Template struct {
	Tag      string                      `json:"tag"`
	Origin   string                      `json:"origin"`
	Name     string                      `json:"name"`
	Config   json.RawMessage             `json:"config"`
	Defaults *json.RawMessage            `json:"defaults"`
	Schema   *map[string]json.RawMessage `json:"schema"`
	Modified time.Time                   `json:"modified"`
}
