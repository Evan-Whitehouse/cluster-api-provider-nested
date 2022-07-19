/*
Copyright 2019 The Kubernetes Authors.

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

package secret

import (
	"crypto/rsa"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	vcpki "sigs.k8s.io/cluster-api-provider-nested/virtualcluster/pkg/controller/pki"
	pkiutil "sigs.k8s.io/cluster-api-provider-nested/virtualcluster/pkg/util/pki"
)

const (
	// RootCASecretName is the name for RootCA secret
	RootCASecretName = "root-ca"
	// APIServerCASecretName name of APIServerCA secret
	APIServerCASecretName = "apiserver-ca"
	// ETCDCASecretName name of ETCDCa secret
	ETCDCASecretName = "etcd-ca"
	// FrontProxyCASecretName name of FrontProxyCA secret
	FrontProxyCASecretName = "front-proxy-ca"
	// ControllerManagerSecretName name of ControllerManager kubeconfig secret
	ControllerManagerSecretName = "controller-manager-kubeconfig"
	// AdminSecretName name of secret with kubeconfig for admin
	AdminSecretName = "admin-kubeconfig" // #nosec G101 -- This is a path to secrets
	// ServiceAccountSecretName name of the secret with ServiceAccount rsa
	ServiceAccountSecretName = "serviceaccount-rsa"
)

// RsaKeyToSecret encapsulates rsaKey into a secret object
func RsaKeyToSecret(name, namespace string, rsaKey *rsa.PrivateKey) (*corev1.Secret, error) {
	encodedPubKey, err := pkiutil.EncodePublicKeyPEM(&rsaKey.PublicKey)
	if err != nil {
		return nil, err
	}
	return &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: corev1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Type: corev1.SecretTypeTLS,
		Data: map[string][]byte{
			corev1.TLSCertKey:       encodedPubKey,
			corev1.TLSPrivateKeyKey: vcpki.EncodePrivateKeyPEM(rsaKey),
		},
	}, nil
}

// CrtKeyPairToSecret encapsulates ca/key pair ckp into a secret object
func CrtKeyPairToSecret(name, namespace string, ckp *vcpki.CrtKeyPair) *corev1.Secret {
	return &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: corev1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Type: corev1.SecretTypeTLS,
		Data: map[string][]byte{
			corev1.TLSCertKey:       pkiutil.EncodeCertPEM(ckp.Crt),
			corev1.TLSPrivateKeyKey: vcpki.EncodePrivateKeyPEM(ckp.Key),
		},
	}
}

// KubeconfigToSecret encapsulates kubeconfig cfgContent into a secret object
func KubeconfigToSecret(name, namespace string, cfgContent string) *corev1.Secret {
	return &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: corev1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Type: corev1.SecretTypeOpaque,
		Data: map[string][]byte{
			name: []byte(cfgContent),
		},
	}
}
