$(function() {
 if (window["WebSocket"]) {
        var width = 0.0;
        {{ if .Dev }}
          conn = new WebSocket("ws://{{.WSHost}}/ws");
        {{ else }}
          conn = new WebSocket("wss://{{.WSHost}}/ws");
        {{ end }}
        conn.onclose = function (evt) {
        };
        conn.onmessage = function (evt) {
          var resp = JSON.parse(evt.data);
          switch (resp.action) {
            // Do some things here
          }
        };
    }
});
