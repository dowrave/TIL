let init = () => {
	return function(a, b) {
		return a - b > 0 ? a - b : b - a;	
	}
}

document.write(init(30, 20));
document.write(init()(30, 20));