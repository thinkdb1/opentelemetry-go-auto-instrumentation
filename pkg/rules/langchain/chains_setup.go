// Copyright (c) 2024 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package langchain

import (
	"context"
	_ "unsafe"

	"github.com/alibaba/opentelemetry-go-auto-instrumentation/pkg/api"
	"github.com/tmc/langchaingo/chains"
)

//go:linkname callChainOnEnter github.com/tmc/langchaingo/chains.callChainOnEnter
func callChainOnEnter(call api.CallContext, ctx context.Context,
	c chains.Chain,
	fullValues map[string]any,
	options ...chains.ChainCallOption) {
	if !langChainEnabler.Enable() {
		return
	}
	request := langChainRequest{
		operationName: MChains,
		system:        "langchain",
	}
	langCtx := langChainCommonInstrument.Start(ctx, request)
	data := make(map[string]interface{})
	data["ctx"] = langCtx
	call.SetData(data)
}

//go:linkname callChainOnExit github.com/tmc/langchaingo/chains.callChainOnExit
func callChainOnExit(call api.CallContext, v map[string]any, err error) {
	if !langChainEnabler.Enable() {
		return
	}
	data := call.GetData().(map[string]interface{})
	ctx, ok := data["ctx"].(context.Context)
	if !ok {
		return
	}
	request := langChainRequest{
		operationName: MChains,
		system:        "langchain",
	}
	langChainCommonInstrument.End(ctx, request, nil, err)
}
