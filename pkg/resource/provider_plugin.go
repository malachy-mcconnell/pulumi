// Licensed to Pulumi Corporation ("Pulumi") under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// Pulumi licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/pulumi/lumi/pkg/tokens"
	"github.com/pulumi/lumi/pkg/util/contract"
	"github.com/pulumi/lumi/pkg/workspace"
	"github.com/pulumi/lumi/sdk/go/pkg/lumirpc"
)

const providerPrefix = "lumi-resource"

// provider reflects a resource plugin, loaded dynamically for a single package.
type provider struct {
	ctx    *Context
	pkg    tokens.Package
	plug   *plugin
	client lumirpc.ResourceProviderClient
}

// NewProvider attempts to bind to a given package's resource plugin and then creates a gRPC connection to it.  If the
// plugin could not be found, or an error occurs while creating the child process, an error is returned.
func NewProvider(ctx *Context, pkg tokens.Package) (Provider, error) {
	// Setup the search paths; first, the naked name (found on the PATH); next, the fully qualified name.
	srvexe := providerPrefix + "-" + strings.Replace(string(pkg), tokens.QNameDelimiter, "_", -1)
	paths := []string{
		srvexe, // naked PATH.
		filepath.Join(
			workspace.InstallRoot(), workspace.InstallRootLibdir, string(pkg), srvexe), // qualified name.
	}

	// Now go ahead and attempt to load the plugin.
	plug, err := newPlugin(ctx, paths, fmt.Sprintf("resource[%v]", pkg))
	if err != nil {
		return nil, err
	}

	return &provider{
		ctx:    ctx,
		pkg:    pkg,
		plug:   plug,
		client: lumirpc.NewResourceProviderClient(plug.Conn),
	}, nil
}

func (p *provider) Pkg() tokens.Package { return p.pkg }

// Check validates that the given property bag is valid for a resource of the given type.
func (p *provider) Check(t tokens.Type, props PropertyMap) ([]CheckFailure, error) {
	glog.V(7).Infof("resource[%v].Check(t=%v,#props=%v) executing", p.pkg, t, len(props))
	pstr, unks := MarshalPropertiesWithUnknowns(p.ctx, props, MarshalOptions{
		OldURNs:      true, // permit old URNs, since this is pre-update.
		RawResources: true, // pre-create, IDs won't be ready, just ship over the URNs.
	})
	req := &lumirpc.CheckRequest{
		Type:       string(t),
		Properties: pstr,
		Unknowns:   unks,
	}

	resp, err := p.client.Check(p.ctx.Request(), req)
	if err != nil {
		glog.V(7).Infof("resource[%v].Check(t=%v,...) failed: err=%v", p.pkg, t, err)
		return nil, err
	}

	var failures []CheckFailure
	for _, failure := range resp.GetFailures() {
		failures = append(failures, CheckFailure{PropertyKey(failure.Property), failure.Reason})
	}
	glog.V(7).Infof("resource[%v].Check(t=%v,...) success: failures=#%v", p.pkg, t, len(failures))
	return failures, nil
}

