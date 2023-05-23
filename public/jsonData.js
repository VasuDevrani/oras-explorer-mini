import jsonFormat from 'json-format';

function updateData(data) {
  var dataContainer = document.getElementById("data-container");
  dataContainer.innerHTML = "<p>" + jsonFormat(data, { type: 'space', size: 2 }) + "</p>";
}

// Fetch data from the API
fetch("http://localhost:8080/api/data")
  .then(function (response) {
    if (response.ok) {
      return response.json();
    }
    throw new Error("Network response was not OK.");
  })
  .then(function (data) {
    console.log(data);
    // Update the HTML with the fetched data
    updateData(data);
  })
  .catch(function (error) {
    console.error("Error:", error);
  });