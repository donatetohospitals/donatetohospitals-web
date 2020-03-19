$(function () {
    console.log('supplier js loaded');

    $("#btnAddItem").click(function(event){
        event.preventDefault();
        var newCard = $('.card-item').last().clone();

        $('.card-item').last().after(newCard);

        $(".card-item").last().find('input').val('');
        $("#target").val($("#target option:first").val());
        $('.card-item').last().find('h3').text("Item " + $(".card-item").length)
    })

    function buildItems(cards){
        return $.map(cards, function(card){
            var deserialized = $(card).find(':input').serializeArray();
            return {
                name: deserialized[0].value,
                count: deserialized[1].value ? 0 : Number(deserialized[1].value),
                condition: deserialized[2].value
            }
        })
    }

    $('#supplierForm').on('submit', function(e) { //use on if jQuery 1.7+
        e.preventDefault();  //prevent form from submitting

        var data = $("#supplierForm :input").serializeArray();

        var normalized = {email: data[0].value, geo: data[1].value, items: buildItems($('.card-item'))};
        sendData(normalized, function(){
            alert('Submission successful. You will be emailed shipping information when a match is made with a hospital in need. Please consider spreading the word on social media. THANK YOU FOR SAVING LIVES.')
            window.location = '/'
        }, function(errMsg){
            console.error('error', errMsg);
            alert(errMsg && errMsg.responseJSON && errMsg.responseJSON.error);
        })

    });

    function sendData(data, handleSuccess, handleError){
        $.ajax({
            type: "POST",
            url: "/suppliers/json",
            data: JSON.stringify(data),
            contentType: "application/json",
            dataType: "json",
            processData: false,
            success: handleSuccess,
            error: handleError
        });
    }
});