package version

import (
	"github.com/go-gormigrate/gormigrate/v2"
	v1 "github.com/optimism-java/dispute-explorer/migration/version/v1"
)

var ModelSchemaList = []*gormigrate.Migration{
	&v1.AddGameLostBondTable,
}
