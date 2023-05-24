import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import dotenv from "dotenv";
dotenv.config();

const GORLI_PRIVATE_KEY = process.env.GORLI_PRIVATE_KEY ? process.env.GORLI_PRIVATE_KEY : "";
const GOERLI_ALCHEMY_URL = process.env.GOERLI_ALCHEMY_URL ? process.env.GOERLI_ALCHEMY_URL : "" ;

const GANACHE_URL = process.env.GANACHE_URL ? process.env.GANACHE_URL : "" ;
const GANACHE_PRIVATE_KEY = process.env.GANACHE_PRIVATE_KEY ? process.env.GANACHE_PRIVATE_KEY : "" ;

const HARDHAT_PRIVATE_KEY = process.env.HARDHAT_PRIVATE_KEY ? process.env.HARDHAT_PRIVATE_KEY : "" ;
const HARDHAT_URL = process.env.HARDHAT_URL ? process.env.HARDHAT_URL : "" ;


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
    ganache:{
      url: GANACHE_URL,
      accounts: [GANACHE_PRIVATE_KEY,],
    }
  },
};

export default config;
