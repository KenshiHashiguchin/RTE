// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./interfaces/ISlashCustomPlugin.sol";
import "../libraries/UniversalERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import {LibReserve} from "../libraries/LibApp.sol";


interface ReserveContract {
    function getReservation(string memory _paymentId) external view returns (LibReserve.Reserve memory);
}

contract ReceivePayment is ISlashCustomPlugin, Ownable {
    using UniversalERC20 for IERC20;

    ReserveContract private reserveContract;

    constructor(address _address) {
        _reservationAddress(_address);
    }

    function updateReservationContractAddress(address _newAddress) external onlyOwner {
        _reservationAddress(_newAddress);
    }

    function _reservationAddress(address _newAddress) internal {
        reserveContract = ReserveContract(_newAddress);
    }

    function receivePayment(
        address receiveToken,
        uint256 amount,
        string memory paymentId,
        string memory optional
    ) external payable override {
        require(amount >= 0, "invalid amount");
        require(receiveToken != address(0), "invalid token");

        LibReserve.Reserve memory reserve = reserveContract.getReservation(paymentId);

        IERC20(receiveToken).transfer(
            reserve.merchant,
            amount
        );
        // 受け取りトークンを保存する(paymentIdをキーに)
    }
}