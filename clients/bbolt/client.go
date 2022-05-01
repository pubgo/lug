package bbolt

import (
	"context"

	"github.com/opentracing/opentracing-go/ext"
	"github.com/pubgo/x/strutil"
	bolt "go.etcd.io/bbolt"

	"github.com/pubgo/lava/core/tracing"
	"github.com/pubgo/lava/logging"
	"github.com/pubgo/lava/logging/logutil"
	"github.com/pubgo/lava/pkg/utils"
)

type Client struct {
	*bolt.DB
	log *logging.Logger
}

func (t *Client) bucket(name string, tx *bolt.Tx) *bolt.Bucket {
	var _, err = tx.CreateBucketIfNotExists(strutil.ToBytes(name))
	logutil.ErrRecord(t.log, err)
	return tx.Bucket([]byte(name))
}

func (t *Client) Set(ctx context.Context, key string, val []byte, names ...string) error {
	return t.Update(ctx, func(bucket *bolt.Bucket) error {
		return bucket.Put(utils.StoB(key), val)
	}, names...)
}

func (t *Client) Get(ctx context.Context, key string, names ...string) (val []byte, err error) {
	return val, t.View(ctx, func(bucket *bolt.Bucket) error {
		val = bucket.Get(utils.StoB(key))
		return nil
	}, names...)
}

func (t *Client) List(ctx context.Context, fn func(k, v []byte) error, names ...string) error {
	return t.View(ctx, func(bucket *bolt.Bucket) error { return bucket.ForEach(fn) }, names...)
}

func (t *Client) Delete(ctx context.Context, key string, names ...string) error {
	return t.Update(ctx, func(bucket *bolt.Bucket) error {
		return bucket.Delete(utils.StoB(key))
	}, names...)
}

func (t *Client) View(ctx context.Context, fn func(*bolt.Bucket) error, names ...string) error {
	name := utils.GetDefault(names...)

	var span = tracing.CreateChild(ctx, name)
	defer span.Finish()
	ext.DBType.Set(span, Name)

	return t.DB.View(func(tx *bolt.Tx) error {
		return fn(t.bucket(name, tx))
	})
}

func (t *Client) Update(ctx context.Context, fn func(*bolt.Bucket) error, names ...string) error {
	name := utils.GetDefault(names...)

	var span = tracing.CreateChild(ctx, name)
	defer span.Finish()
	ext.DBType.Set(span, Name)

	return t.DB.Update(func(tx *bolt.Tx) (err error) {
		return fn(t.bucket(name, tx))
	})
}
