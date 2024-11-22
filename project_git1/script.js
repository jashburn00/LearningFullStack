const words = [
	['the', 'I\'m', 'Oh', 'Don\'t'] // the 
	['Russians', 'not', 'my', 'forget'] // russians
	['are', ] // are
	[] // nuking 
	[] // us
]
/*
 * the russians are nuking us
 * Im not fat im big boned
 * oh my god they killed kenny
 * dont forget to bring a towel
 * */

function generate(){
	for(let i = 0; i < 6; i++){
		pickRandom(i);
	}
}
