function execFn(fileName, funcName, args){
    fileName = require(`./src/${fileName}`);
    exResult = fileName[funcName].apply(this, args);
    return exResult;
}
console.log(process.argv);

fileName = process.argv[2];
funcName = process.argv[3];

result = execFn(fileName, funcName, process.argv[4]);

console.log(result);