const darkIcon = document.querySelector("#theme-switcher span#dark");

darkIcon.addEventListener("click", () => {
  if (document.body.id == "dark") {
    document.body.id = "white";
  } else {
    document.body.id = "dark";
  }
});
