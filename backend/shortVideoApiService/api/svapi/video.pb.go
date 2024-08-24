// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: svapi/video.proto

package svapi

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VideoAuthor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar      string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	IsFollowing bool   `protobuf:"varint,4,opt,name=is_following,json=isFollowing,proto3" json:"is_following,omitempty"`
}

func (x *VideoAuthor) Reset() {
	*x = VideoAuthor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoAuthor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoAuthor) ProtoMessage() {}

func (x *VideoAuthor) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoAuthor.ProtoReflect.Descriptor instead.
func (*VideoAuthor) Descriptor() ([]byte, []int) {
	return file_svapi_video_proto_rawDescGZIP(), []int{0}
}

func (x *VideoAuthor) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VideoAuthor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VideoAuthor) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *VideoAuthor) GetIsFollowing() bool {
	if x != nil {
		return x.IsFollowing
	}
	return false
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                            // 视频唯一标识
	Author        *VideoAuthor `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`                                     // 视频作者信息
	PlayUrl       string       `protobuf:"bytes,3,opt,name=play_url,json=playUrl,proto3" json:"play_url,omitempty"`                    // 视频播放地址
	CoverUrl      string       `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`                 // 视频封面地址
	FavoriteCount int64        `protobuf:"varint,5,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"` // 视频的点赞总数
	CommentCount  int64        `protobuf:"varint,6,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`    // 视频的评论总数
	IsFavorite    bool         `protobuf:"varint,7,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite,omitempty"`          // true-已点赞，false-未点赞
	Title         string       `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`                                       // 视频标题
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_svapi_video_proto_rawDescGZIP(), []int{1}
}

func (x *Video) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetAuthor() *VideoAuthor {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Video) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Video) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type FeedShortVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatestTime int64 `protobuf:"varint,1,opt,name=latest_time,json=latestTime,proto3" json:"latest_time,omitempty"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	UserId     int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *FeedShortVideoRequest) Reset() {
	*x = FeedShortVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedShortVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedShortVideoRequest) ProtoMessage() {}

func (x *FeedShortVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedShortVideoRequest.ProtoReflect.Descriptor instead.
func (*FeedShortVideoRequest) Descriptor() ([]byte, []int) {
	return file_svapi_video_proto_rawDescGZIP(), []int{2}
}

func (x *FeedShortVideoRequest) GetLatestTime() int64 {
	if x != nil {
		return x.LatestTime
	}
	return 0
}

func (x *FeedShortVideoRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type FeedShortVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta     *Metadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Videos   []*Video  `protobuf:"bytes,2,rep,name=videos,proto3" json:"videos,omitempty"`
	NextTime int64     `protobuf:"varint,3,opt,name=next_time,json=nextTime,proto3" json:"next_time,omitempty"` // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

func (x *FeedShortVideoResponse) Reset() {
	*x = FeedShortVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedShortVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedShortVideoResponse) ProtoMessage() {}

func (x *FeedShortVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedShortVideoResponse.ProtoReflect.Descriptor instead.
func (*FeedShortVideoResponse) Descriptor() ([]byte, []int) {
	return file_svapi_video_proto_rawDescGZIP(), []int{3}
}

func (x *FeedShortVideoResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *FeedShortVideoResponse) GetVideos() []*Video {
	if x != nil {
		return x.Videos
	}
	return nil
}

func (x *FeedShortVideoResponse) GetNextTime() int64 {
	if x != nil {
		return x.NextTime
	}
	return 0
}

type GetVideoByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId int64 `protobuf:"varint,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
}

func (x *GetVideoByIdRequest) Reset() {
	*x = GetVideoByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_video_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoByIdRequest) ProtoMessage() {}

