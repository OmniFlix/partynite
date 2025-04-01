set -ex
MONIKER=test
DENOM=uflix
CHAINID=partynite
ACCKEY=partynite12ckr8algymfe30rpthmpp6x4xs4nhklctthea5
partynited version --long
# giggle ivory help urge toilet rhythm toss turtle install town height crater disease melody ivory fork language idle awkward glove bird vault liar sense
# Setup PartyNite
partynited init $MONIKER --chain-id $CHAINID
sed -i '' 's#tcp://127.0.0.1:26657#tcp://0.0.0.0:26657#g' ~/.partynite/config/config.toml
sed -i '' "s/\"stake\"/\"$DENOM\"/g" ~/.partynite/config/genesis.json
sed -i '' 's/enable = false/enable = true/g' ~/.partynite/config/app.toml
partynited keys add validator --keyring-backend test

partynited add-genesis-account $(partynited keys --keyring-backend test show validator -a) 100000000000$DENOM
partynited add-genesis-account $ACCKEY 100000000000$DENOM
partynited gentx validator 900000000$DENOM --keyring-backend test --chain-id $CHAINID
partynited collect-gentxs

partynited start