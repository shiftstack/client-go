// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfigurations

import (
	v1 "github.com/openshift/api/user/v1"
	userv1 "github.com/openshift/client-go/user/applyconfigurations/user/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=user.openshift.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("Group"):
		return &userv1.GroupApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Identity"):
		return &userv1.IdentityApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("User"):
		return &userv1.UserApplyConfiguration{}

	}
	return nil
}