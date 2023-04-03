// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.7;

contract Mempool {
    uint constant MAX_TRANSACTIONS = 5;
    uint public rewardPerRecipient = 0;
    mapping(uint => uint) public counts; // state 변수로 선언
    Transaction[MAX_TRANSACTIONS] public mempoolTransactions;
    uint public numTransactions;
    address payable[] public rewardRecipients;
    uint public totalReward;

    // 이벤트 정의
    event TransactionAdded(address from, uint amount);
    event TransactionsSent(uint totalAmount, uint numRecipients);

    struct Transaction {
        address from;
        uint amount;
    }
    
    // 트랜잭션 추가 함수
    function addTransaction(address from, uint amount) external {
        require(numTransactions < MAX_TRANSACTIONS, "Mempool is full");
        mempoolTransactions[numTransactions] = Transaction(from, amount);
        numTransactions++;
        emit TransactionAdded(from, amount);
        
        // 트랜잭션이 MAX_TRANSACTIONS 개수에 도달하면 보상 수령자 저장
        if (numTransactions == MAX_TRANSACTIONS) {
            saveRewardRecipients();
        }
    }
    
    // 보상 수령자 저장 함수
    function saveRewardRecipients() private {
        totalReward = 0;
        uint maxCount = 0;
        uint numMaxAmounts = 0;
        for(uint i = 0; i < MAX_TRANSACTIONS; i++){
            counts[i] = 0; // 상태 변수에서 초기화
        }
     
        for(uint i = 0; i < MAX_TRANSACTIONS; i++) {
            uint amount = mempoolTransactions[i].amount;
            counts[amount]++;
            if(counts[amount] > maxCount){
                maxCount = counts[amount];
                numMaxAmounts = amount;
            } 
        }

        for(uint i = 0;i < MAX_TRANSACTIONS;i++){
            uint amount = mempoolTransactions[i].amount;
            if(amount == numMaxAmounts){
                rewardRecipients.push(payable(mempoolTransactions[i].from));
                totalReward += amount;
            }
        }

        rewardPerRecipient = totalReward / numMaxAmounts;
 
        // mempool 초기화
        numTransactions = 0;
        delete mempoolTransactions;
       
        emit TransactionsSent(totalReward, rewardRecipients.length);
    }

    function RewardRecipientReset() public{
         for (uint i = 0; i < rewardRecipients.length; i++) {
            delete rewardRecipients[i];
        }
        delete rewardRecipients;
    }
}
