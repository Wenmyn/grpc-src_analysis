/*
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package passthrough implements a pass-through resolver. It sends the target
// name without scheme back to gRPC as resolved address.
package passthrough

import "google.golang.org/grpc/resolver"

const scheme = "passthrough"

//默认自带的解析构建器
type passthroughBuilder struct{}

func (*passthroughBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	/*
		注：用户设置的链接地址，其实就是后端服务的真实地址
		不需要再去远程访问获取后端服务地址列表了。
	*/
	r := &passthroughResolver{
		target: target,	//就是后端对应的服务器地址
		cc:     cc,
	}

	//将获取到的后端服务地址列表信息更新到解析器里
	r.start()
	return r, nil
}

func (*passthroughBuilder) Scheme() string {
	return scheme
}

type passthroughResolver struct {
	target resolver.Target
	cc     resolver.ClientConn
}

func (r *passthroughResolver) start() {
	r.cc.UpdateState(resolver.State{Addresses: []resolver.Address{{Addr: r.target.Endpoint}}})
}

func (*passthroughResolver) ResolveNow(o resolver.ResolveNowOptions) {}

func (*passthroughResolver) Close() {}

func init() {
	resolver.Register(&passthroughBuilder{})
}
