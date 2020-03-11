var _ = require("underscore");

var a = [1, 2, 3];
var b = [   2, 3, 4];

var c = _.union(_.difference(a, b), _.difference(b, a));
console.log(c);
