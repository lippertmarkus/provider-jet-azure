/*
Copyright 2020 The Crossplane Authors.

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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha1"
	rconfig "github.com/crossplane-contrib/provider-jet-azure/apis/rconfig"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this IPGroup.
func (mg *IPGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LoadBalancer.
func (mg *LoadBalancer) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Subnet.
func (mg *Subnet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualNetworkName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VirtualNetworkNameRef,
		Selector:     mg.Spec.ForProvider.VirtualNetworkNameSelector,
		To: reference.To{
			List:    &VirtualNetworkList{},
			Managed: &VirtualNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualNetworkName")
	}
	mg.Spec.ForProvider.VirtualNetworkName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualNetworkNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubnetNATGatewayAssociation.
func (mg *SubnetNATGatewayAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &SubnetList{},
			Managed: &Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubnetNetworkSecurityGroupAssociation.
func (mg *SubnetNetworkSecurityGroupAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &SubnetList{},
			Managed: &Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubnetRouteTableAssociation.
func (mg *SubnetRouteTableAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &SubnetList{},
			Managed: &Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubnetServiceEndpointStoragePolicy.
func (mg *SubnetServiceEndpointStoragePolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualNetwork.
func (mg *VirtualNetwork) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualNetworkGateway.
func (mg *VirtualNetworkGateway) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.IPConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IPConfiguration[i3].SubnetID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.IPConfiguration[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.IPConfiguration[i3].SubnetIDSelector,
			To: reference.To{
				List:    &SubnetList{},
				Managed: &Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.IPConfiguration[i3].SubnetID")
		}
		mg.Spec.ForProvider.IPConfiguration[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.IPConfiguration[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualNetworkGatewayConnection.
func (mg *VirtualNetworkGatewayConnection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PeerVirtualNetworkGatewayID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PeerVirtualNetworkGatewayIDRef,
		Selector:     mg.Spec.ForProvider.PeerVirtualNetworkGatewayIDSelector,
		To: reference.To{
			List:    &VirtualNetworkGatewayList{},
			Managed: &VirtualNetworkGateway{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PeerVirtualNetworkGatewayID")
	}
	mg.Spec.ForProvider.PeerVirtualNetworkGatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PeerVirtualNetworkGatewayIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualNetworkGatewayID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VirtualNetworkGatewayIDRef,
		Selector:     mg.Spec.ForProvider.VirtualNetworkGatewayIDSelector,
		To: reference.To{
			List:    &VirtualNetworkGatewayList{},
			Managed: &VirtualNetworkGateway{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualNetworkGatewayID")
	}
	mg.Spec.ForProvider.VirtualNetworkGatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualNetworkGatewayIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualNetworkPeering.
func (mg *VirtualNetworkPeering) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RemoteVirtualNetworkID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.RemoteVirtualNetworkIDRef,
		Selector:     mg.Spec.ForProvider.RemoteVirtualNetworkIDSelector,
		To: reference.To{
			List:    &VirtualNetworkList{},
			Managed: &VirtualNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RemoteVirtualNetworkID")
	}
	mg.Spec.ForProvider.RemoteVirtualNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RemoteVirtualNetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualNetworkName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VirtualNetworkNameRef,
		Selector:     mg.Spec.ForProvider.VirtualNetworkNameSelector,
		To: reference.To{
			List:    &VirtualNetworkList{},
			Managed: &VirtualNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualNetworkName")
	}
	mg.Spec.ForProvider.VirtualNetworkName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualNetworkNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualWAN.
func (mg *VirtualWAN) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}
