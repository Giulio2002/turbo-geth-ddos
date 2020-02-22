pragma solidity ^0.5.0;

contract ExtcodesizeAttack {
    function loop() public view {
        uint g = gasleft();
        while(g > 30000) {
            assembly {
                pop(extcodesize(address))
            }     
        }
    }
}