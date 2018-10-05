// Code generated by Prisma CLI (https://github.com/prisma/prisma). DO NOT EDIT.

package prisma

import (
	"context"

	"github.com/prisma/prisma-client-lib-go"

	"github.com/machinebox/graphql"
)

func Str(v string) *string { return &v }
func Int32(v int32) *int32 { return &v }
func Bool(v bool) *bool    { return &v }

type BatchPayloadExec struct {
	exec *prisma.BatchPayloadExec
}

func (exec *BatchPayloadExec) Exec(ctx context.Context) (BatchPayload, error) {
	bp, err := exec.exec.Exec(ctx)
	return BatchPayload(bp), err
}

type BatchPayload struct {
	Count int64 `json:"count"`
}

type Aggregate struct {
	Count int64 `json:"count"`
}

type Client struct {
	Client *prisma.Client
}

type Options struct {
	Endpoint string
}

func New(options *Options, opts ...graphql.ClientOption) *Client {
	endpoint := DefaultEndpoint
	if options != nil {
		endpoint = options.Endpoint
	}
	return &Client{
		Client: prisma.New(endpoint, opts...),
	}
}

var DefaultEndpoint = "`http://localhost:4466/go-orm/dev`"

func (client *Client) Todo(params TodoWhereUniqueInput) *TodoExec {
	ret := client.Client.GetOne(
		nil,
		params,
		[2]string{"TodoWhereUniqueInput!", "Todo"},
		"todo",
		[]string{"id", "text", "done"})

	return &TodoExec{ret}
}

