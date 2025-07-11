package usecase

import (
	"context"

	"connectrpc.com/connect"
	"github.com/darwishdev/devkit-api/app/public/adapter"
	"github.com/darwishdev/devkit-api/app/public/repo"
	"github.com/darwishdev/devkit-api/db"
	"github.com/darwishdev/devkit-api/pkg/redisclient"
	"github.com/darwishdev/devkit-api/pkg/resend"
	"github.com/darwishdev/devkit-api/pkg/weaviateclient"
	devkitv1 "github.com/darwishdev/devkit-api/proto_gen/devkit/v1"
	supaapigo "github.com/darwishdev/supaapi-go"
)

type PublicUsecaseInterface interface {
	TranslationList(ctx context.Context) (*devkitv1.TranslationListResponse, error)
	TranslationCreateUpdateBulk(ctx context.Context, req *devkitv1.TranslationCreateUpdateBulkRequest) (*devkitv1.TranslationCreateUpdateBulkResponse, error)
	TranslationFindLocale(ctx context.Context, req *devkitv1.TranslationFindLocaleRequest) (*devkitv1.TranslationFindLocaleResponse, error)
	TranslationDelete(ctx context.Context, req *devkitv1.TranslationDeleteRequest) (*devkitv1.TranslationDeleteResponse, error)
	GalleryList(ctx context.Context, req *devkitv1.GalleryListRequest) (*devkitv1.GalleryListResponse, error)
	FileDeleteByBucket(ctx context.Context, req *devkitv1.FileDeleteByBucketRequest) (*devkitv1.FileDeleteByBucketResponse, error)
	FileDelete(ctx context.Context, req *devkitv1.FileDeleteRequest) (*devkitv1.FileDeleteResponse, error)
	FileList(ctx context.Context, req *devkitv1.FileListRequest) (*devkitv1.FileListResponse, error)
	EmailSend(ctx context.Context, req *devkitv1.EmailSendRequest) (*devkitv1.EmailSendResponse, error)
	BucketList(ctx context.Context, req *devkitv1.BucketListRequest) (*devkitv1.BucketListResponse, error)
	SettingUpdate(ctx context.Context, req *devkitv1.SettingUpdateRequest) error
	SettingFindForUpdate(ctx context.Context, req *devkitv1.SettingFindForUpdateRequest) (*devkitv1.SettingFindForUpdateResponse, error)
	FileCreate(ctx context.Context, req *devkitv1.FileCreateRequest) (*devkitv1.FileCreateResponse, error)
	BucketCreateUpdate(ctx context.Context, req *devkitv1.BucketCreateUpdateRequest) (*devkitv1.BucketCreateUpdateResponse, error)

	IconFind(ctx context.Context, req *devkitv1.IconFindRequest) (*devkitv1.IconFindResponse, error)
	IconCreateUpdateBulk(ctx context.Context, req *devkitv1.IconCreateUpdateBulkRequest) (*devkitv1.IconListResponse, error)
	IconList(ctx context.Context) (*devkitv1.IconListResponse, error)
	FileUploadUrlFind(ctx context.Context, req *devkitv1.FileUploadUrlFindRequest) (*devkitv1.FileUploadUrlFindResponse, error)
	FileCreateBulk(ctx context.Context, req *devkitv1.FileCreateBulkRequest) (*devkitv1.FileCreateBulkResponse, error)

	CommandPalleteSync(ctx context.Context, req *connect.Request[devkitv1.CommandPalleteSyncRequest]) (*devkitv1.CommandPalleteSyncResponse, error)
 CommandPalleteSearch(ctx context.Context, req *connect.Request[devkitv1.CommandPalleteSearchRequest]) (*devkitv1.CommandPalleteSearchResponse, error) 
}

type PublicUsecase struct {
	store          db.Store
	repo           repo.PublicRepoInterface
	adapter        adapter.PublicAdapterInterface
	supaapi        supaapigo.Supaapi
	supaAnonApiKey string
	resendClient   resend.ResendServiceInterface
	weaviateClient weaviateclient.WeaviateClientInterface
	redisClient    redisclient.RedisClientInterface
}

func NewPublicUsecase(store db.Store, supaAnonApiKey string, supaapi supaapigo.Supaapi, redisClient redisclient.RedisClientInterface, resendClient resend.ResendServiceInterface, weaviateClient weaviateclient.WeaviateClientInterface) PublicUsecaseInterface {
	return &PublicUsecase{
		resendClient:   resendClient,
		supaAnonApiKey: supaAnonApiKey,
		supaapi:        supaapi,
		redisClient:    redisClient,
		weaviateClient: weaviateClient,
		adapter:        adapter.NewPublicAdapter(),
		repo:           repo.NewPublicRepo(store),
		store:          store,
	}
}
