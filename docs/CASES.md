# Cases
Here we have some use cases to run the application with the expected input data and output data.
The examples are in the folder: [resources](../resources).

## Use case #1
| Operation | Unit-cost | Quantity | Tax paid | Explanation                    |
|:----------|:----------|:---------|:---------|:-------------------------------|
| buy       | 10.00     | 100      | 0        | Buying shares does not pay tax |
| sell      | 15.00     | 50       | 0        | Total value less than R$20,000 |
| sell      | 15.00     | 50       | 0        | Total value less than R$20,000 |

Input:
```json
[{"operation":"buy", "unit-cost":10.00, "quantity": 100}, {"operation":"sell", "unit-cost":15.00, "quantity": 50}, {"operation":"sell", "unit-cost":15.00, "quantity": 50}]
```
Output:
```json
[{"tax": 0},{"tax": 0},{"tax": 0}]
```

## Use case #2
| Operation | Unit-cost | Quantity | Tax paid | Explanation                                                                            |
|:----------|:----------|:---------|:---------|:---------------------------------------------------------------------------------------|
| buy       | 10.00     | 10000    | 0        | Buying shares does not pay tax                                                         |
| sell      | 20.00     | 5000     | 10000    | Profit of R$50,000: 20% of the profit corresponds to R$10,000 and has no previous loss |
| sell      | 5.00      | 5000     | 0        | Loss of R$25,000: no tax paid                                                          |

Input:
```json
[{"operation":"buy", "unit-cost":10.00, "quantity": 10000}, {"operation":"sell", "unit-cost":20.00, "quantity": 5000}, {"operation":"sell", "unit-cost":5.00, "quantity": 5000}]
```
Output:
```json
[{"tax": 0.00},{"tax": 10000.00},{"tax": 0.00}]
```

## Use case #1+#2
In this use case we are running use cases 1 and 2 in the same file. Resulting in a return of two lists.

Input:
```json
[{"operation":"buy", "unit-cost":10.00, "quantity": 100}, {"operation":"sell", "unit-cost":15.00, "quantity": 50},{"operation":"sell", "unit-cost":15.00, "quantity": 50}]
[{"operation":"buy", "unit-cost":10.00, "quantity": 10000}, {"operation":"sell", "unit-cost":20.00, "quantity": 5000}, {"operation":"sell", "unit-cost":5.00, "quantity": 5000}]
```
Output:
```json
[{"tax": 0.00},{"tax": 0.00},{"tax": 0.00}]
[{"tax": 0.00},{"tax": 10000.00},{"tax": 0.00}]
```

## Use case #3
| Operation | Unit-cost | Quantity | Tax paid | Explanation                                                                                    |
|:----------|:----------|:---------|:---------|:-----------------------------------------------------------------------------------------------|
| buy       | 10.00     | 10000    | 0        | Buying shares does not pay tax                                                                 |
| sell      | 5.00      | 5000     | 0        | Loss of R$25,000: no tax paid                                                                  |
| sell      | 20.00     | 3000     | 1000     | Profit of R$30,000: You must deduct a loss of R$25,000 and pay 20% of R$5,000 in tax (R$1,000) |

Input:
```json
[{"operation":"buy", "unit-cost":10.00, "quantity": 10000}, {"operation":"sell", "unit-cost":5.00, "quantity": 5000}, {"operation":"sell", "unit-cost":20.00, "quantity": 3000}]
```
Output:
```json
[{"tax": 0.00},{"tax": 0.00},{"tax": 1000.00}]
```

## Use case #4
| Operation | Unit-cost | Quantity | Tax paid | Explanation                                                                                              |
|:----------|:----------|:---------|:---------|:---------------------------------------------------------------------------------------------------------|
| buy       | 10.00     | 10000    | 0        | Buying shares does not pay tax                                                                           |
| buy       | 25.00     | 5000     | 0        | Buying shares does not pay tax                                                                           |
| sell      | 15.00     | 10000    | 0        | Considering a weighted average price of R$ 15 ((10×10000 + 25×5000) ÷ 15000) there was no profit or loss |

Input:
```json
[{"operation":"buy", "unit-cost":10.00, "quantity": 100}, {"operation":"sell", "unit-cost":15.00, "quantity": 50}, {"operation":"sell", "unit-cost":15.00, "quantity": 50}]
```
Output:
```json
[{"tax": 0},{"tax": 0},{"tax": 0}]
```

## Use case #5
| Operation | Unit-cost | Quantity | Tax paid   | Explanation                                                                                            |
|:----------|:----------|:---------|:-----------|:-------------------------------------------------------------------------------------------------------|
| buy       | 10.00     | 10000    | 0          | Buying shares does not pay tax                                                                         |
| buy       | 25.00     | 5000     | 0          | Total value less than R$20.000                                                                         |
| sell      | 15.00     | 10000    | 0          | Considering a weighted average price of R$15, there was no profit or loss                              |
| sell      | 25.00     | 5000     | 10000      | Considering a weighted average price of R$15 profit of R$50,000: pay 20% of R$50,000 in tax (R$10,000) |

Input:
```json
[{"operation":"buy", "unit-cost":10.00, "quantity": 10000}, {"operation":"buy", "unit-cost":25.00, "quantity": 5000}, {"operation":"sell", "unit-cost":15.00, "quantity": 10000}, {"operation":"sell", "unit-cost":25.00, "quantity": 5000}]
```
Output:
```json
[{"tax": 0.00},{"tax": 0.00},{"tax": 0.00},{"tax": 10000.00}]
```

