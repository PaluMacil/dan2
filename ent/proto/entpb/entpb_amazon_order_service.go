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
	amazonlist "github.com/PaluMacil/dan2/ent/amazonlist"
	amazonorder "github.com/PaluMacil/dan2/ent/amazonorder"
	empty "github.com/golang/protobuf/ptypes/empty"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	strconv "strconv"
)

// AmazonOrderService implements AmazonOrderServiceServer
type AmazonOrderService struct {
	client *ent.Client
	UnimplementedAmazonOrderServiceServer
}

// NewAmazonOrderService returns a new AmazonOrderService
func NewAmazonOrderService(client *ent.Client) *AmazonOrderService {
	return &AmazonOrderService{
		client: client,
	}
}

// toProtoAmazonOrder transforms the ent type to the pb type
func toProtoAmazonOrder(e *ent.AmazonOrder) (*AmazonOrder, error) {
	v := &AmazonOrder{}
	address1 := e.Address1
	v.Address1 = address1
	address2 := e.Address2
	v.Address2 = address2
	brand := e.Brand
	v.Brand = brand
	category := e.Category
	v.Category = category
	city := e.City
	v.City = city
	created_at := timestamppb.New(e.CreatedAt)
	v.CreatedAt = created_at
	id := int64(e.ID)
	v.Id = id
	name := e.Name
	v.Name = name
	ordered_at := timestamppb.New(e.OrderedAt)
	v.OrderedAt = ordered_at
	price := e.Price
	v.Price = price
	refund := e.Refund
	v.Refund = refund
	seller := e.Seller
	v.Seller = seller
	state := e.State
	v.State = state
	tax := e.Tax
	v.Tax = tax
	zip := e.Zip
	v.Zip = zip
	if edg := e.Edges.AmazonList; edg != nil {
		id := int64(edg.ID)
		v.AmazonList = &AmazonList{
			Id: id,
		}
	}
	return v, nil
}

