package contract

import (
	"context"
	"math"
	"math/big"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/retry"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/optimism-java/dispute-explorer/pkg/log"
	"golang.org/x/time/rate"
)

const maxAttempts = math.MaxInt // Succeed or die trying

type RateAndRetryDisputeGameClient struct {
	disputeGame *DisputeGame
	rl          *rate.Limiter
	strategy    retry.Strategy
}

func NewRateAndRetryDisputeGameClient(game *DisputeGame, limit rate.Limit, burst int) *RateAndRetryDisputeGameClient {
	return &RateAndRetryDisputeGameClient{
		disputeGame: game,
		rl:          rate.NewLimiter(limit, burst),
		strategy:    retry.Exponential(),
	}
}

func (s *RateAndRetryDisputeGameClient) RetryL2BlockNumber(ctx context.Context, opts *bind.CallOpts) (*big.Int, error) {
	return retry.Do(ctx, maxAttempts, s.strategy, func() (*big.Int, error) {
		res, err := s.l2BlockNumber(ctx, opts)
		if err != nil {
			log.Errorf("Failed to retryL2BlockNumber info %s", err)
		}
		return res, err
	})
}

func (s *RateAndRetryDisputeGameClient) l2BlockNumber(ctx context.Context, opts *bind.CallOpts) (*big.Int, error) {
	if err := s.rl.Wait(ctx); err != nil {
		return nil, err
	}
	cCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	opts.Context = cCtx
	return s.disputeGame.L2BlockNumber(opts)
}

func (s *RateAndRetryDisputeGameClient) RetryStatus(ctx context.Context, opts *bind.CallOpts) (uint8, error) {
	return retry.Do(ctx, maxAttempts, s.strategy, func() (uint8, error) {
		res, err := s.status(ctx, opts)
		if err != nil {
			log.Errorf("Failed to RetryStatus info %s", err)
		}
		return res, err
	})
}

func (s *RateAndRetryDisputeGameClient) status(ctx context.Context, opts *bind.CallOpts) (uint8, error) {
	if err := s.rl.Wait(ctx); err != nil {
		return 0, err
	}
	cCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	opts.Context = cCtx
	return s.disputeGame.Status(opts)
}

func (s *RateAndRetryDisputeGameClient) RetryClaimData(ctx context.Context, opts *bind.CallOpts, arg0 *big.Int) (struct {
	ParentIndex uint32
	CounteredBy common.Address
	Claimant    common.Address
	Bond        *big.Int
	Claim       [32]byte
	Position    *big.Int
	Clock       *big.Int
}, error,
) {
	return retry.Do(ctx, maxAttempts, s.strategy, func() (struct {
		ParentIndex uint32
		CounteredBy common.Address
		Claimant    common.Address
		Bond        *big.Int
		Claim       [32]byte
		Position    *big.Int
		Clock       *big.Int
	}, error,
	) {
		res, err := s.claimData(ctx, opts, arg0)
		if err != nil {
			log.Errorf("Failed to RetryClaimData info %s", err)
		}
		return res, err
	})
}

func (s *RateAndRetryDisputeGameClient) claimData(ctx context.Context, opts *bind.CallOpts, arg0 *big.Int) (struct {
	ParentIndex uint32
	CounteredBy common.Address
	Claimant    common.Address
	Bond        *big.Int
	Claim       [32]byte
	Position    *big.Int
	Clock       *big.Int
}, error,
) {
	out := new(struct {
		ParentIndex uint32
		CounteredBy common.Address
		Claimant    common.Address
		Bond        *big.Int
		Claim       [32]byte
		Position    *big.Int
		Clock       *big.Int
	})
	if err := s.rl.Wait(ctx); err != nil {
		return *out, err
	}
	cCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	opts.Context = cCtx
	return s.disputeGame.ClaimData(opts, arg0)
}

func (s *RateAndRetryDisputeGameClient) RetryCredit(ctx context.Context, opts *bind.CallOpts, address common.Address) (*big.Int, error) {
	return retry.Do(ctx, maxAttempts, s.strategy, func() (*big.Int, error) {
		res, err := s.credit(ctx, opts, address)
		if err != nil {
			log.Errorf("Failed to RetryStatus info %s", err)
		}
		return res, err
	})
}

func (s *RateAndRetryDisputeGameClient) credit(ctx context.Context, opts *bind.CallOpts, address common.Address) (*big.Int, error) {
	if err := s.rl.Wait(ctx); err != nil {
		return nil, err
	}
	cCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	opts.Context = cCtx
	return s.disputeGame.Credit(opts, address)
}

func (s *RateAndRetryDisputeGameClient) RetrySplitDepth(ctx context.Context, opts *bind.CallOpts) (*big.Int, error) {
	return retry.Do(ctx, maxAttempts, s.strategy, func() (*big.Int, error) {
		res, err := s.splitDepth(ctx, opts)
		if err != nil {
			log.Errorf("Failed to splitDepth info %s", err)
		}
		return res, err
	})
}

func (s *RateAndRetryDisputeGameClient) splitDepth(ctx context.Context, opts *bind.CallOpts) (*big.Int, error) {
	if err := s.rl.Wait(ctx); err != nil {
		return nil, err
	}
	cCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	opts.Context = cCtx
	return s.disputeGame.SplitDepth(opts)
}

func (s *RateAndRetryDisputeGameClient) RetryStartingBlockNumber(ctx context.Context, opts *bind.CallOpts) (*big.Int, error) {
	return retry.Do(ctx, maxAttempts, s.strategy, func() (*big.Int, error) {
		res, err := s.startingBlockNumber(ctx, opts)
		if err != nil {
			log.Errorf("Failed to splitDepth info %s", err)
		}
		return res, err
	})
}

func (s *RateAndRetryDisputeGameClient) startingBlockNumber(ctx context.Context, opts *bind.CallOpts) (*big.Int, error) {
	if err := s.rl.Wait(ctx); err != nil {
		return nil, err
	}
	cCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	opts.Context = cCtx
	return s.disputeGame.StartingBlockNumber(opts)
}

func (s *RateAndRetryDisputeGameClient) RetryMaxSplitDepth(ctx context.Context, opts *bind.CallOpts) (*big.Int, error) {
	return retry.Do(ctx, maxAttempts, s.strategy, func() (*big.Int, error) {
		res, err := s.maxSplitDepth(ctx, opts)
		if err != nil {
			log.Errorf("Failed to splitDepth info %s", err)
		}
		return res, err
	})
}

func (s *RateAndRetryDisputeGameClient) maxSplitDepth(ctx context.Context, opts *bind.CallOpts) (*big.Int, error) {
	if err := s.rl.Wait(ctx); err != nil {
		return nil, err
	}
	cCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	opts.Context = cCtx
	return s.disputeGame.MaxGameDepth(opts)
}
