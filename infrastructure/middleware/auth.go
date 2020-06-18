package middleware

import (
	"context"
	"crypto/ecdsa"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/ezio1119/fishapp-auth/conf"
	"github.com/ezio1119/fishapp-auth/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// 認証が必要なメソッド
var (
	requireIDTokenMethod = []string{
		"/pb.UserService/CurrentUser",
		"/pb.UserService/UpdateUser",
		"/pb.UserService/UpdatePassword",
	}
	requireRefreshTokenMethod = []string{
		"/pb.UserService/RefreshIDToken",
		"/pb.UserService/Logout",
	}
)

func (m *Middleware) AuthInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var err error
		for _, m := range requireIDTokenMethod {
			if m == info.FullMethod {
				ctx, err = authFunc(ctx, "id_token")
				if err != nil {
					return nil, err
				}
			}
		}
		for _, m := range requireRefreshTokenMethod {
			if m == info.FullMethod {
				ctx, err = authFunc(ctx, "refresh_token")
				if err != nil {
					return nil, err
				}
			}
		}
		return handler(ctx, req)
	}
}

func authFunc(ctx context.Context, tokenType string) (context.Context, error) {
	t := getTokenFromMeta(ctx, "authorization")
	if t == "" {
		return nil, status.Errorf(codes.Unauthenticated, "%s not in metadata", tokenType)
	}
	c, err := getClaimsFromToken(t)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token validation failed: %s", err.Error())
	}
	if c.Subject != tokenType {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token_type: require %s", tokenType)
	}
	return context.WithValue(ctx, domain.JwtCtxKey, *c), nil
}

func getTokenFromMeta(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	if len(md[key]) != 1 {
		return ""
	}
	return md[key][0]
}

func getClaimsFromToken(t string) (*domain.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(t, &domain.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if c, ok := token.Claims.(*domain.JwtClaims); ok && token.Valid {
		return c, nil
	}
	return nil, err
}

var publicKey *ecdsa.PublicKey

func init() {
	var err error
	publicKey, err = jwt.ParseECPublicKeyFromPEM([]byte(conf.C.Auth.PubJwtkey))
	if err != nil {
		log.Fatal(err)
	}
}