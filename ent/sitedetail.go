// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kayprogrammer/socialnet-v4/ent/sitedetail"
)

// SiteDetail is the model entity for the SiteDetail schema.
type SiteDetail struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Fb holds the value of the "fb" field.
	Fb string `json:"fb,omitempty"`
	// Tw holds the value of the "tw" field.
	Tw string `json:"tw,omitempty"`
	// Wh holds the value of the "wh" field.
	Wh string `json:"wh,omitempty"`
	// Ig holds the value of the "ig" field.
	Ig           string `json:"ig,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SiteDetail) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sitedetail.FieldName, sitedetail.FieldEmail, sitedetail.FieldPhone, sitedetail.FieldAddress, sitedetail.FieldFb, sitedetail.FieldTw, sitedetail.FieldWh, sitedetail.FieldIg:
			values[i] = new(sql.NullString)
		case sitedetail.FieldCreatedAt, sitedetail.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case sitedetail.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SiteDetail fields.
func (sd *SiteDetail) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sitedetail.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sd.ID = *value
			}
		case sitedetail.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sd.CreatedAt = value.Time
			}
		case sitedetail.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sd.UpdatedAt = value.Time
			}
		case sitedetail.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sd.Name = value.String
			}
		case sitedetail.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				sd.Email = value.String
			}
		case sitedetail.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				sd.Phone = value.String
			}
		case sitedetail.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				sd.Address = value.String
			}
		case sitedetail.FieldFb:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fb", values[i])
			} else if value.Valid {
				sd.Fb = value.String
			}
		case sitedetail.FieldTw:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tw", values[i])
			} else if value.Valid {
				sd.Tw = value.String
			}
		case sitedetail.FieldWh:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wh", values[i])
			} else if value.Valid {
				sd.Wh = value.String
			}
		case sitedetail.FieldIg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ig", values[i])
			} else if value.Valid {
				sd.Ig = value.String
			}
		default:
			sd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SiteDetail.
// This includes values selected through modifiers, order, etc.
func (sd *SiteDetail) Value(name string) (ent.Value, error) {
	return sd.selectValues.Get(name)
}

// Update returns a builder for updating this SiteDetail.
// Note that you need to call SiteDetail.Unwrap() before calling this method if this SiteDetail
// was returned from a transaction, and the transaction was committed or rolled back.
func (sd *SiteDetail) Update() *SiteDetailUpdateOne {
	return NewSiteDetailClient(sd.config).UpdateOne(sd)
}

// Unwrap unwraps the SiteDetail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sd *SiteDetail) Unwrap() *SiteDetail {
	_tx, ok := sd.config.driver.(*txDriver)
	if !ok {
		panic("ent: SiteDetail is not a transactional entity")
	}
	sd.config.driver = _tx.drv
	return sd
}

// String implements the fmt.Stringer.
func (sd *SiteDetail) String() string {
	var builder strings.Builder
	builder.WriteString("SiteDetail(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sd.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sd.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(sd.Name)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(sd.Email)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(sd.Phone)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(sd.Address)
	builder.WriteString(", ")
	builder.WriteString("fb=")
	builder.WriteString(sd.Fb)
	builder.WriteString(", ")
	builder.WriteString("tw=")
	builder.WriteString(sd.Tw)
	builder.WriteString(", ")
	builder.WriteString("wh=")
	builder.WriteString(sd.Wh)
	builder.WriteString(", ")
	builder.WriteString("ig=")
	builder.WriteString(sd.Ig)
	builder.WriteByte(')')
	return builder.String()
}

// SiteDetails is a parsable slice of SiteDetail.
type SiteDetails []*SiteDetail