func (x *GetVideoByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_video_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoByIdRequest.ProtoReflect.Descriptor instead.
func (*GetVideoByIdRequest) Descriptor() ([]byte, []int) {
	return file_svapi_video_proto_rawDescGZIP(), []int{4}
}

func (x *GetVideoByIdRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type GetVideoByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta  *Metadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Video *Video    `protobuf:"bytes,2,opt,name=video,proto3" json:"video,omitempty"`
}

func (x *GetVideoByIdResponse) Reset() {
	*x = GetVideoByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_video_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoByIdResponse) ProtoMessage() {}

func (x *GetVideoByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_video_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoByIdResponse.ProtoReflect.Descriptor instead.
func (*GetVideoByIdResponse) Descriptor() ([]byte, []int) {
	return file_svapi_video_proto_rawDescGZIP(), []int{5}
}

func (x *GetVideoByIdResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *GetVideoByIdResponse) GetVideo() *Video {
	if x != nil {
		return x.Video
	}
	return nil
}

type PublishVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`                         // 视频数据
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`                       // 视频标题
	CoverUrl string `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"` // 视频封面地址
	UserId   int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`      // 发布视频的user id
}

func (x *PublishVideoRequest) Reset() {
	*x = PublishVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_video_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishVideoRequest) ProtoMessage() {}

func (x *PublishVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_video_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishVideoRequest.ProtoReflect.Descriptor instead.
func (*PublishVideoRequest) Descriptor() ([]byte, []int) {
	return file_svapi_video_proto_rawDescGZIP(), []int{6}
}

func (x *PublishVideoRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PublishVideoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PublishVideoRequest) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *PublishVideoRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type PublishVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Metadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *PublishVideoResponse) Reset() {
	*x = PublishVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_video_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishVideoResponse) ProtoMessage() {}

func (x *PublishVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_video_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishVideoResponse.ProtoReflect.Descriptor instead.
func (*PublishVideoResponse) Descriptor() ([]byte, []int) {
	return file_svapi_video_proto_rawDescGZIP(), []int{7}
}

func (x *PublishVideoResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

type ListPublishedVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64              `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Pagination *PaginationRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListPublishedVideoRequest) Reset() {
	*x = ListPublishedVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_video_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPublishedVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPublishedVideoRequest) ProtoMessage() {}

func (x *ListPublishedVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_video_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPublishedVideoRequest.ProtoReflect.Descriptor instead.
func (*ListPublishedVideoRequest) Descriptor() ([]byte, []int) {
	return file_svapi_video_proto_rawDescGZIP(), []int{8}
}

func (x *ListPublishedVideoRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListPublishedVideoRequest) GetPagination() *PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListPublishedVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta       *Metadata           `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Videos     []*Video            `protobuf:"bytes,2,rep,name=videos,proto3" json:"videos,omitempty"`
	Pagination *PaginationResponse `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListPublishedVideoResponse) Reset() {
	*x = ListPublishedVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_video_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPublishedVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPublishedVideoResponse) ProtoMessage() {}

func (x *ListPublishedVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_video_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPublishedVideoResponse.ProtoReflect.Descriptor instead.
func (*ListPublishedVideoResponse) Descriptor() ([]byte, []int) {
	return file_svapi_video_proto_rawDescGZIP(), []int{9}
}

func (x *ListPublishedVideoResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ListPublishedVideoResponse) GetVideos() []*Video {
	if x != nil {
		return x.Videos
	}
	return nil
}

func (x *ListPublishedVideoResponse) GetPagination() *PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_svapi_video_proto protoreflect.FileDescriptor

var file_svapi_video_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x76, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x0b, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x22, 0xfe, 0x01, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2a, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x5a, 0x0a, 0x15, 0x46, 0x65, 0x65, 0x64, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x07, 0xba, 0x48, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x16, 0x46, 0x65, 0x65, 0x64, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x76,
	0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x12, 0x24, 0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x65, 0x78,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x08,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07,
	0xba, 0x48, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64,
	0x22, 0x5f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x22, 0x0a,
	0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73,
	0x76, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x22, 0x89, 0x01, 0x0a, 0x13, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0x48,
	0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x32, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xba, 0x48,
	0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3b, 0x0a,
	0x14, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x77, 0x0a, 0x19, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xba, 0x48, 0x04, 0x22, 0x02, 0x20,
	0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xa2, 0x01, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x39, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xb5, 0x03, 0x0a, 0x1a, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x72, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x0e, 0x46, 0x65, 0x65, 0x64, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x1c, 0x2e, 0x73, 0x76, 0x61, 0x70,
	0x69, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e,
	0x46, 0x65, 0x65, 0x64, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c,
	0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x3a, 0x01, 0x2a, 0x12,
	0x63, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x1a, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x76,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x12, 0x12, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x2f, 0x7b, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x5f, 0x69, 0x64, 0x7d, 0x12, 0x5b, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x12, 0x1a, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x22, 0x07, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x3a, 0x01,
	0x2a, 0x12, 0x6d, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x20, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x76, 0x61, 0x70,
	0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0c, 0x22, 0x07, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x3a, 0x01, 0x2a,
	0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x7a, 0x65, 0x6e, 0x69, 0x74, 0x68, 0x2f, 0x44, 0x6f, 0x75, 0x54, 0x6f,
	0x6b, 0x2f, 0x2e, 0x2e, 0x2e, 0x3b, 0x73, 0x76, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_svapi_video_proto_rawDescOnce sync.Once
	file_svapi_video_proto_rawDescData = file_svapi_video_proto_rawDesc
)

func file_svapi_video_proto_rawDescGZIP() []byte {
	file_svapi_video_proto_rawDescOnce.Do(func() {
		file_svapi_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_svapi_video_proto_rawDescData)
	})
	return file_svapi_video_proto_rawDescData
}

