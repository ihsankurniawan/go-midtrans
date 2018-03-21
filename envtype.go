package midtrans

import "strings"

type EnvironmentType int8

const (
    _ EnvironmentType = iota
    Sandbox
    Production
    BcaSandbox
)

var typeString = map[EnvironmentType]string {
    Sandbox: "https://api.sandbox.veritrans.co.id",
    Production: "https://api.veritrans.co.id",
    BcaSandbox: "http://117.102.118.210:8080",
}

// implement stringer
func (e EnvironmentType) String() string {
    for k, v := range typeString {
        if k == e {
            return v
        }
    }
    return "undefined"
}

func (e EnvironmentType) SnapUrl() string {
    return strings.Replace(e.String(), "api.", "app.", 1)
}