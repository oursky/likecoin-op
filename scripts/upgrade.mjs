import hardhat from "hardhat";

const { ethers, defender } = hardhat;

async function main() {
  const Box = await ethers.getContractFactory("Box");

  const proposal = await defender.proposeUpgradeWithApproval(process.env.PROXY_ADDRESS, Box);

  console.log(`Upgrade proposed with URL: ${proposal.url}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
