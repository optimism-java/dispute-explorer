package version

import (
	"github.com/go-gormigrate/gormigrate/v2"
	v1 "github.com/optimism-java/dispute-explorer/migration/version/v1"
	v2 "github.com/optimism-java/dispute-explorer/migration/version/v2"
	v3 "github.com/optimism-java/dispute-explorer/migration/version/v3"
)

var ModelSchemaList = []*gormigrate.Migration{
	&v1.AddGameLostBondTable,
	&v2.AddCalculateLostForDisputeGameTable,
	&v3.UpdateLostBondAndClaimDataTable,
}
