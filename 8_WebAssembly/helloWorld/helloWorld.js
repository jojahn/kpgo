var fs = require("fs")
var buffer = fs.readFileSync("./lib.html")

WebAssembly.instantiate(buffer).then(res => {
    console.log(res.instance)
});
