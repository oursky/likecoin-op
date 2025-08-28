import { HardhatRuntimeEnvironment } from "hardhat/types";
import { BaseContract, ContractFactory } from "ethers";
import { encode } from "./encode";

export async function precomputeAddressWithCreate2ableContract<
  A extends Array<any> = Array<any>,
  I = BaseContract,
>(
  ethers: HardhatRuntimeEnvironment["ethers"],
  create2ableContract: BaseContract,
  contractFactory: ContractFactory<A, I>,
  deployArgs: A,
  salt: string,
) {
  const deploy2ableContractAddress = await create2ableContract.getAddress();

  return precomputeAddressWithCreate2ableAddress(
    ethers,
    deploy2ableContractAddress,
    contractFactory,
    deployArgs,
    salt,
  );
}

export function precomputeAddressWithCreate2ableAddress<
  A extends Array<any> = Array<any>,
  I = BaseContract,
>(
  ethers: HardhatRuntimeEnvironment["ethers"],
  create2ableAddress: string,
  contractFactory: ContractFactory<A, I>,
  deployArgs: A,
  salt: string,
) {
  const { getCreate2Address, keccak256 } = ethers;

  // refs: https://www.alchemy.com/docs/create2-an-alternative-to-deriving-contract-addresses
  // - vaultDeploy.js:main
  const byteCode = contractFactory.bytecode;
  const deployFrag = contractFactory.interface.deploy;
  const initCode = byteCode + encode(ethers, deployFrag.inputs, deployArgs);

  const saltHex = ethers.id(salt);

  // refs: https://www.alchemy.com/docs/create2-an-alternative-to-deriving-contract-addresses
  // - utils.js:create2Address
  return getCreate2Address(create2ableAddress, saltHex, keccak256(initCode));
}
