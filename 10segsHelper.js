hard= [365235, 6452581, 1089563, 4956199]
medium=[664647, 2187, 5732938]
easy=[654030, 616554, 3139189]

mediumEfforts = [];		// aggregate segment efforts by difficulty
easyEfforts = [];
hardEfforts = [];

var getEfforts = function(segId) {
	return require('./dist/segmentEfforts_' + segId + '.json');
}

hard.forEach(function(segId) {
	var moreSegs = require('./dist/hard/segmentEfforts_' + segId + '.json');
	hardEfforts = hardEfforts.concat(moreSegs);
});

medium.forEach(function(segId) {
	var moreSegs = require('./dist/medium/segmentEfforts_' + segId + '.json');
	mediumEfforts = mediumEfforts.concat(moreSegs);
});

easy.forEach(function(segId) {
	var moreSegs = require('./dist/easy/segmentEfforts_' + segId + '.json');
	easyEfforts = easyEfforts.concat(moreSegs);
});

var obj = {
	easy: easyEfforts,
	medium: mediumEfforts,
	hard: hardEfforts
}

function mean(values) {
	var sum = values.reduce(function(sum, value){
	  return sum + value;
	}, 0);

	var avg = sum / values.length;
	return avg;
}

function psuedoMedian(values) {
	return values.sort()[Math.round(values.length/2)];
}


function standardDeviation(values){
  var avg = mean(values);
  
  var squareDiffs = values.map(function(value){
    var diff = value - avg;
    var sqrDiff = diff * diff;
    return sqrDiff;
  });
  
  var avgSquareDiff = mean(squareDiffs);

  var stdDev = Math.sqrt(avgSquareDiff);
  return stdDev;
}

var getStuff = function(prop) {
	console.log(prop + ":");

	var processDifficulty = function(difficulty) {
		var arr = obj[difficulty].map(function(nextEffort) {
			return nextEffort[prop];
		});

		var unfilteredCount = arr.length;

		arr = arr.filter(function(propResults) {
			return propResults;
		});

		var filteredCount = arr.length;

		console.log("\t" + difficulty + " (" + filteredCount + "/" + unfilteredCount + "):");


		var avg = mean(arr);
		var median = psuedoMedian(arr);
		var stdDev = standardDeviation(arr);

		console.log("\t\tmean: " + avg.toFixed(2) + ", median: " + median.toFixed(2) + ", std. dev.: " + stdDev.toFixed(2));
	}

	processDifficulty("easy");
	processDifficulty("medium");
	processDifficulty("hard");
}

var props = [
	"average_watts",
	"moving_time",
	"distance",
	"average_cadence",
	"average_watts",
	"average_heartrate",
	"max_heartrate"
]

props.forEach(function(nextProp) {
	getStuff(nextProp);
});

