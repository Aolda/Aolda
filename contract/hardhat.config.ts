import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import dotenv from 'dotenv';

dotenv.config();
const PRIVATE_KEY : string  = process.env.PRIVATE_KEY ? process.env.PRIVATE_KEY : "";
const GOERLI_URL : string = process.env.GOERLI_URL ? process.env.GOERLI_URL : "" ;

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
    }
  }
};

export default config;
