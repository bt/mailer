package emailing

import "github.com/pkg/errors"

type Provider interface {
	// Sends an email out using the provider.
	SendEmail(*Email) bool
	// Will return true if provider is active, otherwise false.
	Active() bool
	// Deactivates the provider.
	Deactivate()
}

// Singleton to hold a registry of providers
var Providers []Provider

// Register a provider into the registry.
func RegisterProvider(p Provider) {
	Providers = append(Providers, p)
}

// Returns an available provider, or error if none available.
func AvailableProvider() (Provider, error) {
	for _, provider := range Providers {
		if provider.Active() {
			return provider, nil
		}
	}
	return nil, errors.New("No available providers.")
}

// Initialise the provider array.
func init() {
	Providers = make([]Provider, 0)
}
