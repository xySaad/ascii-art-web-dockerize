const convertButton = document.querySelector("#convertButton");
const ascii = document.querySelector(".ascii");
var converted = false;

const handleConvertClick = async () => {
  converted = true;
  const inputField = document.querySelector("#inputField");
  const banners = document.querySelector("#banners");
  const res = await fetch(window.location.origin + "/api/v1/ascii/", {
    method: "POST",
    body: `input=${encodeURIComponent(inputField.value)}&banner=${
      banners.value
    }`,
  });
  ascii.textContent = await res.text();
};
const banners = document.querySelector("#banners");
banners.addEventListener("change", () => {
  if (converted) {
    handleConvertClick();
  }
});
convertButton.addEventListener("click", handleConvertClick);
const clear = document.querySelector("#clear");

clear.addEventListener("click", () => {
  converted = false;
  const inputField = document.querySelector("#inputField");
  inputField.value = "";
});
