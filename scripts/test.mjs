import { Defender } from "@openzeppelin/defender-sdk";
import hardhat from "hardhat";

const { ethers } = hardhat;

async function getSigner() {
  const client = new Defender({
    relayerApiKey: process.env.RELAYER_API_KEY,
    relayerApiSecret: process.env.RELAYER_SECRET_KEY,
  });

  const provider = client.relaySigner.getProvider();
  const signer = await client.relaySigner.getSigner(provider, {
    speed: "fast",
  });

  return signer;
}

async function main() {
  const signer = await getSigner();

  const Box = await ethers.getContractFactory("Box", { signer });
  const box = Box.attach(process.env.PROXY_ADDRESS);
//   const box = Box.attach(process.env.IMPLEMENTATION_ADDRESSS);
//   console.log(await box.removeObject());
  console.log(await box.boxVersion());
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
