pragma solidity ^0.5.0;

contract Extcodesize {
    function start() public returns(uint) {
        uint g = gasleft();
        assembly {
            pop(extcodesize(address))
        }     
        return g - gasleft(); 
    }
}

contract ExtcodesizeAttack {
    function loop() public returns(uint) {
        uint g = gasleft();
        while(g > 50000) {
            assembly {
                pop(extcodesize(address))
            }
            g = gasleft();
        }
        return gasleft();
    }
}