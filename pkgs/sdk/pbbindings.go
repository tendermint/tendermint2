package sdk

import (
	amino "github.com/tendermint/tendermint2/pkgs/amino"
	abci "github.com/tendermint/tendermint2/pkgs/bft/abci/types"
	abcipb "github.com/tendermint/tendermint2/pkgs/bft/abci/types/pb"
	sdkpb "github.com/tendermint/tendermint2/pkgs/sdk/pb"
	proto "google.golang.org/protobuf/proto"
)

func (goo Result) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *sdkpb.Result
	{
		if IsResultReprEmpty(goo) {
			var pbov *sdkpb.Result
			msg = pbov
			return
		}
		pbo = new(sdkpb.Result)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.ResponseBase.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.ResponseBase = pbom.(*abcipb.ResponseBase)
		}
		{
			pbo.GasWanted = int64(goo.GasWanted)
		}
		{
			pbo.GasUsed = int64(goo.GasUsed)
		}
	}
	msg = pbo
	return
}

func (goo Result) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(sdkpb.Result)
	msg = pbo
	return
}

func (goo *Result) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *sdkpb.Result = msg.(*sdkpb.Result)
	{
		if pbo != nil {
			{
				if pbo.ResponseBase != nil {
					err = (*goo).ResponseBase.FromPBMessage(cdc, pbo.ResponseBase)
					if err != nil {
						return
					}
				}
			}
			{
				(*goo).GasWanted = int64(pbo.GasWanted)
			}
			{
				(*goo).GasUsed = int64(pbo.GasUsed)
			}
		}
	}
	return
}

func (_ Result) GetTypeURL() (typeURL string) {
	return "/tm.Result"
}

func IsResultReprEmpty(goor Result) (empty bool) {
	{
		empty = true
		{
			e := abci.IsResponseBaseReprEmpty(goor.ResponseBase)
			if e == false {
				return false
			}
		}
		{
			if goor.GasWanted != int64(0) {
				return false
			}
		}
		{
			if goor.GasUsed != int64(0) {
				return false
			}
		}
	}
	return
}
