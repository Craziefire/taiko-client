package prover

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

var (
	l1WsEndpoint   = os.Getenv("L1_NODE_WS_ENDPOINT")
	l1HttpEndpoint = os.Getenv("L1_NODE_HTTP_ENDPOINT")
	l2WsEndpoint   = os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT")
	l2HttpEndpoint = os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT")
	taikoL1        = os.Getenv("TAIKO_L1_ADDRESS")
	taikoL2        = os.Getenv("TAIKO_L2_ADDRESS")
	allowance      = "10000000000000000000000000000000000000000000000000"
	rpcTimeout     = 5 * time.Second
	minTierFee     = 1024
)

func (s *ProverTestSuite) TestNewConfigFromCliContextGuardianProver() {
	app := s.SetupApp()
	app.Action = func(ctx *cli.Context) error {
		c, err := NewConfigFromCliContext(ctx)
		s.Nil(err)
		s.Equal(l1WsEndpoint, c.L1WsEndpoint)
		s.Equal(l1HttpEndpoint, c.L1HttpEndpoint)
		s.Equal(l2WsEndpoint, c.L2WsEndpoint)
		s.Equal(l2HttpEndpoint, c.L2HttpEndpoint)
		s.Equal(taikoL1, c.TaikoL1Address.String())
		s.Equal(taikoL2, c.TaikoL2Address.String())
		s.Equal(
			crypto.PubkeyToAddress(s.p.cfg.L1ProverPrivKey.PublicKey),
			crypto.PubkeyToAddress(c.L1ProverPrivKey.PublicKey),
		)
		s.True(c.Dummy)
		s.Equal("", c.Graffiti)
		s.True(c.ProveUnassignedBlocks)
		s.True(c.ContesterMode)
		s.Equal(rpcTimeout, *c.RPCTimeout)
		s.Equal(uint64(8), c.Capacity)
		s.Equal(uint64(minTierFee), c.MinOptimisticTierFee.Uint64())
		s.Equal(uint64(minTierFee), c.MinSgxTierFee.Uint64())
		s.Equal(uint64(minTierFee), c.MinPseZkevmTierFee.Uint64())
		s.Equal(uint64(3), c.ProveBlockTxReplacementMultiplier)
		s.Equal(uint64(256), c.ProveBlockMaxTxGasTipCap.Uint64())
		s.Nil(new(Prover).InitFromCli(context.Background(), ctx))
		s.True(c.ProveUnassignedBlocks)
		s.Equal("dbPath", c.DatabasePath)
		s.Equal(uint64(128), c.DatabaseCacheSize)
		s.Equal(uint64(100), c.MaxProposedIn)
		s.Equal(os.Getenv("ASSIGNMENT_HOOK_ADDRESS"), c.AssignmentHookAddress.String())
		s.Equal(allowance, c.Allowance.String())

		return err
	}

	s.Nil(app.Run([]string{
		"TestNewConfigFromCliContextGuardianProver",
		"--" + flags.L1WSEndpoint.Name, l1WsEndpoint,
		"--" + flags.L1HTTPEndpoint.Name, l1HttpEndpoint,
		"--" + flags.L2WSEndpoint.Name, l2WsEndpoint,
		"--" + flags.L2HTTPEndpoint.Name, l2HttpEndpoint,
		"--" + flags.TaikoL1Address.Name, taikoL1,
		"--" + flags.TaikoL2Address.Name, taikoL2,
		"--" + flags.L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"--" + flags.StartingBlockID.Name, "0",
		"--" + flags.RPCTimeout.Name, "5s",
		"--" + flags.ProveBlockTxGasLimit.Name, "100000",
		"--" + flags.Dummy.Name,
		"--" + flags.MinOptimisticTierFee.Name, fmt.Sprint(minTierFee),
		"--" + flags.MinSgxTierFee.Name, fmt.Sprint(minTierFee),
		"--" + flags.MinPseZkevmTierFee.Name, fmt.Sprint(minTierFee),
		"--" + flags.ProverCapacity.Name, "8",
		"--" + flags.GuardianProver.Name, os.Getenv("GUARDIAN_PROVER_CONTRACT_ADDRESS"),
		"--" + flags.ProverAssignmentHookAddress.Name, os.Getenv("ASSIGNMENT_HOOK_ADDRESS"),
		"--" + flags.ProveBlockTxReplacementMultiplier.Name, "3",
		"--" + flags.ProveBlockMaxTxGasTipCap.Name, "256",
		"--" + flags.Graffiti.Name, "",
		"--" + flags.ProveUnassignedBlocks.Name,
		"--" + flags.DatabasePath.Name, "dbPath",
		"--" + flags.DatabaseCacheSize.Name, "128",
		"--" + flags.MaxProposedIn.Name, "100",
		"--" + flags.Allowance.Name, allowance,
	}))
}

func (s *ProverTestSuite) TestNewConfigFromCliContextProverKeyError() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"--" + flags.L1ProverPrivKey.Name, "0x",
	}), "invalid L1 prover private key")
}

func (s *ProverTestSuite) SetupApp() *cli.App {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: flags.L1WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L1HTTPEndpoint.Name},
		&cli.StringFlag{Name: flags.L2WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L2HTTPEndpoint.Name},
		&cli.StringFlag{Name: flags.TaikoL1Address.Name},
		&cli.StringFlag{Name: flags.TaikoL2Address.Name},
		&cli.StringFlag{Name: flags.L1ProverPrivKey.Name},
		&cli.Uint64Flag{Name: flags.StartingBlockID.Name},
		&cli.BoolFlag{Name: flags.Dummy.Name},
		&cli.StringFlag{Name: flags.GuardianProver.Name},
		&cli.StringFlag{Name: flags.Graffiti.Name},
		&cli.BoolFlag{Name: flags.ProveUnassignedBlocks.Name},
		&cli.Uint64Flag{Name: flags.ProveBlockTxReplacementMultiplier.Name},
		&cli.Uint64Flag{Name: flags.ProveBlockMaxTxGasTipCap.Name},
		&cli.DurationFlag{Name: flags.RPCTimeout.Name},
		&cli.Uint64Flag{Name: flags.ProverCapacity.Name},
		&cli.Uint64Flag{Name: flags.MinOptimisticTierFee.Name},
		&cli.Uint64Flag{Name: flags.MinSgxTierFee.Name},
		&cli.Uint64Flag{Name: flags.MinPseZkevmTierFee.Name},
		&cli.Uint64Flag{Name: flags.ProveBlockTxGasLimit.Name},
		&cli.StringFlag{Name: flags.DatabasePath.Name},
		&cli.Uint64Flag{Name: flags.DatabaseCacheSize.Name},
		&cli.Uint64Flag{Name: flags.MaxProposedIn.Name},
		&cli.StringFlag{Name: flags.ProverAssignmentHookAddress.Name},
		&cli.StringFlag{Name: flags.Allowance.Name},
		&cli.StringFlag{Name: flags.ContesterMode.Name},
	}
	app.Action = func(ctx *cli.Context) error {
		_, err := NewConfigFromCliContext(ctx)
		s.NotNil(err)
		return err
	}
	return app
}
