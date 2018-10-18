package config

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"gotest.tools/fs"
)

func TestLoadFile(t *testing.T) {
	dir := fs.NewDir(t, "config",
		fs.WithFile("foo.hcl", `name = "foo"
config {
  bar = "baz"
}`))
	defer dir.Remove()
	c, err := LoadFile(dir.Join("foo.hcl"))
	assert.NilError(t, err)
	assert.Assert(t, is.DeepEqual(c, Config{
		Name: "foo",
		Configs: map[string]interface{}{
			"bar": "baz",
		},
	}))
}

func TestLoad(t *testing.T) {
	c, err := Load(`name = "knative"
config {
       openshift-version = "v3.11.0"
       memory = "8GB"
       cpus = 4
       disk-size = "50GB"
       image-caching = true
}`)
	assert.NilError(t, err)
	assert.Assert(t, is.DeepEqual(c, Config{
		Name: "knative",
		Configs: map[string]interface{}{
			"openshift-version": "v3.11.0",
			"memory":            "8GB",
			"cpus":              4,
			"disk-size":         "50GB",
			"image-caching":     true,
		},
	}))
}