type TodoesParams struct {
	Where   *TodoWhereInput   `json:"where,omitempty"`
	OrderBy *TodoOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

func (client *Client) Todoes(params *TodoesParams) *TodoExecArray {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := client.Client.GetMany(
		nil,
		wparams,
		[3]string{"TodoWhereInput", "TodoOrderByInput", "Todo"},
		"todoes",
		[]string{"id", "text", "done"})

	return &TodoExecArray{ret}
}

type TodoesConnectionParams struct {
	Where   *TodoWhereInput   `json:"where,omitempty"`
	OrderBy *TodoOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

func (client *Client) TodoesConnection(params *TodoesConnectionParams) TodoConnectionExec {
	panic("not implemented")
}

func (client *Client) User(params UserWhereUniqueInput) *UserExec {
	ret := client.Client.GetOne(
		nil,
		params,
		[2]string{"UserWhereUniqueInput!", "User"},
		"user",
		[]string{"id", "name"})

	return &UserExec{ret}
}

type UsersParams struct {
	Where   *UserWhereInput   `json:"where,omitempty"`
	OrderBy *UserOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

func (client *Client) Users(params *UsersParams) *UserExecArray {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := client.Client.GetMany(
		nil,
		wparams,
		[3]string{"UserWhereInput", "UserOrderByInput", "User"},
		"users",
		[]string{"id", "name"})

	return &UserExecArray{ret}
}

type UsersConnectionParams struct {
	Where   *UserWhereInput   `json:"where,omitempty"`
	OrderBy *UserOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

func (client *Client) UsersConnection(params *UsersConnectionParams) UserConnectionExec {
	panic("not implemented")
}

func (client *Client) CreateTodo(params TodoCreateInput) *TodoExec {
	ret := client.Client.Create(
		params,
		[2]string{"TodoCreateInput!", "Todo"},
		"createTodo",
		[]string{"id", "text", "done"})

	return &TodoExec{ret}
}

type TodoUpdateParams struct {
	Data  TodoUpdateInput      `json:"data"`
	Where TodoWhereUniqueInput `json:"where"`
}

func (client *Client) UpdateTodo(params TodoUpdateParams) *TodoExec {
	ret := client.Client.Update(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[3]string{"TodoUpdateInput!", "TodoWhereUniqueInput!", "Todo"},
		"updateTodo",
		[]string{"id", "text", "done"})

	return &TodoExec{ret}
}

type TodoUpdateManyParams struct {
	Data  TodoUpdateInput `json:"data"`
	Where *TodoWhereInput `json:"where,omitempty"`
}

func (client *Client) UpdateManyTodoes(params TodoUpdateManyParams) *BatchPayloadExec {
	exec := client.Client.UpdateMany(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[2]string{"TodoUpdateInput!", "TodoWhereInput"},
		"updateManyTodoes")
	return &BatchPayloadExec{exec}
}

type TodoUpsertParams struct {
	Where  TodoWhereUniqueInput `json:"where"`
	Create TodoCreateInput      `json:"create"`
	Update TodoUpdateInput      `json:"update"`
}

func (client *Client) UpsertTodo(params TodoUpsertParams) *TodoExec {
	uparams := &prisma.UpsertParams{
		Where:  params.Where,
		Create: params.Create,
		Update: params.Update,
	}
	ret := client.Client.Upsert(
		uparams,
		[4]string{"TodoWhereUniqueInput!", "TodoCreateInput!", "TodoUpdateInput!", "Todo"},
		"upsertTodo",
		[]string{"id", "text", "done"})

	return &TodoExec{ret}
}

func (client *Client) DeleteTodo(params TodoWhereUniqueInput) *TodoExec {
	ret := client.Client.Delete(
		params,
		[2]string{"TodoWhereUniqueInput!", "Todo"},
		"deleteTodo",
		[]string{"id", "text", "done"})

	return &TodoExec{ret}
}

func (client *Client) DeleteManyTodoes(params *TodoWhereInput) *BatchPayloadExec {
	exec := client.Client.DeleteMany(params, "TodoWhereInput", "deleteManyTodoes")
	return &BatchPayloadExec{exec}
}

func (client *Client) CreateUser(params UserCreateInput) *UserExec {
	ret := client.Client.Create(
		params,
		[2]string{"UserCreateInput!", "User"},
		"createUser",
		[]string{"id", "name"})

	return &UserExec{ret}
}

type UserUpdateParams struct {
	Data  UserUpdateInput      `json:"data"`
	Where UserWhereUniqueInput `json:"where"`
}

func (client *Client) UpdateUser(params UserUpdateParams) *UserExec {
	ret := client.Client.Update(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[3]string{"UserUpdateInput!", "UserWhereUniqueInput!", "User"},
		"updateUser",
		[]string{"id", "name"})

	return &UserExec{ret}
}

type UserUpdateManyParams struct {
	Data  UserUpdateInput `json:"data"`
	Where *UserWhereInput `json:"where,omitempty"`
}

func (client *Client) UpdateManyUsers(params UserUpdateManyParams) *BatchPayloadExec {
	exec := client.Client.UpdateMany(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[2]string{"UserUpdateInput!", "UserWhereInput"},
		"updateManyUsers")
	return &BatchPayloadExec{exec}
}

type UserUpsertParams struct {
	Where  UserWhereUniqueInput `json:"where"`
	Create UserCreateInput      `json:"create"`
	Update UserUpdateInput      `json:"update"`
}

func (client *Client) UpsertUser(params UserUpsertParams) *UserExec {
	uparams := &prisma.UpsertParams{
		Where:  params.Where,
		Create: params.Create,
		Update: params.Update,
	}
	ret := client.Client.Upsert(
		uparams,
		[4]string{"UserWhereUniqueInput!", "UserCreateInput!", "UserUpdateInput!", "User"},
		"upsertUser",
		[]string{"id", "name"})

	return &UserExec{ret}
}

func (client *Client) DeleteUser(params UserWhereUniqueInput) *UserExec {
	ret := client.Client.Delete(
		params,
		[2]string{"UserWhereUniqueInput!", "User"},
		"deleteUser",
		[]string{"id", "name"})

	return &UserExec{ret}
}

func (client *Client) DeleteManyUsers(params *UserWhereInput) *BatchPayloadExec {
	exec := client.Client.DeleteMany(params, "UserWhereInput", "deleteManyUsers")
	return &BatchPayloadExec{exec}
}

type TodoOrderByInput string

const (
	TodoOrderByInputIDAsc         TodoOrderByInput = "id_ASC"
	TodoOrderByInputIDDesc        TodoOrderByInput = "id_DESC"
	TodoOrderByInputTextAsc       TodoOrderByInput = "text_ASC"
	TodoOrderByInputTextDesc      TodoOrderByInput = "text_DESC"
	TodoOrderByInputDoneAsc       TodoOrderByInput = "done_ASC"
	TodoOrderByInputDoneDesc      TodoOrderByInput = "done_DESC"
	TodoOrderByInputCreatedAtAsc  TodoOrderByInput = "createdAt_ASC"
	TodoOrderByInputCreatedAtDesc TodoOrderByInput = "createdAt_DESC"
	TodoOrderByInputUpdatedAtAsc  TodoOrderByInput = "updatedAt_ASC"
	TodoOrderByInputUpdatedAtDesc TodoOrderByInput = "updatedAt_DESC"
)

type UserOrderByInput string

const (
	UserOrderByInputIDAsc         UserOrderByInput = "id_ASC"
	UserOrderByInputIDDesc        UserOrderByInput = "id_DESC"
	UserOrderByInputNameAsc       UserOrderByInput = "name_ASC"
	UserOrderByInputNameDesc      UserOrderByInput = "name_DESC"
	UserOrderByInputCreatedAtAsc  UserOrderByInput = "createdAt_ASC"
	UserOrderByInputCreatedAtDesc UserOrderByInput = "createdAt_DESC"
	UserOrderByInputUpdatedAtAsc  UserOrderByInput = "updatedAt_ASC"
	UserOrderByInputUpdatedAtDesc UserOrderByInput = "updatedAt_DESC"
)

type MutationType string

const (
	MutationTypeCreated MutationType = "CREATED"
	MutationTypeUpdated MutationType = "UPDATED"
	MutationTypeDeleted MutationType = "DELETED"
)

type UserUpdateDataInput struct {
	Name *string `json:"name,omitempty"`
}

type TodoWhereUniqueInput struct {
	ID *string `json:"id,omitempty"`
}

type TodoUpdateInput struct {
	Text *string             `json:"text,omitempty"`
	Done *bool               `json:"done,omitempty"`
	User *UserUpdateOneInput `json:"user,omitempty"`
}

type TodoWhereInput struct {
	ID                *string          `json:"id,omitempty"`
	IDNot             *string          `json:"id_not,omitempty"`
	IDIn              []string         `json:"id_in,omitempty"`
	IDNotIn           []string         `json:"id_not_in,omitempty"`
	IDLt              *string          `json:"id_lt,omitempty"`
	IDLte             *string          `json:"id_lte,omitempty"`
	IDGt              *string          `json:"id_gt,omitempty"`
	IDGte             *string          `json:"id_gte,omitempty"`
	IDContains        *string          `json:"id_contains,omitempty"`
	IDNotContains     *string          `json:"id_not_contains,omitempty"`
	IDStartsWith      *string          `json:"id_starts_with,omitempty"`
	IDNotStartsWith   *string          `json:"id_not_starts_with,omitempty"`
	IDEndsWith        *string          `json:"id_ends_with,omitempty"`
	IDNotEndsWith     *string          `json:"id_not_ends_with,omitempty"`
	Text              *string          `json:"text,omitempty"`
	TextNot           *string          `json:"text_not,omitempty"`
	TextIn            []string         `json:"text_in,omitempty"`
	TextNotIn         []string         `json:"text_not_in,omitempty"`
	TextLt            *string          `json:"text_lt,omitempty"`
	TextLte           *string          `json:"text_lte,omitempty"`
	TextGt            *string          `json:"text_gt,omitempty"`
	TextGte           *string          `json:"text_gte,omitempty"`
	TextContains      *string          `json:"text_contains,omitempty"`
	TextNotContains   *string          `json:"text_not_contains,omitempty"`
	TextStartsWith    *string          `json:"text_starts_with,omitempty"`
	TextNotStartsWith *string          `json:"text_not_starts_with,omitempty"`
	TextEndsWith      *string          `json:"text_ends_with,omitempty"`
	TextNotEndsWith   *string          `json:"text_not_ends_with,omitempty"`
	Done              *bool            `json:"done,omitempty"`
	DoneNot           *bool            `json:"done_not,omitempty"`
	User              *UserWhereInput  `json:"user,omitempty"`
	And               []TodoWhereInput `json:"AND,omitempty"`
	Or                []TodoWhereInput `json:"OR,omitempty"`
	Not               []TodoWhereInput `json:"NOT,omitempty"`
}

type UserCreateInput struct {
	Name string `json:"name"`
}

type UserUpdateInput struct {
	Name *string `json:"name,omitempty"`
}

type TodoCreateInput struct {
	Text string             `json:"text"`
	Done bool               `json:"done"`
	User UserCreateOneInput `json:"user"`
}

type UserCreateOneInput struct {
	Create  *UserCreateInput      `json:"create,omitempty"`
	Connect *UserWhereUniqueInput `json:"connect,omitempty"`
}

type TodoSubscriptionWhereInput struct {
	MutationIn                 []MutationType               `json:"mutation_in,omitempty"`
	UpdatedFieldsContains      *string                      `json:"updatedFields_contains,omitempty"`
	UpdatedFieldsContainsEvery []string                     `json:"updatedFields_contains_every,omitempty"`
	UpdatedFieldsContainsSome  []string                     `json:"updatedFields_contains_some,omitempty"`
	Node                       *TodoWhereInput              `json:"node,omitempty"`
	And                        []TodoSubscriptionWhereInput `json:"AND,omitempty"`
	Or                         []TodoSubscriptionWhereInput `json:"OR,omitempty"`
	Not                        []TodoSubscriptionWhereInput `json:"NOT,omitempty"`
}

type UserWhereUniqueInput struct {
	ID *string `json:"id,omitempty"`
}

type UserUpsertNestedInput struct {
	Update UserUpdateDataInput `json:"update"`
	Create UserCreateInput     `json:"create"`
}

type UserWhereInput struct {
	ID                *string          `json:"id,omitempty"`
	IDNot             *string          `json:"id_not,omitempty"`
	IDIn              []string         `json:"id_in,omitempty"`
	IDNotIn           []string         `json:"id_not_in,omitempty"`
	IDLt              *string          `json:"id_lt,omitempty"`
	IDLte             *string          `json:"id_lte,omitempty"`
	IDGt              *string          `json:"id_gt,omitempty"`
	IDGte             *string          `json:"id_gte,omitempty"`
	IDContains        *string          `json:"id_contains,omitempty"`
	IDNotContains     *string          `json:"id_not_contains,omitempty"`
	IDStartsWith      *string          `json:"id_starts_with,omitempty"`
	IDNotStartsWith   *string          `json:"id_not_starts_with,omitempty"`
	IDEndsWith        *string          `json:"id_ends_with,omitempty"`
	IDNotEndsWith     *string          `json:"id_not_ends_with,omitempty"`
	Name              *string          `json:"name,omitempty"`
	NameNot           *string          `json:"name_not,omitempty"`
	NameIn            []string         `json:"name_in,omitempty"`
	NameNotIn         []string         `json:"name_not_in,omitempty"`
	NameLt            *string          `json:"name_lt,omitempty"`
	NameLte           *string          `json:"name_lte,omitempty"`
	NameGt            *string          `json:"name_gt,omitempty"`
	NameGte           *string          `json:"name_gte,omitempty"`
	NameContains      *string          `json:"name_contains,omitempty"`
	NameNotContains   *string          `json:"name_not_contains,omitempty"`
	NameStartsWith    *string          `json:"name_starts_with,omitempty"`
	NameNotStartsWith *string          `json:"name_not_starts_with,omitempty"`
	NameEndsWith      *string          `json:"name_ends_with,omitempty"`
	NameNotEndsWith   *string          `json:"name_not_ends_with,omitempty"`
	And               []UserWhereInput `json:"AND,omitempty"`
	Or                []UserWhereInput `json:"OR,omitempty"`
	Not               []UserWhereInput `json:"NOT,omitempty"`
}

type UserUpdateOneInput struct {
	Create  *UserCreateInput       `json:"create,omitempty"`
	Update  *UserUpdateDataInput   `json:"update,omitempty"`
	Upsert  *UserUpsertNestedInput `json:"upsert,omitempty"`
	Delete  *bool                  `json:"delete,omitempty"`
	Connect *UserWhereUniqueInput  `json:"connect,omitempty"`
}

type UserSubscriptionWhereInput struct {
	MutationIn                 []MutationType               `json:"mutation_in,omitempty"`
	UpdatedFieldsContains      *string                      `json:"updatedFields_contains,omitempty"`
	UpdatedFieldsContainsEvery []string                     `json:"updatedFields_contains_every,omitempty"`
	UpdatedFieldsContainsSome  []string                     `json:"updatedFields_contains_some,omitempty"`
	Node                       *UserWhereInput              `json:"node,omitempty"`
	And                        []UserSubscriptionWhereInput `json:"AND,omitempty"`
	Or                         []UserSubscriptionWhereInput `json:"OR,omitempty"`
	Not                        []UserSubscriptionWhereInput `json:"NOT,omitempty"`
}

type UserSubscriptionPayloadExec struct {
	exec *prisma.Exec
}

func (instance *UserSubscriptionPayloadExec) Node() *UserExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "User"},
		"node",
		[]string{"id", "name"})

	return &UserExec{ret}
}

func (instance *UserSubscriptionPayloadExec) PreviousValues() *UserPreviousValuesExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "UserPreviousValues"},
		"previousValues",
		[]string{"id", "name"})

	return &UserPreviousValuesExec{ret}
}

