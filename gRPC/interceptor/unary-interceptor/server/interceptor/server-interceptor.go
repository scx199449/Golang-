package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 前置处理逻辑
	ctx = context.WithValue(ctx, "role", "user")

	// 继续处理请求
	m, err := handler(ctx, req)

	//后置处理逻辑
	fmt.Println(m)

	return m, err
}
