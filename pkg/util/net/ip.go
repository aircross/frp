// Copyright 2023 The aircross/frp Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package net

import (
	"fmt"
	"net"
)

// IsIPAddress 判断输入的字符串是否是一个合法的 IP 地址
func IsIPAddress(input string) bool {
	ip := net.ParseIP(input)
	return ip != nil
}

// GetIPAddress 获取域名对应的最优 IP 地址并作为结果返回
// 优先返回 IPv4 地址，如果没有 IPv4 地址再返回 IPv6 地址
func GetDomainIP(domain string) (string, error) {
	if IsIPAddress(domain) {
		return domain, nil
	}
	ips, err := net.LookupIP(domain)
	if err != nil {
		return "", err
	}

	// 遍历 IP 地址，优先返回 IPv4 地址
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			return ipv4.String(), nil
		}
	}

	// 如果没有找到 IPv4 地址，则返回第一个 IPv6 地址
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 == nil {
			fmt.Println(ip.String())
			return ip.String(), nil
		}
	}

	return "", fmt.Errorf("no IP addresses found for domain: %s", domain)
}
