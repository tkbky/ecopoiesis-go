package tf

import (
	"encoding/json"
	"fmt"
)

type tfresource interface {
	Tf() string
}

// Resource describes a terraform resource
type Resource struct {
	Kind string
	Name string
	Obj  interface{}
}

// Tf returns a tf string
func (r *Resource) Tf() string {
	body, _ := json.MarshalIndent(r.Obj, "", "  ")
	return fmt.Sprintf("resource  \"%s\" \"%s\" %s", r.Kind, r.Name, body)
}
