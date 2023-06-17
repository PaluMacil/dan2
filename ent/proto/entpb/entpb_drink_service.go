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
	drink "github.com/PaluMacil/dan2/ent/drink"
	user "github.com/PaluMacil/dan2/ent/user"
	empty "github.com/golang/protobuf/ptypes/empty"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	regexp "regexp"
	strconv "strconv"
	strings "strings"
)

// DrinkService implements DrinkServiceServer
type DrinkService struct {
	client *ent.Client
	UnimplementedDrinkServiceServer
}

// NewDrinkService returns a new DrinkService
func NewDrinkService(client *ent.Client) *DrinkService {
	return &DrinkService{
		client: client,
	}
}

var protoIdentNormalizeRegexpDrink_Type = regexp.MustCompile(`[^a-zA-Z0-9_]+`)

func protoIdentNormalizeDrink_Type(e string) string {
	return protoIdentNormalizeRegexpDrink_Type.ReplaceAllString(e, "_")
}

func toProtoDrink_Type(e drink.Type) Drink_Type {
	if v, ok := Drink_Type_value[strings.ToUpper("TYPE_"+protoIdentNormalizeDrink_Type(string(e)))]; ok {
		return Drink_Type(v)
	}
	return Drink_Type(0)
}

func toEntDrink_Type(e Drink_Type) drink.Type {
	if v, ok := Drink_Type_name[int32(e)]; ok {
		entVal := map[string]string{
			"TYPE_UNKNOWN":    "unknown",
			"TYPE_LIGHT_BEER": "light_beer",
			"TYPE_CRAFT_BEER": "craft_beer",
			"TYPE_WINE":       "wine",
			"TYPE_LIQUOR":     "liquor",
			"TYPE_HIGHBALL":   "highball",
			"TYPE_COCKTAIL":   "cocktail",
		}[v]
		return drink.Type(entVal)
	}
	return ""
}

// toProtoDrink transforms the ent type to the pb type
func toProtoDrink(e *ent.Drink) (*Drink, error) {
	v := &Drink{}
	abv := int32(e.Abv)
	v.Abv = abv
	created_at := timestamppb.New(e.CreatedAt)
	v.CreatedAt = created_at
	day := int64(e.Day)
	v.Day = day
	id := int64(e.ID)
	v.Id = id
	month := int64(e.Month)
	v.Month = month
	note := e.Note
	v.Note = note
	ounces := int32(e.Ounces)
	v.Ounces = ounces
	_type := toProtoDrink_Type(e.Type)
	v.Type = _type
	year := int64(e.Year)
	v.Year = year
	if edg := e.Edges.Owner; edg != nil {
		id := int64(edg.ID)
		v.Owner = &User{
			Id: id,
		}
	}
	return v, nil
}

// toProtoDrinkList transforms a list of ent type to a list of pb type
func toProtoDrinkList(e []*ent.Drink) ([]*Drink, error) {
	var pbList []*Drink
	for _, entEntity := range e {
		pbEntity, err := toProtoDrink(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements DrinkServiceServer.Create
func (svc *DrinkService) Create(ctx context.Context, req *CreateDrinkRequest) (*Drink, error) {
	drink := req.GetDrink()
	m, err := svc.createBuilder(drink)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoDrink(res)
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

// Get implements DrinkServiceServer.Get
func (svc *DrinkService) Get(ctx context.Context, req *GetDrinkRequest) (*Drink, error) {
	var (
		err error
		get *ent.Drink
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetDrinkRequest_VIEW_UNSPECIFIED, GetDrinkRequest_BASIC:
		get, err = svc.client.Drink.Get(ctx, id)
	case GetDrinkRequest_WITH_EDGE_IDS:
		get, err = svc.client.Drink.Query().
			Where(drink.ID(id)).
			WithOwner(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoDrink(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements DrinkServiceServer.Update
func (svc *DrinkService) Update(ctx context.Context, req *UpdateDrinkRequest) (*Drink, error) {
	drink := req.GetDrink()
	drinkID := int(drink.GetId())
	m := svc.client.Drink.UpdateOneID(drinkID)
	drinkAbv := int8(drink.GetAbv())
	m.SetAbv(drinkAbv)
	drinkCreatedAt := runtime.ExtractTime(drink.GetCreatedAt())
	m.SetCreatedAt(drinkCreatedAt)
	drinkDay := int(drink.GetDay())
	m.SetDay(drinkDay)
	drinkMonth := int(drink.GetMonth())
	m.SetMonth(drinkMonth)
	drinkNote := drink.GetNote()
	m.SetNote(drinkNote)
	drinkOunces := int8(drink.GetOunces())
	m.SetOunces(drinkOunces)
	drinkType := toEntDrink_Type(drink.GetType())
	m.SetType(drinkType)
	drinkYear := int(drink.GetYear())
	m.SetYear(drinkYear)
	if drink.GetOwner() != nil {
		drinkOwner := int(drink.GetOwner().GetId())
		m.SetOwnerID(drinkOwner)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoDrink(res)
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

// Delete implements DrinkServiceServer.Delete
func (svc *DrinkService) Delete(ctx context.Context, req *DeleteDrinkRequest) (*empty.Empty, error) {
	var err error
	id := int(req.GetId())
	err = svc.client.Drink.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements DrinkServiceServer.List
func (svc *DrinkService) List(ctx context.Context, req *ListDrinkRequest) (*ListDrinkResponse, error) {
	var (
		err      error
		entList  []*ent.Drink
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Drink.Query().
		Order(ent.Desc(drink.FieldID)).
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
			Where(drink.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListDrinkRequest_VIEW_UNSPECIFIED, ListDrinkRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListDrinkRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithOwner(func(query *ent.UserQuery) {
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
		protoList, err := toProtoDrinkList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListDrinkResponse{
			DrinkList:     protoList,
			NextPageToken: nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements DrinkServiceServer.BatchCreate
func (svc *DrinkService) BatchCreate(ctx context.Context, req *BatchCreateDrinksRequest) (*BatchCreateDrinksResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.DrinkCreate, len(requests))
	for i, req := range requests {
		drink := req.GetDrink()
		var err error
		bulk[i], err = svc.createBuilder(drink)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.Drink.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoDrinkList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreateDrinksResponse{
			Drinks: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *DrinkService) createBuilder(drink *Drink) (*ent.DrinkCreate, error) {
	m := svc.client.Drink.Create()
	drinkAbv := int8(drink.GetAbv())
	m.SetAbv(drinkAbv)
	drinkCreatedAt := runtime.ExtractTime(drink.GetCreatedAt())
	m.SetCreatedAt(drinkCreatedAt)
	drinkDay := int(drink.GetDay())
	m.SetDay(drinkDay)
	drinkMonth := int(drink.GetMonth())
	m.SetMonth(drinkMonth)
	drinkNote := drink.GetNote()
	m.SetNote(drinkNote)
	drinkOunces := int8(drink.GetOunces())
	m.SetOunces(drinkOunces)
	drinkType := toEntDrink_Type(drink.GetType())
	m.SetType(drinkType)
	drinkYear := int(drink.GetYear())
	m.SetYear(drinkYear)
	if drink.GetOwner() != nil {
		drinkOwner := int(drink.GetOwner().GetId())
		m.SetOwnerID(drinkOwner)
	}
	return m, nil
}
