package manifests

import (
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"

	kubeaddon "github.com/coreos/tectonic-config/config/kube-addon"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/rand"
)

const (
	kaoCfgFilename = "kube-addon-operator-config.yml"
)

// kubeAddonOperator generates the network-operator-*.yml files
type kubeAddonOperator struct {
	Config *kubeaddon.OperatorConfig
	File   *asset.File
}

var _ asset.WritableAsset = (*kubeAddonOperator)(nil)

// Name returns a human friendly name for the operator
func (kao *kubeAddonOperator) Name() string {
	return "Kube Addon Operator"
}

// Dependencies returns all of the dependencies directly needed by an
// kubeAddonOperator asset.
func (kao *kubeAddonOperator) Dependencies() []asset.Asset {
	return []asset.Asset{
		&installconfig.InstallConfig{},
	}
}

// Generate generates the network-operator-config.yml and network-operator-manifest.yml files
func (kao *kubeAddonOperator) Generate(dependencies asset.Parents) error {
	installConfig := &installconfig.InstallConfig{}
	dependencies.Get(installConfig)

	kao.Config = &kubeaddon.OperatorConfig{
		TypeMeta: metav1.TypeMeta{
			APIVersion: kubeaddon.APIVersion,
			Kind:       kubeaddon.Kind,
		},
		CloudProvider:      tectonicCloudProvider(installConfig.Config.Platform),
		RegistryHTTPSecret: rand.String(16),
		ClusterConfig: kubeaddon.ClusterConfig{
			APIServerURL: getAPIServerURL(installConfig.Config),
		},
	}

	data, err := yaml.Marshal(kao.Config)
	if err != nil {
		return errors.Wrapf(err, "failed to create %s config from InstallConfig", kao.Name())
	}

	kao.File = &asset.File{
		Filename: kaoCfgFilename,
		Data:     data,
	}

	return nil
}

// Files returns the files generated by the asset.
func (kao *kubeAddonOperator) Files() []*asset.File {
	if kao.File != nil {
		return []*asset.File{kao.File}
	}
	return []*asset.File{}
}

// Load is a no-op because kube-addon-operator manifest is not written to disk.
func (kao *kubeAddonOperator) Load(asset.FileFetcher) (bool, error) {
	return false, nil
}