// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LoadbalanceEntry loadbalance entry
//
// swagger:model LoadbalanceEntry
type LoadbalanceEntry struct {

	// values of End point servers
	Endpoints []*LoadbalanceEntryEndpointsItems0 `json:"endpoints"`

	// values of Secondary IPs
	SecondaryIPs []*LoadbalanceEntrySecondaryIPsItems0 `json:"secondaryIPs"`

	// service arguments
	ServiceArguments *LoadbalanceEntryServiceArguments `json:"serviceArguments,omitempty"`
}

// Validate validates this loadbalance entry
func (m *LoadbalanceEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryIPs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceArguments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoadbalanceEntry) validateEndpoints(formats strfmt.Registry) error {
	if swag.IsZero(m.Endpoints) { // not required
		return nil
	}

	for i := 0; i < len(m.Endpoints); i++ {
		if swag.IsZero(m.Endpoints[i]) { // not required
			continue
		}

		if m.Endpoints[i] != nil {
			if err := m.Endpoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("endpoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("endpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LoadbalanceEntry) validateSecondaryIPs(formats strfmt.Registry) error {
	if swag.IsZero(m.SecondaryIPs) { // not required
		return nil
	}

	for i := 0; i < len(m.SecondaryIPs); i++ {
		if swag.IsZero(m.SecondaryIPs[i]) { // not required
			continue
		}

		if m.SecondaryIPs[i] != nil {
			if err := m.SecondaryIPs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("secondaryIPs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("secondaryIPs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LoadbalanceEntry) validateServiceArguments(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceArguments) { // not required
		return nil
	}

	if m.ServiceArguments != nil {
		if err := m.ServiceArguments.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceArguments")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serviceArguments")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this loadbalance entry based on the context it is used
func (m *LoadbalanceEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndpoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecondaryIPs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceArguments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoadbalanceEntry) contextValidateEndpoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Endpoints); i++ {

		if m.Endpoints[i] != nil {
			if err := m.Endpoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("endpoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("endpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LoadbalanceEntry) contextValidateSecondaryIPs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SecondaryIPs); i++ {

		if m.SecondaryIPs[i] != nil {
			if err := m.SecondaryIPs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("secondaryIPs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("secondaryIPs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LoadbalanceEntry) contextValidateServiceArguments(ctx context.Context, formats strfmt.Registry) error {

	if m.ServiceArguments != nil {
		if err := m.ServiceArguments.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceArguments")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serviceArguments")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoadbalanceEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoadbalanceEntry) UnmarshalBinary(b []byte) error {
	var res LoadbalanceEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LoadbalanceEntryEndpointsItems0 loadbalance entry endpoints items0
//
// swagger:model LoadbalanceEntryEndpointsItems0
type LoadbalanceEntryEndpointsItems0 struct {

	// traffic counters of the endpoint
	Counter string `json:"counter,omitempty"`

	// IP address for external access
	EndpointIP string `json:"endpointIP,omitempty"`

	// state of the endpoint
	State string `json:"state,omitempty"`

	// port number for access service
	TargetPort int64 `json:"targetPort,omitempty"`

	// Weight for the load balancing
	Weight int64 `json:"weight,omitempty"`
}

// Validate validates this loadbalance entry endpoints items0
func (m *LoadbalanceEntryEndpointsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this loadbalance entry endpoints items0 based on context it is used
func (m *LoadbalanceEntryEndpointsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoadbalanceEntryEndpointsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoadbalanceEntryEndpointsItems0) UnmarshalBinary(b []byte) error {
	var res LoadbalanceEntryEndpointsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LoadbalanceEntrySecondaryIPsItems0 loadbalance entry secondary i ps items0
//
// swagger:model LoadbalanceEntrySecondaryIPsItems0
type LoadbalanceEntrySecondaryIPsItems0 struct {

	// IP address for secondary access
	SecondaryIP string `json:"secondaryIP,omitempty"`
}

// Validate validates this loadbalance entry secondary i ps items0
func (m *LoadbalanceEntrySecondaryIPsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this loadbalance entry secondary i ps items0 based on context it is used
func (m *LoadbalanceEntrySecondaryIPsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoadbalanceEntrySecondaryIPsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoadbalanceEntrySecondaryIPsItems0) UnmarshalBinary(b []byte) error {
	var res LoadbalanceEntrySecondaryIPsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LoadbalanceEntryServiceArguments loadbalance entry service arguments
//
// swagger:model LoadbalanceEntryServiceArguments
type LoadbalanceEntryServiceArguments struct {

	// value for BGP enable or not
	Bgp bool `json:"bgp,omitempty"`

	// block-number if any of this LB entry
	Block uint16 `json:"block,omitempty"`

	// IP address for externel access
	ExternalIP string `json:"externalIP,omitempty"`

	// value for inactivity timeout (in seconds)
	InactiveTimeOut int32 `json:"inactiveTimeOut,omitempty"`

	// externally managed rule or not
	Managed bool `json:"managed,omitempty"`

	// value for NAT mode (0-DNAT, 1-oneArm, 2-fullNAT)
	Mode int32 `json:"mode,omitempty"`

	// value for monitoring enabled or not
	Monitor bool `json:"monitor,omitempty"`

	// service name
	Name string `json:"name,omitempty"`

	// end-point specific op (0-create, 1-attachEP, 2-detachEP)
	Oper int32 `json:"oper,omitempty"`

	// port number for the access
	Port int64 `json:"port,omitempty"`

	// value for probe retries
	ProbeRetries int32 `json:"probeRetries,omitempty"`

	// value for probe timer (in seconds)
	ProbeTimeout uint32 `json:"probeTimeout,omitempty"`

	// probe port if probetype is tcp/udp/sctp
	Probeport uint16 `json:"probeport,omitempty"`

	// probe request string
	Probereq string `json:"probereq,omitempty"`

	// probe response string
	Proberesp string `json:"proberesp,omitempty"`

	// probe type for any end-point of this entry
	Probetype string `json:"probetype,omitempty"`

	// value for access protocol
	Protocol string `json:"protocol,omitempty"`

	// value for Security mode (0-Plain, 1-HTTPs)
	Security int32 `json:"security,omitempty"`

	// value for load balance algorithim
	Sel int64 `json:"sel,omitempty"`
}

// Validate validates this loadbalance entry service arguments
func (m *LoadbalanceEntryServiceArguments) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this loadbalance entry service arguments based on context it is used
func (m *LoadbalanceEntryServiceArguments) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoadbalanceEntryServiceArguments) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoadbalanceEntryServiceArguments) UnmarshalBinary(b []byte) error {
	var res LoadbalanceEntryServiceArguments
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