// toProtoAmazonOrderList transforms a list of ent type to a list of pb type
func toProtoAmazonOrderList(e []*ent.AmazonOrder) ([]*AmazonOrder, error) {
	var pbList []*AmazonOrder
	for _, entEntity := range e {
		pbEntity, err := toProtoAmazonOrder(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements AmazonOrderServiceServer.Create
func (svc *AmazonOrderService) Create(ctx context.Context, req *CreateAmazonOrderRequest) (*AmazonOrder, error) {
	amazonorder := req.GetAmazonOrder()
	m, err := svc.createBuilder(amazonorder)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoAmazonOrder(res)
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

// Get implements AmazonOrderServiceServer.Get
func (svc *AmazonOrderService) Get(ctx context.Context, req *GetAmazonOrderRequest) (*AmazonOrder, error) {
	var (
		err error
		get *ent.AmazonOrder
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetAmazonOrderRequest_VIEW_UNSPECIFIED, GetAmazonOrderRequest_BASIC:
		get, err = svc.client.AmazonOrder.Get(ctx, id)
	case GetAmazonOrderRequest_WITH_EDGE_IDS:
		get, err = svc.client.AmazonOrder.Query().
			Where(amazonorder.ID(id)).
			WithAmazonList(func(query *ent.AmazonListQuery) {
				query.Select(amazonlist.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoAmazonOrder(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements AmazonOrderServiceServer.Update
func (svc *AmazonOrderService) Update(ctx context.Context, req *UpdateAmazonOrderRequest) (*AmazonOrder, error) {
	amazonorder := req.GetAmazonOrder()
	amazonorderID := int(amazonorder.GetId())
	m := svc.client.AmazonOrder.UpdateOneID(amazonorderID)
	amazonorderAddress1 := amazonorder.GetAddress1()
	m.SetAddress1(amazonorderAddress1)
	amazonorderAddress2 := amazonorder.GetAddress2()
	m.SetAddress2(amazonorderAddress2)
	amazonorderBrand := amazonorder.GetBrand()
	m.SetBrand(amazonorderBrand)
	amazonorderCategory := amazonorder.GetCategory()
	m.SetCategory(amazonorderCategory)
	amazonorderCity := amazonorder.GetCity()
	m.SetCity(amazonorderCity)
	amazonorderCreatedAt := runtime.ExtractTime(amazonorder.GetCreatedAt())
	m.SetCreatedAt(amazonorderCreatedAt)
	amazonorderName := amazonorder.GetName()
	m.SetName(amazonorderName)
	amazonorderOrderedAt := runtime.ExtractTime(amazonorder.GetOrderedAt())
	m.SetOrderedAt(amazonorderOrderedAt)
	amazonorderPrice := float32(amazonorder.GetPrice())
	m.SetPrice(amazonorderPrice)
	amazonorderRefund := amazonorder.GetRefund()
	m.SetRefund(amazonorderRefund)
	amazonorderSeller := amazonorder.GetSeller()
	m.SetSeller(amazonorderSeller)
	amazonorderState := amazonorder.GetState()
	m.SetState(amazonorderState)
	amazonorderTax := float32(amazonorder.GetTax())
	m.SetTax(amazonorderTax)
	amazonorderZip := amazonorder.GetZip()
	m.SetZip(amazonorderZip)
	if amazonorder.GetAmazonList() != nil {
		amazonorderAmazonList := int(amazonorder.GetAmazonList().GetId())
		m.SetAmazonListID(amazonorderAmazonList)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoAmazonOrder(res)
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

// Delete implements AmazonOrderServiceServer.Delete
func (svc *AmazonOrderService) Delete(ctx context.Context, req *DeleteAmazonOrderRequest) (*empty.Empty, error) {
	var err error
	id := int(req.GetId())
	err = svc.client.AmazonOrder.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements AmazonOrderServiceServer.List
func (svc *AmazonOrderService) List(ctx context.Context, req *ListAmazonOrderRequest) (*ListAmazonOrderResponse, error) {
	var (
		err      error
		entList  []*ent.AmazonOrder
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.AmazonOrder.Query().
		Order(ent.Desc(amazonorder.FieldID)).
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
			Where(amazonorder.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListAmazonOrderRequest_VIEW_UNSPECIFIED, ListAmazonOrderRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListAmazonOrderRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithAmazonList(func(query *ent.AmazonListQuery) {
				query.Select(amazonlist.FieldID)
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
		protoList, err := toProtoAmazonOrderList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListAmazonOrderResponse{
			AmazonOrderList: protoList,
			NextPageToken:   nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements AmazonOrderServiceServer.BatchCreate
func (svc *AmazonOrderService) BatchCreate(ctx context.Context, req *BatchCreateAmazonOrdersRequest) (*BatchCreateAmazonOrdersResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.AmazonOrderCreate, len(requests))
	for i, req := range requests {
		amazonorder := req.GetAmazonOrder()
		var err error
		bulk[i], err = svc.createBuilder(amazonorder)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.AmazonOrder.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoAmazonOrderList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreateAmazonOrdersResponse{
			AmazonOrders: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *AmazonOrderService) createBuilder(amazonorder *AmazonOrder) (*ent.AmazonOrderCreate, error) {
	m := svc.client.AmazonOrder.Create()
	amazonorderAddress1 := amazonorder.GetAddress1()
	m.SetAddress1(amazonorderAddress1)
	amazonorderAddress2 := amazonorder.GetAddress2()
	m.SetAddress2(amazonorderAddress2)
	amazonorderBrand := amazonorder.GetBrand()
	m.SetBrand(amazonorderBrand)
	amazonorderCategory := amazonorder.GetCategory()
	m.SetCategory(amazonorderCategory)
	amazonorderCity := amazonorder.GetCity()
	m.SetCity(amazonorderCity)
	amazonorderCreatedAt := runtime.ExtractTime(amazonorder.GetCreatedAt())
	m.SetCreatedAt(amazonorderCreatedAt)
	amazonorderName := amazonorder.GetName()
	m.SetName(amazonorderName)
	amazonorderOrderedAt := runtime.ExtractTime(amazonorder.GetOrderedAt())
	m.SetOrderedAt(amazonorderOrderedAt)
	amazonorderPrice := float32(amazonorder.GetPrice())
	m.SetPrice(amazonorderPrice)
	amazonorderRefund := amazonorder.GetRefund()
	m.SetRefund(amazonorderRefund)
	amazonorderSeller := amazonorder.GetSeller()
	m.SetSeller(amazonorderSeller)
	amazonorderState := amazonorder.GetState()
	m.SetState(amazonorderState)
	amazonorderTax := float32(amazonorder.GetTax())
	m.SetTax(amazonorderTax)
	amazonorderZip := amazonorder.GetZip()
	m.SetZip(amazonorderZip)
	if amazonorder.GetAmazonList() != nil {
		amazonorderAmazonList := int(amazonorder.GetAmazonList().GetId())
		m.SetAmazonListID(amazonorderAmazonList)
	}
	return m, nil
}
