// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";


//
// using UniversalERC20 for IERC20;
//
library UniversalERC20 {

    using SafeMath for uint256;
    using SafeERC20 for IERC20;

    function universalTransfer(IERC20 token, address to, uint256 amount) internal {
        if (token == IERC20(address(0))) {
            payable(to).transfer(amount);
        } else {
            token.safeTransfer(to, amount);
        }
    }

    function universalApprove(IERC20 token, address to, uint256 amount) internal {
        if (token != IERC20(address(0))) {
            token.safeApprove(to, amount);
        }
    }

    function universalTransferFrom(IERC20 token, address from, address to, uint256 amount) internal {
        if (token == IERC20(address(0))) {
            require(from == msg.sender && msg.value >= amount, "msg.value is zero");
            if (to != address(this)) {
                payable(to).transfer(amount);
            }
            if (msg.value > amount) {
                payable(msg.sender).transfer(msg.value.sub(amount));
            }
        } else {
            token.safeTransferFrom(from, to, amount);
        }
    }

    function universalBalanceOf(IERC20 token, address who) internal view returns (uint256) {
        if (token == IERC20(address(0))) {
            return who.balance;
        } else {
            return token.balanceOf(who);
        }
    }
}