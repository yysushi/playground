var _ = require("underscore");

var a1 = {
    "1": {
        "2": 3,
    },
};

var a2 = _.extend({}, a1);

console.log(a1);
console.log(a2);

a2[1][2] = 4;
console.log(a1);
console.log(a2);

var b1 = {
    "1": 2,
};

var b2 = _.extend({}, b1);

console.log(b1);
console.log(b2);

b2[1] = 4;
console.log(b1);
console.log(b2);

var c1 = {
    "1": {
        "2": 3,
    },
};

var c2 = _.clone(c1);

console.log(c1);
console.log(c2);

c2[1][2] = 4;
console.log(c1);
console.log(c2);
