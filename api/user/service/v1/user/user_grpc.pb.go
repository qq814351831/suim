// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	//hello
	Hello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRsp, error)
	//创建用户
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserRsp, error)
	//修改身份信息
	UpdateIdCard(ctx context.Context, in *UpdateIdCardReq, opts ...grpc.CallOption) (*UpdateIdCardRsp, error)
	//修改电话号码
	UpdatePhone(ctx context.Context, in *UpdatePhoneReq, opts ...grpc.CallOption) (*UpdatePhoneRsp, error)
	//修改密码
	UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordRsp, error)
	//忘记密码
	ForgetPassword(ctx context.Context, in *ForgetPasswordReq, opts ...grpc.CallOption) (*ForgetPasswordRsp, error)
	//修改昵称
	UpdateNickname(ctx context.Context, in *UpdateNicknameReq, opts ...grpc.CallOption) (*UpdateNicknameRsp, error)
	//修改性别
	UpdateSex(ctx context.Context, in *UpdateSexReq, opts ...grpc.CallOption) (*UpdateSexRsp, error)
	//修改头像
	UpdateAvatarUrl(ctx context.Context, in *UpdateAvatarUrlReq, opts ...grpc.CallOption) (*UpdateAvatarUrlRsp, error)
	//修改个性签名
	UpdatePersonalSign(ctx context.Context, in *UpdatePersonalSignReq, opts ...grpc.CallOption) (*UpdatePersonalSignRsp, error)
	//修改个人介绍
	UpdateIntroduce(ctx context.Context, in *UpdateIntroduceReq, opts ...grpc.CallOption) (*UpdateIntroduceRsp, error)
	//修改是否允许临时会话
	UpdateSnapCall(ctx context.Context, in *UpdateSnapCallReq, opts ...grpc.CallOption) (*UpdateSnapCallRsp, error)
	//修改用户添加好友方式
	UpdateFriendPass(ctx context.Context, in *UpdateFriendPassReq, opts ...grpc.CallOption) (*UpdateFriendPassRsp, error)
	//删除帐号
	DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserRsp, error)
	//查询用户基本信息
	InfoUserBase(ctx context.Context, in *InfoUserBaseReq, opts ...grpc.CallOption) (*InfoUserBaseRsp, error)
	//查询用户帐号信息
	InfoAccount(ctx context.Context, in *InfoAccountReq, opts ...grpc.CallOption) (*InfoAccountRsp, error)
	//查询用户是否允许临时会话
	InfoSnapCall(ctx context.Context, in *InfoSnapCallReq, opts ...grpc.CallOption) (*InfoSnapCallRsp, error)
	//查询用户添加好友方式
	InfoFriendPass(ctx context.Context, in *InfoFriendPassReq, opts ...grpc.CallOption) (*InfoFriendPassRsp, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Hello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRsp, error) {
	out := new(HelloRsp)
	err := c.cc.Invoke(ctx, "/User/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserRsp, error) {
	out := new(CreateUserRsp)
	err := c.cc.Invoke(ctx, "/User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateIdCard(ctx context.Context, in *UpdateIdCardReq, opts ...grpc.CallOption) (*UpdateIdCardRsp, error) {
	out := new(UpdateIdCardRsp)
	err := c.cc.Invoke(ctx, "/User/UpdateIdCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdatePhone(ctx context.Context, in *UpdatePhoneReq, opts ...grpc.CallOption) (*UpdatePhoneRsp, error) {
	out := new(UpdatePhoneRsp)
	err := c.cc.Invoke(ctx, "/User/UpdatePhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordRsp, error) {
	out := new(UpdatePasswordRsp)
	err := c.cc.Invoke(ctx, "/User/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ForgetPassword(ctx context.Context, in *ForgetPasswordReq, opts ...grpc.CallOption) (*ForgetPasswordRsp, error) {
	out := new(ForgetPasswordRsp)
	err := c.cc.Invoke(ctx, "/User/ForgetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateNickname(ctx context.Context, in *UpdateNicknameReq, opts ...grpc.CallOption) (*UpdateNicknameRsp, error) {
	out := new(UpdateNicknameRsp)
	err := c.cc.Invoke(ctx, "/User/UpdateNickname", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateSex(ctx context.Context, in *UpdateSexReq, opts ...grpc.CallOption) (*UpdateSexRsp, error) {
	out := new(UpdateSexRsp)
	err := c.cc.Invoke(ctx, "/User/UpdateSex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateAvatarUrl(ctx context.Context, in *UpdateAvatarUrlReq, opts ...grpc.CallOption) (*UpdateAvatarUrlRsp, error) {
	out := new(UpdateAvatarUrlRsp)
	err := c.cc.Invoke(ctx, "/User/UpdateAvatarUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdatePersonalSign(ctx context.Context, in *UpdatePersonalSignReq, opts ...grpc.CallOption) (*UpdatePersonalSignRsp, error) {
	out := new(UpdatePersonalSignRsp)
	err := c.cc.Invoke(ctx, "/User/UpdatePersonalSign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateIntroduce(ctx context.Context, in *UpdateIntroduceReq, opts ...grpc.CallOption) (*UpdateIntroduceRsp, error) {
	out := new(UpdateIntroduceRsp)
	err := c.cc.Invoke(ctx, "/User/UpdateIntroduce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateSnapCall(ctx context.Context, in *UpdateSnapCallReq, opts ...grpc.CallOption) (*UpdateSnapCallRsp, error) {
	out := new(UpdateSnapCallRsp)
	err := c.cc.Invoke(ctx, "/User/UpdateSnapCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateFriendPass(ctx context.Context, in *UpdateFriendPassReq, opts ...grpc.CallOption) (*UpdateFriendPassRsp, error) {
	out := new(UpdateFriendPassRsp)
	err := c.cc.Invoke(ctx, "/User/UpdateFriendPass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserRsp, error) {
	out := new(DeleteUserRsp)
	err := c.cc.Invoke(ctx, "/User/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) InfoUserBase(ctx context.Context, in *InfoUserBaseReq, opts ...grpc.CallOption) (*InfoUserBaseRsp, error) {
	out := new(InfoUserBaseRsp)
	err := c.cc.Invoke(ctx, "/User/InfoUserBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) InfoAccount(ctx context.Context, in *InfoAccountReq, opts ...grpc.CallOption) (*InfoAccountRsp, error) {
	out := new(InfoAccountRsp)
	err := c.cc.Invoke(ctx, "/User/InfoAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) InfoSnapCall(ctx context.Context, in *InfoSnapCallReq, opts ...grpc.CallOption) (*InfoSnapCallRsp, error) {
	out := new(InfoSnapCallRsp)
	err := c.cc.Invoke(ctx, "/User/InfoSnapCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) InfoFriendPass(ctx context.Context, in *InfoFriendPassReq, opts ...grpc.CallOption) (*InfoFriendPassRsp, error) {
	out := new(InfoFriendPassRsp)
	err := c.cc.Invoke(ctx, "/User/InfoFriendPass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	//hello
	Hello(context.Context, *HelloReq) (*HelloRsp, error)
	//创建用户
	CreateUser(context.Context, *CreateUserReq) (*CreateUserRsp, error)
	//修改身份信息
	UpdateIdCard(context.Context, *UpdateIdCardReq) (*UpdateIdCardRsp, error)
	//修改电话号码
	UpdatePhone(context.Context, *UpdatePhoneReq) (*UpdatePhoneRsp, error)
	//修改密码
	UpdatePassword(context.Context, *UpdatePasswordReq) (*UpdatePasswordRsp, error)
	//忘记密码
	ForgetPassword(context.Context, *ForgetPasswordReq) (*ForgetPasswordRsp, error)
	//修改昵称
	UpdateNickname(context.Context, *UpdateNicknameReq) (*UpdateNicknameRsp, error)
	//修改性别
	UpdateSex(context.Context, *UpdateSexReq) (*UpdateSexRsp, error)
	//修改头像
	UpdateAvatarUrl(context.Context, *UpdateAvatarUrlReq) (*UpdateAvatarUrlRsp, error)
	//修改个性签名
	UpdatePersonalSign(context.Context, *UpdatePersonalSignReq) (*UpdatePersonalSignRsp, error)
	//修改个人介绍
	UpdateIntroduce(context.Context, *UpdateIntroduceReq) (*UpdateIntroduceRsp, error)
	//修改是否允许临时会话
	UpdateSnapCall(context.Context, *UpdateSnapCallReq) (*UpdateSnapCallRsp, error)
	//修改用户添加好友方式
	UpdateFriendPass(context.Context, *UpdateFriendPassReq) (*UpdateFriendPassRsp, error)
	//删除帐号
	DeleteUser(context.Context, *DeleteUserReq) (*DeleteUserRsp, error)
	//查询用户基本信息
	InfoUserBase(context.Context, *InfoUserBaseReq) (*InfoUserBaseRsp, error)
	//查询用户帐号信息
	InfoAccount(context.Context, *InfoAccountReq) (*InfoAccountRsp, error)
	//查询用户是否允许临时会话
	InfoSnapCall(context.Context, *InfoSnapCallReq) (*InfoSnapCallRsp, error)
	//查询用户添加好友方式
	InfoFriendPass(context.Context, *InfoFriendPassReq) (*InfoFriendPassRsp, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Hello(context.Context, *HelloReq) (*HelloRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedUserServer) CreateUser(context.Context, *CreateUserReq) (*CreateUserRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServer) UpdateIdCard(context.Context, *UpdateIdCardReq) (*UpdateIdCardRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIdCard not implemented")
}
func (UnimplementedUserServer) UpdatePhone(context.Context, *UpdatePhoneReq) (*UpdatePhoneRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePhone not implemented")
}
func (UnimplementedUserServer) UpdatePassword(context.Context, *UpdatePasswordReq) (*UpdatePasswordRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedUserServer) ForgetPassword(context.Context, *ForgetPasswordReq) (*ForgetPasswordRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgetPassword not implemented")
}
func (UnimplementedUserServer) UpdateNickname(context.Context, *UpdateNicknameReq) (*UpdateNicknameRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNickname not implemented")
}
func (UnimplementedUserServer) UpdateSex(context.Context, *UpdateSexReq) (*UpdateSexRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSex not implemented")
}
func (UnimplementedUserServer) UpdateAvatarUrl(context.Context, *UpdateAvatarUrlReq) (*UpdateAvatarUrlRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAvatarUrl not implemented")
}
func (UnimplementedUserServer) UpdatePersonalSign(context.Context, *UpdatePersonalSignReq) (*UpdatePersonalSignRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePersonalSign not implemented")
}
func (UnimplementedUserServer) UpdateIntroduce(context.Context, *UpdateIntroduceReq) (*UpdateIntroduceRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIntroduce not implemented")
}
func (UnimplementedUserServer) UpdateSnapCall(context.Context, *UpdateSnapCallReq) (*UpdateSnapCallRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSnapCall not implemented")
}
func (UnimplementedUserServer) UpdateFriendPass(context.Context, *UpdateFriendPassReq) (*UpdateFriendPassRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFriendPass not implemented")
}
func (UnimplementedUserServer) DeleteUser(context.Context, *DeleteUserReq) (*DeleteUserRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServer) InfoUserBase(context.Context, *InfoUserBaseReq) (*InfoUserBaseRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoUserBase not implemented")
}
func (UnimplementedUserServer) InfoAccount(context.Context, *InfoAccountReq) (*InfoAccountRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoAccount not implemented")
}
func (UnimplementedUserServer) InfoSnapCall(context.Context, *InfoSnapCallReq) (*InfoSnapCallRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoSnapCall not implemented")
}
func (UnimplementedUserServer) InfoFriendPass(context.Context, *InfoFriendPassReq) (*InfoFriendPassRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoFriendPass not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Hello(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateIdCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIdCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateIdCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UpdateIdCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateIdCard(ctx, req.(*UpdateIdCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdatePhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePhoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdatePhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UpdatePhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdatePhone(ctx, req.(*UpdatePhoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdatePassword(ctx, req.(*UpdatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ForgetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgetPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ForgetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/ForgetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ForgetPassword(ctx, req.(*ForgetPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateNickname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNicknameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateNickname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UpdateNickname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateNickname(ctx, req.(*UpdateNicknameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateSex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateSex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UpdateSex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateSex(ctx, req.(*UpdateSexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateAvatarUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAvatarUrlReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateAvatarUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UpdateAvatarUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateAvatarUrl(ctx, req.(*UpdateAvatarUrlReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdatePersonalSign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePersonalSignReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdatePersonalSign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UpdatePersonalSign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdatePersonalSign(ctx, req.(*UpdatePersonalSignReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateIntroduce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIntroduceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateIntroduce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UpdateIntroduce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateIntroduce(ctx, req.(*UpdateIntroduceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateSnapCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSnapCallReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateSnapCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UpdateSnapCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateSnapCall(ctx, req.(*UpdateSnapCallReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateFriendPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFriendPassReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateFriendPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UpdateFriendPass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateFriendPass(ctx, req.(*UpdateFriendPassReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteUser(ctx, req.(*DeleteUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_InfoUserBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoUserBaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).InfoUserBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/InfoUserBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).InfoUserBase(ctx, req.(*InfoUserBaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_InfoAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).InfoAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/InfoAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).InfoAccount(ctx, req.(*InfoAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_InfoSnapCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoSnapCallReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).InfoSnapCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/InfoSnapCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).InfoSnapCall(ctx, req.(*InfoSnapCallReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_InfoFriendPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoFriendPassReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).InfoFriendPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/InfoFriendPass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).InfoFriendPass(ctx, req.(*InfoFriendPassReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _User_Hello_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "UpdateIdCard",
			Handler:    _User_UpdateIdCard_Handler,
		},
		{
			MethodName: "UpdatePhone",
			Handler:    _User_UpdatePhone_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _User_UpdatePassword_Handler,
		},
		{
			MethodName: "ForgetPassword",
			Handler:    _User_ForgetPassword_Handler,
		},
		{
			MethodName: "UpdateNickname",
			Handler:    _User_UpdateNickname_Handler,
		},
		{
			MethodName: "UpdateSex",
			Handler:    _User_UpdateSex_Handler,
		},
		{
			MethodName: "UpdateAvatarUrl",
			Handler:    _User_UpdateAvatarUrl_Handler,
		},
		{
			MethodName: "UpdatePersonalSign",
			Handler:    _User_UpdatePersonalSign_Handler,
		},
		{
			MethodName: "UpdateIntroduce",
			Handler:    _User_UpdateIntroduce_Handler,
		},
		{
			MethodName: "UpdateSnapCall",
			Handler:    _User_UpdateSnapCall_Handler,
		},
		{
			MethodName: "UpdateFriendPass",
			Handler:    _User_UpdateFriendPass_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _User_DeleteUser_Handler,
		},
		{
			MethodName: "InfoUserBase",
			Handler:    _User_InfoUserBase_Handler,
		},
		{
			MethodName: "InfoAccount",
			Handler:    _User_InfoAccount_Handler,
		},
		{
			MethodName: "InfoSnapCall",
			Handler:    _User_InfoSnapCall_Handler,
		},
		{
			MethodName: "InfoFriendPass",
			Handler:    _User_InfoFriendPass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/user/user.proto",
}
