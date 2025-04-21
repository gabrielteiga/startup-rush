package startup_entity

import "time"

type IStartupRepository interface {
	Create(name, slogan string, foundation time.Time) *Startup
	List() []*Startup
	FindByID(id uint) *Startup
	FindByIDs(ids []uint) []*Startup
}
