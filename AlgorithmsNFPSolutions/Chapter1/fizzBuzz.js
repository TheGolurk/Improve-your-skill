'use strict';

let number = [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15];

for (let num of number) {

	if (num % 3 == 0 && num % 5 == 0) {
		console.log(`${num} fizz buzz`);
	} else if (num % 3 == 0) {
		console.log(`${num} fizz`);
	} else if (num % 5 == 0) {
		console.log(`${num} buzz`);
	} else {
		console.log(num);
	}

}
