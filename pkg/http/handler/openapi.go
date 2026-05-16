package handler

import (
	"reflect"

	"github.com/gorilla/mux"
	"github.com/swaggest/jsonschema-go"
	openapigo "github.com/swaggest/openapi-go"
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/rest/openapi"
)

// OpenAPICollector accumulates OpenAPI operations by walking a gorilla/mux Router.
// Port of SigNoz's pkg/http/handler.OpenAPICollector.
type OpenAPICollector struct {
	collector *openapi.Collector
	reflector *openapi3.Reflector
}

// NewOpenAPICollector creates an OpenAPICollector backed by the given reflector.
// It installs an InterceptDefName hook so that render.SuccessResponse (whose Data
// field is interface{}) is inlined at each call site rather than extracted into a
// shared named schema — the same approach SigNoz uses.
func NewOpenAPICollector(reflector *openapi3.Reflector) *OpenAPICollector {
	reflector.InterceptDefName(func(t reflect.Type, defaultDefName string) string {
		if defaultDefName != "RenderSuccessResponse" {
			return defaultDefName
		}
		dataField, ok := t.FieldByName("Data")
		if !ok {
			return defaultDefName
		}
		// interface{} has no name → swaggest inlines the schema instead of
		// extracting it to components/schemas (defName == "" triggers inline path).
		return dataField.Type.Name()
	})
	return &OpenAPICollector{
		collector: openapi.NewCollector(reflector),
		reflector: reflector,
	}
}

// Walker is a mux.WalkFunc. Pass it to router.Walk() after all routes are registered.
// It inspects each route's handler: if the handler implements Handler (i.e. was created
// with New()), it collects the OpenAPI operation from ServeOpenAPI.
func (c *OpenAPICollector) Walker(route *mux.Route, _ *mux.Router, _ []*mux.Route) error {
	httpHandler := route.GetHandler()
	if httpHandler == nil {
		return nil
	}

	path, err := route.GetPathTemplate()
	if err != nil && path == "" {
		return nil
	}

	methods, err := route.GetMethods()
	if err != nil {
		// No methods constraint — skip (e.g. the /openapi.yaml route registered without .Methods())
		return nil
	}

	h, ok := httpHandler.(Handler)
	if !ok {
		return nil
	}

	for _, method := range methods {
		if err := c.collector.CollectOperation(method, path, c.collect(method, path, h.ServeOpenAPI)); err != nil {
			return err
		}
	}
	return nil
}

func (c *OpenAPICollector) collect(method, path string, serveOpenAPI func(openapigo.OperationContext)) func(openapigo.OperationContext) error {
	return func(oc openapigo.OperationContext) error {
		serveOpenAPI(oc)

		if c.collector.HasAnnotation(method, path) {
			return nil
		}

		// Inject path parameters automatically from {param} segments in the path.
		_, _, pathItems, err := openapigo.SanitizeMethodPath(method, path)
		if err != nil {
			return err
		}
		if len(pathItems) > 0 {
			req := jsonschema.Struct{}
			for _, p := range pathItems {
				req.Fields = append(req.Fields, jsonschema.Field{
					Name:  "F" + p,
					Tag:   reflect.StructTag(`path:"` + p + `"`),
					Value: "",
				})
			}
			oc.AddReqStructure(req)
		}
		return nil
	}
}

// MarshalYAML returns the accumulated OpenAPI 3.0 spec as YAML bytes.
func (c *OpenAPICollector) MarshalYAML() ([]byte, error) {
	return c.reflector.Spec.MarshalYAML()
}
