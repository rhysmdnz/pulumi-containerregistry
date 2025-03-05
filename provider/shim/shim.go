package shim

import (
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/rhysmdnz/terraform-provider-containerregistry/internal/provider"
)

func NewProvider(version string) tfpf.Provider {
	return provider.New(version)()
}
