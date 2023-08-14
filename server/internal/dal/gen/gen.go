// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q             = new(Query)
	CustomerSlack *customerSlack
	SysBot        *sysBot
	SysCommunity  *sysCommunity
	SysCustomer   *sysCustomer
	SysUser       *sysUser
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	CustomerSlack = &Q.CustomerSlack
	SysBot = &Q.SysBot
	SysCommunity = &Q.SysCommunity
	SysCustomer = &Q.SysCustomer
	SysUser = &Q.SysUser
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:            db,
		CustomerSlack: newCustomerSlack(db, opts...),
		SysBot:        newSysBot(db, opts...),
		SysCommunity:  newSysCommunity(db, opts...),
		SysCustomer:   newSysCustomer(db, opts...),
		SysUser:       newSysUser(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	CustomerSlack customerSlack
	SysBot        sysBot
	SysCommunity  sysCommunity
	SysCustomer   sysCustomer
	SysUser       sysUser
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:            db,
		CustomerSlack: q.CustomerSlack.clone(db),
		SysBot:        q.SysBot.clone(db),
		SysCommunity:  q.SysCommunity.clone(db),
		SysCustomer:   q.SysCustomer.clone(db),
		SysUser:       q.SysUser.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:            db,
		CustomerSlack: q.CustomerSlack.replaceDB(db),
		SysBot:        q.SysBot.replaceDB(db),
		SysCommunity:  q.SysCommunity.replaceDB(db),
		SysCustomer:   q.SysCustomer.replaceDB(db),
		SysUser:       q.SysUser.replaceDB(db),
	}
}

type queryCtx struct {
	CustomerSlack ICustomerSlackDo
	SysBot        ISysBotDo
	SysCommunity  ISysCommunityDo
	SysCustomer   ISysCustomerDo
	SysUser       ISysUserDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		CustomerSlack: q.CustomerSlack.WithContext(ctx),
		SysBot:        q.SysBot.WithContext(ctx),
		SysCommunity:  q.SysCommunity.WithContext(ctx),
		SysCustomer:   q.SysCustomer.WithContext(ctx),
		SysUser:       q.SysUser.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
