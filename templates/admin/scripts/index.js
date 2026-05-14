// const websocket = new WebSocket("/ws?password=Hjx33@");

// websocket.onopen = () => {
//     alert("Connected");
// }

// websocket.onmessage = (data) => {
//     const message = JSON.parse(data.data);

//     console.log(message);

//     if (message.type != "UpdateAdminCounter") {
//         alert(`Novo pedido\nVeja no console`);
//     }

// }

// websocket.onclose = (err) => {
//     alert("Disconnected");
//     console.log(err);

// }

const formData = new FormData()

const fileInput = document.getElementById("f");

formData.append(
    "data",
    JSON.stringify({
        id: 1,
        product_name: "crisnet",
        price: 1500,
        image_url: ""
    })
)

function d() {
    formData.append("image", fileInput.files[0])

    fetch("/product/upload", {
        method: "POST",
        body: formData
    }).then(r => r.json()).then(d => console.log(d))
}