// Name names a given resource.
func (p *provider) Name(t tokens.Type, props PropertyMap) (tokens.QName, error) {
	glog.V(7).Infof("resource[%v].Name(t=%v,#props=%v) executing", p.pkg, t, len(props))
	pstr, unks := MarshalPropertiesWithUnknowns(p.ctx, props, MarshalOptions{
		OldURNs:      true, // permit old URNs, since this is pre-update.
		RawResources: true, // pre-create, IDs won't be ready, just ship over the URNs.
	})
	req := &lumirpc.NameRequest{
		Type:       string(t),
		Properties: pstr,
		Unknowns:   unks,
	}

	resp, err := p.client.Name(p.ctx.Request(), req)
	if err != nil {
		glog.V(7).Infof("resource[%v].Name(t=%v,...) failed: err=%v", p.pkg, t, err)
		return "", err
	}

	name := tokens.QName(resp.GetName())
	glog.V(7).Infof("resource[%v].Name(t=%v,...) success: name=%v", p.pkg, t, name)
	return name, nil
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.
func (p *provider) Create(t tokens.Type, props PropertyMap) (ID, State, error) {
	glog.V(7).Infof("resource[%v].Create(t=%v,#props=%v) executing", p.pkg, t, len(props))
	req := &lumirpc.CreateRequest{
		Type:       string(t),
		Properties: MarshalProperties(p.ctx, props, MarshalOptions{}),
	}

	resp, err := p.client.Create(p.ctx.Request(), req)
	if err != nil {
		glog.V(7).Infof("resource[%v].Create(t=%v,...) failed: err=%v", p.pkg, t, err)
		return ID(""), StateUnknown, err
	}

	id := ID(resp.GetId())
	glog.V(7).Infof("resource[%v].Create(t=%v,...) success: id=%v", p.pkg, t, id)
	if id == "" {
		return id, StateUnknown,
			errors.Errorf("plugin for package '%v' returned empty ID from create '%v'", p.pkg, t)
	}
	return id, StateOK, nil
}

// Get reads the instance state identified by id/t, and copies into a bag of properties props.
func (p *provider) Get(id ID, t tokens.Type, props PropertyMap) error {
	glog.V(7).Infof("resource[%v].Get(id=%v,t=%v) executing", p.pkg, id, t)
	req := &lumirpc.GetRequest{
		Id:   string(id),
		Type: string(t),
	}

	resp, err := p.client.Get(p.ctx.Request(), req)
	if err != nil {
		glog.V(7).Infof("resource[%v].Get(id=%v,t=%v) failed: err=%v", p.pkg, id, t, err)
		return err
	}

	UnmarshalPropertiesInto(p.ctx, resp.GetProperties(), props, MarshalOptions{})
	glog.V(7).Infof("resource[%v].Get(id=%v,t=%v) success: #props=%v", p.pkg, t, id, len(props))
	return nil
}

// InspectChange checks what impacts a hypothetical update will have on the resource's properties.
func (p *provider) InspectChange(id ID, t tokens.Type,
	olds PropertyMap, news PropertyMap) ([]string, PropertyMap, error) {
	contract.Requiref(id != "", "id", "not empty")
	contract.Requiref(t != "", "t", "not empty")

	glog.V(7).Infof("resource[%v].InspectChange(id=%v,t=%v,#olds=%v,#news=%v) executing",
		p.pkg, id, t, len(olds), len(news))
	newpstr, newunks := MarshalPropertiesWithUnknowns(p.ctx, news, MarshalOptions{
		RawResources: true, // pre-change, IDs won't be ready, ship over URNs.
	})
	req := &lumirpc.InspectChangeRequest{
		Id:   string(id),
		Type: string(t),
		Olds: MarshalProperties(p.ctx, olds, MarshalOptions{
			RawResources: true, // just leave these as-is, so they match the news.
		}),
		News:     newpstr,
		Unknowns: newunks,
	}

	resp, err := p.client.InspectChange(p.ctx.Request(), req)
	if err != nil {
		glog.V(7).Infof("resource[%v].InspectChange(id=%v,t=%v,...) failed: %v", p.pkg, id, t, err)
		return nil, nil, err
	}

	replaces := resp.GetReplaces()
	changes := UnmarshalProperties(p.ctx, resp.GetChanges(), MarshalOptions{RawResources: true})
	glog.V(7).Infof("resource[%v].Update(id=%v,t=%v,...) success: #replaces=%v #changes=%v",
		p.pkg, id, t, len(replaces), len(changes))
	return replaces, changes, nil
}

// Update updates an existing resource with new values.
func (p *provider) Update(id ID, t tokens.Type, olds PropertyMap, news PropertyMap) (State, error) {
	contract.Requiref(id != "", "id", "not empty")
	contract.Requiref(t != "", "t", "not empty")

	glog.V(7).Infof("resource[%v].Update(id=%v,t=%v,#olds=%v,#news=%v) executing",
		p.pkg, id, t, len(olds), len(news))
	req := &lumirpc.UpdateRequest{
		Id:   string(id),
		Type: string(t),
		Olds: MarshalProperties(p.ctx, olds, MarshalOptions{
			OldURNs: true, // permit old URNs since these are the old values.
		}),
		News: MarshalProperties(p.ctx, news, MarshalOptions{}),
	}

	_, err := p.client.Update(p.ctx.Request(), req)
	if err != nil {
		glog.V(7).Infof("resource[%v].Update(id=%v,t=%v,...) failed: %v", p.pkg, id, t, err)
		return StateUnknown, err
	}

	glog.V(7).Infof("resource[%v].Update(id=%v,t=%v,...) success", p.pkg, id, t)
	return StateOK, nil
}

// Delete tears down an existing resource.
func (p *provider) Delete(id ID, t tokens.Type) (State, error) {
	contract.Requiref(id != "", "id", "not empty")
	contract.Requiref(t != "", "t", "not empty")

	glog.V(7).Infof("resource[%v].Delete(id=%v,t=%v) executing", p.pkg, id, t)
	req := &lumirpc.DeleteRequest{
		Id:   string(id),
		Type: string(t),
	}

	if _, err := p.client.Delete(p.ctx.Request(), req); err != nil {
		glog.V(7).Infof("resource[%v].Delete(id=%v,t=%v) failed: %v", p.pkg, id, t, err)
		return StateUnknown, err
	}

	glog.V(7).Infof("resource[%v].Delete(id=%v,t=%v) success", p.pkg, id, t)
	return StateOK, nil
}

// Close tears down the underlying plugin RPC connection and process.
func (p *provider) Close() error {
	return p.plug.Close()
}
