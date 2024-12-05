# Sample Hardhat Project

## Prerequisite

1. Node `lts/iron`

## Dev

1. Start local dev node

```bash
$ npm run dev
```

2. Lint

```bash
$ npm run format
$ npm run lint [-- --fix]
```

2. Test locally

```bash
$ npm run build
$ npm run deploy:local
```

Retrieve the smart contract address from deploy output for further use

```
$ hardhat run --network localhost scripts/deploy.js
Deploying Box...
Box deployed to: 0x5FbDB2315678afecb367f032d93F642f64180aa3
```
