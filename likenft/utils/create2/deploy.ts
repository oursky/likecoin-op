import { BaseContract, ContractFactory } from "ethers";
import { HardhatRuntimeEnvironment } from "hardhat/types";
import { encode } from "./encode";

import { contractName as create2ableContractName } from "../../artifacts/contracts/utils/create2/Create2able.sol/Create2able.json";
import { precomputeAddressWithCreate2ableContract } from "./address";

export async function deployWithCreate2able<
  A extends Array<any> = Array<any>,
  I = BaseContract,
>(
  ethers: HardhatRuntimeEnvironment["ethers"],
  create2ableContract: BaseContract,
  contractFactory: ContractFactory<A, I>,
  deployArgs: A,
  salt: string,
): Promise<ReturnType<ContractFactory<A, I>["deploy"]>> {
  const { assert } = ethers;

  const byteCode = contractFactory.bytecode;
  const deployFrag = contractFactory.interface.deploy;
  const initCode = byteCode + encode(ethers, deployFrag.inputs, deployArgs);

  const saltHex = ethers.id(salt);

  const address = await precomputeAddressWithCreate2ableContract(
    ethers,
    create2ableContract,
    contractFactory,
    deployArgs,
    salt,
  );

  const transaction = await create2ableContract.create2(initCode, saltHex);
  await transaction.wait();

  const deployedContract = contractFactory.attach(address);
  const deployedCode = await deployedContract.getDeployedCode();
  // @ts-ignore-next-line
  assert(
    deployedCode != null,
    `contract code not found in address ${address}`,
    "VALUE_MISMATCH",
  );

  // Mimicing https://github.com/ethers-io/ethers.js/blob/9fd9d41d017a5e3b329aca47c79786e69cd40b99/src.ts/contract/factory.ts#L104
  // return contract after deployment
  return new (<any>BaseContract)(
    address,
    contractFactory.interface,
    create2ableContract.runner,
    transaction,
  );
}

async function makeCreate2ableContractFromAddress(
  ethers: HardhatRuntimeEnvironment["ethers"],
  create2ableAddress: string,
): Promise<BaseContract> {
  const Create2FactoryContractFactory = await ethers.getContractFactory(
    create2ableContractName,
  );

  return Create2FactoryContractFactory.attach(create2ableAddress);
}

export async function deployWithCreate2ableAddress<
  A extends Array<any> = Array<any>,
  I = BaseContract,
>(
  ethers: HardhatRuntimeEnvironment["ethers"],
  create2ableAddress: string,
  contractFactory: ContractFactory<A, I>,
  deployArgs: A,
  salt: string,
): Promise<ReturnType<ContractFactory<A, I>["deploy"]>> {
  const create2ableContract = await makeCreate2ableContractFromAddress(
    ethers,
    create2ableAddress,
  );

  return deployWithCreate2able(
    ethers,
    create2ableContract,
    contractFactory,
    deployArgs,
    salt,
  );
}
