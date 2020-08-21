package object

import (
	"context"
)

type Service interface {
	Get(context.Context, *GetRequest) (GetObjectStreamer, error)
	Put(context.Context) (PutObjectStreamer, error)
	Head(context.Context, *HeadRequest) (*HeadResponse, error)
	Search(context.Context, *SearchRequest) (SearchObjectStreamer, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	GetRange(context.Context, *GetRangeRequest) (GetRangeObjectStreamer, error)
	GetRangeHash(context.Context, *GetRangeHashRequest) (*GetRangeHashResponse, error)
}
