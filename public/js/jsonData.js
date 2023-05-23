import PrettyJSON from "./file";

function updateData(data) {
  var dataContainer = document.getElementById("data-container");
  // dataContainer.innerHTML = JSON.stringify(data, "", 2);
  var node = new PrettyJSON.view.Node({
    el: dataContainer,
    data: data
  });
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
    // Update the HTML with the fetched data
    updateData(data);
  })
  .catch(function (error) {
    console.error("Error:", error);
  });
