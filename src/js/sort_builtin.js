var fs = require('fs');

/////////////////////
// read input file //
/////////////////////

let inputFilePath = process.argv[2];
let inputString;
try {
  inputString = fs.readFileSync(inputFilePath, 'utf8');
} catch (err) { throw err }
let inputStringSplit = inputString.split(/\n/);

////////////////
// sort input //
////////////////

// drop the trailing newline, so that it doesn't get included in the sort
if (inputStringSplit[inputStringSplit.length - 1] == "") {
  inputStringSplit = inputStringSplit.slice(0, inputStringSplit.length - 1);
}
inputStringSplit.sort();

///////////////////////
// write output file //
///////////////////////

let outputFilePath = process.argv[3];
let outputString = inputStringSplit.join("\n");
outputString += "\n"; // trailing newline
try {
  fs.writeFileSync(outputFilePath, outputString);
} catch (err) { throw err }