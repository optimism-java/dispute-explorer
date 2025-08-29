package version

import (
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	v1 "github.com/optimism-java/dispute-explorer/migration/version/v1"
	v2 "github.com/optimism-java/dispute-explorer/migration/version/v2"
	v3 "github.com/optimism-java/dispute-explorer/migration/version/v3"
	v4 "github.com/optimism-java/dispute-explorer/migration/version/v4"
	v5 "github.com/optimism-java/dispute-explorer/migration/version/v5"
	v6 "github.com/optimism-java/dispute-explorer/migration/version/v6"
	v7 "github.com/optimism-java/dispute-explorer/migration/version/v7"
	v8 "github.com/optimism-java/dispute-explorer/migration/version/v8"
	v9 "github.com/optimism-java/dispute-explorer/migration/version/v9"
)

var ModelSchemaList = []*gormigrate.Migration{
	&v1.AddGameLostBondTable,
	&v2.AddCalculateLostForDisputeGameTable,
	&v3.UpdateLostBondAndClaimDataTable,
	&v4.UpdateClaimDataPositionColumnTable,
	&v5.AddOnChainStatusForDisputeGameTable,
	&v6.UpdateClaimDataClockColumnTable,
	&v7.AddClaimDataLenForDisputeGameTable,
	&v8.AddFrontendMoveTrackingTable,
	&v9.AddIsSyncedFieldTable,
}
