// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/hive/pkg/apis/hive/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterClaimLister helps list ClusterClaims.
type ClusterClaimLister interface {
	// List lists all ClusterClaims in the indexer.
	List(selector labels.Selector) (ret []*v1.ClusterClaim, err error)
	// ClusterClaims returns an object that can list and get ClusterClaims.
	ClusterClaims(namespace string) ClusterClaimNamespaceLister
	ClusterClaimListerExpansion
}

// clusterClaimLister implements the ClusterClaimLister interface.
type clusterClaimLister struct {
	indexer cache.Indexer
}

// NewClusterClaimLister returns a new ClusterClaimLister.
func NewClusterClaimLister(indexer cache.Indexer) ClusterClaimLister {
	return &clusterClaimLister{indexer: indexer}
}

// List lists all ClusterClaims in the indexer.
func (s *clusterClaimLister) List(selector labels.Selector) (ret []*v1.ClusterClaim, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterClaim))
	})
	return ret, err
}

// ClusterClaims returns an object that can list and get ClusterClaims.
func (s *clusterClaimLister) ClusterClaims(namespace string) ClusterClaimNamespaceLister {
	return clusterClaimNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterClaimNamespaceLister helps list and get ClusterClaims.
type ClusterClaimNamespaceLister interface {
	// List lists all ClusterClaims in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.ClusterClaim, err error)
	// Get retrieves the ClusterClaim from the indexer for a given namespace and name.
	Get(name string) (*v1.ClusterClaim, error)
	ClusterClaimNamespaceListerExpansion
}

// clusterClaimNamespaceLister implements the ClusterClaimNamespaceLister
// interface.
type clusterClaimNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterClaims in the indexer for a given namespace.
func (s clusterClaimNamespaceLister) List(selector labels.Selector) (ret []*v1.ClusterClaim, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterClaim))
	})
	return ret, err
}

// Get retrieves the ClusterClaim from the indexer for a given namespace and name.
func (s clusterClaimNamespaceLister) Get(name string) (*v1.ClusterClaim, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clusterclaim"), name)
	}
	return obj.(*v1.ClusterClaim), nil
}
