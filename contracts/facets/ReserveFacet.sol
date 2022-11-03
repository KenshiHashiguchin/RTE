// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import {LibReserve} from "../libraries/LibApp.sol";
import {LibMerchant} from "../libraries/LibMerchant.sol";
import "../libraries/UniversalERC20.sol";
import "../extension/interfaces/ISlashSubmitTransaction.sol";

interface ISwap {
    function swapTokensForExactTokens(
        uint256 amountOut,
        uint256 amountInMax,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external returns (uint256[] memory amounts);
}

contract ReserveFacet {
    using UniversalERC20 for IERC20;
    using SafeMath for uint256;

    event SettleReservation(address to, address token, uint amount);

    function getReservation(string memory _paymentId) external view returns (LibReserve.Reserve memory) {
        LibReserve.Reserve storage reserve = LibReserve.reserveStorage().reserve[_paymentId];
        require(reserve.status != LibReserve.Status.None, "Don't exist paymentId");
        return reserve;
    }

    function reserve(string memory _paymentId, address _merchant, address _token, uint _depositAmount, uint _cancelableTime, uint _withdrawableTime) external {
        LibReserve.validateReserve(_paymentId, _merchant, _token, _cancelableTime, _withdrawableTime);
        LibReserve.Reserve storage reserve = LibReserve.reserveStorage().reserve[_paymentId];
        require(reserve.status == LibReserve.Status.None, "Already exist paymentId");

        reserve.token = _token;
        reserve.merchant = _merchant;
        reserve.subscriber = msg.sender;
        reserve.withdrawableTime = _withdrawableTime;
        reserve.cancelableTime = _cancelableTime;
        reserve.depositAmount = _depositAmount;
        reserve.status = LibReserve.Status.Reserved;

        if (_depositAmount > 0) {
            IERC20(_token).universalTransferFrom(msg.sender, address(this), _depositAmount);
        }
    }

    function settleReservation(
        uint256 _amountIn,
        uint256 _requiredAmountOut,
        uint256 _deadline,
        address[] calldata _path,
        string memory _paymentId
    ) external {
        LibReserve.isSubscriber(_paymentId);
        LibReserve.isStatusReserved(_paymentId);

        LibReserve.ReserveStorage storage reserveStorage = LibReserve.reserveStorage();
        LibReserve.Reserve storage reserve = reserveStorage.reserve[_paymentId];
        require(_amountIn >= reserve.depositAmount, "Excess payments");

        uint256 amount = _amountIn.sub(reserve.depositAmount);
        IERC20(reserve.token).universalTransferFrom(msg.sender, address(this), amount);

        IERC20(reserve.token).approve(reserveStorage.swapSubmitAddress, _amountIn);
        ISwap(reserveStorage.swapSubmitAddress).swapTokensForExactTokens(
            _requiredAmountOut,
            _amountIn,
            _path,
            reserve.merchant,
            _deadline
        );

        reserve.status = LibReserve.Status.Settled;
        reserve.additionalAmount = amount;

        // TODO SBT
    }

    function cancel(string memory _paymentId) external {
        LibReserve.isSubscriber(_paymentId);
        LibReserve.isCancelable(_paymentId);

        LibReserve.Reserve storage reserve = LibReserve.reserveStorage().reserve[_paymentId];
        reserve.status = LibReserve.Status.Canceled;

        IERC20(reserve.token).universalTransfer(
            reserve.subscriber,
            reserve.depositAmount
        );
    }

    function withdrawDeposit(string memory _paymentId) external {
        LibReserve.isWithdrawDeposit(_paymentId);

        LibReserve.Reserve storage reserve = LibReserve.reserveStorage().reserve[_paymentId];
        reserve.status = LibReserve.Status.Canceled;

        IERC20(reserve.token).approve(reserveStorage.swapSubmitAddress, reserve.depositAmount);
        ISwap(reserveStorage.swapSubmitAddress).swapTokensForExactTokens(
            _requiredAmountOut,
            _amountIn,
            _path,
            reserve.merchant,
            _deadline
        );
    }

    function transferSwapSubmitAddress(address _newAddress) external {
        LibReserve.isContractOwner();
        LibReserve.ReserveStorage storage reserveStorage = LibReserve.reserveStorage();
        reserveStorage.swapSubmitAddress = _newAddress;
    }

    function swapSubmitAddress() external view returns (address) {
        return LibReserve.reserveStorage().swapSubmitAddress;
    }

}