var file_svapi_video_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_svapi_video_proto_goTypes = []interface{}{
	(*VideoAuthor)(nil),                // 0: svapi.VideoAuthor
	(*Video)(nil),                      // 1: svapi.Video
	(*FeedShortVideoRequest)(nil),      // 2: svapi.FeedShortVideoRequest
	(*FeedShortVideoResponse)(nil),     // 3: svapi.FeedShortVideoResponse
	(*GetVideoByIdRequest)(nil),        // 4: svapi.GetVideoByIdRequest
	(*GetVideoByIdResponse)(nil),       // 5: svapi.GetVideoByIdResponse
	(*PublishVideoRequest)(nil),        // 6: svapi.PublishVideoRequest
	(*PublishVideoResponse)(nil),       // 7: svapi.PublishVideoResponse
	(*ListPublishedVideoRequest)(nil),  // 8: svapi.ListPublishedVideoRequest
	(*ListPublishedVideoResponse)(nil), // 9: svapi.ListPublishedVideoResponse
	(*Metadata)(nil),                   // 10: svapi.Metadata
	(*PaginationRequest)(nil),          // 11: svapi.PaginationRequest
	(*PaginationResponse)(nil),         // 12: svapi.PaginationResponse
}
var file_svapi_video_proto_depIdxs = []int32{
	0,  // 0: svapi.Video.author:type_name -> svapi.VideoAuthor
	10, // 1: svapi.FeedShortVideoResponse.meta:type_name -> svapi.Metadata
	1,  // 2: svapi.FeedShortVideoResponse.videos:type_name -> svapi.Video
	10, // 3: svapi.GetVideoByIdResponse.meta:type_name -> svapi.Metadata
	1,  // 4: svapi.GetVideoByIdResponse.video:type_name -> svapi.Video
	10, // 5: svapi.PublishVideoResponse.meta:type_name -> svapi.Metadata
	11, // 6: svapi.ListPublishedVideoRequest.pagination:type_name -> svapi.PaginationRequest
	10, // 7: svapi.ListPublishedVideoResponse.meta:type_name -> svapi.Metadata
	1,  // 8: svapi.ListPublishedVideoResponse.videos:type_name -> svapi.Video
	12, // 9: svapi.ListPublishedVideoResponse.pagination:type_name -> svapi.PaginationResponse
	2,  // 10: svapi.ShortVideoCoreVideoService.FeedShortVideo:input_type -> svapi.FeedShortVideoRequest
	4,  // 11: svapi.ShortVideoCoreVideoService.GetVideoById:input_type -> svapi.GetVideoByIdRequest
	6,  // 12: svapi.ShortVideoCoreVideoService.PublishVideo:input_type -> svapi.PublishVideoRequest
	8,  // 13: svapi.ShortVideoCoreVideoService.ListPublishedVideo:input_type -> svapi.ListPublishedVideoRequest
	3,  // 14: svapi.ShortVideoCoreVideoService.FeedShortVideo:output_type -> svapi.FeedShortVideoResponse
	5,  // 15: svapi.ShortVideoCoreVideoService.GetVideoById:output_type -> svapi.GetVideoByIdResponse
	7,  // 16: svapi.ShortVideoCoreVideoService.PublishVideo:output_type -> svapi.PublishVideoResponse
	9,  // 17: svapi.ShortVideoCoreVideoService.ListPublishedVideo:output_type -> svapi.ListPublishedVideoResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_svapi_video_proto_init() }
func file_svapi_video_proto_init() {
	if File_svapi_video_proto != nil {
		return
	}
	file_svapi_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_svapi_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoAuthor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svapi_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svapi_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedShortVideoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svapi_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedShortVideoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svapi_video_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoByIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svapi_video_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoByIdResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svapi_video_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishVideoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svapi_video_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishVideoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svapi_video_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPublishedVideoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svapi_video_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPublishedVideoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svapi_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svapi_video_proto_goTypes,
		DependencyIndexes: file_svapi_video_proto_depIdxs,
		MessageInfos:      file_svapi_video_proto_msgTypes,
	}.Build()
	File_svapi_video_proto = out.File
	file_svapi_video_proto_rawDesc = nil
	file_svapi_video_proto_goTypes = nil
	file_svapi_video_proto_depIdxs = nil
}
