package server

import (
	"context"
	"fmt"
	"github.com/ZQCard/kratos-base-project/pkg/middleware/jwt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	v1 "github.com/ZQCard/kratos-base-project/api/admin/v1"
	"github.com/ZQCard/kratos-base-project/app/admin/service/internal/conf"
	"github.com/ZQCard/kratos-base-project/app/admin/service/internal/service"
)

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/shop.interface.v1.ShopInterface/Login"] = struct{}{}
	whiteList["/shop.interface.v1.ShopInterface/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

func getOperation() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				fmt.Println(tr.Operation())
			}
			return handler(ctx, req)
		}
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, service *service.AdminInterface, tp *tracesdk.TracerProvider,logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(
				tracing.WithTracerProvider(tp)),
			// 日志记录
			logging.Server(logger),
			// 对于需要登录的路由进行jwt中间件验证
			selector.Server(jwt.AuthMiddleware()).
				//Prefix("/").
				Path(

				).Build(),
			getOperation(),
		),
		// 跨域设置
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterAdminHTTPServer(srv, service)
	return srv
}
