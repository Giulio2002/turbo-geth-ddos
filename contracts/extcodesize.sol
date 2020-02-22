pragma solidity ^0.5.0;

contract Extcodesize {
    function start() view public returns(uint) {
        uint g = gasleft();
        assembly {
            pop(extcodesize(address))
        }     
        return g - gasleft(); 
    }
}