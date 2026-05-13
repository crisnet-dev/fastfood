const websocket = new WebSocket("/ws?password=Hjx33@");

websocket.onopen = () => {
    document.querySelector("h1").textContent = "Connected";
}

websocket.onmessage = (data) => {
    const message = JSON.parse(data.data);

    console.log(message);

    if (message.type != "UpdateAdminCounter") {
        document.querySelector("#msg").textContent = `Novo pedido\nVeja no console`;
    }

}

websocket.onclose = () => {
    document.querySelector("h1").textContent = "Disconnected";
}

