const button = document.querySelector("button");
const registry = document.querySelector(".hero .input_fields #registry");
const repo = document.querySelector(".hero .input_fields #repo");
const tag = document.querySelector(".hero .input_fields #tag");
const dataContainer = document.getElementById("data-container");

function updateData(data) {
  dataContainer.innerHTML = data;
}

function fetchData(data) {
  dataContainer.innerText = "loading..."
  // Fetch data from the API
  fetch("http://localhost:8080/api/data", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
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
      dataContainer.innerText = "Nothing to show -_-"
      alert("Error: Cannot fetch OCI content, try changing the artifact details");
    });
}

button.addEventListener("click", () => {
  if (!registry.value || !repo.value || !tag.value)
    alert("Please fill all fields to search content");

  const data = {
    registry: registry.value,
    repo: repo.value,
    tag: tag.value,
  };

  console.log(data)
  fetchData(data);
});
