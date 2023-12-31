/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha11 "github.com/kingsoftcloud/provider-ksyun/apis/slb/v1alpha1"
	v1alpha1 "github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Alb.
func (mg *Alb) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1alpha1.VPCList{},
			Managed: &v1alpha1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this AlbListener.
func (mg *AlbListener) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ALBID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ALBIDRef,
		Selector:     mg.Spec.ForProvider.ALBIDSelector,
		To: reference.To{
			List:    &AlbList{},
			Managed: &Alb{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ALBID")
	}
	mg.Spec.ForProvider.ALBID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ALBIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this AlbListenerCertGroup.
func (mg *AlbListenerCertGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ALBListenerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ALBListenerIDRef,
		Selector:     mg.Spec.ForProvider.ALBListenerIDSelector,
		To: reference.To{
			List:    &AlbListenerList{},
			Managed: &AlbListener{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ALBListenerID")
	}
	mg.Spec.ForProvider.ALBListenerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ALBListenerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this AlbRuleGroup.
func (mg *AlbRuleGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ALBListenerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ALBListenerIDRef,
		Selector:     mg.Spec.ForProvider.ALBListenerIDSelector,
		To: reference.To{
			List:    &AlbListenerList{},
			Managed: &AlbListener{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ALBListenerID")
	}
	mg.Spec.ForProvider.ALBListenerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ALBListenerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BackendServerGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BackendServerGroupIDRef,
		Selector:     mg.Spec.ForProvider.BackendServerGroupIDSelector,
		To: reference.To{
			List:    &v1alpha11.LbBackendServerGroupList{},
			Managed: &v1alpha11.LbBackendServerGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BackendServerGroupID")
	}
	mg.Spec.ForProvider.BackendServerGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackendServerGroupIDRef = rsp.ResolvedReference

	return nil
}
