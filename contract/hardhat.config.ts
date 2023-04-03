import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import dotenv from 'dotenv';

dotenv.config();
const PRIVATE_KEY : string  = process.env.PRIVATE_KEY ? process.env.PRIVATE_KEY : "";
const GOERLI_URL : string = process.env.GOERLI_URL ? process.env.GOERLI_URL : "" ;

console.log(`PRIVATE_KEY:${PRIVATE_KEY}`);
const config: HardhatUserConfig = {
  solidity: {
    version: "0.8.17",
    settings: {
      optimizer: {
        enabled: true,
        runs: 1000,
      },
    },
  },
  networks: {
    hardhat:{},
    goerli:{
      url: GOERLI_URL,
      accounts: [PRIVATE_KEY,],
    },
    dev:{
      url: "HTTP://127.0.0.1:7546",
      accounts: [
        "0x4e29b01fa5f045175987ed4a3b681d05cab00c7f68c0f25063d0138f04410422",
      ],
    }
  }
};

export default config;
