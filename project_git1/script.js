const words = [
    ['The', 'I\'m', 'Oh', 'Don\'t'],
    ['Russians', 'not', 'my', 'forget'],
    ['are', 'fat', 'god', 'to'],
    ['nuking', 'I\'m', 'they', 'bring'],
    ['us', 'big', 'killed', 'a'],
    ['', 'boned', 'Kenny', 'towel']
];
/*
 * the russians are nuking us ''
 * Im not fat im big boned
 * oh my god they killed kenny
 * dont forget to bring a towel
 * */

function generate(){
	let msg = '';
	for(let i = 0; i < 6; i++){
		msg = msg + ' ' + words[i][Math.floor(Math.random()*4)];
	}

	document.getElementById('msg').innerHTML = msg;
}
