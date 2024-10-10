const banners = document.querySelector("#banners");
const inputField = document.querySelector("#inputField");
const convertButton = document.querySelector("#convertButton");
const ascii = document.querySelector(".ascii");

var prevInput = inputField.value;
var prevBanner = banners.value;

const getBanners = async () => {
  const res = await fetch(window.location.origin + "/apis/v1/ascii/banners");
  if (!res.ok) {
    // TODO: notify user
    return;
  }

  banners.innerHTML = "";

  const bannersList = (await res.text()).split(",");

  bannersList.forEach((banner) => {
    if (banner === "") {
      return;
    }
    const option = document.createElement("option");
    option.value = banner;
    option.textContent = banner;
    banners.appendChild(option);
  });
};
getBanners();

const handleConvertClick = async () => {
  if (
    inputField.value === "" ||
    (inputField.value === prevInput && banners.value === prevBanner)
  ) {
    return;
  }

  prevInput = inputField.value;
  prevBanner = banners.value;

  const res = await fetch(window.location.origin + "/api/v1/ascii/", {
    method: "POST",
    body: `input=${encodeURIComponent(inputField.value)}&banner=${
      banners.value
    }`,
  });
  ascii.textContent = await res.text();
};

banners.addEventListener("change", () => {
  handleConvertClick();
});

convertButton.addEventListener("click", handleConvertClick);
const clear = document.querySelector("#clear");

clear.addEventListener("click", () => {
  inputField.value = "";
});
