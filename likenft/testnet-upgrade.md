## Optimism sepolia

This quick calling reference on upgrading in testnet. For full text explination, please refer to [safe-wallet.md](./safe-wallet.md)


```
ERC721_PROXY_ADDRESS=0x67BCd74981c33E95E5e306085754DD0A721183F1 npm run script:optimism-sepolia scripts/queryBookNFTImpl.ts
```

```
npm run script:optimism-sepolia scripts/deployBookNFT.ts
```

It should out put example command for verifyBookNFT
```
BOOKNFT_ADDRESS=0x0068850030F619DA024d2B9B6a2e1211Aa14739D \
    npm run script:optimism-sepolia scripts/verifyBookNFT.ts
```

If we only upgrade the BookNFT, not the LikeProtocol

```
BOOKNFT_ADDRESS=0x0068850030F619DA024d2B9B6a2e1211Aa14739D \
    npm run script:optimism-sepolia scripts/upgradeBookNFT.ts
```

---

In case we upgrade LikeProtocol and BookNFT
```
BOOKNFT_ADDRESS=0x0068850030F619DA024d2B9B6a2e1211Aa14739D \
    npm run script:optimism-sepolia scripts/prepareUpgradeToAndCall.ts
```

Example output, copy down the call data for upgradToAndCall
```
Target new BookNFT implementation address: 0x0068850030F619DA024d2B9B6a2e1211Aa14739D
Upgrade to and call data: 0x3659cfe60000000000000000000000000068850030f619da024d2b9b6a2e1211aa14739d
```

```
npm run script:optimism-sepolia scripts/deployLikeProtocol.ts
CALL_DATA=0x3659cfe60000000000000000000000000068850030f619da024d2b9b6a2e1211aa14739d \
NEW_LIKEPROTOCOL=0x67BCd74981c33E95E5e306085754DD0A721183F1 \
    npm run script:optimism-sepolia scripts/upgradeToAndCall.ts
```

`git push oursky main:optimism-sepolia` for quick reference.