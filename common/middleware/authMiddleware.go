package middleware

//
//import (
//	"fmt"
//	"github.com/pkg/errors"
//	"github.com/zeromicro/go-zero/core/logx"
//	"google.golang.org/grpc/metadata"
//	"net/http"
//	"pj-auxiliary-diagnosis-C/common/ctxdata"
//	"pj-auxiliary-diagnosis-C/common/parserUA"
//	"pj-auxiliary-diagnosis-C/common/result"
//	"pj-auxiliary-diagnosis-C/common/tool"
//	"pj-auxiliary-diagnosis-C/common/userCenter"
//	"pj-auxiliary-diagnosis-C/common/xerr"
//	"pj-auxiliary-diagnosis-C/internal/svc"
//	"strconv"
//	"strings"
//)
//
//type AuthMiddleware struct {
//	svcCtx *svc.ServiceContext
//}
//
//func NewAuthMiddleware(svcCtx *svc.ServiceContext) *AuthMiddleware {
//	return &AuthMiddleware{
//		svcCtx: svcCtx,
//	}
//}
//
////验证请求（设备ip/mac、密钥、令牌）
//func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var err error
//		ctx := r.Context()
//		requestUrlPath := r.URL.Path
//		// 判断请求的api是否要ip/mac验证,验证
//		if tool.IsVerifyForApi(requestUrlPath, "DeviceApis", false) {
//			err = tool.VerifyDevice(&r.Header)
//			if err != nil{
//				result.RequestVerifyErrorResult(r, w, err)
//				return
//			}
//		}
//		// 判断请求的api是否要密钥（clientID,ClientSecret,对参数(不包headers)的签名）验证,验证
//		if tool.IsVerifyForApi(requestUrlPath, "SecretKeyApis", false) {
//			err = tool.VerifySecretKey(&r.Header)
//			if err != nil{
//				result.RequestVerifyErrorResult(r, w, err)
//				return
//			}
//		}
//		userData := map[string]string{ctxdata.CtxKeyUserId: "0"}
//		// 判断请求是否要令牌验证
//		if tool.IsVerifyForApi(requestUrlPath, "TokenApis", true) {
//			var authorization string
//			if authorization = r.Header.Get("Authorization"); authorization == "" {
//				result.RequestVerifyErrorResult(r, w, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_EXPIRE_ERROR), "Authorization字段为空") )
//				return
//			}
//			// 解析token，并验证签名来验证token是否合法
//			//userData, err = tool.VerifyToken(r, m.svcCtx.Config.JwtAuth.AccessSecret)
//			//if err != nil {
//			//	result.AuthHttpResult(r, w,"",err)
//			//	return
//			//}
//			//中间件-向用户中心请求验证token是否过期
//			verifyR, err := m.svcCtx.UserCenter.VerifyToken(&userCenter.VerifyTokenReq{
//				ClientId: m.svcCtx.Config.UserCenterCfg.ClientId,
//				Token:    authorization,
//			})
//			if err != nil {
//				result.AuthHttpResult(r, w,"",errors.Wrap(xerr.NewErrMsg("token错误"), err.Error()))
//				return
//			}
//			userData[ctxdata.CtxKeyUserId] = verifyR.Data.Id
//			//存储token携带的用户信息
//			ctx = metadata.AppendToOutgoingContext(ctx, ctxdata.CtxKeyUserId, userData[ctxdata.CtxKeyUserId])
//		}
//		userId, err := strconv.ParseInt(userData[ctxdata.CtxKeyUserId], 10, 64)
//		logx.WithContext(ctx).Errorf("[%s-*-%s] userId:%d, token:%s", tool.GetIp(r), r.URL, userId, r.Header.Get("Authorization"))
//		//logx.WithContext(ctx).Errorf("userId:%d, NodeCfg:%+v\n", userId, *m.svcCtx.FileCenter.NodeCfg)
//		//验证该用户是否有该接口操作权限
//
//		next(w, r.WithContext(ctx))
//	}
//}
//
////构建验证api请求的入参
//func MakeVerifyApiReq(r *http.Request, userId int64) *map[string]interface{} {
//	uA := r.Header.Get("User-Agent")
//	ua,err := parserUA.GetUA(uA)
//	var brand, unitType string
//	var machineType int64
//	if err != nil && ua != nil {
//		brand_ := strings.Split(ua.Device.Family," ")
//		if len(brand_) == 1{
//			brand = brand_[0]
//		}
//		if len(brand_) > 1 {
//			brand = brand_[0]
//			unitType = brand_[1]
//		}
//		machineType = ctxdata.MachineTypes[ua.Os.Family]
//	}
//	return &map[string]interface{}{
//		"Token":         r.Header.Get("Authorization"),
//		"UserId":        userId,
//		"UserAgent": 	   uA,
//		"MachineType":   machineType,
//		"Brand":         brand,	// r.Header.Get("Brand")
//		"UnitType":      unitType,	// r.Header.Get("UnitType")
//		"SystemVersion": fmt.Sprintf("%s %s.%s.%s",ua.Os.Family, ua.Os.Major, ua.Os.Minor, ua.Os.Patch), // r.Header.Get("SystemVersion")
//		"Browser": fmt.Sprintf("%s %s.%s.%s",ua.UserAgent.Family, ua.UserAgent.Major, ua.UserAgent.Minor, ua.UserAgent.Patch),
//		"Language": strings.Split(r.Header.Get("Accept-Language"),",")[0],
//		"NetType": "",
//		"SdkVersion":    r.Header.Get("SdkVersion"),
//		"ConnIp":        r.Host,
//		"ClientIp":      tool.GetIp(r),
//		"ClientAddr":    r.Header.Get("ClientAddr"),
//	}
//}
