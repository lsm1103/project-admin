package xerr

import "google.golang.org/grpc/codes"

//成功返回
const OK codes.Code = 200

/**(前3位代表业务,后三位代表具体功能)**/

//全局错误码
const SERVER_COMMON_ERROR codes.Code = 100001	// 服务器常见错误
const REUQES_PARAM_ERROR codes.Code = 100002	// 请求参数错误
const TOKEN_EXPIRE_ERROR codes.Code = 100003	// 令牌过期错误
const TOKEN_GENERATE_ERROR codes.Code = 100004	// 令牌生成错误
const DB_ERROR codes.Code = 100005				// 数据库错误
const REQUEST_VERIFY_ERROR codes.Code = 100006	// 请求验证错误
const ALREADY_EXISTS codes.Code = 100401		// 已存在
const DATA_NOT_FIND codes.Code = 100404			// 找不到数据

//用户模块
const USER_EXIST codes.Code = 200001             // 用户已存在
const USER_REGISTRT_FAIL codes.Code = 200002     // 用户注册失败
const USER_LOGIN_ERR codes.Code = 200003         // 用户登入失败
const USER_UPDATE_PWD_ERR codes.Code = 200004    // 用户更新密码失败
const USER_UPDATE_STATUS_ERR codes.Code = 200004 // 用户更新状态失败
const USER_OPERATION_ERR codes.Code = 200004     // 用户操作失败
