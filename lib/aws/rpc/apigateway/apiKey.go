// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

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

/* RPC stubs for APIKey resource provider */

// APIKeyToken is the type token corresponding to the APIKey package type.
const APIKeyToken = tokens.Type("aws:apigateway/apiKey:APIKey")

// APIKeyProviderOps is a pluggable interface for APIKey-related management functionality.
type APIKeyProviderOps interface {
    Check(ctx context.Context, obj *APIKey) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *APIKey) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*APIKey, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *APIKey, new *APIKey, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *APIKey, new *APIKey, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// APIKeyProvider is a dynamic gRPC-based plugin for managing APIKey resources.
type APIKeyProvider struct {
    ops APIKeyProviderOps
}

// NewAPIKeyProvider allocates a resource provider that delegates to a ops instance.
func NewAPIKeyProvider(ops APIKeyProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &APIKeyProvider{ops: ops}
}

func (p *APIKeyProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(APIKeyToken))
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

func (p *APIKeyProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(APIKeyToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == "" {
        if req.Unknowns[APIKey_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: obj.Name}, nil
}

func (p *APIKeyProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(APIKeyToken))
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

func (p *APIKeyProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(APIKeyToken))
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

func (p *APIKeyProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(APIKeyToken))
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
        if diff.Changed("keyName") {
            replaces = append(replaces, "keyName")
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

func (p *APIKeyProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(APIKeyToken))
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

func (p *APIKeyProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(APIKeyToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *APIKeyProvider) Unmarshal(
    v *pbstruct.Struct) (*APIKey, resource.PropertyMap, mapper.DecodeError) {
    var obj APIKey
    props := resource.UnmarshalProperties(nil, v, resource.MarshalOptions{})
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable APIKey structure(s) */

// APIKey is a marshalable representation of its corresponding IDL type.
type APIKey struct {
    Name string `json:"name"`
    KeyName *string `json:"keyName,omitempty"`
    Description *string `json:"description,omitempty"`
    Enabled *bool `json:"enabled,omitempty"`
    StageKeys *StageKey `json:"stageKeys,omitempty"`
}

// APIKey's properties have constants to make dealing with diffs and property bags easier.
const (
    APIKey_Name = "name"
    APIKey_KeyName = "keyName"
    APIKey_Description = "description"
    APIKey_Enabled = "enabled"
    APIKey_StageKeys = "stageKeys"
)

/* Marshalable StageKey structure(s) */

// StageKey is a marshalable representation of its corresponding IDL type.
type StageKey struct {
    RestAPI *resource.ID `json:"restAPI,omitempty"`
    Stage *resource.ID `json:"stage,omitempty"`
}

// StageKey's properties have constants to make dealing with diffs and property bags easier.
const (
    StageKey_RestAPI = "restAPI"
    StageKey_Stage = "stage"
)


