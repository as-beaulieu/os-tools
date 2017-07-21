var os = require('os');
var fs = require('fs');
var path = require('path');

//can use os to get number of processors, computer info, etc.
console.log(os.homedir());

var memInfo = os.freemem(); //Displays free memory of system
var a = os.networkInterfaces(); //an object of the available network interfaces
var plat = os.platform(); //the platform being run
var b = os.userInfo(); //info about the current user logged in
var d = [memInfo, a, plat, b];



//make a new file, to the desktop, naming it 'newfile.txt', passing data 'hi there', and checking for error
fs.writeFile(path.join(os.homedir(), 'Desktop', 'newfile.txt'), d, function(err) {
    if (err) {
        console.log(err);
    }
});