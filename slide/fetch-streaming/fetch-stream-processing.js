var query = 'gopher';
fetch('/big-data.csv').then(function(response) { // HL
  var reader = response.body.getReader(); // HL
  var partialRecord = '';
  var decoder = new TextDecoder();

  function search() {
    return reader.read().then(function(result) { // HL
      partialRecord += decoder.decode(result.value || new Uint8Array, { stream: !result.done }); // HL
      // query logic ...
      // Call reader.cancel("No more reading needed."); when result found early. // HL
      if (result.done) {
        throw Error("Could not find value after " + query);
      }
      return search(); // HL
    })
  }

  return search(); // HL
}).then(function(result) {
  console.log("Got the result! It's '" + result + "'");
}).catch(function(err) {
  console.log(err.message);
});
