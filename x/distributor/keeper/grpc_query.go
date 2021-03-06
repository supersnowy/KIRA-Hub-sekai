package keeper

import (
	"context"

	"github.com/KiraCore/sekai/x/distributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Querier struct {
	keeper Keeper
}

func NewQuerier(keeper Keeper) types.QueryServer {
	return &Querier{keeper: keeper}
}

var _ types.QueryServer = Querier{}

func (q Querier) FeesTreasury(c context.Context, request *types.QueryFeesTreasuryRequest) (*types.QueryFeesTreasuryResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryFeesTreasuryResponse{
		Coins: q.keeper.GetFeesTreasury(ctx),
	}, nil
}

func (q Querier) FeesCollected(c context.Context, request *types.QueryFeesCollectedRequest) (*types.QueryFeesCollectedResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryFeesCollectedResponse{
		Coins: q.keeper.GetFeesCollected(ctx),
	}, nil
}

func (q Querier) SnapshotPeriod(c context.Context, request *types.QuerySnapshotPeriodRequest) (*types.QuerySnapshotPeriodResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QuerySnapshotPeriodResponse{
		SnapshotPeriod: q.keeper.GetSnapPeriod(ctx),
	}, nil
}

func (q Querier) SnapshotPeriodPerformance(c context.Context, request *types.QuerySnapshotPeriodPerformanceRequest) (*types.QuerySnapshotPeriodPerformanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	valAddr, err := sdk.ValAddressFromBech32(request.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	performance, err := q.keeper.GetValidatorPerformance(ctx, valAddr)
	return &types.QuerySnapshotPeriodPerformanceResponse{
		SnapshotPeriod: q.keeper.GetSnapPeriod(ctx),
		Performance:    performance,
	}, nil
}
