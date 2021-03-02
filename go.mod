module github.com/sourcenetwork/defradb

go 1.12

require (
	github.com/SierraSoftworks/connor v1.0.2
	github.com/davecgh/go-spew v1.1.1
	github.com/fxamacker/cbor/v2 v2.2.0
	github.com/go-chi/chi v1.5.2
	github.com/gogo/protobuf v1.2.1
	github.com/graphql-go/graphql v0.7.9
	github.com/ipfs/go-block-format v0.0.2
	github.com/ipfs/go-blockservice v0.1.0
	github.com/ipfs/go-cid v0.0.3
	github.com/ipfs/go-datastore v0.4.4
	github.com/ipfs/go-ds-badger v0.2.1
	github.com/ipfs/go-ipfs-blockstore v0.0.1
	github.com/ipfs/go-ipfs-ds-help v0.0.1
	github.com/ipfs/go-ipfs-exchange-offline v0.0.1
	github.com/ipfs/go-ipld-format v0.0.2
	github.com/ipfs/go-log v0.0.1
	github.com/ipfs/go-log/v2 v2.1.1
	github.com/ipfs/go-merkledag v0.2.3
	github.com/jbenet/goprocess v0.1.4
	github.com/kr/pretty v0.2.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/multiformats/go-multibase v0.0.1
	github.com/multiformats/go-multihash v0.0.6
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pkg/errors v0.8.1
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.4.0
	github.com/ugorji/go/codec v1.1.7
	github.com/whyrusleeping/go-logging v0.0.0-20170515211332-0457bb6b88fc
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 // indirect
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80 // indirect
	golang.org/x/sys v0.0.0-20190726091711-fc99dfbffb4e // indirect
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/SierraSoftworks/connor => ../connor // local replace

	// temp bug fixing
	github.com/graphql-go/graphql => ../../graphql-go/graphql
// github.com/graphql-go/graphql => github.com/sourcenetwork/graphql v0.7.10-0.20210211004004-07fce0d1409f
)
