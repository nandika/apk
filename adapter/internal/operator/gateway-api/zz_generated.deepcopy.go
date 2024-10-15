//go:build !ignore_autogenerated

/*
 *  Copyright (c) 2023, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

// Code generated by controller-gen. DO NOT EDIT.

package gatewayapi

import (
	dpv1alpha2 "github.com/wso2/apk/common-go-libs/apis/dp/v1alpha2"
	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/gateway-api/apis/v1"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"
	"sigs.k8s.io/gateway-api/apis/v1alpha3"
	"sigs.k8s.io/gateway-api/apis/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	if in.GatewayClass != nil {
		in, out := &in.GatewayClass, &out.GatewayClass
		*out = new(v1.GatewayClass)
		(*in).DeepCopyInto(*out)
	}
	if in.Gateways != nil {
		in, out := &in.Gateways, &out.Gateways
		*out = make([]*v1.Gateway, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.Gateway)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.HTTPRoutes != nil {
		in, out := &in.HTTPRoutes, &out.HTTPRoutes
		*out = make([]*v1.HTTPRoute, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.HTTPRoute)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.GRPCRoutes != nil {
		in, out := &in.GRPCRoutes, &out.GRPCRoutes
		*out = make([]*v1.GRPCRoute, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.GRPCRoute)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.TLSRoutes != nil {
		in, out := &in.TLSRoutes, &out.TLSRoutes
		*out = make([]*v1alpha2.TLSRoute, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1alpha2.TLSRoute)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.TCPRoutes != nil {
		in, out := &in.TCPRoutes, &out.TCPRoutes
		*out = make([]*v1alpha2.TCPRoute, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1alpha2.TCPRoute)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.UDPRoutes != nil {
		in, out := &in.UDPRoutes, &out.UDPRoutes
		*out = make([]*v1alpha2.UDPRoute, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1alpha2.UDPRoute)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ReferenceGrants != nil {
		in, out := &in.ReferenceGrants, &out.ReferenceGrants
		*out = make([]*v1beta1.ReferenceGrant, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1beta1.ReferenceGrant)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]*corev1.Namespace, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1.Namespace)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]*corev1.Service, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1.Service)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Backends != nil {
		in, out := &in.Backends, &out.Backends
		*out = make([]*dpv1alpha2.Backend, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(dpv1alpha2.Backend)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.EndpointSlices != nil {
		in, out := &in.EndpointSlices, &out.EndpointSlices
		*out = make([]*discoveryv1.EndpointSlice, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(discoveryv1.EndpointSlice)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]*corev1.Secret, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1.Secret)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]*corev1.ConfigMap, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1.ConfigMap)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ExtensionRefFilters != nil {
		in, out := &in.ExtensionRefFilters, &out.ExtensionRefFilters
		*out = make([]unstructured.Unstructured, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BackendTLSPolicies != nil {
		in, out := &in.BackendTLSPolicies, &out.BackendTLSPolicies
		*out = make([]*v1alpha3.BackendTLSPolicy, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1alpha3.BackendTLSPolicy)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.APIs != nil {
		in, out := &in.APIs, &out.APIs
		*out = make([]*dpv1alpha2.API, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(dpv1alpha2.API)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}
