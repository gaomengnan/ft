// #DModel 跟随鼠标旋转功能
htmx.onLoad(function (content) {
  const astronaut = document.getElementById("astronaut")
  document.addEventListener("mousemove", (r) => {
    // astronaut.cameraOrbit = `${azimuth}deg ${elevation}deg ${initialDistance}m`;
    let t = -(r.x / window.innerWidth - .5) * Math.PI / 2
      , e = Math.max(Math.PI / 4, -(r.y / window.innerHeight - .5) * Math.PI / 2 + Math.PI / 3);
    astronaut.cameraOrbit = `${t}rad ${e}rad 12m`
  })




})
htmx.toggleTab = function toggleActiveTab(t) {
  document.querySelectorAll('.header_nav__links>a').forEach(item => {
    item.classList.remove('nav-active');
  });
  if (!t.classList.contains('nav-active')) {
    t.classList.toggle('nav-active');
  }
}
