package main

var INPUT = `
				let five = 5;
				let ten = 10;
				let add = fn(x, y) {
					x + y;
				};

				let result = add(five, ten);  
				5 < 10 > 5;

				if (5 < 10) {
					return true;
				} else {
					return false;
				}

				10 == 10;

				class Pt(x: Int, y: Int) {
					thisx = x;
					thisy = y;
		
					def _x() : Int { return thisy; }
				}`