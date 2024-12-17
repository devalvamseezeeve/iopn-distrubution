package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"code.zeeve.net/client-projects/iopn/v2/x/cronos/client/cli"
)

// ProposalHandler is the token mapping change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitTokenMappingChangeProposalTxCmd)
