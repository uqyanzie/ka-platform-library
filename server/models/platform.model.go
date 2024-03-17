
package models


type Platform struct {
	Id              string  `json:"id" bson:"_id,omitempty" structs:"id,omitempty" mapstructure:"id,omitempty"`
	UnitName        string  `json:"unitName" bson:"unitName,omitempty" structs:"unitName,omitempty" mapstructure:"unitName,omitempty"`
	UnitClass       string  `json:"unitClass" bson:"unitClass,omitempty" structs:"unitClass,omitempty" mapstructure:"unitClass,omitempty"`
	OperationField  string  `json:"operationField" bson:"operationField,omitempty" structs:"operationField,omitempty" mapstructure:"operationField,omitempty"`
	GeneralCategory string  `json:"generalCategory" bson:"generalCategory,omitempty" structs:"generalCategory,omitempty" mapstructure:"generalCategory,omitempty"`
	GeneralType     string  `json:"generalType" bson:"generalType,omitempty" structs:"generalType,omitempty" mapstructure:"generalType,omitempty"`
	LethalityLevel  string  `json:"lethalityLevel" bson:"lethalityLevel,omitempty" structs:"lethalityLevel,omitempty" mapstructure:"lethalityLevel,omitempty"`
	CruisingSpeed   float32 `json:"cruisingSpeed" bson:"cruisingSpeed,omitempty" structs:"cruisingSpeed,omitempty" mapstructure:"cruisingSpeed,omitempty"`
	MaxSpeed        float32 `json:"maxSpeed" bson:"maxSpeed,omitempty" structs:"maxSpeed,omitempty" mapstructure:"maxSpeed,omitempty"`
	CombatRange     float32 `json:"combatRange" bson:"combatRange,omitempty" structs:"combatRange,omitempty" mapstructure:"combatRange,omitempty"`
	FuelLoad        float32 `json:"fuelLoad" bson:"fuelLoad,omitempty" structs:"fuelLoad,omitempty" mapstructure:"fuelLoad,omitempty"`
	CountryOrigin   string  `json:"countryOrigin" bson:"countryOrigin,omitempty" structs:"countryOrigin,omitempty" mapstructure:"countryOrigin,omitempty"`
}

type NewPlatform struct {
	UnitName        string  `json:"unitName" bson:"unitName" binding:"required"`
	UnitClass       string  `json:"unitClass" bson:"unitClass" binding:"required"`
	OperationField  string  `json:"operationField" bson:"operationField" binding:"required"`
	GeneralCategory string  `json:"generalCategory" bson:"generalCategory" binding:"required"`
	GeneralType     string  `json:"generalType" bson:"generalType" binding:"required"`
	LethalityLevel  string  `json:"lethalityLevel" bson:"lethalityLevel" binding:"required"`
	CruisingSpeed   float32 `json:"cruisingSpeed" bson:"cruisingSpeed" binding:"required"`
	MaxSpeed        float32 `json:"maxSpeed" bson:"maxSpeed" binding:"required"`
	CombatRange     float32 `json:"combatRange" bson:"combatRange" binding:"required"`
	FuelLoad        float32 `json:"fuelLoad" bson:"fuelLoad" binding:"required"`
	CountryOrigin   string  `json:"countryOrigin" bson:"countryOrigin" binding:"required"`
}
