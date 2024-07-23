// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package orm

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Product is an object representing the database table.
type Product struct {
	ProductID   string    `boil:"product_id" json:"product_id" toml:"product_id" yaml:"product_id"`
	Name        string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Description string    `boil:"description" json:"description" toml:"description" yaml:"description"`
	Price       float32   `boil:"price" json:"price" toml:"price" yaml:"price"`
	Stock       int64     `boil:"stock" json:"stock" toml:"stock" yaml:"stock"`
	CategoryID  string    `boil:"category_id" json:"category_id" toml:"category_id" yaml:"category_id"`
	CreatedAt   null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt   null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *productR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L productL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProductColumns = struct {
	ProductID   string
	Name        string
	Description string
	Price       string
	Stock       string
	CategoryID  string
	CreatedAt   string
	UpdatedAt   string
}{
	ProductID:   "product_id",
	Name:        "name",
	Description: "description",
	Price:       "price",
	Stock:       "stock",
	CategoryID:  "category_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

var ProductTableColumns = struct {
	ProductID   string
	Name        string
	Description string
	Price       string
	Stock       string
	CategoryID  string
	CreatedAt   string
	UpdatedAt   string
}{
	ProductID:   "products.product_id",
	Name:        "products.name",
	Description: "products.description",
	Price:       "products.price",
	Stock:       "products.stock",
	CategoryID:  "products.category_id",
	CreatedAt:   "products.created_at",
	UpdatedAt:   "products.updated_at",
}

// Generated where

type whereHelperfloat32 struct{ field string }

func (w whereHelperfloat32) EQ(x float32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat32) NEQ(x float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat32) LT(x float32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat32) LTE(x float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat32) GT(x float32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat32) GTE(x float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperfloat32) IN(slice []float32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperfloat32) NIN(slice []float32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var ProductWhere = struct {
	ProductID   whereHelperstring
	Name        whereHelperstring
	Description whereHelperstring
	Price       whereHelperfloat32
	Stock       whereHelperint64
	CategoryID  whereHelperstring
	CreatedAt   whereHelpernull_Time
	UpdatedAt   whereHelpernull_Time
}{
	ProductID:   whereHelperstring{field: "\"products\".\"product_id\""},
	Name:        whereHelperstring{field: "\"products\".\"name\""},
	Description: whereHelperstring{field: "\"products\".\"description\""},
	Price:       whereHelperfloat32{field: "\"products\".\"price\""},
	Stock:       whereHelperint64{field: "\"products\".\"stock\""},
	CategoryID:  whereHelperstring{field: "\"products\".\"category_id\""},
	CreatedAt:   whereHelpernull_Time{field: "\"products\".\"created_at\""},
	UpdatedAt:   whereHelpernull_Time{field: "\"products\".\"updated_at\""},
}

// ProductRels is where relationship names are stored.
var ProductRels = struct {
	OrderItems string
}{
	OrderItems: "OrderItems",
}

// productR is where relationships are stored.
type productR struct {
	OrderItems OrderItemSlice `boil:"OrderItems" json:"OrderItems" toml:"OrderItems" yaml:"OrderItems"`
}

// NewStruct creates a new relationship struct
func (*productR) NewStruct() *productR {
	return &productR{}
}

func (r *productR) GetOrderItems() OrderItemSlice {
	if r == nil {
		return nil
	}
	return r.OrderItems
}

// productL is where Load methods for each relationship are stored.
type productL struct{}

var (
	productAllColumns            = []string{"product_id", "name", "description", "price", "stock", "category_id", "created_at", "updated_at"}
	productColumnsWithoutDefault = []string{"product_id", "name", "description", "price", "stock", "category_id"}
	productColumnsWithDefault    = []string{"created_at", "updated_at"}
	productPrimaryKeyColumns     = []string{"product_id"}
	productGeneratedColumns      = []string{}
)

type (
	// ProductSlice is an alias for a slice of pointers to Product.
	// This should almost always be used instead of []Product.
	ProductSlice []*Product
	// ProductHook is the signature for custom Product hook methods
	ProductHook func(context.Context, boil.ContextExecutor, *Product) error

	productQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	productType                 = reflect.TypeOf(&Product{})
	productMapping              = queries.MakeStructMapping(productType)
	productPrimaryKeyMapping, _ = queries.BindMapping(productType, productMapping, productPrimaryKeyColumns)
	productInsertCacheMut       sync.RWMutex
	productInsertCache          = make(map[string]insertCache)
	productUpdateCacheMut       sync.RWMutex
	productUpdateCache          = make(map[string]updateCache)
	productUpsertCacheMut       sync.RWMutex
	productUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var productAfterSelectMu sync.Mutex
var productAfterSelectHooks []ProductHook

var productBeforeInsertMu sync.Mutex
var productBeforeInsertHooks []ProductHook
var productAfterInsertMu sync.Mutex
var productAfterInsertHooks []ProductHook

var productBeforeUpdateMu sync.Mutex
var productBeforeUpdateHooks []ProductHook
var productAfterUpdateMu sync.Mutex
var productAfterUpdateHooks []ProductHook

var productBeforeDeleteMu sync.Mutex
var productBeforeDeleteHooks []ProductHook
var productAfterDeleteMu sync.Mutex
var productAfterDeleteHooks []ProductHook

var productBeforeUpsertMu sync.Mutex
var productBeforeUpsertHooks []ProductHook
var productAfterUpsertMu sync.Mutex
var productAfterUpsertHooks []ProductHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Product) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Product) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Product) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Product) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Product) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Product) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Product) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Product) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Product) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProductHook registers your hook function for all future operations.
func AddProductHook(hookPoint boil.HookPoint, productHook ProductHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		productAfterSelectMu.Lock()
		productAfterSelectHooks = append(productAfterSelectHooks, productHook)
		productAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		productBeforeInsertMu.Lock()
		productBeforeInsertHooks = append(productBeforeInsertHooks, productHook)
		productBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		productAfterInsertMu.Lock()
		productAfterInsertHooks = append(productAfterInsertHooks, productHook)
		productAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		productBeforeUpdateMu.Lock()
		productBeforeUpdateHooks = append(productBeforeUpdateHooks, productHook)
		productBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		productAfterUpdateMu.Lock()
		productAfterUpdateHooks = append(productAfterUpdateHooks, productHook)
		productAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		productBeforeDeleteMu.Lock()
		productBeforeDeleteHooks = append(productBeforeDeleteHooks, productHook)
		productBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		productAfterDeleteMu.Lock()
		productAfterDeleteHooks = append(productAfterDeleteHooks, productHook)
		productAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		productBeforeUpsertMu.Lock()
		productBeforeUpsertHooks = append(productBeforeUpsertHooks, productHook)
		productBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		productAfterUpsertMu.Lock()
		productAfterUpsertHooks = append(productAfterUpsertHooks, productHook)
		productAfterUpsertMu.Unlock()
	}
}

