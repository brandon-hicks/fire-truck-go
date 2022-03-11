# fire-truck-go
This app was designed to do calculations that are vital to fire ground operations. These include calculating friction loss and total pump discharge pressure that a fire apparatus is to be pumped at.

This app will prompt the user to enter in the calculation that they wish to make.

From here the user will be prompted to enter in the values needed for the calculation.
If the user enters in an incorrect coefficient for the nozzle it will default to the coefficient for a smooth bore hand-line which is 50.
If the user enters that they are using a fog nozzle then the app will skip over the pieces that require the user to enter in values that only pertain to smooth-bore nozzles.

To run this app simply go to the terminal and type `go run .`