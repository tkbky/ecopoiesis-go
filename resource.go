package main

import (
	"encoding/json"
	"fmt"
)

type resource interface {
	tf() string
}

// Resource describes a terraform resource
type Resource struct {
	Kind string
	Name string
	Obj  interface{}
}

func (r *Resource) tf() string {
	body, _ := json.MarshalIndent(r.Obj, "", "  ")
	return fmt.Sprintf("resource  \"%s\" \"%s\" %s", r.Kind, r.Name, body)
}
