// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	base64 "encoding/base64"
	entproto "entgo.io/contrib/entproto"
	runtime "entgo.io/contrib/entproto/runtime"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	fmt "fmt"
	ent "github.com/PaluMacil/dan2/ent"
	grocerylist "github.com/PaluMacil/dan2/ent/grocerylist"
	grocerylistshare "github.com/PaluMacil/dan2/ent/grocerylistshare"
	user "github.com/PaluMacil/dan2/ent/user"
	empty "github.com/golang/protobuf/ptypes/empty"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	strconv "strconv"
)

// GroceryListShareService implements GroceryListShareServiceServer
type GroceryListShareService struct {
	client *ent.Client
	UnimplementedGroceryListShareServiceServer
}

// NewGroceryListShareService returns a new GroceryListShareService
func NewGroceryListShareService(client *ent.Client) *GroceryListShareService {
	return &GroceryListShareService{
		client: client,
	}
}

// toProtoGroceryListShare transforms the ent type to the pb type
func toProtoGroceryListShare(e *ent.GroceryListShare) (*GroceryListShare, error) {
	v := &GroceryListShare{}
	can_edit := e.CanEdit
	v.CanEdit = can_edit
	created_at := timestamppb.New(e.CreatedAt)
	v.CreatedAt = created_at
	id := int64(e.ID)
	v.Id = id
	if edg := e.Edges.GroceryList; edg != nil {
		id := int64(edg.ID)
		v.GroceryList = &GroceryList{
			Id: id,
		}
	}
	if edg := e.Edges.User; edg != nil {
		id := int64(edg.ID)
		v.User = &User{
			Id: id,
		}
	}
	return v, nil
}

// toProtoGroceryListShareList transforms a list of ent type to a list of pb type
func toProtoGroceryListShareList(e []*ent.GroceryListShare) ([]*GroceryListShare, error) {
	var pbList []*GroceryListShare
	for _, entEntity := range e {
		pbEntity, err := toProtoGroceryListShare(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements GroceryListShareServiceServer.Create
func (svc *GroceryListShareService) Create(ctx context.Context, req *CreateGroceryListShareRequest) (*GroceryListShare, error) {
	grocerylistshare := req.GetGroceryListShare()
	m, err := svc.createBuilder(grocerylistshare)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoGroceryListShare(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Get implements GroceryListShareServiceServer.Get
func (svc *GroceryListShareService) Get(ctx context.Context, req *GetGroceryListShareRequest) (*GroceryListShare, error) {
	var (
		err error
		get *ent.GroceryListShare
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetGroceryListShareRequest_VIEW_UNSPECIFIED, GetGroceryListShareRequest_BASIC:
		get, err = svc.client.GroceryListShare.Get(ctx, id)
	case GetGroceryListShareRequest_WITH_EDGE_IDS:
		get, err = svc.client.GroceryListShare.Query().
			Where(grocerylistshare.ID(id)).
			WithGroceryList(func(query *ent.GroceryListQuery) {
				query.Select(grocerylist.FieldID)
			}).
			WithUser(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoGroceryListShare(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements GroceryListShareServiceServer.Update
func (svc *GroceryListShareService) Update(ctx context.Context, req *UpdateGroceryListShareRequest) (*GroceryListShare, error) {
	grocerylistshare := req.GetGroceryListShare()
	grocerylistshareID := int(grocerylistshare.GetId())
	m := svc.client.GroceryListShare.UpdateOneID(grocerylistshareID)
	grocerylistshareCanEdit := grocerylistshare.GetCanEdit()
	m.SetCanEdit(grocerylistshareCanEdit)
	grocerylistshareCreatedAt := runtime.ExtractTime(grocerylistshare.GetCreatedAt())
	m.SetCreatedAt(grocerylistshareCreatedAt)
	if grocerylistshare.GetGroceryList() != nil {
		grocerylistshareGroceryList := int(grocerylistshare.GetGroceryList().GetId())
		m.SetGroceryListID(grocerylistshareGroceryList)
	}
	if grocerylistshare.GetUser() != nil {
		grocerylistshareUser := int(grocerylistshare.GetUser().GetId())
		m.SetUserID(grocerylistshareUser)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoGroceryListShare(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Delete implements GroceryListShareServiceServer.Delete
func (svc *GroceryListShareService) Delete(ctx context.Context, req *DeleteGroceryListShareRequest) (*empty.Empty, error) {
	var err error
	id := int(req.GetId())
	err = svc.client.GroceryListShare.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements GroceryListShareServiceServer.List
func (svc *GroceryListShareService) List(ctx context.Context, req *ListGroceryListShareRequest) (*ListGroceryListShareResponse, error) {
	var (
		err      error
		entList  []*ent.GroceryListShare
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.GroceryListShare.Query().
		Order(ent.Desc(grocerylistshare.FieldID)).
		Limit(pageSize + 1)
	if req.GetPageToken() != "" {
		bytes, err := base64.StdEncoding.DecodeString(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		token, err := strconv.ParseInt(string(bytes), 10, 32)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		pageToken := int(token)
		listQuery = listQuery.
			Where(grocerylistshare.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListGroceryListShareRequest_VIEW_UNSPECIFIED, ListGroceryListShareRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListGroceryListShareRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithGroceryList(func(query *ent.GroceryListQuery) {
				query.Select(grocerylist.FieldID)
			}).
			WithUser(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
			}).
			All(ctx)
	}
	switch {
	case err == nil:
		var nextPageToken string
		if len(entList) == pageSize+1 {
			nextPageToken = base64.StdEncoding.EncodeToString(
				[]byte(fmt.Sprintf("%v", entList[len(entList)-1].ID)))
			entList = entList[:len(entList)-1]
		}
		protoList, err := toProtoGroceryListShareList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListGroceryListShareResponse{
			GroceryListShareList: protoList,
			NextPageToken:        nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements GroceryListShareServiceServer.BatchCreate
func (svc *GroceryListShareService) BatchCreate(ctx context.Context, req *BatchCreateGroceryListSharesRequest) (*BatchCreateGroceryListSharesResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.GroceryListShareCreate, len(requests))
	for i, req := range requests {
		grocerylistshare := req.GetGroceryListShare()
		var err error
		bulk[i], err = svc.createBuilder(grocerylistshare)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.GroceryListShare.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoGroceryListShareList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreateGroceryListSharesResponse{
			GroceryListShares: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *GroceryListShareService) createBuilder(grocerylistshare *GroceryListShare) (*ent.GroceryListShareCreate, error) {
	m := svc.client.GroceryListShare.Create()
	grocerylistshareCanEdit := grocerylistshare.GetCanEdit()
	m.SetCanEdit(grocerylistshareCanEdit)
	grocerylistshareCreatedAt := runtime.ExtractTime(grocerylistshare.GetCreatedAt())
	m.SetCreatedAt(grocerylistshareCreatedAt)
	if grocerylistshare.GetGroceryList() != nil {
		grocerylistshareGroceryList := int(grocerylistshare.GetGroceryList().GetId())
		m.SetGroceryListID(grocerylistshareGroceryList)
	}
	if grocerylistshare.GetUser() != nil {
		grocerylistshareUser := int(grocerylistshare.GetUser().GetId())
		m.SetUserID(grocerylistshareUser)
	}
	return m, nil
}
