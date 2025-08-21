// SPDX-License-Identifier: Apache-2.0
// Copyright 2022 Authors of KubeDig

// Package cert is responsible for generating client and server certificates using KubeDig cert pkg.
package cert

import (
	"fmt"
	"time"

	certutil "github.com/zfz-725/KubeDig/KubeDig/cert"
	"github.com/zfz-725/KubeDig/pkg/KubeDigOperator/common"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

var CACert *certutil.CertKeyPair

func GetRelayServiceDnsName(namespace string) []string {
	res := []string{
		fmt.Sprintf("kubedig.%s", namespace),
		fmt.Sprintf("kubedig.%s.svc", namespace),
		fmt.Sprintf("kubedig.%s.svc.cluster.local", namespace),
	}
	return res
}

func GenerateKubeDigCACert() (certutil.CertBytes, error) {
	caCertConfig := certutil.DefaultKubeDigCAConfig
	caCertConfig.NotAfter = time.Now().AddDate(1, 0, 0) // valid for one year
	crtBytes, err := certutil.GenerateCA(&caCertConfig)
	if err != nil {
		klog.Errorf("error generating ca cert: %s", err)
		return certutil.CertBytes{}, err
	}
	crtKeyPair, err := certutil.GetCertKeyPairFromCertBytes(crtBytes)
	if err != nil {
		return certutil.CertBytes{}, err
	}
	CACert = crtKeyPair
	return *crtBytes, nil
}

func GenerateKubeDigClientCert() (certutil.CertBytes, error) {
	if CACert == nil {
		_, err := GenerateKubeDigCACert()
		if err != nil {
			return certutil.CertBytes{}, err
		}
	}
	certCfg := certutil.DefaultKubeDigClientConfig
	certCfg.NotAfter = time.Now().AddDate(1, 0, 0)
	crtBytes, err := certutil.GenerateSelfSignedCert(CACert, &certCfg)
	if err != nil {
		klog.Errorf("error generating kubedig client cert: %s", err)
		return certutil.CertBytes{}, err
	}
	return *crtBytes, nil
}

func GenerateKubeDigRelayCert() (certutil.CertBytes, error) {
	if CACert == nil {
		_, err := GenerateKubeDigCACert()
		if err != nil {
			return certutil.CertBytes{}, err
		}
	}
	certCfg := certutil.DefaultKubeDigServerConfig
	certCfg.NotAfter = time.Now().AddDate(1, 0, 0)
	dnsList := GetRelayServiceDnsName(common.Namespace)
	dnsList = append(dnsList, common.ExtraDnsNames...)
	certCfg.DNS = append(certCfg.DNS, dnsList...)
	certCfg.IPs = append(certCfg.IPs, common.ExtraIpAddresses...)
	klog.Infof("relay cert extUsage: %v", certCfg.ExtKeyUsage)
	crtBytes, err := certutil.GenerateSelfSignedCert(CACert, &certCfg)
	if err != nil {
		klog.Errorf("error generating kubedig relay cert: %s", err)
		return certutil.CertBytes{}, err
	}
	return *crtBytes, nil
}

func GetCertSecret(crt *[]byte, key *[]byte, name, namespace string, labels *map[string]string) *corev1.Secret {

	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    *labels,
		},
		Type: corev1.SecretTypeTLS,
		Data: map[string][]byte{
			"tls.crt": *crt,
			"tls.key": *key,
		},
	}
}

func GetCertWithCaSecret(ca, crt, key *[]byte, name, namespace string, labels *map[string]string) *corev1.Secret {

	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    *labels,
		},
		Type: corev1.SecretTypeTLS,
		Data: map[string][]byte{
			"tls.crt": *crt,
			"tls.key": *key,
			"ca.crt":  *ca,
		},
	}

}

func GetAllTlsCertSecrets() ([]*corev1.Secret, error) {
	fmt.Println("Prepairing all the tls secrets")
	secrets := []*corev1.Secret{}
	var certGenErr, err error
	var kaCaCert, kaClientCert, kaRelayCert certutil.CertBytes
	for i := 0; i < 3; i++ {
		// generate kubedig-ca certs
		kaCaCert, err = GenerateKubeDigCACert()
		if err != nil {
			certGenErr = err
		}
		// generate kubedig-client certs
		kaClientCert, err = GenerateKubeDigClientCert()
		if err != nil {
			certGenErr = err
		}
		// generate kubedig-relay certs
		kaRelayCert, err = GenerateKubeDigRelayCert()
		if err != nil {
			certGenErr = err
		}
		if certGenErr == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
	if certGenErr != nil {
		return secrets, certGenErr
	} else {
		// create secrets
		secrets = append(secrets,
			GetCertSecret(&kaCaCert.Crt, &kaCaCert.Key, common.KubeDigCaSecretName, common.Namespace, &map[string]string{"kubedig-app": common.KubeDigCaSecretName}),
			GetCertWithCaSecret(&kaCaCert.Crt, &kaClientCert.Crt, &kaClientCert.Key, common.KubeDigClientSecretName, common.Namespace, &map[string]string{"kubedig-app": common.KubeDigClientSecretName}),
			GetCertWithCaSecret(&kaCaCert.Crt, &kaRelayCert.Crt, &kaRelayCert.Key, common.KubeDigRelayServerSecretName, common.Namespace, &map[string]string{"kubedig-app": common.KubeDigRelayServerSecretName}))
	}
	return secrets, nil
}
