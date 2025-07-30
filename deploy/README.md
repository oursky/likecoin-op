# Deployment Notes

We rely on operation to maintains the initial contract address of the evm smart contract.

Testnet operator keys are stored in the `env.operator` file. [Blackbox](https://github.com/StackExchange/blackbox) encrypts the keys.

To decrypt the keys, run `blackbox_decrypt_all_files`, then `make operator-key-link` at the root of the repo.

## Expected procedure

Current encrypted operator wallet: `0xC71fe89e4C0e5458a793fc6548EF6B392417A7Fb` (For production, we use ledger `0xB0318A8f049b625dA5DdD184FfFF668Aa6E96261`, see `likenft/.env.optimism` for now)

Production
LikeProtocol proxy: `0x526237a676444A67bc79E9009756df128Ca9a619`
LikeCoin proxy: `pending`

Optimism-sepolia
LikeProtocol proxy: `0x67BCd74981c33E95E5e306085754DD0A721183F1`
EkilCoin proxy: `pending`

Once the initial deployment is done, we should updated the expected proxy address in the `env.operator` file for reference. On new L2 chain launch, we should operate the following steps to obtain same proxy address.

0. Checkout a specific commit
1. Someone sends ETH to the operator's wallet (!!failed operation becasue of lack of fund will make the address different from the expected one)
2. Use operator's wallet to run the deployment on `likenft` and `likecoin` (!! Order matters)
   - 2.1. `likenft` come second, itself as the owner.
   - 2.2. (Pending on production)`likenft` transfer owner come second, simulating transfer to community wallet.
   - 2.3. (Pending on production)`likecoin` should run first, with itself as the owner and minter
3. Upgrade the implementation to latest implementation
4. Update the owner and minter accordingly

Formula refs:
- newAddressCREATE = `keccak256(deployingAddress ++ nonce)[12:]`
- newAddressCREATE2 = `keccak256(0xff ++ deployingAddress ++ salt ++ initCodeHash)[12:]`


## Migration signer hotwallet

Because of the natural of the project, some credential are manage independently for segregation.

For the signer private key, please consult @rickmak/@williamchong  For the htpassword, we generate separately
Please refer to helm/{env}/Makefile, it provide helper and default for illustration.

For deployment to likeco

```
APP_VERSION=8c9106e4 ENV=likeco K8S_NAMESPACE=optimism make deploy
```
