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
	grocerylistitem "github.com/PaluMacil/dan2/ent/grocerylistitem"
	empty "github.com/golang/protobuf/ptypes/empty"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	strconv "strconv"
)

// GroceryListItemService implements GroceryListItemServiceServer
type GroceryListItemService struct {
	client *ent.Client
	UnimplementedGroceryListItemServiceServer
}

// NewGroceryListItemService returns a new GroceryListItemService
func NewGroceryListItemService(client *ent.Client) *GroceryListItemService {
	return &GroceryListItemService{
		client: client,
	}
}

// toProtoGroceryListItem transforms the ent type to the pb type
func toProtoGroceryListItem(e *ent.GroceryListItem) (*GroceryListItem, error) {
	v := &GroceryListItem{}
	created_at := timestamppb.New(e.CreatedAt)
	v.CreatedAt = created_at
	id := int64(e.ID)
	v.Id = id
	name := e.Name
	v.Name = name
	note := e.Note
	v.Note = note
	quantity := int64(e.Quantity)
	v.Quantity = quantity
	if edg := e.Edges.GroceryList; edg != nil {
		id := int64(edg.ID)
		v.GroceryList = &GroceryList{
			Id: id,
		}
	}
	return v, nil
}

// toProtoGroceryListItemList transforms a list of ent type to a list of pb type
func toProtoGroceryListItemList(e []*ent.GroceryListItem) ([]*GroceryListItem, error) {
	var pbList []*GroceryListItem
	for _, entEntity := range e {
		pbEntity, err := toProtoGroceryListItem(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements GroceryListItemServiceServer.Create
func (svc *GroceryListItemService) Create(ctx context.Context, req *CreateGroceryListItemRequest) (*GroceryListItem, error) {
	grocerylistitem := req.GetGroceryListItem()
	m, err := svc.createBuilder(grocerylistitem)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoGroceryListItem(res)
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

// Get implements GroceryListItemServiceServer.Get
func (svc *GroceryListItemService) Get(ctx context.Context, req *GetGroceryListItemRequest) (*GroceryListItem, error) {
	var (
		err error
		get *ent.GroceryListItem
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetGroceryListItemRequest_VIEW_UNSPECIFIED, GetGroceryListItemRequest_BASIC:
		get, err = svc.client.GroceryListItem.Get(ctx, id)
	case GetGroceryListItemRequest_WITH_EDGE_IDS:
		get, err = svc.client.GroceryListItem.Query().
			Where(grocerylistitem.ID(id)).
			WithGroceryList(func(query *ent.GroceryListQuery) {
				query.Select(grocerylist.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoGroceryListItem(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements GroceryListItemServiceServer.Update
func (svc *GroceryListItemService) Update(ctx context.Context, req *UpdateGroceryListItemRequest) (*GroceryListItem, error) {
	grocerylistitem := req.GetGroceryListItem()
	grocerylistitemID := int(grocerylistitem.GetId())
	m := svc.client.GroceryListItem.UpdateOneID(grocerylistitemID)
	grocerylistitemCreatedAt := runtime.ExtractTime(grocerylistitem.GetCreatedAt())
	m.SetCreatedAt(grocerylistitemCreatedAt)
	grocerylistitemName := grocerylistitem.GetName()
	m.SetName(grocerylistitemName)
	grocerylistitemNote := grocerylistitem.GetNote()
	m.SetNote(grocerylistitemNote)
	grocerylistitemQuantity := int(grocerylistitem.GetQuantity())
	m.SetQuantity(grocerylistitemQuantity)
	if grocerylistitem.GetGroceryList() != nil {
		grocerylistitemGroceryList := int(grocerylistitem.GetGroceryList().GetId())
		m.SetGroceryListID(grocerylistitemGroceryList)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoGroceryListItem(res)
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

// Delete implements GroceryListItemServiceServer.Delete
func (svc *GroceryListItemService) Delete(ctx context.Context, req *DeleteGroceryListItemRequest) (*empty.Empty, error) {
	var err error
	id := int(req.GetId())
	err = svc.client.GroceryListItem.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements GroceryListItemServiceServer.List
func (svc *GroceryListItemService) List(ctx context.Context, req *ListGroceryListItemRequest) (*ListGroceryListItemResponse, error) {
	var (
		err      error
		entList  []*ent.GroceryListItem
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.GroceryListItem.Query().
		Order(ent.Desc(grocerylistitem.FieldID)).
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
			Where(grocerylistitem.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListGroceryListItemRequest_VIEW_UNSPECIFIED, ListGroceryListItemRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListGroceryListItemRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithGroceryList(func(query *ent.GroceryListQuery) {
				query.Select(grocerylist.FieldID)
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
		protoList, err := toProtoGroceryListItemList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListGroceryListItemResponse{
			GroceryListItemList: protoList,
			NextPageToken:       nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements GroceryListItemServiceServer.BatchCreate
func (svc *GroceryListItemService) BatchCreate(ctx context.Context, req *BatchCreateGroceryListItemsRequest) (*BatchCreateGroceryListItemsResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.GroceryListItemCreate, len(requests))
	for i, req := range requests {
		grocerylistitem := req.GetGroceryListItem()
		var err error
		bulk[i], err = svc.createBuilder(grocerylistitem)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.GroceryListItem.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoGroceryListItemList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreateGroceryListItemsResponse{
			GroceryListItems: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *GroceryListItemService) createBuilder(grocerylistitem *GroceryListItem) (*ent.GroceryListItemCreate, error) {
	m := svc.client.GroceryListItem.Create()
	grocerylistitemCreatedAt := runtime.ExtractTime(grocerylistitem.GetCreatedAt())
	m.SetCreatedAt(grocerylistitemCreatedAt)
	grocerylistitemName := grocerylistitem.GetName()
	m.SetName(grocerylistitemName)
	grocerylistitemNote := grocerylistitem.GetNote()
	m.SetNote(grocerylistitemNote)
	grocerylistitemQuantity := int(grocerylistitem.GetQuantity())
	m.SetQuantity(grocerylistitemQuantity)
	if grocerylistitem.GetGroceryList() != nil {
		grocerylistitemGroceryList := int(grocerylistitem.GetGroceryList().GetId())
		m.SetGroceryListID(grocerylistitemGroceryList)
	}
	return m, nil
}