var _ = require("underscore");

var a1 = {
    "1": {
        "2": {
            "3": 4,
        },
    },
};

var a2 = _.extend({}, a1, a1);

console.log(a1);
console.log(a2);

a2[1][2][3] = 5;
console.log(a1);
console.log(a2);
