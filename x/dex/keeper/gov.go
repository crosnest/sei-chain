package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	dexcache "github.com/sei-protocol/sei-chain/x/dex/cache"
	"github.com/sei-protocol/sei-chain/x/dex/types"
)

func (k Keeper) HandleRegisterPairsProposal(ctx sdk.Context, p *types.RegisterPairsProposal) error {
	// Loop through each batch contract pair an individual contract pair, token pair
	// tuple and register them individually
	for _, batchContractPair := range p.Batchcontractpair {
		var contractAddress string = batchContractPair.ContractAddr
		for _, pair := range batchContractPair.Pairs {
			k.AddRegisteredPair(ctx, contractAddress, *pair)
			k.SetTickSizeForPair(ctx, contractAddress, *pair, pair.Ticksize)
			contractAddr := types.ContractAddress(contractAddress)
			pairStr := types.PairString(pair.String())
			newBlockOrders := dexcache.BlockOrders([]types.Order{})
			k.BlockOrders[contractAddr][pairStr] = &newBlockOrders
			newBlockCancels := dexcache.BlockCancellations([]types.Cancellation{})
			k.BlockCancels[contractAddr][pairStr] = &newBlockCancels
		}
	}

	return nil
}

func (k Keeper) HandleUpdateTickSizeProposal(ctx sdk.Context, p *types.UpdateTickSizeProposal) error {
	for _, ticksize := range p.TickSizeList {
		k.SetTickSizeForPair(ctx, ticksize.ContractAddr, *ticksize.Pair, ticksize.Ticksize)
	}
	return nil
}
