// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {LibDiamond} from "./LibDiamond.sol";

library LibMerchant {
    bytes32 constant MERCHANT_STORAGE_POSITION = keccak256("diamond.standard.merchant.storage");
//
    struct Merchant {
        bool isApprove; // 承認済みの事業者か(Earn対象の事業者か？)

//        uint256 requireDepositAmount; // デポジットに必要な金額
//        uint256 timeLock; // 予約履行時間 - timelock = キャンセル可能な閾値
    }
//
    struct MerchantStorage {
        mapping(address => Merchant) merchant;
    }
//
    function merchantStorage() internal pure returns (MerchantStorage storage rs) {
        bytes32 position = MERCHANT_STORAGE_POSITION;
        assembly {
            rs.slot := position
        }
    }
//
//    function create(uint256 _requireDepositAmount, uint256 _timeLock) public {
//        Merchant storage merchant = merchantStorage().merchant[msg.sender];
//        require(merchant.requireDepositAmount == 0, "Already Exist");
//
//        merchant.requireDepositAmount = _requireDepositAmount;
//        merchant.timeLock = _timeLock;
//    }
//
//    function changeRequireDepositAmount(uint256 _requireDepositAmount) public {
//        Merchant storage merchant = merchantStorage().merchant[msg.sender];
//        require(merchant.requireDepositAmount != 0, "Not Exist");
//
//        merchant.requireDepositAmount = _requireDepositAmount;
//    }
//
//    function changeTimeLock(uint256 _timeLock) public {
//        Merchant storage merchant = merchantStorage().merchant[msg.sender];
//        require(merchant.timeLock != 0, "Not Exist");
//
//        merchant.requireDepositAmount = _requireDepositAmount;
//    }


}