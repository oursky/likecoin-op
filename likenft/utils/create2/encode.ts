import { ParamType } from "ethers";
import { HardhatRuntimeEnvironment } from "hardhat/types";

// refs: https://www.alchemy.com/docs/create2-an-alternative-to-deriving-contract-addresses
// - utils.js:encoder
export function encode(
  ethers: HardhatRuntimeEnvironment["ethers"],
  types: ReadonlyArray<string | ParamType>,
  values: ReadonlyArray<any>,
) {
  const { AbiCoder } = ethers;
  const abiCoder = AbiCoder.defaultAbiCoder();
  const encodedParams = abiCoder.encode(types, values);
  return encodedParams.slice(2);
}
