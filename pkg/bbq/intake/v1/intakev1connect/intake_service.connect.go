// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: bbq/intake/v1/intake_service.proto

package intakev1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/ssargent/apis/pkg/bbq/intake/v1"
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
	// IntakeServiceName is the fully-qualified name of the IntakeService service.
	IntakeServiceName = "bbq.intake.v1.IntakeService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// IntakeServiceRecordProcedure is the fully-qualified name of the IntakeService's Record RPC.
	IntakeServiceRecordProcedure = "/bbq.intake.v1.IntakeService/Record"
	// IntakeServiceSessionProcedure is the fully-qualified name of the IntakeService's Session RPC.
	IntakeServiceSessionProcedure = "/bbq.intake.v1.IntakeService/Session"
)

// IntakeServiceClient is a client for the bbq.intake.v1.IntakeService service.
type IntakeServiceClient interface {
	Record(context.Context, *connect_go.Request[v1.RecordRequest]) (*connect_go.Response[v1.RecordResponse], error)
	Session(context.Context, *connect_go.Request[v1.SessionRequest]) (*connect_go.Response[v1.SessionResponse], error)
}

// NewIntakeServiceClient constructs a client for the bbq.intake.v1.IntakeService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIntakeServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) IntakeServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &intakeServiceClient{
		record: connect_go.NewClient[v1.RecordRequest, v1.RecordResponse](
			httpClient,
			baseURL+IntakeServiceRecordProcedure,
			opts...,
		),
		session: connect_go.NewClient[v1.SessionRequest, v1.SessionResponse](
			httpClient,
			baseURL+IntakeServiceSessionProcedure,
			opts...,
		),
	}
}

// intakeServiceClient implements IntakeServiceClient.
type intakeServiceClient struct {
	record  *connect_go.Client[v1.RecordRequest, v1.RecordResponse]
	session *connect_go.Client[v1.SessionRequest, v1.SessionResponse]
}

// Record calls bbq.intake.v1.IntakeService.Record.
func (c *intakeServiceClient) Record(ctx context.Context, req *connect_go.Request[v1.RecordRequest]) (*connect_go.Response[v1.RecordResponse], error) {
	return c.record.CallUnary(ctx, req)
}

// Session calls bbq.intake.v1.IntakeService.Session.
func (c *intakeServiceClient) Session(ctx context.Context, req *connect_go.Request[v1.SessionRequest]) (*connect_go.Response[v1.SessionResponse], error) {
	return c.session.CallUnary(ctx, req)
}

// IntakeServiceHandler is an implementation of the bbq.intake.v1.IntakeService service.
type IntakeServiceHandler interface {
	Record(context.Context, *connect_go.Request[v1.RecordRequest]) (*connect_go.Response[v1.RecordResponse], error)
	Session(context.Context, *connect_go.Request[v1.SessionRequest]) (*connect_go.Response[v1.SessionResponse], error)
}

// NewIntakeServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIntakeServiceHandler(svc IntakeServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	intakeServiceRecordHandler := connect_go.NewUnaryHandler(
		IntakeServiceRecordProcedure,
		svc.Record,
		opts...,
	)
	intakeServiceSessionHandler := connect_go.NewUnaryHandler(
		IntakeServiceSessionProcedure,
		svc.Session,
		opts...,
	)
	return "/bbq.intake.v1.IntakeService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case IntakeServiceRecordProcedure:
			intakeServiceRecordHandler.ServeHTTP(w, r)
		case IntakeServiceSessionProcedure:
			intakeServiceSessionHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedIntakeServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedIntakeServiceHandler struct{}

func (UnimplementedIntakeServiceHandler) Record(context.Context, *connect_go.Request[v1.RecordRequest]) (*connect_go.Response[v1.RecordResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bbq.intake.v1.IntakeService.Record is not implemented"))
}

func (UnimplementedIntakeServiceHandler) Session(context.Context, *connect_go.Request[v1.SessionRequest]) (*connect_go.Response[v1.SessionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bbq.intake.v1.IntakeService.Session is not implemented"))
}