## [1.5.2](https://github.com/mercari/datastore/compare/v1.5.1...v1.5.2) (2019-08-15)


### Bug Fixes

* **cloud:** fix encoding issue on embedded entity. thx @kailuo-qb ! ([d94e3f5](https://github.com/mercari/datastore/commit/d94e3f5))



## [1.5.1](https://github.com/mercari/datastore/compare/v1.5.0...v1.5.1) (2019-07-31)


### Bug Fixes

* **dsmiddleware/rpcretry:** do not retry when context is already done ([d193c83](https://github.com/mercari/datastore/commit/d193c83))



# [1.5.0](https://github.com/mercari/datastore/compare/v1.4.2...v1.5.0) (2019-07-16)


### Features

* **cloud:** add WithGRPCDialOption option ([834e0eb](https://github.com/mercari/datastore/commit/834e0eb))



## [1.4.2](https://github.com/mercari/datastore/compare/v1.4.1...v1.4.2) (2019-05-11)


### Bug Fixes

* **clouddatastore:** execute PostCommit hook after RunInTransaction succeed ([7fa9a91](https://github.com/mercari/datastore/commit/7fa9a91))



## [1.4.1](https://github.com/mercari/datastore/compare/v1.4.0...v1.4.1) (2019-05-07)


### Bug Fixes

* **aedatastore:** fix dead lock in multiple tx method call ([740be27](https://github.com/mercari/datastore/commit/740be27))


### Features

* **core:** pull current Cloud Datastore implementations ([2c50ac8](https://github.com/mercari/datastore/commit/2c50ac8))
* **core:** use go.mod instead of dep ([7759ca6](https://github.com/mercari/datastore/commit/7759ca6))



<a name="1.4.0"></a>
# [1.4.0](https://github.com/mercari/datastore/compare/v1.3.0...v1.4.0) (2019-02-19)


### Bug Fixes

* **boom:** Fix to prevent an error when a field implementing PropertyTranslator returns incomplete Key on KeyError ([80193be](https://github.com/mercari/datastore/commit/80193be))


### Features

* **boom:** Make Kind method doesn't depend to KeyError method ([a1db603](https://github.com/mercari/datastore/commit/a1db603))
* **core:** pull current Cloud Datastore implementations ([85b9a4d](https://github.com/mercari/datastore/commit/85b9a4d))
* **core:** update dependencies ([75e7af3](https://github.com/mercari/datastore/commit/75e7af3))



<a name="1.3.0"></a>
# [1.3.0](https://github.com/mercari/datastore/compare/v1.2.0...v1.3.0) (2018-10-23)


### Features

* **cloud:** add FromClient function ([33b8714](https://github.com/mercari/datastore/commit/33b8714))
* **core:** update dependencies ([273cb3c](https://github.com/mercari/datastore/commit/273cb3c))



<a name="1.2.0"></a>
# [1.2.0](https://github.com/mercari/datastore/compare/v1.1.0...v1.2.0) (2018-09-13)


### Features

* **boom:** add DatastoreTransaction function to boom package ([53c84a3](https://github.com/mercari/datastore/commit/53c84a3))



<a name="1.1.0"></a>
# [1.1.0](https://github.com/mercari/datastore/compare/v1.0.1...v1.1.0) (2018-08-12)


### Features

* **core:** add lock to *Batch#Exec ([5c6aa54](https://github.com/mercari/datastore/commit/5c6aa54))



<a name="1.0.1"></a>
## [1.0.1](https://github.com/mercari/datastore/compare/v1.0.0...v1.0.1) (2018-08-12)


### Bug Fixes

* **ci:** fix broken test ([2fa8d46](https://github.com/mercari/datastore/commit/2fa8d46))
* **dsmiddleware/storagecache:** fix panic occurred when DeleteMulti returns error ([f4e5501](https://github.com/mercari/datastore/commit/f4e5501))


### Features

* **core:** update dependencies ([dffe09e](https://github.com/mercari/datastore/commit/dffe09e))
* **dsmiddleware/rediscache:** move to github.com/gomodule/redigo/redis from github.com/garyburd/redigo/redis ([7be2e88](https://github.com/mercari/datastore/commit/7be2e88))



<a name="1.0.0"></a>
# [1.0.0](https://github.com/mercari/datastore/compare/v0.19.0...v1.0.0) (2018-06-06)


### Features

* **dsmiddleware/storagecache:** go to private about Tx* type & struct ([b06c9a7](https://github.com/mercari/datastore/commit/b06c9a7))


### BREAKING CHANGES

Go to private about TxOps, TxOpLog in `go.mercari.io/datastore/dsmiddleware/storagecache`.


<a name="0.19.0"></a>
# [0.19.0](https://github.com/mercari/datastore/compare/v0.18.0...v0.19.0) (2018-05-30)


### Bug Fixes

* **ci:** fix ci failed ([ffeefe6](https://github.com/mercari/datastore/commit/ffeefe6))


### Features

* **core:** deprecate datastore.FromContext and related API ([7e6fd79](https://github.com/mercari/datastore/commit/7e6fd79))
* **dsmiddleware/rpcretry:** rename WithLogf function to WithLogger function ([2453e9b](https://github.com/mercari/datastore/commit/2453e9b))



<a name="0.18.0"></a>
# [0.18.0](https://github.com/mercari/datastore/compare/v0.17.0...v0.18.0) (2018-05-30)


### Bug Fixes

* **clouddatastore:** set MaxAttempts to 1 by default ([8d3adc1](https://github.com/mercari/datastore/commit/8d3adc1))


### Features

* **ci:** change golint repository ([2ec03e4](https://github.com/mercari/datastore/commit/2ec03e4))
* **core:** update dependencies ([23c9dd3](https://github.com/mercari/datastore/commit/23c9dd3))
* **core:** update dependencies ([751c049](https://github.com/mercari/datastore/commit/751c049))



<a name="0.17.0"></a>
# [0.17.0](https://github.com/mercari/datastore/compare/v0.16.0...v0.17.0) (2018-03-27)


### Bug Fixes

* **all:** fix method name AllocatedIDs to AllocateID 🙇 ([68408f8](https://github.com/mercari/datastore/commit/68408f8))


### Features

* **ci:** update CI and local testing environment ([35c8f7a](https://github.com/mercari/datastore/commit/35c8f7a))


### BREAKING CHANGES

Replace AllocatedIDs to AllocateID. align to original libraries 🙇


<a name="0.16.0"></a>
# [0.16.0](https://github.com/mercari/datastore/compare/v0.15.0...v0.16.0) (2018-01-24)


### Features

* **core:** add Key#SetNamespace method ([56f6294](https://github.com/mercari/datastore/commit/56f6294))
* **dsmiddleware/aememcache,dsmiddleware/localcache,dsmiddleware/rediscache,dsmiddleware/storagecache:** add context.Context parameter to Key filter function ([7f8d7f7](https://github.com/mercari/datastore/commit/7f8d7f7))

### BREAKING CHANGES

Change KeyFilter function signature `func(key datastore.Key) bool` to `func(ctx context.Context, key datastore.Key) bool` .



<a name="0.15.0"></a>
# [0.15.0](https://github.com/mercari/datastore/compare/v0.14.0...v0.15.0) (2018-01-09)


### Features

* **dsmiddleware/aememcache,dsmiddleware/localcache,dsmiddleware/rediscache,dsmiddleware/storagecache:** change options format ([5af7561](https://github.com/mercari/datastore/commit/5af7561))

### BREAKING CHANGES

Change cache middleware signatures.

<a name="0.14.0"></a>
# [0.14.0](https://github.com/mercari/datastore/compare/v0.13.0...v0.14.0) (2018-01-09)


### Features

* **boom:** add Boom() and Transaction() method to each boom objects ([3c680d1](https://github.com/mercari/datastore/commit/3c680d1))
* **core:** add AllocateIDs & Count method to Middleware interface ([f548cca](https://github.com/mercari/datastore/commit/f548cca))
* **core:** replace SwapContext to Context & SetContext ([4b9ccaa](https://github.com/mercari/datastore/commit/4b9ccaa))

### BREAKING CHANGES

replace datastore.Client#SwapContext to datastore.Client#Context & datastore.Client#SetContext.


<a name="0.13.0"></a>
# [0.13.0](https://github.com/mercari/datastore/compare/v0.12.0...v0.13.0) (2017-12-19)


### Features

* **ci:** add redis sidecar container ([bc9908a](https://github.com/mercari/datastore/commit/bc9908a))
* **dsmiddleware/aememcache:** change to display both hit and miss to logging ([257064b](https://github.com/mercari/datastore/commit/257064b))
* **dsmiddleware/rediscache:** add dsmiddleware/rediscache package ([04cf0cb](https://github.com/mercari/datastore/commit/04cf0cb))



<a name="0.12.0"></a>
# [0.12.0](https://github.com/mercari/datastore/compare/v0.11.0...v0.12.0) (2017-12-13)


### Features

* **dsmiddleware/chaosrpc:** add dsmiddleware/chaosrpc middleware for testing ([7da792f](https://github.com/mercari/datastore/commit/7da792f))
* **dsmiddleware/noop:** add dsmiddleware/noop middleware ([5c5af95](https://github.com/mercari/datastore/commit/5c5af95))
* **dsmiddleware/rpcretry:** add dsmiddleware/rpcretry middleware ([17c5b17](https://github.com/mercari/datastore/commit/17c5b17))



<a name="0.11.0"></a>
# [0.11.0](https://github.com/mercari/datastore/compare/v0.10.1...v0.11.0) (2017-12-13)


### Features

* **middleware:** rename CacheStrategy to Middleware & move cache dir to dsmiddleware dir ([ae339b9](https://github.com/mercari/datastore/commit/ae339b9))

### BREAKING CHANGES

refactoring cache layer to middleware layer.


<a name="0.10.1"></a>
## [0.10.1](https://github.com/mercari/datastore/compare/v0.10.0...v0.10.1) (2017-12-12)


### Bug Fixes

* **core:** fix deadlock when recursive batch calling ([5162647](https://github.com/mercari/datastore/commit/5162647))



<a name="0.10.0"></a>
# [0.10.0](https://github.com/mercari/datastore/compare/v0.9.0...v0.10.0) (2017-12-07)


### Bug Fixes

* **cache/aememcache:** skip entity when gob encode & decode error occured ([2c3f8da](https://github.com/mercari/datastore/commit/2c3f8da))
* **core:** change order of application about CacheStrategy to first in - first apply ([231f40b](https://github.com/mercari/datastore/commit/231f40b))

### BREAKING CHANGES

Change the order of application of CacheStrategy first in - last apply to first in - first apply.


<a name="0.9.0"></a>
# [0.9.0](https://github.com/mercari/datastore/compare/v0.8.2...v0.9.0) (2017-12-06)

### Features

* **core,boom:** change batch operation signatures ([51da3ba](https://github.com/mercari/datastore/commit/51da3ba))

### BREAKING CHANGES

For batch processing, we stopped asynchronous processing using chan and switched to synchronous processing using callback function.


<a name="0.8.2"></a>
## [0.8.2](https://github.com/mercari/datastore/compare/v0.8.1...v0.8.2) (2017-12-06)


### Bug Fixes

* **boom:** fix PendingKey handling fixes [#30](https://github.com/mercari/datastore/issues/30) thanks [@sinmetal](https://github.com/sinmetal) ([eaa5729](https://github.com/mercari/datastore/commit/eaa5729))
* **cache/storagecache:** fix MultiError handling that ErrNoSuchEntity contaminated ([d42850b](https://github.com/mercari/datastore/commit/d42850b))



<a name="0.8.1"></a>
## [0.8.1](https://github.com/mercari/datastore/compare/v0.8.0...v0.8.1) (2017-12-05)


### Bug Fixes

* **core:** fix time.Time's default location. fit to Cloud Datastore behaviour ([4226d8f](https://github.com/mercari/datastore/commit/4226d8f))



<a name="0.8.0"></a>
# [0.8.0](https://github.com/mercari/datastore/compare/v0.7.0...v0.8.0) (2017-12-04)


### Features

* **cache/storagecache:** implement WithIncludeKinds, WithExcludeKinds, WithKeyFilter options ([a8b5857](https://github.com/mercari/datastore/commit/a8b5857))



<a name="0.7.0"></a>
# [0.7.0](https://github.com/mercari/datastore/compare/v0.6.0...v0.7.0) (2017-12-04)


### Features

* **cache** implement cache layer & cache strategies ([203ab21](https://github.com/mercari/datastore/commit/203ab21))
* **core,ae,cloud:** add datastore#Client.DecodeKey method ([42fa040](https://github.com/mercari/datastore/commit/42fa040))



<a name="0.6.0"></a>
# [0.6.0](https://github.com/mercari/datastore/compare/v0.5.3...v0.6.0) (2017-11-24)


### Features

* **boom:** add NewQuery method ([a31adb0](https://github.com/mercari/datastore/commit/a31adb0)) thanks @timakin !



<a name="0.5.3"></a>
## [0.5.3](https://github.com/mercari/datastore/compare/v0.5.2...v0.5.3) (2017-11-24)


### Bug Fixes

* **ae,cloud:** fix datastore.PropertyList handling when Put & Get ([0355f35](https://github.com/mercari/datastore/commit/0355f35))
* **ae,cloud:** fix struct (without pointer) handling when Put & Get ([de3eb4c](https://github.com/mercari/datastore/commit/de3eb4c))
* **boom:** fix nil parent key handling ([7dc317b](https://github.com/mercari/datastore/commit/7dc317b))



<a name="0.5.2"></a>
## [0.5.2](https://github.com/mercari/datastore/compare/v0.5.1...v0.5.2) (2017-11-22)


### Bug Fixes

* **core:** fix datastore.Key or []datastore.Key Save & Load handling ([29f465d](https://github.com/mercari/datastore/commit/29f465d))



<a name="0.5.1"></a>
## [0.5.1](https://github.com/mercari/datastore/compare/v0.5.0...v0.5.1) (2017-11-21)


### Bug Fixes

* **boom:** fix *boom.Boom#GetAll with KeysOnly query ([420bb37](https://github.com/mercari/datastore/commit/420bb37))
* **boom:** fix *boom.Iterator#Next with KeysOnly query ([e8bbeed](https://github.com/mercari/datastore/commit/e8bbeed))



<a name="0.5.0"></a>
# [0.5.0](https://github.com/mercari/datastore/compare/v0.4.0...v0.5.0) (2017-11-21)


### Features

* **boom:** add boom.ToAECompatibleTransaction and *boom.AECompatibleTransaction ([dedb72a](https://github.com/mercari/datastore/commit/dedb72a))
* **boom:** add Kind, Key, KeyError method to *boom.Transaction ([5d5da7d](https://github.com/mercari/datastore/commit/5d5da7d))
* **core:** add Equal and Incomplete methods to datastore.Key ([5668f1b](https://github.com/mercari/datastore/commit/5668f1b))



<a name="0.4.0"></a>
# [0.4.0](https://github.com/mercari/datastore/compare/v0.3.0...v0.4.0) (2017-11-20)


### Features

* **boom:** implements AllocateID and AllocateIDs ([014e321](https://github.com/mercari/datastore/commit/014e321))
* **core:** add datastore.Client#SwapContext ([eb26e60](https://github.com/mercari/datastore/commit/eb26e60))



<a name="0.3.0"></a>
# [0.3.0](https://github.com/mercari/datastore/compare/v0.2.0...v0.3.0) (2017-11-14)


### Bug Fixes

* **boom:** improve goon compatibility ([03beb64](https://github.com/mercari/datastore/commit/03beb64))



<a name="0.2.0"></a>
# [0.2.0](https://github.com/mercari/datastore/compare/v0.1.0...v0.2.0) (2017-11-14)


### Features

* **aedatastore:** add custom import path checking ([801299f](https://github.com/mercari/datastore/commit/801299f))
* **ci:** add .circleci/config.yml ([cfc3877](https://github.com/mercari/datastore/commit/cfc3877))
* **clouddatastore:** add custom import path checking ([5585c22](https://github.com/mercari/datastore/commit/5585c22))
* **core:** align TransactionBatch api to Batch api ([3f49066](https://github.com/mercari/datastore/commit/3f49066))
* **boom:** implement boom package ([8c2ed5e](https://github.com/mercari/datastore/commit/8c2ed5e))
