package gateways

import "gorm.io/gorm"


// >>>>>>>>>>>>>>>>>>>>>>>>
// contet:
// GatewaysInterface
// >>>>>>>>>>>>>>>>>>>>>>>>
type GatewaysInterface interface {
	Connect() *gorm.DB
}


