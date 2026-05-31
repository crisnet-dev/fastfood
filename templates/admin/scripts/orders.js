const websocket = new WebSocket("/ws?password=Hj$Xx33@");

function searchOrder() {
    const input = document.querySelector('');
}

function listOrders(data) {
    const orderListContainer = document.querySelector('.order-list-container');

    data.forEach((order) => {
        const btnDetails = document.createElement("button");
        const orderCard = document.createElement("div");
        const orderInfo = document.createElement("div");
        const avatar = document.createElement("div");
        const name = document.createElement("span");
        const time = document.createElement("span");

        const now = new Date();

        name.textContent = order.name;
        time.textContent = `${now.getHours()}:${now.getMinutes()}`;
        btnDetails.textContent = "Ver Detalhes";

        orderCard.classList.add('order-card');
        orderInfo.classList.add('order-info');
        btnDetails.classList.add('btn-info');
        avatar.classList.add('avatar');

        orderInfo.appendChild(avatar);
        orderInfo.appendChild(name);
        orderInfo.appendChild(time);

        orderCard.appendChild(orderInfo);
        orderCard.appendChild(btnDetails);

        orderListContainer.appendChild(orderCard);
    });
}

function addOrder(order) {
    const orderListContainer = document.querySelector('.order-list-container');

    const btnDetails = document.createElement("button");
    const orderCard = document.createElement("div");
    const orderInfo = document.createElement("div");
    const avatar = document.createElement("div");
    const name = document.createElement("span");
    const time = document.createElement("span");

    const now = new Date();

    name.textContent = order.name;
    time.textContent = `${now.getHours()}:${now.getMinutes()}`;
    btnDetails.textContent = "Ver Detalhes";

    orderCard.classList.add('order-card');
    orderInfo.classList.add('order-info');
    btnDetails.classList.add('btn-info');
    avatar.classList.add('avatar');

    orderInfo.appendChild(avatar);
    orderInfo.appendChild(name);
    orderInfo.appendChild(time);

    orderCard.appendChild(orderInfo);
    orderCard.appendChild(btnDetails);

    orderListContainer.appendChild(orderCard);
}

websocket.onopen = () => {
    // alert("Connected");
    document.querySelector('#status').textContent = "Connectado";
}

websocket.onmessage = (data) => {
    const message = JSON.parse(data.data);

    if (message.type != "UpdateAdminCounter") {
        // alert(`Novo pedido\nVeja no console`);

        if (message?.orders) {
            listOrders(message?.orders);
        } else {
            addOrder(message);
        }
    }

    if (message.type === "UpdateAdminCounter") {
        document.querySelector('#admin-counter').textContent = `${message.admin_counter} Admins`;
    }

}

websocket.onclose = (err) => {
    document.querySelector('#status').textContent = "Não connectedo";
    console.log(err);
}
