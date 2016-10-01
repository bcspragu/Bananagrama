$(function() {
  $(".pass-submit").submit(function(e) {
    e.preventDefault()
    var form = $(this);
    $.post("/pass", form.serialize(), function(data) {
      if (data.Valid) {
        var body = $("body");
        body.html(data.Game);
        body.append("<script>" + data.JS + "</script>");
      } else {
        alert(data.Error);
      }
    });
  });
});
