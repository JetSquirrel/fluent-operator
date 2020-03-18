package filter

import (
	"kubesphere.io/fluentbit-operator/api/v1alpha1/plugins"
)

// +kubebuilder:object:generate:=true

// The Nest Filter plugin allows you to operate on or with nested data. Its modes of operation are
type Nest struct {
	// Select the operation nest or lift
	// +kubebuilder:validation:Enum:=nest;list
	Operation string `json:"operation,omitempty"`
	// Nest records which field matches the wildcard
	Wildcard []string `json:"wildcard,omitempty"`
	// Nest records matching the Wildcard under this key
	NestUnder string `json:"nestUnder,omitempty"`
	// Lift records nested under the Nested_under key
	NestedUnder string `json:"nestedUnder,omitempty"`
	// Prefix affected keys with this string
	AddPrefix string `json:"addPrefix,omitempty"`
	// Remove prefix from affected keys if it matches this string
	RemovePrefix string `json:"removePrefix,omitempty"`
}

func (_ *Nest) Name() string {
	return "nest"
}

func (n *Nest) Params(_ plugins.SecretLoader) (*plugins.KVs, error) {
	kvs := plugins.NewKVs()
	if n.Operation != "" {
		kvs.Insert("Operation", n.Operation)
	}
	for _, wc := range n.Wildcard {
		kvs.Insert("Wildcard", wc)
	}
	if n.NestUnder != "" {
		kvs.Insert("Nest_under", n.NestUnder)
	}
	if n.NestedUnder != "" {
		kvs.Insert("Nested_under", n.NestedUnder)
	}
	if n.AddPrefix != "" {
		kvs.Insert("Add_prefix", n.AddPrefix)
	}
	if n.RemovePrefix != "" {
		kvs.Insert("Remove_prefix", n.RemovePrefix)
	}
	return kvs, nil
}