func NewSpamMigitationAnteDecorator(pylonsmodulekeeper PylonsKeeper) AnteSpamMigitationDecorator {
	return AnteSpamMigitationDecorator{
		pk: pylonsmodulekeeper,
	}
}
	}

	return next(ctx, tx, simulate)


func (ad AnteSpamMigitationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if (ctx.IsCheckTx() || ctx.IsReCheckTx()) && !simulate {}


	func TestOtherTransactionIsValid(t *testing.T) {
	// Set MaxTxsInBlock  = 2
	numberTxsinBlocks := 2
	config := network.ConfigWithMaxTxsInBlock(uint64(numberTxsinBlocks))
	net := network.New(t, config)
			args = append(args, common...)
	res, _ := clitestutil.ExecTestCLICmd(ctx, bank.NewSendTxCmd(), args)
	return res

			require.NoError(t, ctx.Codec.UnmarshalJSON(validRes.Bytes(), &resp))
	require.Equal(t, successCode, resp.Code)
}
