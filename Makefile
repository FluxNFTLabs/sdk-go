all:

chain-types:
	rm -rf chain && mkdir chain chain/types chain/app chain/indexer chain/modules \
	chain/modules/fnft chain/modules/fnft/types \
	chain/modules/bazaar chain/modules/bazaar/types \
	chain/modules/evm chain/modules/evm/types chain/modules/evm/evmone \
	chain/modules/svm chain/modules/svm/types chain/modules/svm/golana \
	chain/modules/astromesh chain/modules/astromesh/types \
	chain/stream chain/stream/types

	cp -r ../fluxd/chain/types/ chain/types
	cp -r ../fluxd/chain/app/ante/ chain/app/ante
	cp -r ../fluxd/chain/stream/types/ chain/stream/types
	cp -r ../fluxd/chain/indexer/ chain/indexer

	cp -r ../fluxd/chain/modules/fnft/types/ chain/modules/fnft/types

	cp -r ../fluxd/chain/modules/bazaar/types/ chain/modules/bazaar/types

	cp -r ../fluxd/chain/modules/evm/types/ chain/modules/evm/types
	cp ../fluxd/chain/modules/evm/evmone/types.go chain/modules/evm/evmone/types.go

	cp -r ../fluxd/chain/modules/svm/types/ chain/modules/svm/types
	cp -r ../fluxd/chain/modules/svm/golana/ chain/modules/svm/golana

	cp -r ../fluxd/chain/modules/astromesh/types/ chain/modules/astromesh/types

	./scripts/replace_path.sh

	cp -r ../fluxd/chain/modules/evm/lib/ chain/modules/evm/lib
	cp -r ../fluxd/chain/modules/svm/lib/ chain/modules/svm/lib
