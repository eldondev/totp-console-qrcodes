** This is ALPHA software. Before wiping/removing/relying on this software to generate your codes, test them repeatedly with each device to ensure compatibility ** 
This little go code creates QR codes that Google Authenticator can scan from an existing sqlite3 database.
It uses mdp's library to write a QR code to the terminal, and waits for the user to press enter to get the next code.
It totally punts by shelling out and piping to the sqlite3 executable.

To use it you must pass a flag promising to test your codes before removing them from other devices or etc.

You have been warned.