func (instance UserSubscriptionPayloadExec) Exec(ctx context.Context) (UserSubscriptionPayload, error) {
	var v UserSubscriptionPayload
	err := instance.exec.Exec(ctx, &v)
	return v, err
}

func (instance UserSubscriptionPayloadExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserSubscriptionPayloadExecArray struct {
	exec *prisma.Exec
}

func (instance UserSubscriptionPayloadExecArray) Exec(ctx context.Context) ([]UserSubscriptionPayload, error) {
	var v []UserSubscriptionPayload
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserSubscriptionPayload struct {
	UpdatedFields []string `json:"updatedFields,omitempty"`
}

type TodoSubscriptionPayloadExec struct {
	exec *prisma.Exec
}

func (instance *TodoSubscriptionPayloadExec) Node() *TodoExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "Todo"},
		"node",
		[]string{"id", "text", "done"})

	return &TodoExec{ret}
}

func (instance *TodoSubscriptionPayloadExec) PreviousValues() *TodoPreviousValuesExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "TodoPreviousValues"},
		"previousValues",
		[]string{"id", "text", "done"})

	return &TodoPreviousValuesExec{ret}
}

func (instance TodoSubscriptionPayloadExec) Exec(ctx context.Context) (TodoSubscriptionPayload, error) {
	var v TodoSubscriptionPayload
	err := instance.exec.Exec(ctx, &v)
	return v, err
}

func (instance TodoSubscriptionPayloadExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type TodoSubscriptionPayloadExecArray struct {
	exec *prisma.Exec
}

func (instance TodoSubscriptionPayloadExecArray) Exec(ctx context.Context) ([]TodoSubscriptionPayload, error) {
	var v []TodoSubscriptionPayload
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type TodoSubscriptionPayload struct {
	UpdatedFields []string `json:"updatedFields,omitempty"`
}

type UserPreviousValuesExec struct {
	exec *prisma.Exec
}

func (instance UserPreviousValuesExec) Exec(ctx context.Context) (UserPreviousValues, error) {
	var v UserPreviousValues
	err := instance.exec.Exec(ctx, &v)
	return v, err
}

func (instance UserPreviousValuesExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserPreviousValuesExecArray struct {
	exec *prisma.Exec
}

func (instance UserPreviousValuesExecArray) Exec(ctx context.Context) ([]UserPreviousValues, error) {
	var v []UserPreviousValues
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserPreviousValues struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TodoPreviousValuesExec struct {
	exec *prisma.Exec
}

func (instance TodoPreviousValuesExec) Exec(ctx context.Context) (TodoPreviousValues, error) {
	var v TodoPreviousValues
	err := instance.exec.Exec(ctx, &v)
	return v, err
}

func (instance TodoPreviousValuesExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type TodoPreviousValuesExecArray struct {
	exec *prisma.Exec
}

func (instance TodoPreviousValuesExecArray) Exec(ctx context.Context) ([]TodoPreviousValues, error) {
	var v []TodoPreviousValues
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type TodoPreviousValues struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type TodoExec struct {
	exec *prisma.Exec
}

func (instance *TodoExec) User() *UserExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "User"},
		"user",
		[]string{"id", "name"})

	return &UserExec{ret}
}

func (instance TodoExec) Exec(ctx context.Context) (Todo, error) {
	var v Todo
	err := instance.exec.Exec(ctx, &v)
	return v, err
}

func (instance TodoExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type TodoExecArray struct {
	exec *prisma.Exec
}

func (instance TodoExecArray) Exec(ctx context.Context) ([]Todo, error) {
	var v []Todo
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type UserExec struct {
	exec *prisma.Exec
}

func (instance UserExec) Exec(ctx context.Context) (User, error) {
	var v User
	err := instance.exec.Exec(ctx, &v)
	return v, err
}

func (instance UserExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserExecArray struct {
	exec *prisma.Exec
}

func (instance UserExecArray) Exec(ctx context.Context) ([]User, error) {
	var v []User
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserEdgeExec struct {
	exec *prisma.Exec
}

func (instance *UserEdgeExec) Node() *UserExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "User"},
		"node",
		[]string{"id", "name"})

	return &UserExec{ret}
}

func (instance UserEdgeExec) Exec(ctx context.Context) (UserEdge, error) {
	var v UserEdge
	err := instance.exec.Exec(ctx, &v)
	return v, err
}

func (instance UserEdgeExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserEdgeExecArray struct {
	exec *prisma.Exec
}

func (instance UserEdgeExecArray) Exec(ctx context.Context) ([]UserEdge, error) {
	var v []UserEdge
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserEdge struct {
	Cursor string `json:"cursor"`
}

type TodoEdgeExec struct {
	exec *prisma.Exec
}

func (instance *TodoEdgeExec) Node() *TodoExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "Todo"},
		"node",
		[]string{"id", "text", "done"})

	return &TodoExec{ret}
}

func (instance TodoEdgeExec) Exec(ctx context.Context) (TodoEdge, error) {
	var v TodoEdge
	err := instance.exec.Exec(ctx, &v)
	return v, err
}

func (instance TodoEdgeExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type TodoEdgeExecArray struct {
	exec *prisma.Exec
}

func (instance TodoEdgeExecArray) Exec(ctx context.Context) ([]TodoEdge, error) {
	var v []TodoEdge
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type TodoEdge struct {
	Cursor string `json:"cursor"`
}

type UserConnectionExec struct {
	exec *prisma.Exec
}

func (instance *UserConnectionExec) PageInfo() *PageInfoExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "PageInfo"},
		"pageInfo",
		[]string{"hasNextPage", "hasPreviousPage", "startCursor", "endCursor"})

	return &PageInfoExec{ret}
}

func (instance *UserConnectionExec) Edges() *UserEdgeExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "UserEdge"},
		"edges",
		[]string{"cursor"})

	return &UserEdgeExec{ret}
}

