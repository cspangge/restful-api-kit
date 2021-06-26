package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Claims struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

var JwtSecret=[]byte(
	`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA0NVnB1TZ15xma7+wRYcRLTQ1Zcjw07xIRGHYS5BpgxofsKg0
vYQ9PCs4oXDiR5YPiK0VMlIDx11YmjuLNfJH2g+NIU+XS3wGywQUfLZ+cPLv5HAZ
BFLCGlQtVKwKkhLHOjy/R7JX7JpjHw+M+IJXVUtmrud0yRDd2rHZrGfv76CXYpI5
K0P0Wi1UzzX0Bn+hMQ8WV1FlBjriph/NnX9qv7Va43bHDJOHHjtCAygyRcGHs3k6
Zukvjl3cgo02AF6aAukYLh0caJmY4UHzIKr2vnj+6j5WabAXQpkkSCBy+rOrlwx8
AN2m+jIr+db6Z0sfDEQsIksBN0EJjwvwDIMy1QIDAQABAoIBADlshPEiJui/6/Xy
9uCXfGSmt4wujfJbErAjC5vzsN/I/i/QgPgf9cqmw+K4rLAslkbp+SsmHD3k8Ckw
+F20KofOwOsI/LxX9XI7w/prICek1cxF+ssSiqc8wF5JFHsISNo9vAdWBhFcfqca
2pO3G5arrieoI5fqWGb7grSCvBpy2cNZMlXkr1t+LXt8YCJuSnjFX5bmT2ARetfg
T6aG+xn79kgfJYGgdScqlFIvId7/E5rJeYMWarkO4kIOZJ1v4PserUpnpJwn7VFA
rvQ3/+CrOfasyCaSosHUOq/qBTaLoQGJJLx6YTlgVtSsbLrpClQWFJ48mnrd6xYs
z7sKS6ECgYEA+dS4wcxmQcN+XGDC1jKqAO8vTA9q+UUA1zMvYqtGuKkdUamZwUsX
8yaq3A6444rr/HPwF5eefS/lRZORA57MYKrd/loOIGHT5dfeiPwaHkVCRz3RUa1+
LdtMrySDm+N2v2WR6od/7s/2VVQF1mocp6onAXV/TJc9wFOnPnSBzDkCgYEA1f2F
ScFEq7mJWq630gdlGrzP45AjIcfVnK623Poq6+Q75PjWiSlIrqowj14eTl9zMI+u
zhuQG7YcfWgma7Brc2qoWgIg3oKNextCUAF+Iqj8YU354lo4vW4wR8rqr3LwDEuz
Rfn32tvPgNTKmanwjgLUsZTeyavP5APzPAtMU30CgYAy4x5i8yRc/tntLRRilt1D
5XRH1NggP+rk6psbSPHVyTScnqQf2BDfyR+CRUlpp7Fbsa3U0GBK9YeIvX8LMrxE
cGlZVrSL6gFETs0YL+mOAuN2KWRKc8FAy+t5vUFFbnaQk7q9/62SUi/Fv6CUxPZj
PGhHhsxQKHo+g4uMUCTkGQKBgBgGdLaJKmqVVCSIRT8hcWqFFFoaTrd2Njd3LjS1
Kah+YXMdtQiFSBHELduK16A4+zfSie++DHnwAlbaIKYqkoXMzcX9Qy94POY7c0CE
SmMd0egi1xFgy8oI2wbUc3DrWKQow6HxpLA/yZZPTcfe2pE3JCYj7rd5wMd64g41
`)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"msg":    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		// parseToken 解析token包含的信息
		claims, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
