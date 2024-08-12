# supermarket_checkout

## To run on windows:
Run `checkout\checkout.exe`

The program expects the user to 'scan' products by entering a sequence of product names when prompted. For example:
`A A B B A C`

After pressing the Enter key, a total price for the 'scanned' products will be outputted.

SKUs can be manually updated by updating the file `conf.yaml`. The program expects this file to exist and stay in the correct format.

### To run unit tests:
Run the following:
```
cd .\checkout
```
```
go test
```