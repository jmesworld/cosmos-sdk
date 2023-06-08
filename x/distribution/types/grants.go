package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type WinningGrantExpiration struct {
	AtTime   int64 `json:"at_time"`
	AtHeight int64 `json:"at_height"`
}
type WinningGrants []WinningGrant

type WinningGrant struct {
	DAO sdk.AccAddress `json:"dao"`
	// could be sdk.Uint
	Amount         sdk.Uint `json:"amount"`
	ExpireAtHeight sdk.Uint `json:"expire_at_height"`
	YesRatio       sdk.Dec  `json:"yes_ratio"`
}

//func (wg *WinningGrants) Marshal() ([]byte, error) {
//	return ModuleCdc.MarshalJSON(wg)
//	//TODO implement me
//	//panic("implement me")
//}
//
//func (wg *WinningGrants) MarshalTo(data []byte) (n int, err error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (wg *WinningGrants) MarshalToSizedBuffer(dAtA []byte) (int, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (wg *WinningGrants) Size() int {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (wg *WinningGrants) Unmarshal(data []byte) error {
//	//TODO implement me
//	panic("implement me")
//}
//

//
//func (wg *WinningGrant) Reset() {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (wg *WinningGrant) String() string {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (wg *WinningGrant) ProtoMessage() {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (wg *WinningGrants) Reset() {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (wg *WinningGrants) String() string {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (wg *WinningGrants) ProtoMessage() {
//	//TODO implement me
//	panic("implement me")
//}
//
//// UnmarshalJSON unmarshals the JSON bytes into a WinningGrants struct
//func (wgs *WinningGrants) UnmarshalJSON(bz []byte) error {
//	var wgsJSON WinningGrants
//	if err := ModuleCdc.UnmarshalJSON(bz, &wgsJSON); err != nil {
//		return err
//	}
//	*wgs = wgsJSON
//	return nil
//}
//
//func (wg *WinningGrant) UnmarshalJSON(bz []byte) error {
//	var wgJSON WinningGrant
//	if err := ModuleCdc.UnmarshalJSON(bz, &wgJSON); err != nil {
//		return err
//	}
//	wg.DAO = wgJSON.DAO
//	wg.Amount = wgJSON.Amount
//	wg.Expiration = WinningGrantExpiration{
//		AtTime:   wgJSON.Expiration.AtTime,
//		AtHeight: wgJSON.Expiration.AtHeight,
//	}
//	wg.YesRatio = wgJSON.YesRatio
//	return nil
//}
