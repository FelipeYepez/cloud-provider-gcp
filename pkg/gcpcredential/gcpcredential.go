/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gcpcredential

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"k8s.io/cloud-provider-gcp/pkg/credentialconfig"
	"k8s.io/klog/v2"
)

const (
	metadataURL        = "http://metadata.google.internal./computeMetadata/v1/"
	metadataAttributes = metadataURL + "instance/attributes/"
	// DockerConfigKey is the URL of the dockercfg metadata key used by DockerConfigKeyProvider.
	DockerConfigKey = metadataAttributes + "google-dockercfg"
	// DockerConfigURLKey is the URL of the dockercfg metadata key used by DockerConfigURLKeyProvider.
	DockerConfigURLKey = metadataAttributes + "google-dockercfg-url"
	serviceAccounts    = metadataURL + "instance/service-accounts/"
	metadataScopes     = serviceAccounts + "default/scopes"
	metadataToken      = serviceAccounts + "default/token"
	metadataEmail      = serviceAccounts + "default/email"
	// StorageScopePrefix is the prefix checked by ContainerRegistryProvider.Enabled.
	StorageScopePrefix       = "https://www.googleapis.com/auth/devstorage"
	cloudPlatformScopePrefix = "https://www.googleapis.com/auth/cloud-platform"
	defaultServiceAccount    = "default/"
)

// GCEProductNameFile is the product file path that contains the cloud service name.
// This is a variable instead of a const to enable testing.
var GCEProductNameFile = "/sys/class/dmi/id/product_name"

// For these urls, the parts of the host name can be glob, for example '*.gcr.io" will match
// "foo.gcr.io" and "bar.gcr.io".
var containerRegistryUrls = []string{"container.cloud.google.com", "gcr.io", "*.gcr.io", "*.pkg.dev"}

var metadataHeader = &http.Header{
	"Metadata-Flavor": []string{"Google"},
}

// MetadataProvider is a DockerConfigProvider that reads its configuration from Google
// Compute Engine metadata.
type MetadataProvider struct {
	Client *http.Client
}

// DockerConfigKeyProvider is a DockerConfigProvider that reads its configuration from a specific
// Google Compute Engine metadata key: 'google-dockercfg'.
type DockerConfigKeyProvider struct {
	MetadataProvider
}

// DockerConfigURLKeyProvider is a DockerConfigProvider that reads its configuration from a URL read from
// a specific Google Compute Engine metadata key: 'google-dockercfg-url'.
type DockerConfigURLKeyProvider struct {
	MetadataProvider
}

// ContainerRegistryProvider is a DockerConfigProvider that provides a dockercfg with:
//
//	Username: "_token"
//	Password: "{access token from metadata}"
type ContainerRegistryProvider struct {
	MetadataProvider
	UseRegistryFromImage bool
}

// Returns true if it finds a local GCE VM.
// Looks at a product file that is an undocumented API.
func onGCEVM() bool {
	var name string

	if runtime.GOOS == "windows" {
		data, err := exec.Command("wmic", "computersystem", "get", "model").Output()
		if err != nil {
			return false
		}
		fields := strings.Split(strings.TrimSpace(string(data)), "\r\n")
		if len(fields) != 2 {
			klog.V(2).Infof("Received unexpected value retrieving system model: %q", string(data))
			return false
		}
		name = fields[1]
	} else {
		data, err := ioutil.ReadFile(GCEProductNameFile)
		if err != nil {
			klog.V(2).Infof("Error while reading product_name: %v", err)
			return false
		}
		name = strings.TrimSpace(string(data))
	}
	return name == "Google" || name == "Google Compute Engine"
}

// Enabled implements DockerConfigProvider for all of the Google implementations.
func (g *MetadataProvider) Enabled() bool {
	return onGCEVM()
}

// Provide implements DockerConfigProvider
func (g *DockerConfigKeyProvider) Provide(image string) credentialconfig.DockerConfig {
	// Read the contents of the google-dockercfg metadata key and
	// parse them as an alternate .dockercfg
	if cfg, err := credentialconfig.ReadDockerConfigFileFromURL(DockerConfigKey, g.Client, metadataHeader); err != nil {
		klog.Errorf("while reading 'google-dockercfg' metadata: %v", err)
	} else {
		return cfg
	}

	return credentialconfig.DockerConfig{}
}

// Provide implements DockerConfigProvider
func (g *DockerConfigURLKeyProvider) Provide(image string) credentialconfig.DockerConfig {
	// Read the contents of the google-dockercfg-url key and load a .dockercfg from there
	if url, err := credentialconfig.ReadURL(DockerConfigURLKey, g.Client, metadataHeader); err != nil {
		klog.Errorf("while reading 'google-dockercfg-url' metadata: %v", err)
	} else {
		if strings.HasPrefix(string(url), "http") {
			if cfg, err := credentialconfig.ReadDockerConfigFileFromURL(string(url), g.Client, nil); err != nil {
				klog.Errorf("while reading 'google-dockercfg-url'-specified url: %s, %v", string(url), err)
			} else {
				return cfg
			}
		} else {
			// TODO(mattmoor): support reading alternate scheme URLs (e.g. gs:// or s3://)
			klog.Errorf("Unsupported URL scheme: %s", string(url))
		}
	}

	return credentialconfig.DockerConfig{}
}

