package v1

import (
	"time"
)

// CrUpdTime - struct CrUpdTime
type CrUpdTime struct {
	CreatedAt time.Time `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt time.Time `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

// Location - struct Location
type Location struct {
	Id                      uint32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid4                   []byte    `protobuf:"bytes,2,opt,name=uuid4,proto3" json:"uuid4,omitempty"`
	IdS                     string    `protobuf:"bytes,3,opt,name=id_s,json=idS,proto3" json:"id_s,omitempty"`
	LocId                   string    `protobuf:"bytes,4,opt,name=loc_id,json=locId,proto3" json:"loc_id,omitempty"`
	Description             string    `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Conditions              string    `protobuf:"bytes,6,opt,name=conditions,proto3" json:"conditions,omitempty"`
	CountrySubentity        string    `protobuf:"bytes,7,opt,name=country_subentity,json=countrySubentity,proto3" json:"country_subentity,omitempty"`
	CountrySubentityCode    string    `protobuf:"bytes,8,opt,name=country_subentity_code,json=countrySubentityCode,proto3" json:"country_subentity_code,omitempty"`
	LocationTypeCode        string    `protobuf:"bytes,9,opt,name=location_type_code,json=locationTypeCode,proto3" json:"location_type_code,omitempty"`
	InformationURI          string    `protobuf:"bytes,10,opt,name=information_u_r_i,json=informationURI,proto3" json:"information_u_r_i,omitempty"`
	LocName                 string    `protobuf:"bytes,11,opt,name=loc_name,json=locName,proto3" json:"loc_name,omitempty"`
	ValidityPeriodStartDate time.Time `protobuf:"bytes,12,opt,name=validity_period_start_date,json=validityPeriodStartDate,proto3" json:"validity_period_start_date,omitempty"`
	ValidityPeriodEndDate   time.Time `protobuf:"bytes,13,opt,name=validity_period_end_date,json=validityPeriodEndDate,proto3" json:"validity_period_end_date,omitempty"`
	LocationCoordLat        float64   `protobuf:"fixed64,14,opt,name=location_coord_lat,json=locationCoordLat,proto3" json:"location_coord_lat,omitempty"`
	LocationCoordLon        float64   `protobuf:"fixed64,15,opt,name=location_coord_lon,json=locationCoordLon,proto3" json:"location_coord_lon,omitempty"`
	AltitudeMeasure         float64   `protobuf:"fixed64,16,opt,name=altitude_measure,json=altitudeMeasure,proto3" json:"altitude_measure,omitempty"`
}
