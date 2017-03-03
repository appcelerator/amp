package auth

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// Keys used in context metadata
const (
	TokenKey     = "amp.token"
	RequesterKey = "amp.requester"
)

var (
	// TODO: this MUST NOT be public
	// TODO: find a way to store this key secretly
	secretKey = []byte("&kv@l3go-f=@^*@ush0(o5*5utxe6932j9di+ume=$mkj%d&&9*%k53(bmpksf&!c2&zpw$z=8ndi6ib)&nxms0ia7rf*sj9g8r4")

	anonymousAllowed = []string{
		"/account.Account/SignUp",
		"/account.Account/Verify",
		"/account.Account/Login",
		"/account.Account/PasswordReset",
		"/account.Account/PasswordSet",
		"/account.Account/ForgotLogin",
		"/account.Account/GetUser",
		"/account.Account/ListUsers",
		"/account.Account/GetOrganization",
		"/account.Account/ListOrganizations",
		"/account.Account/GetTeam",
		"/account.Account/ListTeams",

		"/version.Version/List",
	}
)

// LoginCredentials represents login credentials
type LoginCredentials struct {
	Token string
}

// GetRequestMetadata implements credentials.PerRPCCredentials
func (c *LoginCredentials) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		TokenKey: c.Token,
	}, nil
}

// RequireTransportSecurity implements credentials.PerRPCCredentials
func (c *LoginCredentials) RequireTransportSecurity() bool {
	return false
}

func isAnonymous(elem string) bool {
	for _, e := range anonymousAllowed {
		if e == elem {
			return true
		}
	}
	return false
}

// StreamInterceptor is an interceptor checking for authentication tokens
func StreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if !isAnonymous(info.FullMethod) {
		if _, err := authorize(stream.Context()); err != nil {
			return err
		}
	}
	return handler(srv, stream)
}

// Interceptor is an interceptor checking for authentication tokens
func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (i interface{}, err error) {
	if !isAnonymous(info.FullMethod) {
		if ctx, err = authorize(ctx); err != nil {
			return nil, err
		}
	}
	return handler(ctx, req)
}

func authorize(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "credentials required")
	}
	tokens := md[TokenKey]
	if len(tokens) == 0 {
		return nil, grpc.Errorf(codes.Unauthenticated, "credentials required")
	}
	token := tokens[0]
	if token == "" {
		return nil, grpc.Errorf(codes.Unauthenticated, "credentials required")
	}
	claims, err := ValidateToken(token, TokenTypeLogin)
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid credentials")
	}
	// Enrich the context with the requester
	ctx = metadata.NewContext(ctx, metadata.Pairs(RequesterKey, claims.AccountName))
	return ctx, nil
}

// GetRequester gets the requester from context metadata
func GetRequesterName(ctx context.Context) (string, error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return "", fmt.Errorf("unable to get metadata from context")
	}
	users := md[RequesterKey]
	if len(users) == 0 {
		return "", fmt.Errorf("context metadata has no requester field")
	}
	user := users[0]
	return user, nil
}
