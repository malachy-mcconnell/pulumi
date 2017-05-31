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

/* RPC stubs for RouteTable resource provider */

// RouteTableToken is the type token corresponding to the RouteTable package type.
const RouteTableToken = tokens.Type("aws:ec2/routeTable:RouteTable")

// RouteTableProviderOps is a pluggable interface for RouteTable-related management functionality.
type RouteTableProviderOps interface {
    Check(ctx context.Context, obj *RouteTable) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *RouteTable) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*RouteTable, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *RouteTable, new *RouteTable, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *RouteTable, new *RouteTable, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// RouteTableProvider is a dynamic gRPC-based plugin for managing RouteTable resources.
type RouteTableProvider struct {
    ops RouteTableProviderOps
}

// NewRouteTableProvider allocates a resource provider that delegates to a ops instance.
func NewRouteTableProvider(ops RouteTableProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &RouteTableProvider{ops: ops}
}

func (p *RouteTableProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(RouteTableToken))
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

func (p *RouteTableProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(RouteTableToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == "" {
        if req.Unknowns[RouteTable_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: obj.Name}, nil
}

func (p *RouteTableProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(RouteTableToken))
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

func (p *RouteTableProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(RouteTableToken))
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

func (p *RouteTableProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(RouteTableToken))
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
        if diff.Changed("vpc") {
            replaces = append(replaces, "vpc")
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

func (p *RouteTableProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(RouteTableToken))
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

func (p *RouteTableProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(RouteTableToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *RouteTableProvider) Unmarshal(
    v *pbstruct.Struct) (*RouteTable, resource.PropertyMap, mapper.DecodeError) {
    var obj RouteTable
    props := resource.UnmarshalProperties(nil, v, resource.MarshalOptions{})
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable RouteTable structure(s) */

// RouteTable is a marshalable representation of its corresponding IDL type.
type RouteTable struct {
    Name string `json:"name"`
    VPC resource.ID `json:"vpc"`
}

// RouteTable's properties have constants to make dealing with diffs and property bags easier.
const (
    RouteTable_Name = "name"
    RouteTable_VPC = "vpc"
)


