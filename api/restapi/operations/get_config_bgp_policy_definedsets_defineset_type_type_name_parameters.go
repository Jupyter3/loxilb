// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameParams creates a new GetConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameParams object
//
// There are no default values defined in the spec.
func NewGetConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameParams() GetConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameParams {

	return GetConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameParams{}
}

// GetConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameParams contains all the bound params for the get config bgp policy definedsets defineset type type name operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetConfigBgpPolicyDefinedsetsDefinesetTypeTypeName
type GetConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*defineset type one of prefix/neighbor/community/extcommunity/aspath/largecommunity
	  Required: true
	  In: path
	*/
	DefinesetType string
	/*type name
	  Required: true
	  In: path
	*/
	TypeName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameParams() beforehand.
func (o *GetConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rDefinesetType, rhkDefinesetType, _ := route.Params.GetOK("defineset_type")
	if err := o.bindDefinesetType(rDefinesetType, rhkDefinesetType, route.Formats); err != nil {
		res = append(res, err)
	}

	rTypeName, rhkTypeName, _ := route.Params.GetOK("type_name")
	if err := o.bindTypeName(rTypeName, rhkTypeName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDefinesetType binds and validates parameter DefinesetType from path.
func (o *GetConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameParams) bindDefinesetType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.DefinesetType = raw

	return nil
}

// bindTypeName binds and validates parameter TypeName from path.
func (o *GetConfigBgpPolicyDefinedsetsDefinesetTypeTypeNameParams) bindTypeName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.TypeName = raw

	return nil
}
