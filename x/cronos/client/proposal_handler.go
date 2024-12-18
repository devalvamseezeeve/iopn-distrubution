package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/devalvamseezeeve/iopn-distrubution/v2/x/iopn/client/cli"
)

// ProposalHandler is the token mapping change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitTokenMappingChangeProposalTxCmd)
