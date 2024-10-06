const convertButton = document.querySelector("#convertButton");
const ascii = document.querySelector(".ascii");
convertButton.addEventListener("click", async () => {
  const inputField = document.querySelector("#inputField");
  const res = await fetch(window.location.origin + "/api/ascii", {
    method: "POST",
    body: `input=${encodeURIComponent(inputField.value)}&banner=standard`,
  });
  console.log(ascii);
  ascii.textContent = await res.text();
});
