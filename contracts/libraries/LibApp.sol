// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {LibDiamond} from "./LibDiamond.sol";

library LibReserve {
    bytes32 constant RESERVE_STORAGE_POSITION = keccak256("diamond.standard.reserve.storage");

    enum Status {None, Reserved, Canceled, Settled}

    struct Reserve {
        address merchant; // 事業者のアドレス
        address subscriber; // 予約者のアドレス
        address token; // トークン
        uint256 depositAmount; // デポジット量
        uint256 additionalAmount; // 追加量
        uint256 cancelableTime; // キャンセル可能なブロック高 0ならいつでもキャンセル可
        uint256 withdrawableTime; // 事業者が予約不履行になった場合に引き出せる時間
        Status status; // reservation status
    }

    struct ReserveStorage {
        address swapSubmitAddress;
        mapping(string => Reserve) reserve; // maps paymentId to reserve struct
    }

    function reserveStorage() internal pure returns (ReserveStorage storage rs) {
        bytes32 position = RESERVE_STORAGE_POSITION;
        assembly {
            rs.slot := position
        }
    }

    function validateReserve(string memory _paymentId, address _merchant, address _token, uint _cancelableTime, uint _withdrawableTime) internal view {
        require(bytes(_paymentId).length != 0, "Payment id is required");
        require(_merchant != address(0), "Merchant address is required");
        require(_token != address(0), "Token address is required");
        require(_cancelableTime < _withdrawableTime, "Incorrect time");
        require(block.timestamp < _withdrawableTime, "Incorrect time");
    }

    function isContractOwner() internal view {
        require(msg.sender == LibDiamond.diamondStorage().contractOwner, "Must be contract owner");
    }

    function isSwapSubmitAddress() internal view {
        require(msg.sender == reserveStorage().swapSubmitAddress, "Must be swap submit address");
    }

    function isMerchant(string memory _paymentId) internal view {
        require(msg.sender == reserveStorage().reserve[_paymentId].merchant, "Must be merchant owner");
    }

    function isSubscriber(string memory _paymentId) internal view {
        require(msg.sender == reserveStorage().reserve[_paymentId].subscriber, "Must be subscriber owner");
    }

    function isStatusReserved(string memory _paymentId) internal view {
        require(reserveStorage().reserve[_paymentId].status == Status.Reserved, "This transaction has been closed");
    }

    function isCancelable(string memory _paymentId) internal view {
        isStatusReserved(_paymentId);
        isSubscriber(_paymentId);
        require(block.timestamp <= reserveStorage().reserve[_paymentId].cancelableTime, "Cancellation period has passed");
    }

    function isWithdrawDeposit(string memory _paymentId) internal view {
        isStatusReserved(_paymentId);
        isMerchant(_paymentId);
        require(block.timestamp > reserveStorage().reserve[_paymentId].withdrawableTime, "Still can't pull it out");
    }

}


