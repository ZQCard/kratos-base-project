package administratorError

import "github.com/go-kratos/kratos/v2/errors"


var AdministratorNotExist = errors.New(400, "administrator not exist", "管理员账号不存在")

var AdministratorForbid = errors.New(401, "administrator forbid", "管理员已被冻结")

var AdministratorDeleted = errors.New(401, "administrator deleted", "管理员已删除")

var AdministratorPasswordWrong = errors.New(401, "administrator password wrong", "管理员已删除")


