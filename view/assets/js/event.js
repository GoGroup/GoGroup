function readURL(input) {
    if (input.files && input.files[0]) {
      var reader = new FileReader();
      
      reader.onload = function(e) {
        $('#img').attr('src', e.target.result);
      }
      
      reader.readAsDataURL(input.files[0]);
    }
  }
  $("#imgfile").change(function() {
    readURL(this);
  });
  $(function () {
    $('#datetimepicker10').datetimepicker({
        viewMode: 'years'
    });
});


  
