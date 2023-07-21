// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cosmos/staking/v1beta1/query.proto

package stakingv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Query_Validators_FullMethodName                    = "/cosmos.staking.v1beta1.Query/Validators"
	Query_Validator_FullMethodName                     = "/cosmos.staking.v1beta1.Query/Validator"
	Query_ValidatorDelegations_FullMethodName          = "/cosmos.staking.v1beta1.Query/ValidatorDelegations"
	Query_ValidatorUnbondingDelegations_FullMethodName = "/cosmos.staking.v1beta1.Query/ValidatorUnbondingDelegations"
	Query_Delegation_FullMethodName                    = "/cosmos.staking.v1beta1.Query/Delegation"
	Query_UnbondingDelegation_FullMethodName           = "/cosmos.staking.v1beta1.Query/UnbondingDelegation"
	Query_DelegatorDelegations_FullMethodName          = "/cosmos.staking.v1beta1.Query/DelegatorDelegations"
	Query_DelegatorUnbondingDelegations_FullMethodName = "/cosmos.staking.v1beta1.Query/DelegatorUnbondingDelegations"
	Query_Redelegations_FullMethodName                 = "/cosmos.staking.v1beta1.Query/Redelegations"
	Query_DelegatorValidators_FullMethodName           = "/cosmos.staking.v1beta1.Query/DelegatorValidators"
	Query_DelegatorValidator_FullMethodName            = "/cosmos.staking.v1beta1.Query/DelegatorValidator"
	Query_HistoricalInfo_FullMethodName                = "/cosmos.staking.v1beta1.Query/HistoricalInfo"
	Query_Pool_FullMethodName                          = "/cosmos.staking.v1beta1.Query/Pool"
	Query_Params_FullMethodName                        = "/cosmos.staking.v1beta1.Query/Params"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Validators queries all validators that match the given status.
	Validators(ctx context.Context, in *QueryValidatorsRequest, opts ...grpc.CallOption) (*QueryValidatorsResponse, error)
	// Validator queries validator info for given validator address.
	Validator(ctx context.Context, in *QueryValidatorRequest, opts ...grpc.CallOption) (*QueryValidatorResponse, error)
	// ValidatorDelegations queries delegate info for given validator.
	ValidatorDelegations(ctx context.Context, in *QueryValidatorDelegationsRequest, opts ...grpc.CallOption) (*QueryValidatorDelegationsResponse, error)
	// ValidatorUnbondingDelegations queries unbonding delegations of a validator.
	ValidatorUnbondingDelegations(ctx context.Context, in *QueryValidatorUnbondingDelegationsRequest, opts ...grpc.CallOption) (*QueryValidatorUnbondingDelegationsResponse, error)
	// Delegation queries delegate info for given validator delegator pair.
	Delegation(ctx context.Context, in *QueryDelegationRequest, opts ...grpc.CallOption) (*QueryDelegationResponse, error)
	// UnbondingDelegation queries unbonding info for given validator delegator
	// pair.
	UnbondingDelegation(ctx context.Context, in *QueryUnbondingDelegationRequest, opts ...grpc.CallOption) (*QueryUnbondingDelegationResponse, error)
	// DelegatorDelegations queries all delegations of a given delegator address.
	DelegatorDelegations(ctx context.Context, in *QueryDelegatorDelegationsRequest, opts ...grpc.CallOption) (*QueryDelegatorDelegationsResponse, error)
	// DelegatorUnbondingDelegations queries all unbonding delegations of a given
	// delegator address.
	DelegatorUnbondingDelegations(ctx context.Context, in *QueryDelegatorUnbondingDelegationsRequest, opts ...grpc.CallOption) (*QueryDelegatorUnbondingDelegationsResponse, error)
	// Redelegations queries redelegations of given address.
	Redelegations(ctx context.Context, in *QueryRedelegationsRequest, opts ...grpc.CallOption) (*QueryRedelegationsResponse, error)
	// DelegatorValidators queries all validators info for given delegator
	// address.
	DelegatorValidators(ctx context.Context, in *QueryDelegatorValidatorsRequest, opts ...grpc.CallOption) (*QueryDelegatorValidatorsResponse, error)
	// DelegatorValidator queries validator info for given delegator validator
	// pair.
	DelegatorValidator(ctx context.Context, in *QueryDelegatorValidatorRequest, opts ...grpc.CallOption) (*QueryDelegatorValidatorResponse, error)
	// HistoricalInfo queries the historical info for given height.
	HistoricalInfo(ctx context.Context, in *QueryHistoricalInfoRequest, opts ...grpc.CallOption) (*QueryHistoricalInfoResponse, error)
	// Pool queries the pool info.
	Pool(ctx context.Context, in *QueryPoolRequest, opts ...grpc.CallOption) (*QueryPoolResponse, error)
	// Parameters queries the staking parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Validators(ctx context.Context, in *QueryValidatorsRequest, opts ...grpc.CallOption) (*QueryValidatorsResponse, error) {
	out := new(QueryValidatorsResponse)
	err := c.cc.Invoke(ctx, Query_Validators_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Validator(ctx context.Context, in *QueryValidatorRequest, opts ...grpc.CallOption) (*QueryValidatorResponse, error) {
	out := new(QueryValidatorResponse)
	err := c.cc.Invoke(ctx, Query_Validator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ValidatorDelegations(ctx context.Context, in *QueryValidatorDelegationsRequest, opts ...grpc.CallOption) (*QueryValidatorDelegationsResponse, error) {
	out := new(QueryValidatorDelegationsResponse)
	err := c.cc.Invoke(ctx, Query_ValidatorDelegations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ValidatorUnbondingDelegations(ctx context.Context, in *QueryValidatorUnbondingDelegationsRequest, opts ...grpc.CallOption) (*QueryValidatorUnbondingDelegationsResponse, error) {
	out := new(QueryValidatorUnbondingDelegationsResponse)
	err := c.cc.Invoke(ctx, Query_ValidatorUnbondingDelegations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Delegation(ctx context.Context, in *QueryDelegationRequest, opts ...grpc.CallOption) (*QueryDelegationResponse, error) {
	out := new(QueryDelegationResponse)
	err := c.cc.Invoke(ctx, Query_Delegation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) UnbondingDelegation(ctx context.Context, in *QueryUnbondingDelegationRequest, opts ...grpc.CallOption) (*QueryUnbondingDelegationResponse, error) {
	out := new(QueryUnbondingDelegationResponse)
	err := c.cc.Invoke(ctx, Query_UnbondingDelegation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegatorDelegations(ctx context.Context, in *QueryDelegatorDelegationsRequest, opts ...grpc.CallOption) (*QueryDelegatorDelegationsResponse, error) {
	out := new(QueryDelegatorDelegationsResponse)
	err := c.cc.Invoke(ctx, Query_DelegatorDelegations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegatorUnbondingDelegations(ctx context.Context, in *QueryDelegatorUnbondingDelegationsRequest, opts ...grpc.CallOption) (*QueryDelegatorUnbondingDelegationsResponse, error) {
	out := new(QueryDelegatorUnbondingDelegationsResponse)
	err := c.cc.Invoke(ctx, Query_DelegatorUnbondingDelegations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Redelegations(ctx context.Context, in *QueryRedelegationsRequest, opts ...grpc.CallOption) (*QueryRedelegationsResponse, error) {
	out := new(QueryRedelegationsResponse)
	err := c.cc.Invoke(ctx, Query_Redelegations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegatorValidators(ctx context.Context, in *QueryDelegatorValidatorsRequest, opts ...grpc.CallOption) (*QueryDelegatorValidatorsResponse, error) {
	out := new(QueryDelegatorValidatorsResponse)
	err := c.cc.Invoke(ctx, Query_DelegatorValidators_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegatorValidator(ctx context.Context, in *QueryDelegatorValidatorRequest, opts ...grpc.CallOption) (*QueryDelegatorValidatorResponse, error) {
	out := new(QueryDelegatorValidatorResponse)
	err := c.cc.Invoke(ctx, Query_DelegatorValidator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HistoricalInfo(ctx context.Context, in *QueryHistoricalInfoRequest, opts ...grpc.CallOption) (*QueryHistoricalInfoResponse, error) {
	out := new(QueryHistoricalInfoResponse)
	err := c.cc.Invoke(ctx, Query_HistoricalInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Pool(ctx context.Context, in *QueryPoolRequest, opts ...grpc.CallOption) (*QueryPoolResponse, error) {
	out := new(QueryPoolResponse)
	err := c.cc.Invoke(ctx, Query_Pool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Validators queries all validators that match the given status.
	Validators(context.Context, *QueryValidatorsRequest) (*QueryValidatorsResponse, error)
	// Validator queries validator info for given validator address.
	Validator(context.Context, *QueryValidatorRequest) (*QueryValidatorResponse, error)
	// ValidatorDelegations queries delegate info for given validator.
	ValidatorDelegations(context.Context, *QueryValidatorDelegationsRequest) (*QueryValidatorDelegationsResponse, error)
	// ValidatorUnbondingDelegations queries unbonding delegations of a validator.
	ValidatorUnbondingDelegations(context.Context, *QueryValidatorUnbondingDelegationsRequest) (*QueryValidatorUnbondingDelegationsResponse, error)
	// Delegation queries delegate info for given validator delegator pair.
	Delegation(context.Context, *QueryDelegationRequest) (*QueryDelegationResponse, error)
	// UnbondingDelegation queries unbonding info for given validator delegator
	// pair.
	UnbondingDelegation(context.Context, *QueryUnbondingDelegationRequest) (*QueryUnbondingDelegationResponse, error)
	// DelegatorDelegations queries all delegations of a given delegator address.
	DelegatorDelegations(context.Context, *QueryDelegatorDelegationsRequest) (*QueryDelegatorDelegationsResponse, error)
	// DelegatorUnbondingDelegations queries all unbonding delegations of a given
	// delegator address.
	DelegatorUnbondingDelegations(context.Context, *QueryDelegatorUnbondingDelegationsRequest) (*QueryDelegatorUnbondingDelegationsResponse, error)
	// Redelegations queries redelegations of given address.
	Redelegations(context.Context, *QueryRedelegationsRequest) (*QueryRedelegationsResponse, error)
	// DelegatorValidators queries all validators info for given delegator
	// address.
	DelegatorValidators(context.Context, *QueryDelegatorValidatorsRequest) (*QueryDelegatorValidatorsResponse, error)
	// DelegatorValidator queries validator info for given delegator validator
	// pair.
	DelegatorValidator(context.Context, *QueryDelegatorValidatorRequest) (*QueryDelegatorValidatorResponse, error)
	// HistoricalInfo queries the historical info for given height.
	HistoricalInfo(context.Context, *QueryHistoricalInfoRequest) (*QueryHistoricalInfoResponse, error)
	// Pool queries the pool info.
	Pool(context.Context, *QueryPoolRequest) (*QueryPoolResponse, error)
	// Parameters queries the staking parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Validators(context.Context, *QueryValidatorsRequest) (*QueryValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validators not implemented")
}
func (UnimplementedQueryServer) Validator(context.Context, *QueryValidatorRequest) (*QueryValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validator not implemented")
}
func (UnimplementedQueryServer) ValidatorDelegations(context.Context, *QueryValidatorDelegationsRequest) (*QueryValidatorDelegationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorDelegations not implemented")
}
func (UnimplementedQueryServer) ValidatorUnbondingDelegations(context.Context, *QueryValidatorUnbondingDelegationsRequest) (*QueryValidatorUnbondingDelegationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorUnbondingDelegations not implemented")
}
func (UnimplementedQueryServer) Delegation(context.Context, *QueryDelegationRequest) (*QueryDelegationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delegation not implemented")
}
func (UnimplementedQueryServer) UnbondingDelegation(context.Context, *QueryUnbondingDelegationRequest) (*QueryUnbondingDelegationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnbondingDelegation not implemented")
}
func (UnimplementedQueryServer) DelegatorDelegations(context.Context, *QueryDelegatorDelegationsRequest) (*QueryDelegatorDelegationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegatorDelegations not implemented")
}
func (UnimplementedQueryServer) DelegatorUnbondingDelegations(context.Context, *QueryDelegatorUnbondingDelegationsRequest) (*QueryDelegatorUnbondingDelegationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegatorUnbondingDelegations not implemented")
}
func (UnimplementedQueryServer) Redelegations(context.Context, *QueryRedelegationsRequest) (*QueryRedelegationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Redelegations not implemented")
}
func (UnimplementedQueryServer) DelegatorValidators(context.Context, *QueryDelegatorValidatorsRequest) (*QueryDelegatorValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegatorValidators not implemented")
}
func (UnimplementedQueryServer) DelegatorValidator(context.Context, *QueryDelegatorValidatorRequest) (*QueryDelegatorValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegatorValidator not implemented")
}
func (UnimplementedQueryServer) HistoricalInfo(context.Context, *QueryHistoricalInfoRequest) (*QueryHistoricalInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HistoricalInfo not implemented")
}
func (UnimplementedQueryServer) Pool(context.Context, *QueryPoolRequest) (*QueryPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pool not implemented")
}
func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Validators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Validators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Validators_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Validators(ctx, req.(*QueryValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Validator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Validator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Validator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Validator(ctx, req.(*QueryValidatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ValidatorDelegations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorDelegationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorDelegations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ValidatorDelegations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorDelegations(ctx, req.(*QueryValidatorDelegationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ValidatorUnbondingDelegations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorUnbondingDelegationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorUnbondingDelegations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ValidatorUnbondingDelegations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorUnbondingDelegations(ctx, req.(*QueryValidatorUnbondingDelegationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Delegation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Delegation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Delegation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Delegation(ctx, req.(*QueryDelegationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_UnbondingDelegation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUnbondingDelegationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UnbondingDelegation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_UnbondingDelegation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UnbondingDelegation(ctx, req.(*QueryUnbondingDelegationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegatorDelegations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorDelegationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorDelegations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_DelegatorDelegations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorDelegations(ctx, req.(*QueryDelegatorDelegationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegatorUnbondingDelegations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorUnbondingDelegationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorUnbondingDelegations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_DelegatorUnbondingDelegations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorUnbondingDelegations(ctx, req.(*QueryDelegatorUnbondingDelegationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Redelegations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRedelegationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Redelegations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Redelegations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Redelegations(ctx, req.(*QueryRedelegationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegatorValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_DelegatorValidators_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorValidators(ctx, req.(*QueryDelegatorValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegatorValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorValidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_DelegatorValidator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorValidator(ctx, req.(*QueryDelegatorValidatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HistoricalInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHistoricalInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HistoricalInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_HistoricalInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HistoricalInfo(ctx, req.(*QueryHistoricalInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Pool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Pool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Pool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Pool(ctx, req.(*QueryPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.staking.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validators",
			Handler:    _Query_Validators_Handler,
		},
		{
			MethodName: "Validator",
			Handler:    _Query_Validator_Handler,
		},
		{
			MethodName: "ValidatorDelegations",
			Handler:    _Query_ValidatorDelegations_Handler,
		},
		{
			MethodName: "ValidatorUnbondingDelegations",
			Handler:    _Query_ValidatorUnbondingDelegations_Handler,
		},
		{
			MethodName: "Delegation",
			Handler:    _Query_Delegation_Handler,
		},
		{
			MethodName: "UnbondingDelegation",
			Handler:    _Query_UnbondingDelegation_Handler,
		},
		{
			MethodName: "DelegatorDelegations",
			Handler:    _Query_DelegatorDelegations_Handler,
		},
		{
			MethodName: "DelegatorUnbondingDelegations",
			Handler:    _Query_DelegatorUnbondingDelegations_Handler,
		},
		{
			MethodName: "Redelegations",
			Handler:    _Query_Redelegations_Handler,
		},
		{
			MethodName: "DelegatorValidators",
			Handler:    _Query_DelegatorValidators_Handler,
		},
		{
			MethodName: "DelegatorValidator",
			Handler:    _Query_DelegatorValidator_Handler,
		},
		{
			MethodName: "HistoricalInfo",
			Handler:    _Query_HistoricalInfo_Handler,
		},
		{
			MethodName: "Pool",
			Handler:    _Query_Pool_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/staking/v1beta1/query.proto",
}
