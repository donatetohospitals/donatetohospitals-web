$(function () {
    console.log('supplier js loaded');

    // Initialize the Amazon Cognito credentials provider
    // AWS.config.region = 'us-east-2'; // Region
    // AWS.config.credentials = new AWS.CognitoIdentityCredentials({
    //     IdentityPoolId: 'us-east-2:827f8e02-915d-458e-be4f-0e89644771d6',
    // });

    var albumBucketName = "dthbucket";
    var bucketRegion = "us-east-2";
    var IdentityPoolId = "us-east-2:827f8e02-915d-458e-be4f-0e89644771d6";

    AWS.config.update({
        region: bucketRegion,
        credentials: new AWS.CognitoIdentityCredentials({
            IdentityPoolId: IdentityPoolId
        })
    });

    var s3 = new AWS.S3({
        apiVersion: "2006-03-01",
        params: { Bucket: albumBucketName }
    });

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

    function addPhoto() {
        var files = document.getElementById("supplyImage1").files;
        if (!files.length) {
            return alert("Please choose a file to upload first.");
        }
        var file = files[0];
        var fileName = file.name;
        // var albumPhotosKey = encodeURIComponent('supplies') + "//";

        // var photoKey = albumPhotosKey + fileName;

        // Use S3 ManagedUpload class as it supports multipart uploads
        var upload = new AWS.S3.ManagedUpload({
            params: {
                Bucket: albumBucketName,
                Key: fileName,
                Body: file,
                ACL: "public-read"
            }
        });

        var promise = upload.promise();

        promise.then(
            function(data) {
                alert("Successfully uploaded photo.");
                console.log('success', data)
            },
            function(err) {
                console.log('err', err)
                return alert("There was an error uploading your photo: ", err);
            }
        );
    }

    $('#supplierForm').on('submit', function(e) { //use on if jQuery 1.7+
        e.preventDefault();  //prevent form from submitting

        var data = $("#supplierForm :input").serializeArray();

        var normalized = {email: data[0].value, geo: data[1].value, items: buildItems($('.card-item'))};
        // sendData(normalized, function(){
        //     alert('Submission successful. You will be emailed shipping information when a match is made with a hospital in need. Please consider spreading the word on social media. THANK YOU FOR SAVING LIVES.')
        //     window.location = '/'
        // }, function(errMsg){
        //     console.error('error', errMsg);
        //     alert(errMsg && errMsg.responseJSON && errMsg.responseJSON.error);
        // })

        addPhoto()
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