## Extcodesize Attack
```
Gas 471 per extcodesize
go run one.go // for checking
```

## Extcodesize Attack
```
Gas = Cost of smart contract creation per suicide (13K gas in this case)
go run one.go // for checking
```
## Exhaustion Attack vs Extcodesize Attack
```
Effectivenes: 2:1 for exhaustion vs extcodesize
time go run two.go // for checking for exhaustion => 1 minute
time go run three.go // for checking for extcodesize => 30 seconds
```
