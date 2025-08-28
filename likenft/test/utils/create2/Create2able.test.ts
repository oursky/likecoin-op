import { SignerWithAddress } from "@nomicfoundation/hardhat-ethers/signers";
import { expect } from "chai";
import { ContractFactory, TransactionResponse } from "ethers";
import hardhat, { ethers } from "hardhat";
import {
  deployWithCreate2able,
  precomputeAddressWithCreate2ableContract,
} from "../../../utils/create2";

describe("Create2able", () => {
  let Create2ableContractFactory: ContractFactory;
  let BookNFTContractFactory: ContractFactory;
  let signer: SignerWithAddress;

  before(async function () {
    const { ethers } = hardhat;
    Create2ableContractFactory =
      await ethers.getContractFactory("Create2ableMock");
    BookNFTContractFactory = await ethers.getContractFactory("BookNFT");

    [signer] = await ethers.getSigners();
  });

  it("should deploy a BookNFT with a deterministic address", async function () {
    const create2ableContract = await Create2ableContractFactory.deploy();

    const salt = "1234";

    const addressPrecomputed = await precomputeAddressWithCreate2ableContract(
      ethers,
      create2ableContract,
      BookNFTContractFactory,
      [],
      salt,
    );

    const bookNFT = BookNFTContractFactory.attach(addressPrecomputed);

    expect(await bookNFT.getDeployedCode()).to.equal(null);

    const deployedContract = await deployWithCreate2able(
      ethers,
      create2ableContract,
      BookNFTContractFactory,
      [],
      salt,
    );
    await deployedContract.waitForDeployment();

    expect(await bookNFT.getDeployedCode()).to.not.equal(null);

    expect(
      deployWithCreate2able(
        ethers,
        create2ableContract,
        BookNFTContractFactory,
        [],
        salt,
      ),
      "The same address cannot be deployed again",
    ).to.be.rejectedWith("Error: Transaction reverted without a reason string");

    // The deployed address should contain the contract
    const bookNFTInitializeTx: TransactionResponse = await bookNFT.initialize({
      creator: signer.address,
      updaters: [],
      minters: [],
      config: {
        name: "BookNFT Implementation",
        symbol: "BOOKNFTV0",
        metadata: "{}",
        max_supply: 1n,
      },
    });

    const txReceipt = await bookNFTInitializeTx.wait();

    expect(txReceipt!.status).to.equal(1);
  });
});
