$(function () {
    console.log('supplier js loaded')

    $("#btnAddItem").click(function(event){
        event.preventDefault()
        var newCard = $('.card-item').last().clone()

        $('.card-item').last().after(newCard)

        $(".card-item").last().find('input').val('')
        $("#target").val($("#target option:first").val());
        $('.card-item').last().find('h3').text("Item " + $(".card-item").length)
    })
});