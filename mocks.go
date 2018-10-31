package main

var (
	INPUT1 = `
			/* comment should not be scanned */
			let five = "test";
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
			}
			`
	INPUT2 = `let five = "te
	st";`

	INPUT3 = `/*`
	INPUT4 = `""" this is also a comment?
				"""`
	INPUT5 = `"invalid \q escape character"`
)
