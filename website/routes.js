var idCount = 0;



async function getFullConvo() {
    var url = "http://localhost:8000/messages";
    fetch(url)
        .then(res => res.json())
        .then( data => {
            for( i=0 ; i < data.length ; i++ ) {
                x = JSON.parse(JSON.stringify(data[i]));
                if( x.ID > idCount){
                    idCount ++; 
                    document.getElementById("chatbox")
                        .value += x.Message.From + ": " + x.Message.Message + "\n";
                }
            }
        })

}

function sendMessage(messageBox, name){
    var message = messageBox.value;
    var from = name.value || "Anonymous"
    var completeMessage = `{"From":"${from}", "Message":"${message}"}`
    console.log(completeMessage);
    

    fetch("http://localhost:8000/newMessage", {
        method: 'POST',
        body: completeMessage,
        headers: { "Content-Type": "application/json"}
    })
}
