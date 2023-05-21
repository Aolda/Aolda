import { ethers, network } from "hardhat";
const { setConfig } = require("./config.js");

async function main() {
  const networkName = network.name;
  const AoldaClient = await ethers.getContractFactory("AoldaClient");
  const aoldaClient = await AoldaClient.deploy();
  await aoldaClient.deployed();

  setConfig("deployed." + networkName + ".Deposit", aoldaClient.address);
  console.log("aoldaClient:", aoldaClient.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
