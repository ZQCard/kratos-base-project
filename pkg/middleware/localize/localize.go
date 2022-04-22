package localize

import (
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type localizerKey struct{}

func I18N() middleware.Middleware {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	// 指定翻译文件路径
	bundle.MustLoadMessageFile("../../internal/pkg/middleware/localize/active.zh.toml")
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			reply, err = handler(ctx, req)
			se := errors.FromError(err)
			if err != nil {
				return nil, errors.Newf(400, "请求参数错误", se.Message)
			}

			fmt.Println(se.Message)

			//if tr, ok := transport.FromServerContext(ctx); ok {
			//	se := errors.FromError(err)
			//	if  se != nil {
			//		// 判断
			//		se.Message = "1111"
			//	}
			//
			//	tr = tr
			//
			//	//
			//	//
			//	//lang := tr.RequestHeader().Get("Accept-Language")
			//	//if lang == ""{
			//	//	return handler(ctx, req)
			//	//}
			//	//localizer := i18n.NewLocalizer(bundle, lang)
			//	//message := &i18n.Message{
			//	//	ID: "sayHello",
			//	//	Other: "Hello {{.Name}}",
			//	//}
			//	//
			//	//templateDaTa := map[string]interface{}{
			//	//	"Name": "Nick",
			//	//}
			//	//res,_ := localizer.Localize(&i18n.LocalizeConfig{
			//	//	DefaultMessage: message,
			//	//	TemplateData: templateDaTa,
			//	//})
			//	//fmt.Println(res)
			//	//fmt.Println(req)
			//	//fmt.Println(reply)
			//	//fmt.Println(tr)
			//
			//
			//	//ctx = context.WithValue(ctx, localizerKey{}, localizer)
			//}
			return handler(ctx, req)
		}
	}
}

func FromContext(ctx context.Context) *i18n.Localizer {
	return ctx.Value(localizerKey{}).(*i18n.Localizer)
}
