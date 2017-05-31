// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
    "errors"

    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/lumi/pkg/resource"
    "github.com/pulumi/lumi/pkg/tokens"
    "github.com/pulumi/lumi/pkg/util/contract"
    "github.com/pulumi/lumi/pkg/util/mapper"
    "github.com/pulumi/lumi/sdk/go/pkg/lumirpc"
)

/* RPC stubs for SecurityGroupEgress resource provider */

// SecurityGroupEgressToken is the type token corresponding to the SecurityGroupEgress package type.
const SecurityGroupEgressToken = tokens.Type("aws:ec2/securityGroupEgress:SecurityGroupEgress")

// SecurityGroupEgressProviderOps is a pluggable interface for SecurityGroupEgress-related management functionality.
type SecurityGroupEgressProviderOps interface {
    Check(ctx context.Context, obj *SecurityGroupEgress) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *SecurityGroupEgress) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*SecurityGroupEgress, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *SecurityGroupEgress, new *SecurityGroupEgress, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *SecurityGroupEgress, new *SecurityGroupEgress, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// SecurityGroupEgressProvider is a dynamic gRPC-based plugin for managing SecurityGroupEgress resources.
type SecurityGroupEgressProvider struct {
    ops SecurityGroupEgressProviderOps
}

// NewSecurityGroupEgressProvider allocates a resource provider that delegates to a ops instance.
func NewSecurityGroupEgressProvider(ops SecurityGroupEgressProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &SecurityGroupEgressProvider{ops: ops}
}

func (p *SecurityGroupEgressProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupEgressToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr == nil || len(decerr.Failures()) == 0 {
        failures, err := p.ops.Check(ctx, obj)
        if err != nil {
            return nil, err
        }
        if len(failures) > 0 {
            decerr = mapper.NewDecodeErr(failures)
        }
    }
    return resource.NewCheckResponse(decerr), nil
}

func (p *SecurityGroupEgressProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupEgressToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == "" {
        if req.Unknowns[SecurityGroupEgress_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: obj.Name}, nil
}

func (p *SecurityGroupEgressProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupEgressToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    id, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &lumirpc.CreateResponse{Id: string(id)}, nil
}

func (p *SecurityGroupEgressProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupEgressToken))
    id := resource.ID(req.GetId())
    obj, err := p.ops.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return &lumirpc.GetResponse{
        Properties: resource.MarshalProperties(
            nil, resource.NewPropertyMap(obj), resource.MarshalOptions{}),
    }, nil
}

func (p *SecurityGroupEgressProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupEgressToken))
    id := resource.ID(req.GetId())
    old, oldprops, decerr := p.Unmarshal(req.GetOlds())
    if decerr != nil {
        return nil, decerr
    }
    new, newprops, decerr := p.Unmarshal(req.GetNews())
    if decerr != nil {
        return nil, decerr
    }
    var replaces []string
    diff := oldprops.Diff(newprops)
    if diff != nil {
        if diff.Changed("name") {
            replaces = append(replaces, "name")
        }
        if diff.Changed("fromPort") {
            replaces = append(replaces, "fromPort")
        }
        if diff.Changed("group") {
            replaces = append(replaces, "group")
        }
        if diff.Changed("ipProtocol") {
            replaces = append(replaces, "ipProtocol")
        }
        if diff.Changed("toPort") {
            replaces = append(replaces, "toPort")
        }
        if diff.Changed("cidrIp") {
            replaces = append(replaces, "cidrIp")
        }
        if diff.Changed("cidrIpv6") {
            replaces = append(replaces, "cidrIpv6")
        }
        if diff.Changed("destinationPrefixListId") {
            replaces = append(replaces, "destinationPrefixListId")
        }
        if diff.Changed("destinationSecurityGroup") {
            replaces = append(replaces, "destinationSecurityGroup")
        }
    }
    more, err := p.ops.InspectChange(ctx, id, old, new, diff)
    if err != nil {
        return nil, err
    }
    return &lumirpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *SecurityGroupEgressProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(SecurityGroupEgressToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, id, old, new, diff); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *SecurityGroupEgressProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(SecurityGroupEgressToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *SecurityGroupEgressProvider) Unmarshal(
    v *pbstruct.Struct) (*SecurityGroupEgress, resource.PropertyMap, mapper.DecodeError) {
    var obj SecurityGroupEgress
    props := resource.UnmarshalProperties(nil, v, resource.MarshalOptions{})
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable SecurityGroupEgress structure(s) */

// SecurityGroupEgress is a marshalable representation of its corresponding IDL type.
type SecurityGroupEgress struct {
    Name string `json:"name"`
    FromPort float64 `json:"fromPort"`
    Group resource.ID `json:"group"`
    IPProtocol string `json:"ipProtocol"`
    ToPort float64 `json:"toPort"`
    CIDRIP *string `json:"cidrIp,omitempty"`
    CIDRIPv6 *string `json:"cidrIpv6,omitempty"`
    DestinationPrefixListId *string `json:"destinationPrefixListId,omitempty"`
    DestinationSecurityGroup *resource.ID `json:"destinationSecurityGroup,omitempty"`
}

// SecurityGroupEgress's properties have constants to make dealing with diffs and property bags easier.
const (
    SecurityGroupEgress_Name = "name"
    SecurityGroupEgress_FromPort = "fromPort"
    SecurityGroupEgress_Group = "group"
    SecurityGroupEgress_IPProtocol = "ipProtocol"
    SecurityGroupEgress_ToPort = "toPort"
    SecurityGroupEgress_CIDRIP = "cidrIp"
    SecurityGroupEgress_CIDRIPv6 = "cidrIpv6"
    SecurityGroupEgress_DestinationPrefixListId = "destinationPrefixListId"
    SecurityGroupEgress_DestinationSecurityGroup = "destinationSecurityGroup"
)


