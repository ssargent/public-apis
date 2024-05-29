// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: recipe/parser/v1/recipe.proto

package parserv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/ssargent/apis/pkg/recipe/parser/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ParserServiceName is the fully-qualified name of the ParserService service.
	ParserServiceName = "recipe.parser.v1.ParserService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ParserServiceParseRecipeProcedure is the fully-qualified name of the ParserService's ParseRecipe
	// RPC.
	ParserServiceParseRecipeProcedure = "/recipe.parser.v1.ParserService/ParseRecipe"
)

// ParserServiceClient is a client for the recipe.parser.v1.ParserService service.
type ParserServiceClient interface {
	ParseRecipe(context.Context, *connect_go.Request[v1.ParseRecipeRequest]) (*connect_go.Response[v1.ParseRecipeResponse], error)
}

// NewParserServiceClient constructs a client for the recipe.parser.v1.ParserService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewParserServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ParserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &parserServiceClient{
		parseRecipe: connect_go.NewClient[v1.ParseRecipeRequest, v1.ParseRecipeResponse](
			httpClient,
			baseURL+ParserServiceParseRecipeProcedure,
			opts...,
		),
	}
}

// parserServiceClient implements ParserServiceClient.
type parserServiceClient struct {
	parseRecipe *connect_go.Client[v1.ParseRecipeRequest, v1.ParseRecipeResponse]
}

// ParseRecipe calls recipe.parser.v1.ParserService.ParseRecipe.
func (c *parserServiceClient) ParseRecipe(ctx context.Context, req *connect_go.Request[v1.ParseRecipeRequest]) (*connect_go.Response[v1.ParseRecipeResponse], error) {
	return c.parseRecipe.CallUnary(ctx, req)
}

// ParserServiceHandler is an implementation of the recipe.parser.v1.ParserService service.
type ParserServiceHandler interface {
	ParseRecipe(context.Context, *connect_go.Request[v1.ParseRecipeRequest]) (*connect_go.Response[v1.ParseRecipeResponse], error)
}

// NewParserServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewParserServiceHandler(svc ParserServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	parserServiceParseRecipeHandler := connect_go.NewUnaryHandler(
		ParserServiceParseRecipeProcedure,
		svc.ParseRecipe,
		opts...,
	)
	return "/recipe.parser.v1.ParserService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ParserServiceParseRecipeProcedure:
			parserServiceParseRecipeHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedParserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedParserServiceHandler struct{}

func (UnimplementedParserServiceHandler) ParseRecipe(context.Context, *connect_go.Request[v1.ParseRecipeRequest]) (*connect_go.Response[v1.ParseRecipeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("recipe.parser.v1.ParserService.ParseRecipe is not implemented"))
}