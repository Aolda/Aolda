function add(a,b) {
    a = Number(a);
    b = Number(b);
    return a+b;
}
function mod(a,b){
    a = Number(a);
    b = Number(b);
    return a%b;
}
function div(a,b){
    a = Number(a);
    b = Number(b);
    return a/b;
}
function mul(a,b){
    a = Number(a);
    b = Number(b);
    return a*b;
}
function sub(a,b){
    a = Number(a);
    b = Number(b);
    return a-b;
}
module.exports = { add, mod, div, mul, sub };