/*
Copyright 2022 The Knative Authors

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	scheme "knative.dev/serving/pkg/client/certmanager/clientset/versioned/scheme"
)

// IssuersGetter has a method to return a IssuerInterface.
// A group's client should implement this interface.
type IssuersGetter interface {
	Issuers(namespace string) IssuerInterface
}

// IssuerInterface has methods to work with Issuer resources.
type IssuerInterface interface {
	Create(ctx context.Context, issuer *v1.Issuer, opts metav1.CreateOptions) (*v1.Issuer, error)
	Update(ctx context.Context, issuer *v1.Issuer, opts metav1.UpdateOptions) (*v1.Issuer, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, issuer *v1.Issuer, opts metav1.UpdateOptions) (*v1.Issuer, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Issuer, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.IssuerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Issuer, err error)
	IssuerExpansion
}

// issuers implements IssuerInterface
type issuers struct {
	*gentype.ClientWithList[*v1.Issuer, *v1.IssuerList]
}

// newIssuers returns a Issuers
func newIssuers(c *CertmanagerV1Client, namespace string) *issuers {
	return &issuers{
		gentype.NewClientWithList[*v1.Issuer, *v1.IssuerList](
			"issuers",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.Issuer { return &v1.Issuer{} },
			func() *v1.IssuerList { return &v1.IssuerList{} }),
	}
}
