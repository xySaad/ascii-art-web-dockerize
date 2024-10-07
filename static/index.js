const convertButton = document.querySelector("#convertButton");
const ascii = document.querySelector(".ascii");
convertButton.addEventListener("click", async () => {
  const inputField = document.querySelector("#inputField");
  const banners = document.querySelector("#banners");
  const res = await fetch(window.location.origin + "/api/v1/ascii", {
    method: "POST",
    body: `input=${encodeURIComponent(inputField.value)}&banner=${
      banners.value
    }`,
  });
  console.log(ascii);
  ascii.textContent = await res.text();
});
