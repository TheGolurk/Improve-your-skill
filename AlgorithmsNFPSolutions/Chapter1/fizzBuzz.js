'use strict';

let number = [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,20,21,22,23,24,25,26,27];

for (let num of number) {

	// (num % 3 == 0 && num % 5 == 0) Not Optimized
	if (num % 15 == 0) {
		console.log(`${num} fizz buzz`);
	} else if (num % 3 == 0) {
		console.log(`${num} fizz`);
	} else if (num % 5 == 0) {
		console.log(`${num} buzz`);
	} else {
		console.log(num);
	}

}
