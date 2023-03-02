import { ethers } from "hardhat";
import { AoldaClient } from '../contracts/AoldaClient';

async function main() {
  const AoldaClient = await ethers.getContractFactory("AoldaClient");
  const aoldaClient = await AoldaClient.deploy();
  await aoldaClient.deployed();

  console.log("aoldaClient : ", aoldaClient.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