## Use case #6
| Operation | Unit-cost | Quantity | Tax paid | Explanation                                                                                                                    |
|:----------|:----------|:---------|:---------|:-------------------------------------------------------------------------------------------------------------------------------|
| buy       | 10.00     | 10000    | 0        | Buying shares does not pay tax                                                                                                 |
| sell      | 2.00      | 5000     | 0        | Loss of R$40,000: total value is less than R$20,000, but we must deduct the loss regardless of this                            |
| sell      | 20.00     | 2000     | 0        | Profit of R$20000: if you deduct the loss, your profit is R$0. You still have R$20000 of loss to deduct from your next profits |
| sell      | 20.00     | 2000     | 0        | Profit of R$20,000: if you deduct the loss, your profit is R$0. Now there is no more loss to deduct from your next profits     |
| sell      | 25.00     | 1000     | 3000     | Profit of R$15,000 and no losses: pay 20% of R$15,000 in tax (R$3,000)                                                         |

Input:
```json
[{"operation":"buy", "unit-cost":10.00, "quantity": 10000}, {"operation":"sell", "unit-cost":2.00, "quantity": 5000}, {"operation":"sell", "unit-cost":20.00, "quantity": 2000}, {"operation":"sell", "unit-cost":20.00, "quantity": 2000}, {"operation":"sell", "unit-cost":25.00, "quantity": 1000}]
```
Output:
```json
[{"tax": 0.00},{"tax": 0.00},{"tax": 0.00},{"tax": 0.00},{"tax": 3000.00}]
```

## Use case #7
| Operation | Unit-cost | Quantity | Tax paid | Explanation                                                                                                                            |
|:----------|:----------|:---------|:---------|:---------------------------------------------------------------------------------------------------------------------------------------|
| buy       | 10.00     | 10000    | 0        | Buying shares does not pay tax                                                                                                         |
| sell      | 2.00      | 5000     | 0        | Loss of R$40,000: total value is less than R$20,000, but we must deduct the loss regardless of this                                    |
| sell      | 20.00     | 2000     | 0        | Profit of R$20000: if you deduct the loss, your profit is R$0. You still have R$20000 of loss to deduct from your next profits         |
| sell      | 20.00     | 2000     | 0        | Profit of R$20,000: if you deduct the loss, your profit is R$0. Now there is no more loss to deduct from your next profits             |
| sell      | 25.00     | 1000     | 3000     | Profit of R$15,000 and no losses: pay 20% of R$15,000 in tax (R$3,000)                                                                 |
| buy       | 20.00     | 10000    | 0        | All previous shares were sold. Buying new shares changes the weighted average to the amount paid for them (R$20)                       |
| sell      | 15.00     | 5000     | 0        | Loss of R$25,000                                                                                                                       |
| sell      | 30.00     | 4350     | 3700     | Profit of R$43,500: if you deduct the loss of R$25,000, you will be left with R$18,500 in profit. Pay 20% of R$18,500 in tax (R$3,700) |
| sell      | 30.00     | 650      | 0        | Profit of R$6500, no loss to deduct, but the total value is less than R$20000, so no tax is paid                                       |

Input:
```json
[{"operation":"buy", "unit-cost":10.00, "quantity": 10000}, {"operation":"sell", "unit-cost":2.00, "quantity": 5000}, {"operation":"sell", "unit-cost":20.00, "quantity": 2000}, {"operation":"sell", "unit-cost":20.00, "quantity": 2000}, {"operation":"sell", "unit-cost":25.00, "quantity": 1000}, {"operation":"buy", "unit-cost":20.00, "quantity": 10000}, {"operation":"sell", "unit-cost":15.00, "quantity": 5000}, {"operation":"sell", "unit-cost":30.00, "quantity": 4350}, {"operation":"sell", "unit-cost":30.00, "quantity": 650}]
```
Output:
```json
[{"tax":0.00}, {"tax":0.00}, {"tax":0.00}, {"tax":0.00}, {"tax":3000.00}, {"tax":0.00}, {"tax":0.00}, {"tax":3700.00}, {"tax":0.00}]
```

## Use case #8
| Operation | Unit-cost | Quantity | Tax paid | Explanation                                                  |
|:----------|:----------|:---------|:---------|:-------------------------------------------------------------|
| buy       | 10.00     | 10000    | 0        | Buying shares does not pay tax                               |
| sell      | 50.00     | 10000    | 80000    | Profit of R$400,000: pays 20% of R$400,000 in tax (R$80,000) |
| buy       | 20.00     | 10000    | 0        | Buying shares does not pay tax                               |
| sell      | 50.00     | 10000    | 60000    | Profit of R$300,000: pays 20% of R$300,000 in tax (R$60,000) |

Input:
```json
[{"operation":"buy", "unit-cost":10.00, "quantity": 10000}, {"operation":"sell", "unit-cost":50.00, "quantity": 10000}, {"operation":"buy", "unit-cost":20.00, "quantity": 10000}, {"operation":"sell", "unit-cost":50.00, "quantity": 10000}]
```
Output:
```json
[{"tax":0.00},{"tax":80000.00},{"tax":0.00},{"tax":60000.00}]
```