func (instance *UserConnectionExec) Aggregate(ctx context.Context) (Aggregate, error) {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "AggregateUser"},
		"aggregate",
		[]string{"count"})

	var v Aggregate
	err := ret.Exec(ctx, &v)
	return v, err
}

func (instance UserConnectionExec) Exec(ctx context.Context) (UserConnection, error) {
	var v UserConnection
	err := instance.exec.Exec(ctx, &v)
	return v, err
}

func (instance UserConnectionExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserConnectionExecArray struct {
	exec *prisma.Exec
}

func (instance UserConnectionExecArray) Exec(ctx context.Context) ([]UserConnection, error) {
	var v []UserConnection
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserConnection struct {
}

type TodoConnectionExec struct {
	exec *prisma.Exec
}

func (instance *TodoConnectionExec) PageInfo() *PageInfoExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "PageInfo"},
		"pageInfo",
		[]string{"hasNextPage", "hasPreviousPage", "startCursor", "endCursor"})

	return &PageInfoExec{ret}
}

func (instance *TodoConnectionExec) Edges() *TodoEdgeExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "TodoEdge"},
		"edges",
		[]string{"cursor"})

	return &TodoEdgeExec{ret}
}

func (instance *TodoConnectionExec) Aggregate(ctx context.Context) (Aggregate, error) {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "AggregateTodo"},
		"aggregate",
		[]string{"count"})

	var v Aggregate
	err := ret.Exec(ctx, &v)
	return v, err
}

func (instance TodoConnectionExec) Exec(ctx context.Context) (TodoConnection, error) {
	var v TodoConnection
	err := instance.exec.Exec(ctx, &v)
	return v, err
}

func (instance TodoConnectionExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type TodoConnectionExecArray struct {
	exec *prisma.Exec
}

func (instance TodoConnectionExecArray) Exec(ctx context.Context) ([]TodoConnection, error) {
	var v []TodoConnection
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type TodoConnection struct {
}

type PageInfoExec struct {
	exec *prisma.Exec
}

func (instance PageInfoExec) Exec(ctx context.Context) (PageInfo, error) {
	var v PageInfo
	err := instance.exec.Exec(ctx, &v)
	return v, err
}

func (instance PageInfoExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type PageInfoExecArray struct {
	exec *prisma.Exec
}

func (instance PageInfoExecArray) Exec(ctx context.Context) ([]PageInfo, error) {
	var v []PageInfo
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
	EndCursor       *string `json:"endCursor,omitempty"`
}
