package systemError

import "github.com/go-kratos/kratos/v2/errors"


var SystemError = errors.New(500, "System busy", "系统繁忙,请稍后再试")
