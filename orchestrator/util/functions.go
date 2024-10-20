package util

import (
	"net"
	"strings"
)

var localIP string

func Map[A any, B any](vs []A, f func(A) B) []B {
	vsm := make([]B, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func GetLocalIP() string {
	if localIP != "" {
		return localIP
	}

	addresses, err := net.InterfaceAddrs()
	if err != nil {
		localIP = ""
		return localIP
	}

	for _, addr := range addresses {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIP = strings.Split(ipnet.IP.String(), "/")[0]
				break
			}
		}
	}

	return localIP
}

func GetApiUrl() string {
	localIP := GetLocalIP()
	return "http://" + localIP + ":9091"
}

func Filter[T any](input []T, filterFunction func(T) bool) []T {
	filtered := make([]T, 0)

	for _, v := range input {
		if filterFunction(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func FilterNotNull[T any](input []*T) []T {
	return Map(Filter(input, func(v *T) bool {
		return v != nil
	}), func(v *T) T {
		return *v
	})
}