// One returns a single product record from the query.
func (q productQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Product, error) {
	o := &Product{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: failed to execute a one query for products")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Product records from the query.
func (q productQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProductSlice, error) {
	var o []*Product

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "orm: failed to assign all query results to Product slice")
	}

	if len(productAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Product records in the query.
func (q productQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to count products rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q productQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "orm: failed to check if products exists")
	}

	return count > 0, nil
}

// OrderItems retrieves all the order_item's OrderItems with an executor.
func (o *Product) OrderItems(mods ...qm.QueryMod) orderItemQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"order_items\".\"product_id\"=?", o.ProductID),
	)

	return OrderItems(queryMods...)
}

// LoadOrderItems allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (productL) LoadOrderItems(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProduct interface{}, mods queries.Applicator) error {
	var slice []*Product
	var object *Product

	if singular {
		var ok bool
		object, ok = maybeProduct.(*Product)
		if !ok {
			object = new(Product)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeProduct)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeProduct))
			}
		}
	} else {
		s, ok := maybeProduct.(*[]*Product)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeProduct)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeProduct))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &productR{}
		}
		args[object.ProductID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &productR{}
			}
			args[obj.ProductID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`order_items`),
		qm.WhereIn(`order_items.product_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load order_items")
	}

	var resultSlice []*OrderItem
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice order_items")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on order_items")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for order_items")
	}

	if len(orderItemAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.OrderItems = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &orderItemR{}
			}
			foreign.R.Product = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ProductID == foreign.ProductID {
				local.R.OrderItems = append(local.R.OrderItems, foreign)
				if foreign.R == nil {
					foreign.R = &orderItemR{}
				}
				foreign.R.Product = local
				break
			}
		}
	}

	return nil
}

// AddOrderItems adds the given related objects to the existing relationships
// of the product, optionally inserting them as new records.
// Appends related to o.R.OrderItems.
// Sets related.R.Product appropriately.
func (o *Product) AddOrderItems(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*OrderItem) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ProductID = o.ProductID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"order_items\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"product_id"}),
				strmangle.WhereClause("\"", "\"", 2, orderItemPrimaryKeyColumns),
			)
			values := []interface{}{o.ProductID, rel.OrderID, rel.ProductID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ProductID = o.ProductID
		}
	}

	if o.R == nil {
		o.R = &productR{
			OrderItems: related,
		}
	} else {
		o.R.OrderItems = append(o.R.OrderItems, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &orderItemR{
				Product: o,
			}
		} else {
			rel.R.Product = o
		}
	}
	return nil
}

// Products retrieves all the records using an executor.
func Products(mods ...qm.QueryMod) productQuery {
	mods = append(mods, qm.From("\"products\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"products\".*"})
	}

	return productQuery{q}
}

// FindProduct retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProduct(ctx context.Context, exec boil.ContextExecutor, productID string, selectCols ...string) (*Product, error) {
	productObj := &Product{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"products\" where \"product_id\"=$1", sel,
	)

	q := queries.Raw(query, productID)

	err := q.Bind(ctx, exec, productObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: unable to select from products")
	}

	if err = productObj.doAfterSelectHooks(ctx, exec); err != nil {
		return productObj, err
	}

	return productObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Product) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no products provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(productColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	productInsertCacheMut.RLock()
	cache, cached := productInsertCache[key]
	productInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			productAllColumns,
			productColumnsWithDefault,
			productColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(productType, productMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(productType, productMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"products\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"products\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "orm: unable to insert into products")
	}

	if !cached {
		productInsertCacheMut.Lock()
		productInsertCache[key] = cache
		productInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Product.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Product) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	productUpdateCacheMut.RLock()
	cache, cached := productUpdateCache[key]
	productUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			productAllColumns,
			productPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("orm: unable to update products, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"products\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, productPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(productType, productMapping, append(wl, productPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update products row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by update for products")
	}

	if !cached {
		productUpdateCacheMut.Lock()
		productUpdateCache[key] = cache
		productUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q productQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all for products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected for products")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProductSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("orm: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"products\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, productPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all in product slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected all in update all product")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Product) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("orm: no products provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(productColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	productUpsertCacheMut.RLock()
	cache, cached := productUpsertCache[key]
	productUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			productAllColumns,
			productColumnsWithDefault,
			productColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			productAllColumns,
			productPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("orm: unable to upsert products, could not build update column list")
		}

		ret := strmangle.SetComplement(productAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(productPrimaryKeyColumns) == 0 {
				return errors.New("orm: unable to upsert products, could not build conflict column list")
			}

			conflict = make([]string, len(productPrimaryKeyColumns))
			copy(conflict, productPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"products\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(productType, productMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(productType, productMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "orm: unable to upsert products")
	}

	if !cached {
		productUpsertCacheMut.Lock()
		productUpsertCache[key] = cache
		productUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Product record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Product) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no Product provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), productPrimaryKeyMapping)
	sql := "DELETE FROM \"products\" WHERE \"product_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete from products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by delete for products")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q productQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("orm: no productQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for products")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProductSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(productBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"products\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, productPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from product slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for products")
	}

	if len(productAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Product) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProduct(ctx, exec, o.ProductID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProductSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProductSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"products\".* FROM \"products\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, productPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "orm: unable to reload all in ProductSlice")
	}

	*o = slice

	return nil
}

// ProductExists checks if the Product row exists.
func ProductExists(ctx context.Context, exec boil.ContextExecutor, productID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"products\" where \"product_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, productID)
	}
	row := exec.QueryRowContext(ctx, sql, productID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "orm: unable to check if products exists")
	}

	return exists, nil
}

// Exists checks if the Product row exists.
func (o *Product) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ProductExists(ctx, exec, o.ProductID)
}