// runWithBackoff runs input function `f` with an exponential backoff.
// Note that this method can block indefinitely.
func runWithBackoff(f func() ([]byte, error)) []byte {
	var backoff = 100 * time.Millisecond
	const maxBackoff = time.Minute
	for {
		value, err := f()
		if err == nil {
			return value
		}
		time.Sleep(backoff)
		backoff = backoff * 2
		if backoff > maxBackoff {
			backoff = maxBackoff
		}
	}
}

// Enabled implements a special metadata-based check, which verifies the
// storage scope is available on the GCE VM.
// If running on a GCE VM, check if 'default' service account exists.
// If it does not exist, assume that registry is not enabled.
// If default service account exists, check if relevant scopes exist in the default service account.
// The metadata service can become temporarily inaccesible. Hence all requests to the metadata
// service will be retried until the metadata server returns a `200`.
// It is expected that "http://metadata.google.internal./computeMetadata/v1/instance/service-accounts/" will return a `200`
// and "http://metadata.google.internal./computeMetadata/v1/instance/service-accounts/default/scopes" will also return `200`.
// More information on metadata service can be found here - https://cloud.google.com/compute/docs/storing-retrieving-metadata
func (g *ContainerRegistryProvider) Enabled() bool {
	// Given that we are on GCE, we should keep retrying until the metadata server responds.
	value := runWithBackoff(func() ([]byte, error) {
		value, err := credentialconfig.ReadURL(serviceAccounts, g.Client, metadataHeader)
		if err != nil {
			klog.V(2).Infof("Failed to Get service accounts from gce metadata server: %v", err)
		}
		return value, err
	})
	// We expect the service account to return a list of account directories separated by newlines, e.g.,
	//   sv-account-name1/
	//   sv-account-name2/
	// ref: https://cloud.google.com/compute/docs/storing-retrieving-metadata
	defaultServiceAccountExists := false
	for _, sa := range strings.Split(string(value), "\n") {
		if strings.TrimSpace(sa) == defaultServiceAccount {
			defaultServiceAccountExists = true
			break
		}
	}
	if !defaultServiceAccountExists {
		klog.V(2).Infof("'default' service account does not exist. Found following service accounts: %q", string(value))
		return false
	}
	url := metadataScopes + "?alt=json"
	value = runWithBackoff(func() ([]byte, error) {
		value, err := credentialconfig.ReadURL(url, g.Client, metadataHeader)
		if err != nil {
			klog.V(2).Infof("Failed to Get scopes in default service account from gce metadata server: %v", err)
		}
		return value, err
	})
	var scopes []string
	if err := json.Unmarshal(value, &scopes); err != nil {
		klog.Errorf("Failed to unmarshal scopes: %v", err)
		return false
	}
	for _, v := range scopes {
		// cloudPlatformScope implies storage scope.
		if strings.HasPrefix(v, StorageScopePrefix) || strings.HasPrefix(v, cloudPlatformScopePrefix) {
			return true
		}
	}
	klog.Warningf("Google container registry is disabled, no storage scope is available: %s", value)
	return false
}

// TokenBlob is used to decode the JSON blob containing an access token
// that is returned by GCE metadata.
type TokenBlob struct {
	AccessToken string `json:"access_token"`
}

// Provide implements DockerConfigProvider
func (g *ContainerRegistryProvider) Provide(image string) credentialconfig.DockerConfig {
	cfg := credentialconfig.DockerConfig{}

	tokenJSONBlob, err := credentialconfig.ReadURL(metadataToken, g.Client, metadataHeader)
	if err != nil {
		klog.Errorf("while reading access token endpoint: %v", err)
		return cfg
	}

	email, err := credentialconfig.ReadURL(metadataEmail, g.Client, metadataHeader)
	if err != nil {
		klog.Errorf("while reading email endpoint: %v", err)
		return cfg
	}

	var parsedBlob TokenBlob
	if err := json.Unmarshal([]byte(tokenJSONBlob), &parsedBlob); err != nil {
		klog.Errorf("while parsing json blob %s: %v", tokenJSONBlob, err)
		return cfg
	}

	entry := credentialconfig.DockerConfigEntry{
		Username: "_token",
		Password: parsedBlob.AccessToken,
		Email:    string(email),
	}

	// If UseRegistryFromImage is true, we will directly give the credential to the registry of the image.
	// Currently, this is only used by auth-provider-gcp.
	if g.UseRegistryFromImage {
		if registry, _, found := strings.Cut(image, "/"); found {
			cfg[registry] = entry
		}
	}

	// Add our entry for each of the supported container registry URLs
	for _, k := range containerRegistryUrls {
		cfg[k] = entry
	}
	return cfg
}
