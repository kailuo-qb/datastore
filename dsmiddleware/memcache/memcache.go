package memcache

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/dsmiddleware/storagecache"
	"time"
)

var _ storagecache.Storage = &cacheHandler{}
var _ datastore.Middleware = &cacheHandler{}

const defaultExpiration = 15 * time.Minute

// New memcache middleware creates & returns.
func New(memcacheAddr string, opts ...CacheOption) interface {
	datastore.Middleware
	storagecache.Storage
} {
	memcacheClient := memcache.New(memcacheAddr)

	ch := &cacheHandler{
		client:           memcacheClient,
		stOpts:         &storagecache.Options{},
		expireDuration: defaultExpiration,
	}

	for _, opt := range opts {
		opt.Apply(ch)
	}

	s := storagecache.New(ch, ch.stOpts)
	ch.Middleware = s

	if ch.logf == nil {
		ch.logf = func(ctx context.Context, format string, args ...interface{}) {}
	}
	if ch.cacheKey == nil {
		ch.cacheKey = func(key datastore.Key) string {
			return "mercari:memcache:" + key.Encode()
		}
	}

	return ch
}

type cacheHandler struct {
	datastore.Middleware
	stOpts *storagecache.Options

	client         *memcache.Client
	expireDuration time.Duration
	logf           func(ctx context.Context, format string, args ...interface{})
	cacheKey       func(key datastore.Key) string
}

// A CacheOption is an cache option for a memcache middleware.
type CacheOption interface {
	Apply(*cacheHandler)
}

func (ch *cacheHandler) SetMulti(ctx context.Context, cis []*storagecache.CacheItem) error {

	ch.logf(ctx, "dsmiddleware/memcache.SetMulti: incoming len=%d", len(cis))

	for _, ci := range cis {
		if ci.Key.Incomplete() {
			panic("incomplete key incoming")
		}
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		err := enc.Encode(ci.PropertyList)
		if err != nil {
			ch.logf(ctx, "dsmiddleware/memcache.SetMulti: gob.Encode error key=%s err=%s", ci.Key.String(), err.Error())
			continue
		}
		item := &memcache.Item{
			Key:        ch.cacheKey(ci.Key),
			Value:      buf.Bytes(),
			Expiration: int32(ch.expireDuration.Seconds()),
		}
		err = ch.client.Set(item)
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

func (ch *cacheHandler) GetMulti(ctx context.Context, keys []datastore.Key) ([]*storagecache.CacheItem, error) {
	ch.logf(ctx, "dsmiddleware/memcache.GetMulti: incoming len=%d", len(keys))

	resultList := make([]*storagecache.CacheItem, len(keys))

	cacheKeys := make([]string, 0, len(keys))
	for _, key := range keys {
		cacheKeys = append(cacheKeys, ch.cacheKey(key))
	}
	itemMap, err := ch.client.GetMulti(cacheKeys)

	if err != nil {
		ch.logf(ctx, "dsmiddleware/memcache: error on memcache.GetMulti %s", err.Error())
	}

	hit, miss := 0, 0
	for idx, key := range keys {
		item, ok := itemMap[ch.cacheKey(key)]
		if !ok {
			resultList[idx] = nil
			miss++
			continue
		}
		buf := bytes.NewBuffer(item.Value)
		dec := gob.NewDecoder(buf)
		var ps datastore.PropertyList
		err = dec.Decode(&ps)
		if err != nil {
			resultList[idx] = nil
			ch.logf(ctx, "dsmiddleware/memcache.GetMulti: gob.Decode error key=%s err=%s", key.String(), err.Error())
			miss++
			continue
		}

		resultList[idx] = &storagecache.CacheItem{
			Key:          key,
			PropertyList: ps,
		}
		hit++
	}

	ch.logf(ctx, "dsmiddleware/aememcache.GetMulti: hit=%d miss=%d", hit, miss)

	return resultList, nil
}

func (ch *cacheHandler) DeleteMulti(ctx context.Context, keys []datastore.Key) error {
	for _, key := range keys {
		err := ch.client.Delete(ch.cacheKey(key))
		if err != nil {
			ch.logf(ctx, "dsmiddleware/memcache: error on memcache.DeleteMulti %s", err.Error())
		}
	}

	return nil
}
