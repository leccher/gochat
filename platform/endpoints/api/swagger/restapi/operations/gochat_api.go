// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"gochat/platform/endpoints/api/swagger/models"
	"gochat/platform/endpoints/api/swagger/restapi/operations/stable"
)

// NewGochatAPI creates a new Gochat instance
func NewGochatAPI(spec *loads.Document) *GochatAPI {
	return &GochatAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		StableGetAccountHandler: stable.GetAccountHandlerFunc(func(params stable.GetAccountParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation stable.GetAccount has not yet been implemented")
		}),
		StableGetConversationIDHandler: stable.GetConversationIDHandlerFunc(func(params stable.GetConversationIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation stable.GetConversationID has not yet been implemented")
		}),
		StablePostAccountHandler: stable.PostAccountHandlerFunc(func(params stable.PostAccountParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation stable.PostAccount has not yet been implemented")
		}),
		StablePostAuthLoginHandler: stable.PostAuthLoginHandlerFunc(func(params stable.PostAuthLoginParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation stable.PostAuthLogin has not yet been implemented")
		}),
		StablePostConversationHandler: stable.PostConversationHandlerFunc(func(params stable.PostConversationParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation stable.PostConversation has not yet been implemented")
		}),
		StablePostMessageHandler: stable.PostMessageHandlerFunc(func(params stable.PostMessageParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation stable.PostMessage has not yet been implemented")
		}),
		StablePutAccountHandler: stable.PutAccountHandlerFunc(func(params stable.PutAccountParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation stable.PutAccount has not yet been implemented")
		}),
		StablePutConversationHandler: stable.PutConversationHandlerFunc(func(params stable.PutConversationParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation stable.PutConversation has not yet been implemented")
		}),
		StablePutMessageHandler: stable.PutMessageHandlerFunc(func(params stable.PutMessageParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation stable.PutMessage has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		JwtAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (jwt) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*GochatAPI the gochat API */
type GochatAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// JwtAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	JwtAuth func(string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// StableGetAccountHandler sets the operation handler for the get account operation
	StableGetAccountHandler stable.GetAccountHandler
	// StableGetConversationIDHandler sets the operation handler for the get conversation ID operation
	StableGetConversationIDHandler stable.GetConversationIDHandler
	// StablePostAccountHandler sets the operation handler for the post account operation
	StablePostAccountHandler stable.PostAccountHandler
	// StablePostAuthLoginHandler sets the operation handler for the post auth login operation
	StablePostAuthLoginHandler stable.PostAuthLoginHandler
	// StablePostConversationHandler sets the operation handler for the post conversation operation
	StablePostConversationHandler stable.PostConversationHandler
	// StablePostMessageHandler sets the operation handler for the post message operation
	StablePostMessageHandler stable.PostMessageHandler
	// StablePutAccountHandler sets the operation handler for the put account operation
	StablePutAccountHandler stable.PutAccountHandler
	// StablePutConversationHandler sets the operation handler for the put conversation operation
	StablePutConversationHandler stable.PutConversationHandler
	// StablePutMessageHandler sets the operation handler for the put message operation
	StablePutMessageHandler stable.PutMessageHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *GochatAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *GochatAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *GochatAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *GochatAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *GochatAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *GochatAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *GochatAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *GochatAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *GochatAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the GochatAPI
func (o *GochatAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.JwtAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.StableGetAccountHandler == nil {
		unregistered = append(unregistered, "stable.GetAccountHandler")
	}
	if o.StableGetConversationIDHandler == nil {
		unregistered = append(unregistered, "stable.GetConversationIDHandler")
	}
	if o.StablePostAccountHandler == nil {
		unregistered = append(unregistered, "stable.PostAccountHandler")
	}
	if o.StablePostAuthLoginHandler == nil {
		unregistered = append(unregistered, "stable.PostAuthLoginHandler")
	}
	if o.StablePostConversationHandler == nil {
		unregistered = append(unregistered, "stable.PostConversationHandler")
	}
	if o.StablePostMessageHandler == nil {
		unregistered = append(unregistered, "stable.PostMessageHandler")
	}
	if o.StablePutAccountHandler == nil {
		unregistered = append(unregistered, "stable.PutAccountHandler")
	}
	if o.StablePutConversationHandler == nil {
		unregistered = append(unregistered, "stable.PutConversationHandler")
	}
	if o.StablePutMessageHandler == nil {
		unregistered = append(unregistered, "stable.PutMessageHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *GochatAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *GochatAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "jwt":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.JwtAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *GochatAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *GochatAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *GochatAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *GochatAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the gochat API
func (o *GochatAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *GochatAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/account"] = stable.NewGetAccount(o.context, o.StableGetAccountHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/conversation/{id}"] = stable.NewGetConversationID(o.context, o.StableGetConversationIDHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/account"] = stable.NewPostAccount(o.context, o.StablePostAccountHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth/login"] = stable.NewPostAuthLogin(o.context, o.StablePostAuthLoginHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/conversation"] = stable.NewPostConversation(o.context, o.StablePostConversationHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/message"] = stable.NewPostMessage(o.context, o.StablePostMessageHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/account"] = stable.NewPutAccount(o.context, o.StablePutAccountHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/conversation"] = stable.NewPutConversation(o.context, o.StablePutConversationHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/message"] = stable.NewPutMessage(o.context, o.StablePutMessageHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *GochatAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *GochatAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *GochatAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *GochatAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *GochatAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
